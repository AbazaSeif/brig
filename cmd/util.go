package cmd

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	log "github.com/Sirupsen/logrus"
	"github.com/VividCortex/godaemon"
	"github.com/disorganizer/brig/brigd/client"
	"github.com/disorganizer/brig/cmd/pwd"
	"github.com/urfave/cli"
)

// ExitCode is an error that maps the error interface to a specific error
// message and a unix exit code
type ExitCode struct {
	Code    int
	Message string
}

func (err ExitCode) Error() string {
	return err.Message
}

// guessRepoFolder tries to find the repository path
// by using a number of sources.
func guessRepoFolder() string {
	path := os.Getenv("BRIG_PATH")
	if path == "" {
		return "."
	}

	return path
}

func readPasswordFromArgs(ctx *cli.Context) string {
	for curr := ctx; curr != nil; {
		if curr.Bool("no-pass") {
			return "no-pass"
		}

		if password := curr.String("password"); password != "" {
			return password
		}

		curr = curr.Parent()
	}

	return ""
}

func readPassword(ctx *cli.Context, repoPath string) (string, error) {
	if !repoIsInitialized(repoPath) {
		return "", nil
	}

	// Try to read the password from -x or fallback to the default
	// password if requested by the --no-pass switch.
	if password := readPasswordFromArgs(ctx); password != "" {
		return password, nil
	}

	// Read the password from stdin:
	password, err := pwd.PromptPassword()
	if err != nil {
		return "", err
	}

	return password, nil
}

func prefixSlash(s string) string {
	if !strings.HasPrefix(s, "/") {
		return "/" + s
	}

	return s
}

type cmdHandlerWithClient func(ctx *cli.Context, ctl *client.Client) error

func startDaemon(ctx *cli.Context, repoPath string, port int) (*client.Client, error) {
	exePath, err := godaemon.GetExecutablePath()
	if err != nil {
		return nil, err
	}

	pwd, err := readPassword(ctx, repoPath)
	if err != nil {
		return nil, err
	}

	// Start a new daemon process:
	log.Info("Starting daemon from: ", exePath)
	proc := exec.Command(
		exePath, "-l", "/tmp/brig.log", "-x", pwd, "daemon", "launch",
	)

	if err := proc.Start(); err != nil {
		log.Infof("Failed to start the daemon: %v", err)
		return nil, err
	}

	// This will likely suffice for most cases:
	time.Sleep(100 * time.Millisecond)

	warningPrinted := false
	for i := 0; i < 15; i++ {
		ctl, err := client.Dial(context.Background(), port)
		if err != nil {
			// Only print this warning once...
			if !warningPrinted {
				log.Warnf("Waiting for daemon to bootup... :/")
				warningPrinted = true
			}
			time.Sleep(50 * time.Millisecond)
			continue
		}

		return ctl, nil
	}

	return nil, fmt.Errorf("Daemon could not be started or took to long. Wrong password maybe?")
}

func withDaemon(handler cmdHandlerWithClient, startNew bool) cli.ActionFunc {
	// If not, make sure we start a new one:
	// TODO: Make use of cli's error returning signatures.
	return withExit(func(ctx *cli.Context) error {
		port := guessPort()

		// Check if the daemon is running:
		ctl, err := client.Dial(context.Background(), port)
		if err == nil {
			return handler(ctx, ctl)
		}

		if !startNew {
			// Daemon was not running and we may not start a new one.
			return ExitCode{DaemonNotResponding, "Daemon not running"}
		}

		// Start the server & pass the password:
		ctl, err = startDaemon(ctx, guessRepoFolder(), port)
		if err != nil {
			return ExitCode{
				DaemonNotResponding,
				fmt.Sprintf("Unable to start daemon: %v", err),
			}
		}

		// Run the actual handler:
		return handler(ctx, ctl)
	})
}

type checkFunc func(ctx *cli.Context) int

func withArgCheck(checker checkFunc, handler cli.ActionFunc) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		if checker(ctx) != Success {
			os.Exit(BadArgs)
		}

		return handler(ctx)
	}
}

func withExit(handler cli.ActionFunc) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		if err := handler(ctx); err != nil {
			log.Error(err.Error())
			cerr, ok := err.(ExitCode)
			if !ok {
				os.Exit(UnknownError)
			}

			os.Exit(cerr.Code)
		}

		os.Exit(Success)
		return nil
	}
}

func needAtLeast(min int) checkFunc {
	return func(ctx *cli.Context) int {
		if ctx.NArg() < min {
			if min == 1 {
				log.Warningf("Need at least %d argument.", min)
			} else {
				log.Warningf("Need at least %d arguments.", min)
			}
			cli.ShowCommandHelp(ctx, ctx.Command.Name)
			return BadArgs
		}

		return Success
	}
}

func guessPort() int {
	envPort := os.Getenv("BRIG_PORT")
	if envPort != "" {
		// Somebody tried to set BRIG_PORT.
		// Try to parse and spit errors if wrong.
		port, err := strconv.Atoi(envPort)
		if err != nil {
			log.Fatalf("Could not parse $BRIG_PORT: %v", err)
		}

		return port
	}

	// Guess the default port.
	log.Warning("BRIG_PORT not given, assuming :6666")
	return 6666
}

func repoIsInitialized(path string) bool {
	data, err := ioutil.ReadFile(filepath.Join(path, "meta.yml"))
	if err != nil {
		return false
	}

	return len(data) > 0
}

// tempFileWithSuffix works the same as ioutil.TempFile(),
// but allows for the addition of a suffix to the filepath.
// This has the nice side effect that some editors can recognize
// the filetype based on the ending and provide you syntax highlighting.
// (this is used in edit() below)
func tempFileWithSuffix(dir, prefix, suffix string) (f *os.File, err error) {
	if dir == "" {
		dir = os.TempDir()
	}

	for i := 0; i < 10000; i++ {
		mid := strconv.Itoa(rand.Int())
		name := filepath.Join(dir, prefix+mid+suffix)
		f, err = os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_EXCL, 0600)
		if os.IsExist(err) {
			continue
		}
		break
	}
	return
}

// edit opens up $EDITOR with `data` and returns the edited data.
func edit(data []byte, suffix string) ([]byte, error) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		// It makes my heart bleed, but assume that vi is too hard
		// for the majority I've met & that might use brig.
		editor = "nano"
	}

	fd, err := tempFileWithSuffix("", "brig-cmd-buffer-", ".yml")
	if err != nil {
		return nil, err
	}

	// Make sure it gets cleaned up.
	defer func() {
		if err := fd.Close(); err != nil {
			fmt.Printf("Failed to close file: %v\n", err)
		}
		if err := os.Remove(fd.Name()); err != nil {
			fmt.Printf("Failed to remove file: %v\n", err)
		}
	}()

	if _, err := fd.Write(data); err != nil {
		return nil, err
	}

	// Launch editor and hook it up with all necessary fds:
	cmd := exec.Command(editor, fd.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		return nil, fmt.Errorf("Running $EDITOR (%s) failed: %v", editor, err)
	}

	if _, err := fd.Seek(0, os.SEEK_SET); err != nil {
		return nil, err
	}

	newData, err := ioutil.ReadAll(fd)
	if err != nil {
		return nil, err
	}

	// Some editors (including vim) might add a trailing newline:
	return bytes.TrimRight(newData, "\n"), nil
}
