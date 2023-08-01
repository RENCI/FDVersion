package tests

import (
	"strings"
	"testing"

	"FDVersion/pkg/FileSystemWalker"
	"github.com/stretchr/testify/assert"
)

func Test_EmptyDir(t *testing.T) {
	fsw := FileSystemWalker.New()

	//check that it works with an empty dir
	res := fsw.Visit("./data/case01_emptydir")

	assert.Equal(t, 0, res.Nodes.Size())
}

func Test_3files(t *testing.T) {
	fsw := FileSystemWalker.New()

	//check that it works with an empty dir
	res := fsw.Visit("./data/case02_3files")

	dirNode := res
	assert.Equal(t, true, dirNode.IsDir)
	assert.Equal(t, false, dirNode.IsFile)
	assert.Equal(t, int64(0), dirNode.FileSize)
	assert.Equal(t, 3, dirNode.Nodes.Size())

	file1Node := dirNode.Nodes.Get(0)
	assert.Equal(t, "file1.txt", file1Node.Name)
	assert.True(t, strings.HasSuffix(file1Node.Path, file1Node.Name))
	assert.Equal(t, false, file1Node.IsDir)
	assert.Equal(t, true, file1Node.IsFile)
	assert.Equal(t, int64(5), file1Node.FileSize)
	assert.Equal(t, 0, file1Node.Nodes.Size())
}

func Test_case03_filesanddirs(t *testing.T) {
	fsw := FileSystemWalker.New()

	//check that it works with an empty dir
	res := fsw.Visit("./data/case03_filesanddirs")

	assert.Equal(t, 2, res.Nodes.Size())

	dirNode := res
	assert.Equal(t, true, dirNode.IsDir)
	assert.Equal(t, false, dirNode.IsFile)
	assert.Equal(t, int64(0), dirNode.FileSize)
	assert.Equal(t, 2, dirNode.Nodes.Size())

	file1Node := dirNode.Nodes.Get(0)
	assert.Equal(t, "file1.txt", file1Node.Name)
	assert.True(t, strings.HasSuffix(file1Node.Path, file1Node.Name))
	assert.Equal(t, false, file1Node.IsDir)
	assert.Equal(t, true, file1Node.IsFile)
	assert.Equal(t, int64(28), file1Node.FileSize)
	assert.Equal(t, 0, file1Node.Nodes.Size())
}

func Test_case04_ordering(t *testing.T) {
	fsw := FileSystemWalker.New()

	//check that it works with an empty dir
	res := fsw.Visit("./data/case04_ordering")

	assert.Equal(t, 6, res.Nodes.Size())

	assert.Equal(t, "file1.txt", res.Nodes.Get(0).Name)
	assert.Equal(t, "file10.txt", res.Nodes.Get(1).Name)
	assert.Equal(t, "file2.txt", res.Nodes.Get(2).Name)
	assert.Equal(t, "dir1", res.Nodes.Get(3).Name)
	assert.Equal(t, "dir10", res.Nodes.Get(4).Name)
	assert.Equal(t, "dir2", res.Nodes.Get(5).Name)
}

func Test_case05_hash(t *testing.T) {
	res := FileSystemWalker.CalculateHash("./data/case05_hash/file1.txt")
	assert.Equal(t, "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855", res)

}

func Test_case06_hashes(t *testing.T) {
	fsw := FileSystemWalker.New()

	//check that it works with an empty dir
	res := fsw.Visit("./data/case06_hashes")
	fsw.CalculateHashes(res)

	dirNode := res

	file1Node := dirNode.Nodes.Get(0)
	assert.Equal(t, "47f8cf7a375b87e5b1ad96a2b82b594b064b97900196820f65741cd2ebe9caae", file1Node.Hash)
	assert.Equal(t, "c3ffa5a6977564bda0e8264b9f63e2366b495396d051aa1e10d7b107adc809ce", dirNode.Hash)

}

func Test_case07_hashes(t *testing.T) {
	fsw := FileSystemWalker.New()

	//check that it works with an empty dir
	res := fsw.Visit("./data/case07_hashes")
	fsw.CalculateHashes(res)

	dirNode := res

	file1Node := dirNode.Nodes.Get(0)
	file2Node := dirNode.Nodes.Get(1)
	assert.Equal(t, "47f8cf7a375b87e5b1ad96a2b82b594b064b97900196820f65741cd2ebe9caae", file1Node.Hash)
	assert.Equal(t, "85f3e2903210f99e678a3c2cf48be0be2b9a4c0c08403f00535f21c1301743c8", file2Node.Hash)
	assert.Equal(t, "63171381bbf6710fbde74ba3d53ff6df03b2bb853d565d0096960c5fb31fb493", dirNode.Hash)

}
