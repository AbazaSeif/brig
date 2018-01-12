// Code generated by capnpc-go. DO NOT EDIT.

package capnp

import (
	strconv "strconv"
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

// Commit is a set of changes to nodes
type Commit struct{ capnp.Struct }
type Commit_merge Commit

// Commit_TypeID is the unique identifier for the type Commit.
const Commit_TypeID = 0x8da013c66e545daf

func NewCommit(s *capnp.Segment) (Commit, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 6})
	return Commit{st}, err
}

func NewRootCommit(s *capnp.Segment) (Commit, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 6})
	return Commit{st}, err
}

func ReadRootCommit(msg *capnp.Message) (Commit, error) {
	root, err := msg.RootPtr()
	return Commit{root.Struct()}, err
}

func (s Commit) String() string {
	str, _ := text.Marshal(0x8da013c66e545daf, s.Struct)
	return str
}

func (s Commit) Message() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Commit) HasMessage() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Commit) MessageBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Commit) SetMessage(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Commit) Author() (string, error) {
	p, err := s.Struct.Ptr(1)
	return p.Text(), err
}

func (s Commit) HasAuthor() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Commit) AuthorBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return p.TextBytes(), err
}

func (s Commit) SetAuthor(v string) error {
	return s.Struct.SetText(1, v)
}

func (s Commit) Parent() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return []byte(p.Data()), err
}

func (s Commit) HasParent() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Commit) SetParent(v []byte) error {
	return s.Struct.SetData(2, v)
}

func (s Commit) Root() ([]byte, error) {
	p, err := s.Struct.Ptr(3)
	return []byte(p.Data()), err
}

func (s Commit) HasRoot() bool {
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Commit) SetRoot(v []byte) error {
	return s.Struct.SetData(3, v)
}

func (s Commit) Merge() Commit_merge { return Commit_merge(s) }

func (s Commit_merge) With() (string, error) {
	p, err := s.Struct.Ptr(4)
	return p.Text(), err
}

func (s Commit_merge) HasWith() bool {
	p, err := s.Struct.Ptr(4)
	return p.IsValid() || err != nil
}

func (s Commit_merge) WithBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(4)
	return p.TextBytes(), err
}

func (s Commit_merge) SetWith(v string) error {
	return s.Struct.SetText(4, v)
}

func (s Commit_merge) Head() ([]byte, error) {
	p, err := s.Struct.Ptr(5)
	return []byte(p.Data()), err
}

func (s Commit_merge) HasHead() bool {
	p, err := s.Struct.Ptr(5)
	return p.IsValid() || err != nil
}

func (s Commit_merge) SetHead(v []byte) error {
	return s.Struct.SetData(5, v)
}

// Commit_List is a list of Commit.
type Commit_List struct{ capnp.List }

// NewCommit creates a new list of Commit.
func NewCommit_List(s *capnp.Segment, sz int32) (Commit_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 6}, sz)
	return Commit_List{l}, err
}

func (s Commit_List) At(i int) Commit { return Commit{s.List.Struct(i)} }

func (s Commit_List) Set(i int, v Commit) error { return s.List.SetStruct(i, v.Struct) }

func (s Commit_List) String() string {
	str, _ := text.MarshalList(0x8da013c66e545daf, s.List)
	return str
}

// Commit_Promise is a wrapper for a Commit promised by a client call.
type Commit_Promise struct{ *capnp.Pipeline }

func (p Commit_Promise) Struct() (Commit, error) {
	s, err := p.Pipeline.Struct()
	return Commit{s}, err
}

func (p Commit_Promise) Merge() Commit_merge_Promise { return Commit_merge_Promise{p.Pipeline} }

// Commit_merge_Promise is a wrapper for a Commit_merge promised by a client call.
type Commit_merge_Promise struct{ *capnp.Pipeline }

func (p Commit_merge_Promise) Struct() (Commit_merge, error) {
	s, err := p.Pipeline.Struct()
	return Commit_merge{s}, err
}

// A single directory entry
type DirEntry struct{ capnp.Struct }

