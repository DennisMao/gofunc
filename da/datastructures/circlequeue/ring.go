// 线性表-双向链表
package Ring

const (
	DefaultSize = 128
)

type Item string

type ring struct {
	front, rear *Element

	size, used int
}

type Element struct {
	data Item

	next, prev *Element
}

func (this *Element) Data() Item {
	return this.data
}

func New(size int) *ring {
	if size == 0 {
		size = DefaultSize
	}

	return &ring{nil, nil, size, 0}
}

func (this *ring) Size() int {
	return this.size
}

func (this *ring) Len() int {
	return this.used
}

func (this *ring) Search(it Item) *Element {
	return search(this.front, it)
}

func search(curElement *Element, it Item) *Element {
	if curElement == nil {
		return nil
	}

	if curElement.data == it {
		return curElement
	}

	return search(curElement.next, it)
}

func (this *ring) Front() *Element {
	return this.front
}

func (this *ring) Back() *Element {
	return this.rear
}

func (this *ring) RangeFrom(it Item) []Item {
	ret := make([]Item, 0)
	ele := this.Search(it)
	if ele == nil {
		return ret
	}

	p := ele
	for p != nil {
		ret = append(ret, p.Data())
		p = p.next
	}

	return ret

}

func (this *ring) Insert(it Item, ele *Element) {
	if this.used >= this.size {
		return
	}

	if ele == nil {
		return
	}

	if eleExist := this.Search(it); eleExist != nil {
		return
	}

	newElement := &Element{
		data: it,
		prev: ele,
		next: ele.next,
	}

	if ele.next != nil {
		ele.next.prev = newElement
	}

	ele.next = newElement

	if this.rear == ele {
		this.rear = newElement
	}

	this.used++
}

func (this *ring) PushFront(it Item) {
	if this.used >= this.size {
		return
	}

	if ele := this.Search(it); ele != nil {
		return
	}

	newElement := &Element{
		data: it,
		next: nil,
	}

	if this.front == nil {
		this.front = newElement
		this.rear = newElement
		this.used++
		return
	}

	newElement.next = this.front
	this.front.prev = newElement
	this.front = newElement
	this.used++
}

func (this *ring) PushBack(it Item) {
	if this.used >= this.size {
		return
	}

	if ele := this.Search(it); ele != nil {
		return
	}

	newElement := &Element{
		data: it,
		next: nil,
	}

	if this.front == nil {
		this.front = newElement
		this.rear = newElement
		this.used++
		return
	}

	newElement.prev = this.rear
	this.rear.next = newElement
	this.rear = newElement
	this.used++
}

func (this *ring) Reverse() {
	p := this.front
	this.rear, this.front = this.front, this.rear
	for p != nil {
		p.next, p.prev = p.prev, p.next
		p = p.prev
	}
	return
}

func (this *ring) Remove(it Item) {
	ele := this.Search(it)
	if ele == nil {
		return
	}

	if this.rear == ele {
		this.rear = ele.prev
	}

	if this.front == ele {
		this.front = ele.next
	}

	if ele.prev != nil {
		ele.prev.next = ele.next
	}

	this.used--

}

/////////////// Element operation /////////////////
func removeEle(ele *Element) {
	if ele.prev != nil {
		ele.prev.next = ele.next
	}
}
