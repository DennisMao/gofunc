package HashMap

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/bradleyjkemp/memviz"
	"github.com/stretchr/testify/assert"
)

// 生成gv图
func TestHashMap(t *testing.T) {

	hashmap := New(0)

	st := []string{"D", "B", "F", "A", "C", "E"}
	for i := 0; i < 6; i++ {
		hashmap.Insert(st[i], st[i])
	}

	buf := &bytes.Buffer{}
	memviz.Map(buf, hashmap)

	ioutil.WriteFile("./HashMap.gv", buf.Bytes(), os.ModePerm)
	exec.Command("dot", "-Tpng", "./HashMap.gv", "-o", "./HashMap.png").Run()
}

// 搜索
func TestSearch(t *testing.T) {
	hashmap := New(0)

	st := []string{"D", "B", "F", "A", "C", "E"}
	for i := 0; i < 6; i++ {
		hashmap.Insert(st[i], st[i])
	}

	result, isFind := hashmap.Search("B")
	if !isFind || result == nil {
		t.Fatal("Search failed")
	}

	assert.Equal(t, "B", result.(string))
}

// 删除
func TestDelete(t *testing.T) {
	hashmap := New(0)

	st := []string{"D", "B", "F", "A", "C", "E"}
	for i := 0; i < 6; i++ {
		hashmap.Insert(st[i], st[i])
	}

	hashmap.Delete("B")
	_, isFind := hashmap.Search("B")

	assert.Equal(t, false, isFind)

}
