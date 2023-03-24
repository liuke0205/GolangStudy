package arraylist

type Iterator struct {
	list  *List
	index int
}

func (list *List) Iterator() Iterator {
	return Iterator{list: list, index: -1}
}

func (iterator *Iterator) Next() bool {
	if iterator.index < iterator.list.size {
		iterator.index++
	}
	return iterator.list.withinRange(iterator.index)
}

func (iterator *Iterator) Prev() bool {
	if iterator.index >= 0 {
		iterator.index--
	}
	return iterator.list.withinRange(iterator.index)
}

func (iterator *Iterator) Value() interface{} {
	return iterator.list.elements[iterator.index]
}

func (iterator *Iterator) Index() int {
	return iterator.index
}

func (iterator *Iterator) Begin() {
	iterator.index = -1
}

func (iterator *Iterator) End() {
	iterator.index = iterator.list.size
}

func (iterator *Iterator) First() bool {
	iterator.Begin()
	return iterator.Next()
}

func (iterator *Iterator) Last() bool {
	iterator.End()
	return iterator.Prev()
}

//NextTo 将迭代器从当前位置移动到下一个元素，该元素满足
//传递了函数，如果容器中有下一个元素，则返回true。
//如果NextTo返回true，那么index和value可以检索下一个元素的索引和值。
//修改迭代器的状态。
func (iterator *Iterator) NextTo(f func(index int, value interface{}) bool) bool {
	for iterator.Next() {
		index, value := iterator.Index(), iterator.Value()
		if f(index, value) {
			return true
		}
	}
	return false
}

func (iterator *Iterator) PrevTo(f func(index int, value interface{}) bool) bool {
	for iterator.Prev() {
		index, value := iterator.Index(), iterator.Value()
		if f(index, value) {
			return true
		}
	}
	return false
}
