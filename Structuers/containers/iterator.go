package containers

// IteratorWithIndex 是有序容器的有状态迭代器，其值可以通过索引获取。
type IteratorWithIndex interface {

	//Next 将迭代器移动到下一个元素，如果容器中有下一个元件，则返回true。
	//如果Next()返回true，那么index()和value()可以检索下一个元素的索引和值。
	//如果第一次调用Next()，那么它会将迭代器指向第一个元素（如果存在的话）。
	Next() bool

	//Value 返回当前元素的值，不修改迭代器的状态
	Value() interface{}

	//Index 返回当前元素的索引，不修改迭代器的状态
	Index() int

	//Begin 将迭代器重置为初始状态,调用Next来获取第一个元素（如果有的话）
	Begin()

	//First 将迭代器移动到第一个元素，如果容器中有第一个元素则返回true
	//如果First返回true，那么index和value可以检索第一个元素的索引和值
	//修改迭代器的状态
	First() bool

	//NextTo 将迭代器从当前位置移动到下一个元素，该元素满足
	//传递了函数，如果容器中有下一个元素，则返回true。
	//如果NextTo返回true，那么index和value可以检索下一个元素的索引和值。
	//修改迭代器的状态
	NextTo(func(index int, value interface{}) bool) bool
}

// IteratorWithKey 是一个有状态迭代器，用于元素为键值对的有序容器
type IteratorWithKey interface {
	//Next 将迭代器移动到下一个元素，如果容器中有下一个元件，则返回true。
	//如果Next返回true，则可以通过key和value检索下一个元素的键和值。
	//如果第一次调用Next，那么它会将迭代器指向第一个元素（如果存在的话）。
	//修改迭代器的状态。
	Next() bool

	//Value 返回当前元素的值。
	//不修改迭代器的状态。
	Value() interface{}

	// Key returns the current element's key.
	// Does not modify the state of the iterator.
	Key() interface{}

	// Begin resets the iterator to its initial state (one-before-first)
	// Call Next() to fetch the first element if any.
	Begin()

	// First moves the iterator to the first element and returns true if there was a first element in the container.
	// If First() returns true, then first element's key and value can be retrieved by Key() and Value().
	// Modifies the state of the iterator.
	First() bool

	// NextTo moves the iterator to the next element from current position that satisfies the condition given by the
	// passed function, and returns true if there was a next element in the container.
	// If NextTo() returns true, then next element's key and value can be retrieved by Key() and Value().
	// Modifies the state of the iterator.
	NextTo(func(key interface{}, value interface{}) bool) bool
}

// ReverseIteratorWithIndex is stateful iterator for ordered containers whose values can be fetched by an index.
//
// Essentially it is the same as IteratorWithIndex, but provides additional:
//
// Prev() function to enable traversal in reverse
//
// Last() function to move the iterator to the last element.
//
// End() function to move the iterator past the last element (one-past-the-end).
type ReverseIteratorWithIndex interface {
	// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
	// If Prev() returns true, then previous element's index and value can be retrieved by Index() and Value().
	// Modifies the state of the iterator.
	Prev() bool

	// End moves the iterator past the last element (one-past-the-end).
	// Call Prev() to fetch the last element if any.
	End()

	// Last moves the iterator to the last element and returns true if there was a last element in the container.
	// If Last() returns true, then last element's index and value can be retrieved by Index() and Value().
	// Modifies the state of the iterator.
	Last() bool

	// PrevTo moves the iterator to the previous element from current position that satisfies the condition given by the
	// passed function, and returns true if there was a next element in the container.
	// If PrevTo() returns true, then next element's index and value can be retrieved by Index() and Value().
	// Modifies the state of the iterator.
	PrevTo(func(index int, value interface{}) bool) bool

	IteratorWithIndex
}

// ReverseIteratorWithKey is a stateful iterator for ordered containers whose elements are key value pairs.
//
// Essentially it is the same as IteratorWithKey, but provides additional:
//
// Prev() function to enable traversal in reverse
//
// Last() function to move the iterator to the last element.
type ReverseIteratorWithKey interface {
	// Prev moves the iterator to the previous element and returns true if there was a previous element in the container.
	// If Prev() returns true, then previous element's key and value can be retrieved by Key() and Value().
	// Modifies the state of the iterator.
	Prev() bool

	// End moves the iterator past the last element (one-past-the-end).
	// Call Prev() to fetch the last element if any.
	End()

	// Last moves the iterator to the last element and returns true if there was a last element in the container.
	// If Last() returns true, then last element's key and value can be retrieved by Key() and Value().
	// Modifies the state of the iterator.
	Last() bool

	// PrevTo moves the iterator to the previous element from current position that satisfies the condition given by the
	// passed function, and returns true if there was a next element in the container.
	// If PrevTo() returns true, then next element's key and value can be retrieved by Key() and Value().
	// Modifies the state of the iterator.
	PrevTo(func(key interface{}, value interface{}) bool) bool

	IteratorWithKey
}
