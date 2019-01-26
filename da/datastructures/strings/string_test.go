package strings

import (
	"log"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var strTest *String

func init() {
	strTest = newString5M()
	log.Printf("测试字符串:%s\n", string(strTest.arr))
}

func newString5M() *String {
	random := rand.New(rand.NewSource(time.Now().Unix()))

	ret := ""
	for i := 0; i < 1024; i++ {
		ret += strconv.Itoa(random.Int())
	}

	return New(ret)
}

func TestCommonMatch(t *testing.T) {
	assert.Equal(t, 7, indexCommon(New("12444434456789"), New("4456")))
}

func TestKmpMatch(t *testing.T) {
	assert.Equal(t, 7, indexKmp(New("12444434456789"), New("4456")))
}

func BenchmarkCommonMatch(b *testing.B) {
	b.StopTimer()
	b.Logf("Target:000 \n")

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		indexCommon(strTest, New("000"))
	}
}

func BenchmarkKmpMatch(b *testing.B) {

	b.StopTimer()
	b.Logf("Target:000 \n")

	b.StartTimer()

	for i := 0; i < b.N; i++ {
		indexKmp(strTest, New("000"))
	}

}
