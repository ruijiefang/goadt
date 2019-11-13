package storage

type node struct {
	next *node
	prev *node
	val interface{}
}

type LinkedList struct {
	first *node
	last *node
	size uint32
}

func NewLinkedList() * LinkedList {
	return &(LinkedList{nil,nil,0})
}

func (list *LinkedList) Size() uint32 {
	return list.size
}

func (list *LinkedList) Add(element interface{}) {
	if list.size == 0 {
		list.first = &node{prev: nil, next: nil, val: element}
		list.last = list.first
		list.size ++
		return
	}
	list.last.next = &node{prev: list.last, next: nil, val: element}
	list.last = list.last.next
	list.size ++

}

func (list *LinkedList) Exists(element interface{}) bool {
	if list.size == 0 {
		return false
	}
	for nodeptr := list.first; nodeptr != nil; nodeptr = nodeptr.next {
		if nodeptr.val == element {
			return true
		}
	}
	return false
}

func (list *LinkedList) Remove(element interface{}) bool {
	if list.size == 0 {
		return false
	}
	for nodeptr := list.first; nodeptr != nil; nodeptr = nodeptr.next {
		if nodeptr.val == element {
			if nodeptr.prev != nil {
				nodeptr.prev.next = nodeptr.next
			} else {
				// list.first == nodeptr
				list.first = nodeptr.next
			}

			if nodeptr.next != nil {
				nodeptr.next.prev = nodeptr.prev
			} else {
				// list.last == nodeptr
				list.last = nodeptr.prev
			}
			list.size--
			return true
		}
	}
	return false
}

