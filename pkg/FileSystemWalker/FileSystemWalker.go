package FileSystemWalker

import (
	"crypto/sha256"
	"encoding/hex"
	"io"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/RENCI/GoUtils/Collections/List"
)

type PathWalker interface {
	Visit(path string)
}

type FileSystemWalker struct {
	NamePrefixToIgnore List.List[string]
}

func New() *FileSystemWalker {
	res := &FileSystemWalker{
		NamePrefixToIgnore: List.New[string](),
	}

	return res
}

func (this *FileSystemWalker) Visit(path string) *Node {
	res := NewNode()
	absPath, err := filepath.Abs(path)

	if err != nil {
		panic(err)
	}

	this.visitFD(absPath, &res)
	return res.Nodes.First()
}

func (this *FileSystemWalker) visitFD(path string, parent *Node) {

	if HasPathToIgnore(path, this.NamePrefixToIgnore) {
		return
	}

	fd, err := os.Stat(path)
	if err != nil {
		log.Fatal(err)
	}

	node := CreateNode(path, fd)

	parent.Nodes.Add(&node)

	//println(node.JSON())

	if fd.IsDir() {
		direntry := ListDir(path)

		direntry.ForEach(func(item fs.DirEntry) {
			this.visitFD(filepath.Join(path, item.Name()), &node)
		})

	}

}

func (this *FileSystemWalker) CalculateHashes(root *Node) {
	CalculateHashForNodes(root)
}

func CalculateHashForNodes(node *Node) {
	if node.IsDir {

		node.Nodes.ForEach(func(n *Node) {
			CalculateHashForNodes(n)
		})

		var sb strings.Builder
		node.Nodes.ForEach(func(n *Node) {
			sb.WriteString(n.Hash)
		})

		node.Hash = CalculateHashForString(sb.String())

	} else {
		node.Hash = CalculateHash(node.Path)
	}

}

func CreateNode(path string, fd os.FileInfo) Node {
	node := NewNode()
	node.IsDir = fd.IsDir()
	node.IsFile = !fd.IsDir()
	node.FileSize = fd.Size()
	node.Name = fd.Name()
	node.Path = path
	node.Id = path
	node.LastChanged = fd.ModTime()
	return node
}

// ListDir reurns List of entries in the directory. It is sorted and grouped by directories and files.
func ListDir(path string) List.List[fs.DirEntry] {
	direntry, err := os.ReadDir(path)

	if err != nil {
		log.Fatal(err)
	}

	dirlist := List.New[fs.DirEntry]()
	dirlist.AddRange(direntry)
	dirlist.Sort(func(item1 fs.DirEntry, item2 fs.DirEntry) int {
		if item1.IsDir() == item2.IsDir() {
			return strings.Compare(item1.Name(), item2.Name())
		} else {
			if item1.IsDir() {
				return 1
			} else {
				return -1
			}
		}
	})
	return dirlist
}

// HasPathToIgnore checks if path should be ignored
func HasPathToIgnore(name string, pathToIgnore List.List[string]) bool {
	return pathToIgnore.Any(func(prefix string) bool {
		return strings.HasPrefix(name, prefix)
	})
}

// CalculateHash calculates SHA256 hash for a file
func CalculateHash(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	hasher := sha256.New()
	if _, err := io.Copy(hasher, f); err != nil {
		log.Fatal(err)
	}
	value := hex.EncodeToString(hasher.Sum(nil))
	return value
}

// CalculateHashForString calculates SHA256 hash for a string
func CalculateHashForString(content string) string {
	hasher := sha256.New()
	hasher.Write([]byte(content))
	value := hex.EncodeToString(hasher.Sum(nil))
	return value
}
