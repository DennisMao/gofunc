//使用Go实现C++的STL库
//数据结构练习
//抽象工厂模式
package stl

type Stl interface {
	Len()
	Insert()
	Remove()
	Push()
	Swap()
	Clear()
}

type ElementAccess interface {
	At()
	Front()
	Back()
}
