package cmd

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/sahib/brig/cmd/tabwriter"

	"github.com/pksunkara/pygments"
	"github.com/sahib/brig/client"
	"github.com/sahib/brig/util/colors"
	"github.com/urfave/cli"
	yml "gopkg.in/yaml.v2"
)

func handleOffline(ctx *cli.Context, ctl *client.Client) error {
	return ctl.Disconnect()
}

func handleOnline(ctx *cli.Context, ctl *client.Client) error {
	return ctl.Connect()
}

func handleIsOnline(ctx *cli.Context, ctl *client.Client) error {
	self, err := ctl.Whoami()
	if err != nil {
		return err
	}

	if self.IsOnline {
		fmt.Println(colors.Colorize("online", colors.Green))
	} else {
		fmt.Println(colors.Colorize("offline", colors.Red))
	}

	return nil
}

func handleOnlinePeers(ctx *cli.Context, ctl *client.Client) error {
	infos, err := ctl.OnlinePeers()
	if err != nil {
		return err
	}

	tabW := tabwriter.NewWriter(
		os.Stdout, 0, 0, 2, ' ',
		tabwriter.StripEscape,
	)

	if len(infos) == 0 {
		fmt.Println("Remote list is empty. Nobody there to ping.")
		return nil
	}

	fmt.Fprintln(tabW, "NAME\tADDR\tROUNDTRIP\tLASTSEEN\t")

	for _, info := range infos {
		suffix := ""
		if info.Err == nil {
			suffix = fmt.Sprintf(
				"%s\t%s",
				info.Roundtrip,
				colors.Colorize(
					"✔ "+info.LastSeen.Format(time.Stamp),
					colors.Green,
				),
			)
		} else {
			suffix = fmt.Sprintf(
				"∞\t%s",
				colors.Colorize("✘ "+info.Err.Error(), colors.Red),
			)
		}

		shortAddr := info.Addr
		if len(shortAddr) > 9 {
			shortAddr = shortAddr[:9]
		}

		fmt.Fprintf(tabW, "%s\t%s\t%s\t\n", info.Name, shortAddr, suffix)
	}

	return tabW.Flush()
}

const (
	RemoteHelpText = `# No remotes yet. Uncomment the next lines for an example:
# - Name: alice@wonderland.com
#   Fingerprint: QmVA5j2JHPkDTHgZ[...]:SEfXUDeJA1toVnP[...]
`
)

func remoteListToYml(remotes []client.Remote) ([]byte, error) {
	if len(remotes) == 0 {
		// Provide a helpful description, instead of an empty list.
		return []byte(RemoteHelpText), nil
	}

	return yml.Marshal(remotes)
}

func ymlToRemoteList(data []byte) ([]client.Remote, error) {
	remotes := []client.Remote{}

	if err := yml.Unmarshal(data, &remotes); err != nil {
		return nil, err
	}

	return remotes, nil
}

func handleRemoteAdd(ctx *cli.Context, ctl *client.Client) error {
	remote := client.Remote{
		Name:        ctx.Args().Get(0),
		Fingerprint: ctx.Args().Get(1),
		Folders:     nil,
	}

	if err := ctl.RemoteAdd(remote); err != nil {
		return fmt.Errorf("remote add: %v", err)
	}

	return nil
}

func handleRemoteRemove(ctx *cli.Context, ctl *client.Client) error {
	name := ctx.Args().First()
	if err := ctl.RemoteRm(name); err != nil {
		return fmt.Errorf("remote rm: %v", err)
	}

	return nil
}

func handleRemoteList(ctx *cli.Context, ctl *client.Client) error {
	remotes, err := ctl.RemoteLs()
	if err != nil {
		return fmt.Errorf("remote ls: %v", err)
	}

	if len(remotes) == 0 {
		fmt.Println("None yet. Use `brig remote add <user> <id>` to add some.")
		return nil
	}

	data, err := remoteListToYml(remotes)
	if err != nil {
		return fmt.Errorf("Failed to convert to yml: %v", err)
	}

	// Highlight the yml output (That's more of a joke currently):
	highlighted := pygments.Highlight(string(data), "YAML", "terminal256", "utf-8")
	highlighted = strings.TrimSpace(highlighted)
	fmt.Println(highlighted)
	return nil
}

func handleRemoteEdit(ctx *cli.Context, ctl *client.Client) error {
	remotes, err := ctl.RemoteLs()
	if err != nil {
		return fmt.Errorf("remote ls: %v", err)
	}

	data, err := remoteListToYml(remotes)
	if err != nil {
		return fmt.Errorf("Failed to convert to yml: %v", err)
	}

	// Launch an editor on the received data:
	newData, err := edit(data, "yml")
	if err != nil {
		return fmt.Errorf("Failed to launch editor: %v", err)
	}

	// Save a few network roundtrips if nothing was changed:
	if bytes.Equal(data, newData) {
		fmt.Println("Nothing changed.")
		return nil
	}

	newRemotes, err := ymlToRemoteList(newData)
	if err != nil {
		return err
	}

	if err := ctl.RemoteSave(newRemotes); err != nil {
		return fmt.Errorf("Saving back remotes failed: %v", err)
	}

	return nil
}

func handleRemoteLocate(ctx *cli.Context, ctl *client.Client) error {
	who := ctx.Args().First()
	candidates, err := ctl.RemoteLocate(who)
	if err != nil {
		return fmt.Errorf("Failed to locate peers: %v", err)
	}

	for _, candidate := range candidates {
		fmt.Println(candidate.Name, candidate.Fingerprint)
	}

	return nil
}

func handleRemotePing(ctx *cli.Context, ctl *client.Client) error {
	who := ctx.Args().First()

	msg := fmt.Sprintf("ping to %s: ", colors.Colorize(who, colors.Magenta))
	roundtrip, err := ctl.RemotePing(who)
	if err != nil {
		msg += colors.Colorize("✘", colors.Red)
		msg += fmt.Sprintf(" (%v)", err)
	} else {
		msg += colors.Colorize("✔", colors.Green)
		msg += fmt.Sprintf(" (%3.5fms)", roundtrip)
	}

	fmt.Println(msg)
	return nil
}

func handlePin(ctx *cli.Context, ctl *client.Client) error {
	path := ctx.Args().First()
	return ctl.Pin(path)
}

func handleUnpin(ctx *cli.Context, ctl *client.Client) error {
	path := ctx.Args().First()
	return ctl.Unpin(path)
}

func handleWhoami(ctx *cli.Context, ctl *client.Client) error {
	self, err := ctl.Whoami()
	if err != nil {
		return err
	}

	if !ctx.Bool("fingerprint") {
		userName := colors.Colorize(self.CurrentUser, colors.Yellow)
		ownerName := colors.Colorize(self.Owner, colors.Green)
		fmt.Printf("%s", ownerName)
		if self.CurrentUser != self.Owner {
			fmt.Printf(" (viewing %s's data)", userName)
		}

		fmt.Printf(" ")
	}

	fmt.Printf("%s\n", self.Fingerprint)
	return nil
}