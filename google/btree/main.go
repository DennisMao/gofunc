package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/btree"
)

type Data struct {
	Id      int64
	Content string
}

func (this *Data) Less(item btree.Item) bool {
	if v, ok := item.(*Data); ok {
		return this.Id < v.Id
	}

	fmt.Println("Error: not valid item")
	return false
}

func main() {

	lenStorage := 10
	r := rand.New(rand.NewSource(time.Now().Unix()))

	// 初始化树
	tree := btree.New(2)

	for i := 0; i < lenStorage; i++ {
		d := Data{
			Id:      int64(i),
			Content: fmt.Sprintf("%s%d", "data", r.Int63()),
		}

		fmt.Printf("id:%d name:%s\n", d.Id, d.Content)
		tree.ReplaceOrInsert(&d)
	}
	fmt.Println("====================")

	// 迭代器
	var st []*Data
	f := func(item btree.Item) bool {
		if v, ok := item.(*Data); ok {
			st = append(st, v)
			return true
		}
		return false
	}
	tree.AscendLessThan(&Data{3, ""}, f)

	if len(st) > 0 {

		fmt.Print("less than 3 ")
		for i, _ := range st {
			fmt.Print(st[i], ",")
		}
		fmt.Printf("\n")
	}

	// 树操作
	fmt.Println("len:       ", tree.Len())
	fmt.Println("get3:      ", tree.Get(&Data{3, ""}))
	fmt.Println("replace3:  ", tree.ReplaceOrInsert(&Data{3, "test replace"}))
	fmt.Println("delete4:   ", tree.Delete(&Data{4, ""}))
	fmt.Println("replace5:  ", tree.ReplaceOrInsert(&Data{5, "test replace"}))
	fmt.Println("has:       ", tree.Has(&Data{5, ""}))
	fmt.Println("min:       ", tree.Min())
	fmt.Println("delmin:    ", tree.DeleteMin())
	fmt.Println("max:       ", tree.Max())
	fmt.Println("delmax:    ", tree.DeleteMax())
	fmt.Println("len:       ", tree.Len())

}
