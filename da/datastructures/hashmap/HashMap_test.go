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
		hashmap.Set(st[i], st[i])
	}

	buf := &bytes.Buffer{}
	memviz.Map(buf, hashmap)

	ioutil.WriteFile("./HashMap.gv", buf.Bytes(), os.ModePerm)
	exec.Command("dot", "-Tpng", "./HashMap.gv", "-o", "./HashMap.png").Run()
}

// 搜索
func TestGet(t *testing.T) {
	hashmap := New(0)

	st := []string{"D", "B", "F", "A", "C", "E"}
	for i := 0; i < 6; i++ {
		hashmap.Set(st[i], st[i])
	}

	result, isFind := hashmap.Get("B")
	if !isFind || result == nil {
		t.Fatal("Get failed")
	}

	assert.Equal(t, "B", result.(string))
}

// 删除
func TestDel(t *testing.T) {
	hashmap := New(0)

	st := []string{"D", "B", "F", "A", "C", "E"}
	for i := 0; i < 6; i++ {
		hashmap.Set(st[i], st[i])
	}

	hashmap.Del("B")
	_, isFind := hashmap.Get("B")

	assert.Equal(t, false, isFind)

}
