package FileSystemWalker

import (
	"encoding/json"
	"time"

	"github.com/RENCI/GoUtils/Collections/List"
)

type Node struct {
	Id          string
	Name        string
	Path        string
	IsDir       bool
	IsFile      bool
	LastChanged time.Time
	FileSize    int64
	Hash        string
	Nodes       List.List[*Node]
}

// NewNode creates an empty instance of Node
func NewNode() Node {
	res := Node{
		Nodes: List.New[*Node](),
	}

	return res
}

func (this *Node) JSON() string {
	s, _ := json.MarshalIndent(this, "", "    ")
	return string(s)
}
