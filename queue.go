package queue

type Node struct {
	Data interface{}
	Next *Node
}

func NewNode(v interface{}) *Node {
	node := new(Node)
	node.Data = v
	node.Next = nil
	return node
}

type Queue struct {
	header *Node
	footer *Node
}

func NewQueue() *Queue {
	queue := new(Queue)
	queue.header = nil
	queue.footer = nil
	return queue
}

func (q *Queue) Push(v interface{}) {
	node := NewNode(v)
	if q.header == nil {
		q.header = node
	} else {
		q.footer.Next = node
	}
	q.footer = node
}

func (q *Queue) Pull() interface{} {
	if q.header == nil {
		print("Queue Empty")
		return nil
	}
	aux := q.header.Data
	q.header = q.header.Next
	return aux
}

func (q *Queue) Header() interface{} {
	if q.header == nil {
		return nil
	}
	return q.header.Data
}

func (q *Queue) IsEmpty() bool {
	if q.header == nil {
		print("Queue Empty")
		return true
	}
	return false
}
