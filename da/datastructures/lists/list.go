package lists

// 队列接口
type ListInterface interface {
	Init()
	IsEmpty() bool
	Clear()
	GetElem()
	LocateElem()
	Insert()
	Delete()
	Length()
	Reverse()
}

// 栈接口
type StackInterface interface {
	Init()
	Destorty()
	Clear()
	IsEmpty() bool
	GetTop()
	Push()
	Pop()
	Length()
}