// DirEntry_TypeID is the unique identifier for the type DirEntry.
const DirEntry_TypeID = 0x8b15ee76774b1f9d

func NewDirEntry(s *capnp.Segment) (DirEntry, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return DirEntry{st}, err
}

func NewRootDirEntry(s *capnp.Segment) (DirEntry, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2})
	return DirEntry{st}, err
}

func ReadRootDirEntry(msg *capnp.Message) (DirEntry, error) {
	root, err := msg.RootPtr()
	return DirEntry{root.Struct()}, err
}

func (s DirEntry) String() string {
	str, _ := text.Marshal(0x8b15ee76774b1f9d, s.Struct)
	return str
}

func (s DirEntry) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s DirEntry) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s DirEntry) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s DirEntry) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

func (s DirEntry) Hash() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s DirEntry) HasHash() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s DirEntry) SetHash(v []byte) error {
	return s.Struct.SetData(1, v)
}

// DirEntry_List is a list of DirEntry.
type DirEntry_List struct{ capnp.List }

// NewDirEntry creates a new list of DirEntry.
func NewDirEntry_List(s *capnp.Segment, sz int32) (DirEntry_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 0, PointerCount: 2}, sz)
	return DirEntry_List{l}, err
}

func (s DirEntry_List) At(i int) DirEntry { return DirEntry{s.List.Struct(i)} }

func (s DirEntry_List) Set(i int, v DirEntry) error { return s.List.SetStruct(i, v.Struct) }

func (s DirEntry_List) String() string {
	str, _ := text.MarshalList(0x8b15ee76774b1f9d, s.List)
	return str
}

// DirEntry_Promise is a wrapper for a DirEntry promised by a client call.
type DirEntry_Promise struct{ *capnp.Pipeline }

func (p DirEntry_Promise) Struct() (DirEntry, error) {
	s, err := p.Pipeline.Struct()
	return DirEntry{s}, err
}

// Directory contains one or more directories or files
type Directory struct{ capnp.Struct }

// Directory_TypeID is the unique identifier for the type Directory.
const Directory_TypeID = 0xe24c59306c829c01

func NewDirectory(s *capnp.Segment) (Directory, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Directory{st}, err
}

func NewRootDirectory(s *capnp.Segment) (Directory, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2})
	return Directory{st}, err
}

func ReadRootDirectory(msg *capnp.Message) (Directory, error) {
	root, err := msg.RootPtr()
	return Directory{root.Struct()}, err
}

func (s Directory) String() string {
	str, _ := text.Marshal(0xe24c59306c829c01, s.Struct)
	return str
}

func (s Directory) Size() uint64 {
	return s.Struct.Uint64(0)
}

func (s Directory) SetSize(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Directory) Parent() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Directory) HasParent() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Directory) ParentBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Directory) SetParent(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Directory) Children() (DirEntry_List, error) {
	p, err := s.Struct.Ptr(1)
	return DirEntry_List{List: p.List()}, err
}

func (s Directory) HasChildren() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Directory) SetChildren(v DirEntry_List) error {
	return s.Struct.SetPtr(1, v.List.ToPtr())
}

// NewChildren sets the children field to a newly
// allocated DirEntry_List, preferring placement in s's segment.
func (s Directory) NewChildren(n int32) (DirEntry_List, error) {
	l, err := NewDirEntry_List(s.Struct.Segment(), n)
	if err != nil {
		return DirEntry_List{}, err
	}
	err = s.Struct.SetPtr(1, l.List.ToPtr())
	return l, err
}

// Directory_List is a list of Directory.
type Directory_List struct{ capnp.List }

// NewDirectory creates a new list of Directory.
func NewDirectory_List(s *capnp.Segment, sz int32) (Directory_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 2}, sz)
	return Directory_List{l}, err
}

func (s Directory_List) At(i int) Directory { return Directory{s.List.Struct(i)} }

func (s Directory_List) Set(i int, v Directory) error { return s.List.SetStruct(i, v.Struct) }

