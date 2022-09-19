/*
  Get a queue using slices.
*/

/* Package */
package queue

/* Imports */
import "fmt"

/* Glocals */
var QueueSize uint = 10

/* Types */
type Queue struct {
  size uint
  head uint
  tail uint
  queue []interface {}
}

/* Interface */
func (q *Queue) Enqueue(i interface {}) (bool) {

  if (q == nil) {
    return false
  }

  if (q.tail < q.size) {
    q.queue[q.tail] = i

  } else {
    q.queue = append(q.queue, i)
  }

  q.tail++
  return true
}

func (q *Queue) Dequeue() (interface {}) {

  var aux interface {}


  if (q == nil) {
    return nil
  }

  if (q.head >= q.tail) {
    return nil
  }
  aux = q.queue[q.head]
  q.head++
  return aux
}

func (q *Queue) GetQueue() (*[]interface{}){

  if (q == nil) {
    return nil
  }

  return &(q.queue)
}

func (q *Queue) String() (string) {

  var ret string

  if (q == nil) {
    return ""
  }

  ret = fmt.Sprintf("Head: %v | Tail: %v | [", q.head, q.tail)
  for i, v := range q.queue {
    if (i < len(q.queue) - 1) {
      ret += fmt.Sprintf("%v, ", v)

    } else {
      ret += fmt.Sprintf("%v]", v)
    }
  }

  return ret
}

func (q *Queue) GetHead() (int) {

  if (q == nil) {
    return -1
  }

  return int(q.head)
}

func (q *Queue) GetTail() (int) {

  if (q == nil) {
    return -1
  }

  return int(q.tail)
}

func (q *Queue) Search(k interface {}) (int) {

  if (q == nil) {
    return -1
  }

  for i, v := range q.queue {
    if (v == k) {
      return i
    }
  }

  return -1
}

func (q *Queue) IsEmpty() (bool) {
  if (q.head >= q.tail) {
    return true
  }
  return false
}

/* Functions */
func InitQueue(queueSize uint) (*Queue){

  if (queueSize == 0) {
    queueSize = QueueSize
  }

  q := Queue{}
  q.size = queueSize
  q.head = 0
  q.tail = 0
  q.queue = make([]interface{}, queueSize)
  return &q
}
