package Vector

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/bradleyjkemp/memviz"
	"github.com/stretchr/testify/assert"
)

// 生成gv图
func TestVector(t *testing.T) {

	list := New()

	buf := &bytes.Buffer{}
	memviz.Map(buf, tree)

	ioutil.WriteFile("./Vector.gv", buf.Bytes(), os.ModePerm)
	exec.Command("dot", "./Vector.gv", "-Tpng", "./Vector.png").Run()
}

// 镜像翻转
func TestMirror(t *testing.T) {

	assert.Equal(t, Item("DBACFE"), result)
	assert.Equal(t, Item("DFEBCA"), resultMirror)
}

// 查找
func TestPreorderTraversal(t *testing.T) {

	assert.Equal(t, Item("DBACFE"), result)

}