func (s Directory_List) String() string {
	str, _ := text.MarshalList(0xe24c59306c829c01, s.List)
	return str
}

// Directory_Promise is a wrapper for a Directory promised by a client call.
type Directory_Promise struct{ *capnp.Pipeline }

func (p Directory_Promise) Struct() (Directory, error) {
	s, err := p.Pipeline.Struct()
	return Directory{s}, err
}

// A leaf node in the MDAG
type File struct{ capnp.Struct }

// File_TypeID is the unique identifier for the type File.
const File_TypeID = 0x8ea7393d37893155

func NewFile(s *capnp.Segment) (File, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return File{st}, err
}

func NewRootFile(s *capnp.Segment) (File, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3})
	return File{st}, err
}

func ReadRootFile(msg *capnp.Message) (File, error) {
	root, err := msg.RootPtr()
	return File{root.Struct()}, err
}

func (s File) String() string {
	str, _ := text.Marshal(0x8ea7393d37893155, s.Struct)
	return str
}

func (s File) Size() uint64 {
	return s.Struct.Uint64(0)
}

func (s File) SetSize(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s File) Parent() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s File) HasParent() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s File) ParentBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s File) SetParent(v string) error {
	return s.Struct.SetText(0, v)
}

func (s File) Key() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s File) HasKey() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s File) SetKey(v []byte) error {
	return s.Struct.SetData(1, v)
}

func (s File) Content() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return []byte(p.Data()), err
}

func (s File) HasContent() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s File) SetContent(v []byte) error {
	return s.Struct.SetData(2, v)
}

// File_List is a list of File.
type File_List struct{ capnp.List }

// NewFile creates a new list of File.
func NewFile_List(s *capnp.Segment, sz int32) (File_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 3}, sz)
	return File_List{l}, err
}

func (s File_List) At(i int) File { return File{s.List.Struct(i)} }

func (s File_List) Set(i int, v File) error { return s.List.SetStruct(i, v.Struct) }

func (s File_List) String() string {
	str, _ := text.MarshalList(0x8ea7393d37893155, s.List)
	return str
}

// File_Promise is a wrapper for a File promised by a client call.
type File_Promise struct{ *capnp.Pipeline }

func (p File_Promise) Struct() (File, error) {
	s, err := p.Pipeline.Struct()
	return File{s}, err
}

// Ghost indicates that a certain node was at this path once
type Ghost struct{ capnp.Struct }
type Ghost_Which uint16

const (
	Ghost_Which_commit    Ghost_Which = 0
	Ghost_Which_directory Ghost_Which = 1
	Ghost_Which_file      Ghost_Which = 2
)

func (w Ghost_Which) String() string {
	const s = "commitdirectoryfile"
	switch w {
	case Ghost_Which_commit:
		return s[0:6]
	case Ghost_Which_directory:
		return s[6:15]
	case Ghost_Which_file:
		return s[15:19]

	}
	return "Ghost_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Ghost_TypeID is the unique identifier for the type Ghost.
const Ghost_TypeID = 0x80c828d7e89c12ea

func NewGhost(s *capnp.Segment) (Ghost, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2})
	return Ghost{st}, err
}

func NewRootGhost(s *capnp.Segment) (Ghost, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2})
	return Ghost{st}, err
}

func ReadRootGhost(msg *capnp.Message) (Ghost, error) {
	root, err := msg.RootPtr()
	return Ghost{root.Struct()}, err
}

func (s Ghost) String() string {
	str, _ := text.Marshal(0x80c828d7e89c12ea, s.Struct)
	return str
}

func (s Ghost) Which() Ghost_Which {
	return Ghost_Which(s.Struct.Uint16(8))
}
func (s Ghost) GhostInode() uint64 {
	return s.Struct.Uint64(0)
}

func (s Ghost) SetGhostInode(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Ghost) GhostPath() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Ghost) HasGhostPath() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Ghost) GhostPathBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Ghost) SetGhostPath(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Ghost) Commit() (Commit, error) {
	if s.Struct.Uint16(8) != 0 {
		panic("Which() != commit")
	}
	p, err := s.Struct.Ptr(1)
	return Commit{Struct: p.Struct()}, err
}

