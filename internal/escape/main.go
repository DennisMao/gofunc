

// 逃逸分析
// Go版本: 1.11
// 指令: go tool compile -m main.go |grep escape
//
// 1.make,new,匿名函数 会直接导致逃逸
// 2.& 取引用操作,且外部有引用或作为返回值
// 3.已指向对象的interface{}  (即 interface{}.(value)已赋值的对象),且外部有引用或作为返回值
//
// 内联
// go tool compile -m main.go |grep inline
// 1.同func内的匿名函数,且匿名函数不含其他包的引用
// 2.同package下的命名函数,且函数内不含有其他包引用

package main

import (
	"fmt"
)

func TestMake() []string {
	makeI := make([]string, 0)      // 逃逸
	makeE := make([]string, 0)      // 逃逸
	makeI = []string{"1", "2", "3"} // 逃逸
	makeE = []string{"4", "5", "6"} // 逃逸

	makeI = makeE
	makeE = makeI
	return makeE
}

func TestNew() *string {
	newI := new(string) // 逃逸
	newE := new(string) // 逃逸

	newI = newE
	newE = newI
	return newE
}

func TestInt() int {
	intI := 123
	intE := 456

	intI = intE
	intE = intI
	return intE
}

func TestIntReference() *int {
	intI := 123
	intE := 456

	intI = intE
	intE = intI

	return &intE // 逃逸
}

func TestString() string {
	stringI := "123" //
	stringE := "456" // 逃逸

	stringI = stringE
	stringE = stringI

	return stringE
}

type StructStruct struct {
	Name string
}

func TestStruct() StructStruct {
	structI := StructStruct{}
	structE := StructStruct{}

	structI = structE
	structE = structI
	return structE
}

func TestInterface() interface{} {
	var interfaceNilI interface{} = nil
	var interfaceNilE interface{} = nil
	var interfaceI interface{} = 123 // 逃逸
	var interfaceE interface{} = 456 // 逃逸

	interfaceNilI = interfaceNilE
	interfaceNilE = interfaceNilI

	interfaceI = interfaceE
	interfaceE = interfaceI
	return interfaceE
}

func TestFunc() func() {
	funcI := func() {} // 逃逸
	funcE := func() {} // 逃逸

	funcI = funcE
	funcE = funcI

	return funcI
}

func main() {

	// inline
	inlineFunc := func() string {
		return "123"
	}()

	inlineFunc2 := func() string {
		fmt.Println("123")
		return "123"
	}()

	fmt.Println(inlineFunc)
	fmt.Println(inlineFunc2)

}

