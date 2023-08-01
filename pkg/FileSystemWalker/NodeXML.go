package FileSystemWalker

import (
	"encoding/xml"

	"github.com/RENCI/GoUtils/Collections/List"
)

type NodeXML struct {
	Id       string `xml:"id,attr"`
	Name     string
	Path     string
	IsDir    bool
	IsFile   bool
	FileSize int64
	Hash     string
	Nodes    []*NodeXML `xml:"Nodes>NodeXML"`
}

func (this *Node) XML() string {

	s, _ := xml.MarshalIndent(this.ConvertToNodeXML(), "", "    ")
	return string(s)
}

func (this *Node) ConvertToNodeXML() *NodeXML {
	res := NewNodeXML(this)

	return res
}

func NewNodeXML(this *Node) *NodeXML {
	res := &NodeXML{
		Id:       this.Id,
		Name:     this.Name,
		Path:     this.Path,
		IsDir:    this.IsDir,
		IsFile:   this.IsFile,
		FileSize: this.FileSize,
		Hash:     this.Hash,
	}

	nodes := List.New[*NodeXML]()

	this.Nodes.ForEach(func(item *Node) {
		nodes.Add(NewNodeXML(item))
	})

	res.Nodes = nodes.GetSlice()
	return res
}
