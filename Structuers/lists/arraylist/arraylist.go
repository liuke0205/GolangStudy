package arraylist

import (
	"GolangStudy/Structuers/utils"
	"fmt"
	"strings"
)

//var _ lists.List = (*List)(nil)

type List struct {
	elements []interface{}
	size     int
}

const (
	growthFactor = float32(2.0)  // growth by 100%
	shrinkFactor = float32(0.25) // shrink when size is 25% of capacity (0 means never shrink)
)

// New 实例化一个新列表，并将传递的值（如果有的话）添加到列表中
func New(values ...interface{}) *List {
	list := &List{}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

// Add 在列表末尾添加一个值
func (list *List) Add(values ...interface{}) {
	// 检查列表是否需要扩容
	list.growBy(len(values))
	for _, value := range values {
		list.elements[list.size] = value
		list.size++
	}
}

// Get 根据下标获取列表的值
func (list *List) Get(index int) (interface{}, bool) {
	if !list.withinRange(index) {
		return nil, false
	}
	return list.elements[index], true
}

// Remove 根据下标移除列表中的元素
func (list *List) Remove(index int) {
	if !list.withinRange(index) {
		return
	}
	list.elements[index] = nil
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size--
	list.shrink()
}

// Contains
// 包含检查集合中是否存在元素（一个或多个）。
// 所有元素都必须存在于集合中，方法才能返回true。
// 性能时间复杂度为n^2。
// 如果根本没有传递任何参数，则返回true，即set始终是空集的超集。
func (list *List) Contains(values ...interface{}) bool {
	for _, val := range values {
		found := false
		for i := 0; i < list.size; i++ {
			if list.elements[i] == val {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Values 返回列表中的所有元素
func (list *List) Values() []interface{} {
	newElements := make([]interface{}, list.size, list.size)
	copy(newElements, list.elements[:list.size])
	return newElements
}

// IndexOf 获取列表中值的下标，如果没有返回-1
func (list *List) IndexOf(value interface{}) int {
	if list.size == 0 {
		return -1
	}
	for idx, element := range list.elements {
		if value == element {
			return idx
		}
	}
	return -1
}

// Empty 如果列表不包含任何元素，则返回true。
func (list *List) Empty() bool {
	return list.size == 0
}

// Size 返回列表中元素的数量
func (list *List) Size() int {
	return list.size
}

// Clear 清除列表中所有元素
func (list *List) Clear() {
	list.size = 0
	list.elements = []interface{}{}
}

// Sort 使用对值（就地）进行排序。
func (list *List) Sort(comparator utils.Comparator) {
	if len(list.elements) < 2 {
		return
	}
	utils.Sort(list.elements[:list.size], comparator)
}

// Swap swaps the two values at the specified positions.
func (list *List) Swap(i, j int) {
	if list.withinRange(i) && list.withinRange(j) {
		list.elements[i], list.elements[j] = list.elements[j], list.elements[i]
	}
}

// Insert
// Insert在指定的索引位置插入值，将该位置的值（如果有）和任何后续元素向右移动。
// 如果位置为负数或大于列表大小，则不执行任何操作
// 注意：等于列表大小的位置是有效的，即追加。
func (list *List) Insert(index int, values ...interface{}) {
	if !list.withinRange(index) {
		if list.size == index {
			list.Add(values...)
		}
		return
	}
	len := len(values)
	list.growBy(len)
	list.size += len
	copy(list.elements[index+len:], list.elements[index:list.size-len])
	copy(list.elements[index:], values)
}

// Set
// 在指定索引处设置值
// 如果位置为负数或大于列表大小，则不执行任何操作
// 注意：等于列表大小的位置是有效的，即追加。
func (list *List) Set(index int, value interface{}) {
	if !list.withinRange(index) {
		if list.size == index {
			list.Add(value)
		}
		return
	}
	list.elements[index] = value
}

// String returns a string representation of container
func (list *List) String() string {
	str := "ArrayList\n"
	values := []string{}
	for _, value := range list.elements[:list.size] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// 必要时扩展列表，即如果我们添加n个元素，将达到容量
func (list *List) growBy(len int) {
	curCapacity := cap(list.elements)
	if list.size+len >= curCapacity {
		newCapacity := int(float32(curCapacity+len) * growthFactor)
		list.reSize(newCapacity)
	}
}

func (list *List) reSize(capacity int) {
	newElements := make([]interface{}, capacity, capacity)
	copy(newElements, list.elements)
	list.elements = newElements
}

// withinRange 检查索引是否在列表的范围内
func (list *List) withinRange(index int) bool {
	return index >= 0 && index < list.size
}

// shrink 必要时收缩列表
func (list *List) shrink() {
	if shrinkFactor == 0.0 {
		return
	}
	curCapacity := cap(list.elements)
	if list.size <= int(float32(curCapacity)*shrinkFactor) {
		list.reSize(list.size)
	}
}