func (s Ghost) HasCommit() bool {
	if s.Struct.Uint16(8) != 0 {
		return false
	}
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Ghost) SetCommit(v Commit) error {
	s.Struct.SetUint16(8, 0)
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewCommit sets the commit field to a newly
// allocated Commit struct, preferring placement in s's segment.
func (s Ghost) NewCommit() (Commit, error) {
	s.Struct.SetUint16(8, 0)
	ss, err := NewCommit(s.Struct.Segment())
	if err != nil {
		return Commit{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Ghost) Directory() (Directory, error) {
	if s.Struct.Uint16(8) != 1 {
		panic("Which() != directory")
	}
	p, err := s.Struct.Ptr(1)
	return Directory{Struct: p.Struct()}, err
}

func (s Ghost) HasDirectory() bool {
	if s.Struct.Uint16(8) != 1 {
		return false
	}
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Ghost) SetDirectory(v Directory) error {
	s.Struct.SetUint16(8, 1)
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewDirectory sets the directory field to a newly
// allocated Directory struct, preferring placement in s's segment.
func (s Ghost) NewDirectory() (Directory, error) {
	s.Struct.SetUint16(8, 1)
	ss, err := NewDirectory(s.Struct.Segment())
	if err != nil {
		return Directory{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

func (s Ghost) File() (File, error) {
	if s.Struct.Uint16(8) != 2 {
		panic("Which() != file")
	}
	p, err := s.Struct.Ptr(1)
	return File{Struct: p.Struct()}, err
}

func (s Ghost) HasFile() bool {
	if s.Struct.Uint16(8) != 2 {
		return false
	}
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Ghost) SetFile(v File) error {
	s.Struct.SetUint16(8, 2)
	return s.Struct.SetPtr(1, v.Struct.ToPtr())
}

// NewFile sets the file field to a newly
// allocated File struct, preferring placement in s's segment.
func (s Ghost) NewFile() (File, error) {
	s.Struct.SetUint16(8, 2)
	ss, err := NewFile(s.Struct.Segment())
	if err != nil {
		return File{}, err
	}
	err = s.Struct.SetPtr(1, ss.Struct.ToPtr())
	return ss, err
}

// Ghost_List is a list of Ghost.
type Ghost_List struct{ capnp.List }

// NewGhost creates a new list of Ghost.
func NewGhost_List(s *capnp.Segment, sz int32) (Ghost_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 2}, sz)
	return Ghost_List{l}, err
}

func (s Ghost_List) At(i int) Ghost { return Ghost{s.List.Struct(i)} }

func (s Ghost_List) Set(i int, v Ghost) error { return s.List.SetStruct(i, v.Struct) }

func (s Ghost_List) String() string {
	str, _ := text.MarshalList(0x80c828d7e89c12ea, s.List)
	return str
}

// Ghost_Promise is a wrapper for a Ghost promised by a client call.
type Ghost_Promise struct{ *capnp.Pipeline }

func (p Ghost_Promise) Struct() (Ghost, error) {
	s, err := p.Pipeline.Struct()
	return Ghost{s}, err
}

func (p Ghost_Promise) Commit() Commit_Promise {
	return Commit_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p Ghost_Promise) Directory() Directory_Promise {
	return Directory_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

func (p Ghost_Promise) File() File_Promise {
	return File_Promise{Pipeline: p.Pipeline.GetPipeline(1)}
}

// Node is a node in the merkle dag of brig
type Node struct{ capnp.Struct }
type Node_Which uint16

const (
	Node_Which_commit    Node_Which = 0
	Node_Which_directory Node_Which = 1
	Node_Which_file      Node_Which = 2
	Node_Which_ghost     Node_Which = 3
)

func (w Node_Which) String() string {
	const s = "commitdirectoryfileghost"
	switch w {
	case Node_Which_commit:
		return s[0:6]
	case Node_Which_directory:
		return s[6:15]
	case Node_Which_file:
		return s[15:19]
	case Node_Which_ghost:
		return s[19:24]

	}
	return "Node_Which(" + strconv.FormatUint(uint64(w), 10) + ")"
}

// Node_TypeID is the unique identifier for the type Node.
const Node_TypeID = 0xa629eb7f7066fae3

func NewNode(s *capnp.Segment) (Node, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 4})
	return Node{st}, err
}

func NewRootNode(s *capnp.Segment) (Node, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 16, PointerCount: 4})
	return Node{st}, err
}

func ReadRootNode(msg *capnp.Message) (Node, error) {
	root, err := msg.RootPtr()
	return Node{root.Struct()}, err
}

func (s Node) String() string {
	str, _ := text.Marshal(0xa629eb7f7066fae3, s.Struct)
	return str
}

func (s Node) Which() Node_Which {
	return Node_Which(s.Struct.Uint16(8))
}
func (s Node) Name() (string, error) {
	p, err := s.Struct.Ptr(0)
	return p.Text(), err
}

func (s Node) HasName() bool {
	p, err := s.Struct.Ptr(0)
	return p.IsValid() || err != nil
}

func (s Node) NameBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(0)
	return p.TextBytes(), err
}

func (s Node) SetName(v string) error {
	return s.Struct.SetText(0, v)
}

func (s Node) Hash() ([]byte, error) {
	p, err := s.Struct.Ptr(1)
	return []byte(p.Data()), err
}

func (s Node) HasHash() bool {
	p, err := s.Struct.Ptr(1)
	return p.IsValid() || err != nil
}

func (s Node) SetHash(v []byte) error {
	return s.Struct.SetData(1, v)
}

func (s Node) ModTime() (string, error) {
	p, err := s.Struct.Ptr(2)
	return p.Text(), err
}

func (s Node) HasModTime() bool {
	p, err := s.Struct.Ptr(2)
	return p.IsValid() || err != nil
}

func (s Node) ModTimeBytes() ([]byte, error) {
	p, err := s.Struct.Ptr(2)
	return p.TextBytes(), err
}

func (s Node) SetModTime(v string) error {
	return s.Struct.SetText(2, v)
}

func (s Node) Inode() uint64 {
	return s.Struct.Uint64(0)
}

func (s Node) SetInode(v uint64) {
	s.Struct.SetUint64(0, v)
}

func (s Node) Commit() (Commit, error) {
	if s.Struct.Uint16(8) != 0 {
		panic("Which() != commit")
	}
	p, err := s.Struct.Ptr(3)
	return Commit{Struct: p.Struct()}, err
}

func (s Node) HasCommit() bool {
	if s.Struct.Uint16(8) != 0 {
		return false
	}
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Node) SetCommit(v Commit) error {
	s.Struct.SetUint16(8, 0)
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewCommit sets the commit field to a newly
// allocated Commit struct, preferring placement in s's segment.
func (s Node) NewCommit() (Commit, error) {
	s.Struct.SetUint16(8, 0)
	ss, err := NewCommit(s.Struct.Segment())
	if err != nil {
		return Commit{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

func (s Node) Directory() (Directory, error) {
	if s.Struct.Uint16(8) != 1 {
		panic("Which() != directory")
	}
	p, err := s.Struct.Ptr(3)
	return Directory{Struct: p.Struct()}, err
}

func (s Node) HasDirectory() bool {
	if s.Struct.Uint16(8) != 1 {
		return false
	}
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Node) SetDirectory(v Directory) error {
	s.Struct.SetUint16(8, 1)
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewDirectory sets the directory field to a newly
// allocated Directory struct, preferring placement in s's segment.
func (s Node) NewDirectory() (Directory, error) {
	s.Struct.SetUint16(8, 1)
	ss, err := NewDirectory(s.Struct.Segment())
	if err != nil {
		return Directory{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

func (s Node) File() (File, error) {
	if s.Struct.Uint16(8) != 2 {
		panic("Which() != file")
	}
	p, err := s.Struct.Ptr(3)
	return File{Struct: p.Struct()}, err
}

func (s Node) HasFile() bool {
	if s.Struct.Uint16(8) != 2 {
		return false
	}
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Node) SetFile(v File) error {
	s.Struct.SetUint16(8, 2)
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewFile sets the file field to a newly
// allocated File struct, preferring placement in s's segment.
func (s Node) NewFile() (File, error) {
	s.Struct.SetUint16(8, 2)
	ss, err := NewFile(s.Struct.Segment())
	if err != nil {
		return File{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

func (s Node) Ghost() (Ghost, error) {
	if s.Struct.Uint16(8) != 3 {
		panic("Which() != ghost")
	}
	p, err := s.Struct.Ptr(3)
	return Ghost{Struct: p.Struct()}, err
}

func (s Node) HasGhost() bool {
	if s.Struct.Uint16(8) != 3 {
		return false
	}
	p, err := s.Struct.Ptr(3)
	return p.IsValid() || err != nil
}

func (s Node) SetGhost(v Ghost) error {
	s.Struct.SetUint16(8, 3)
	return s.Struct.SetPtr(3, v.Struct.ToPtr())
}

// NewGhost sets the ghost field to a newly
// allocated Ghost struct, preferring placement in s's segment.
func (s Node) NewGhost() (Ghost, error) {
	s.Struct.SetUint16(8, 3)
	ss, err := NewGhost(s.Struct.Segment())
	if err != nil {
		return Ghost{}, err
	}
	err = s.Struct.SetPtr(3, ss.Struct.ToPtr())
	return ss, err
}

// Node_List is a list of Node.
type Node_List struct{ capnp.List }

// NewNode creates a new list of Node.
func NewNode_List(s *capnp.Segment, sz int32) (Node_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 16, PointerCount: 4}, sz)
	return Node_List{l}, err
}

func (s Node_List) At(i int) Node { return Node{s.List.Struct(i)} }

func (s Node_List) Set(i int, v Node) error { return s.List.SetStruct(i, v.Struct) }

func (s Node_List) String() string {
	str, _ := text.MarshalList(0xa629eb7f7066fae3, s.List)
	return str
}

// Node_Promise is a wrapper for a Node promised by a client call.
type Node_Promise struct{ *capnp.Pipeline }

func (p Node_Promise) Struct() (Node, error) {
	s, err := p.Pipeline.Struct()
	return Node{s}, err
}

func (p Node_Promise) Commit() Commit_Promise {
	return Commit_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

func (p Node_Promise) Directory() Directory_Promise {
	return Directory_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

func (p Node_Promise) File() File_Promise {
	return File_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

func (p Node_Promise) Ghost() Ghost_Promise {
	return Ghost_Promise{Pipeline: p.Pipeline.GetPipeline(3)}
}

const schema_9195d073cb5c5953 = "x\xda\xacU\xed\x8b\x15\xd5\x1f\xff~\xce\x99{g/" +
	"\xac\xbf\xbd\xf3\x1b\x85\x90\xd6{\x10#\x95\xd4\xd5\x1b\xd4" +
	".,\xba\xea&k*{J_(\x16\x9d\xee={" +
	"g\xf0\xde\x99ufj\xdbH\xd6BA\xab\x0d%{" +
	"\x11(I/\x0cz\xe3_\x10A\xd0\x03\xd5\x0b\x8b\x1e" +
	"\x0c\x0cz0\x92\x9e(\xe8EJ:q\xee\xf3.\x9b" +
	"\xfb\xa6w\xf7~\xe6\x9c9\xe7\xf3\xf0\xfd\xcc\xc0\xdbl" +
	"\x0b\xdb\x98\xb9\xdb\"\x92\x03\x99l\xfa\xd3\xff\xcf^\xfb" +
	"r\xf5\x07GI\xae\x00K\x1f\xde\x7f\xf0\xa3\xf8\xd2+" +
	"\xa7i\x94\xd9\x1cV\xf1:V\xc2\xcd1\xdb\xcd\xb1B" +
	"q\x94\x15@H\xcf\x15\x1e\x9cz\xf2\xb7e/\x90\xb3" +
	"\x02\x9d\x0d\x19f\x13\x15\x15\x1f\x82{\x98\xdb\xeea^" +
	"p\xcf\xf1)Bz\xf1\x91\xbd\xc1{\xee\xf9\xd9y\xcb" +
	"\xb3f\xf9u\xbe\x16n\xce\xb2\xdd\x9cU(\x0eZ\xf5" +
	"\xd7\xef\xdbx\xf2\xbe\xe1\xc17^2\x17\xea\xde\xc0\xcd" +
	"\x86u\x99\xe5p\x873\xb6;\x9c)\xb8:\xf3#!" +
	"\xfd\xfe\xc6\xc4\xe4\xcc\xcfk.\xcc'`\xd9\x16\xac\xe2" +
	"`v9\xdc\xb1\xac\xed\x8ee\x0b\xc5#\xd9\xf7\xcd\x09" +
	"8\xfb\\u`\xff\xae\xef\xe6\x9fPg0\xda\xb3\x15" +
	"\xee\xbe\x1e\xdb\xdd\xd7Spg{.\x12\xd2\xf3\xb97" +
	"/\\\x1d\x8en\x90sW\x17\x9deY\x1bD\xc5\xc1" +
	"\xdc\x01\x10\xdc\xd1\xdc\x14\xdd\x9f\x96T2\x11o\x08B" +
	"^\xd6\xf1\x86\x92\x9a\x0c&7\xd4\xc2\xb2\xae\xae\xaf\xff" +
	"\x1e\xda\xe1\xd9a\x9c\x8c\x03\xd2\x02K\x1f}\xf95\xf9" +
	"\xd6\x17\xcf\xbfK\xd2b\x18\xb9\x07\xe8%\xda\x88O\x91" +
	"\xee\xf0\xc28\x11~\x90-\xfb%\x95\xe8X$\x9eJ" +
	"\x84\x12%\x1d%\xca\x0fD\x10\x96\xb5\x98R\xb1P\x89" +
	"H<?\x16\x93*\xf1D\x18\x94\xa0\x89\xe4Rn\x11" +
	"Y r\x8e\x1c \x92\xcfp\xc8\x13\x0c\xc0R\x18\xec" +
	"\xf8CD\xf2\x18\x87<\xc5\xd0\xcf\xd2\x14K\xc1\x88\x9c" +
	"\xd9!\"y\x82C\x9ea\xe8\xe7\xb7\x0c\xcc\x89\x9c\xd3" +
	"f\xf5)\x0ey\x96\xa1\xdf\xbai`\x8b\xc8yu-" +
	"\x91<\xc3!\xcf3\xa4\x15s\xdb\xb1 $^\xd6\xc8" +
	"\x11C\x8e\x9a\xe0\xb8J\x08\x1ez\x89\xa1\x97\xb0\xb9\x14" +
	"\xd6j~\x82|GB\x02\xf2\x84\xb4\xecG\xba\x94\x84" +
	"\x11a\x1a\xf9\x8e;\x8d\xa7}\x13~U#\xdf\x89E" +
	"s\xd3\"Ro\xf77G\xa3A\x12M/\xac\xf6\x9d" +
	"u\xb5\x1d|\x98\x8e\x88\xd8\x0f*U\xcdD\xeb\x1a\xd3" +
	"B\x9b\x8d\x04\xd9\xd3\x96r\x8da\xbc\x8aC\x0e08" +
	"--\xd7\x19p5\x87\xbc\x97\xa1/P5\xdd\xa2\xda" +
	"\xe7\xa9\xd8\xc3\x12bX\xb2\xf8M\xb7\x85}F\x97\x85" +
	"\xef)\x9a\xa9X\x89t[]>\xe1\xf3X(\x11\xeb" +
	"D\x84\x13\xa2\xe4\xa9\xa0b\x02\x12\x8a \xb4\xcb:\x9e" +
	"\xeb\xffV\"\xf9\x14\x87<\xd6u\xe9g\x87:\xa1p" +
	"\x18k\xd8\x7f\xdc\x80G9\xe4\x8b\x0c\x0e\xe7\x0d\xf3O" +
	"\xae\xedD\x05\x16\xba\xe7`v\x13\xb1\x99\x9a\x8ecU" +
	"i\x93\xde\xac\x9eH\xbc0j\xff\x9dT\x91\x0e\x92\x96" +
	"\x0a}Q\x18\xb6\xff\x14j:\xaa\xe8\xc5\x84y\xc0\xe7" +
	"U\xbd\xb0,w4\xed{'\x1d\x11U\xad&D\xc0" +
	"\xccL\xf8\x81H<-vo\x1f\xd9AD2\xdfV" +
	"B\x19*\x079\xa4\xd7\x99\x04m8?\xc6!\xabF" +
	"\x88\xe6\x1c\xf8+\x89d\x99CN\x1a!XC\x88\x9a" +
	"\xd1\xd1\xe3\x90\x09C_\xec?\xdd\x8ey\x8bc\x93\xb2" +
	"}HO\xb7(\xce\x94\xc2 \xe9\xe2\xbf\x18\xd9=\xe6" +
	"\xc1\xc2dW53\xb0\x13\xe9\x9e:\xcbXX\xaaQ" +
	"\x02M\xc25\x1d\x1d\xaajQV\x15\x13\x8a\xc7#\xbf" +
	"B\x90\xa2\xcd\xfe\x13\xc3\xfec\x0ey\xb9+\x07\x9f\x1b" +
	"\xf0\x12\x87\xbc\xd2\x95\x83\xaf\x0c\xd3\xcf8\xe47\x0ch" +
	"\xc6\xe0\xebMD\xf22\x87\xbcj: My\xa3\x03" +
	"\xbe5\xf2]\xe1\x90\xd7\x18\xfa3\xb7\x0c\x9c!r~" +
	"0\x8dq\x95C\xfe\xce\xd0\x9f\xbdi\xe0,\x91\xf3\xab" +
	"9\xed\x1a\x87\xfc\x93\xa1\xdf\xfe\xdb\xc06\x91\xf3\x87y" +
	"\xf7/\x1c\xf2\xaf\xdbM\xd0L-,\xef\xf5;\x0f\x0b" +
	"\xbe!\xdf6\xe1?\xed\x95B\xbd\xb9\x90\xef|\x17\xe7" +
	"\xf5\x8d\xf5o}\xd3,\x0fZ\xd8\xc4\xd5M\x13_G" +
	"\xdaZ\x9a\x99\x16&$\xca\x0fb\x11\x06Z\x84\x91\xa8" +
	"\x85\x91n\xf7\x90\xafc\x83M\xf8v\xb5>\xd8\xbdm" +
	"CG\x8d\x9a[8\xe4\xaeN\x9c\xc7\x8c\x1f\xdb9\xe4" +
	"xW\x9cw\xef$\x92\xbb\x1a\xb9\xbf]r\xd3\x92\xe7" +
	"W\xcb\x91\x0e\x88\x08\xff#\x8cs \xdf\xf9\xd0\x13\x0c" +
	"\xb8\x98\x04\x8d\x82Z_\xd3\x11\xafh\xd3\x9dy\xab\x11" +
	"\x95\xb9\xe5\x99i\x04eNyN\xf9\x89\xd7\xb1^\xab" +
	"r\xcb\xfa\x7f\x02\x00\x00\xff\xffu\xe9)\xbf"

func init() {
	schemas.Register(schema_9195d073cb5c5953,
		0x80c828d7e89c12ea,
		0x8b15ee76774b1f9d,
		0x8da013c66e545daf,
		0x8ea7393d37893155,
		0xa629eb7f7066fae3,
		0xe24c59306c829c01,
		0xfa723de4a6aa09a0)
}