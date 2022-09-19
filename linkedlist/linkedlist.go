/* Package */
package linkedlist

/* Imports */

/* Glocals */

/* Types */
type Node struct {
  Next *Node
  Data interface{}
}

type List struct {
  Head *Node
  Stack bool
  OutputFormat FormatType
}

/* Interface 'List' */
func (l *List) Insert(data interface {}) (bool) {

  if (l == nil) {
    return false
  }

  node := Node{nil, data}
  current := l.Head

  if (current == nil) {
    l.Head = &node
    return true
  }

  if (l.Stack) {
    node.Next = l.Head
    l.Head = &node
    return true
  }

  for (current != nil) {
    if (current.Next == nil) {
      break
    }
    current = current.Next
  }

  current.Next = &node
  return true
}

func (l *List) Delete(pos uint) (bool) {

  it := l.Head

  if (it == nil) {
    return true
  }

  if (pos == 0) {
    l.Head = it.Next
    it.Next = nil
    return true
  }

  for (it != nil && pos > 1) {
    it = it.Next
    pos--
  }

  if (it == nil || it.Next == nil) {
    return false
  }

  aux := it.Next
  it.Next = aux.Next
  aux.Next = nil
  return true
}

func (l *List) Search(i interface {}) (int, interface {}) {

  var pos int

  pos = 0
  if (l == nil) {
    return -1, nil
  }

  for it := l.Head; it != nil; it = it.Next {
    if (it.Data == i) {
      return pos, it.Data
    }
    pos++
  }
  return -2, nil
}

func (l *List) Reverse() {

  var prev *Node

  if (l == nil) {
    return
  }

  it := l.Head
  act := l.Head
  prev = nil

  for (it != nil) {
    it = it.Next
    act.Next = prev
    prev = act
    act = it
  }

  l.Head = prev
}

func (l *List) ReverseRec(current *Node, nextPtr *Node) {

  if (current == nil) {
    return
  }

  l.ReverseRec(current.Next, current)

  if (current.Next == nil) {
    l.Head = current
  }

  current.Next = nextPtr
}

/* Interface 'Stringer' */
func (l *List) String() (string) {

  switch (l.OutputFormat) {
    case FMT1:
      return format1(l)

    case FMT2:
      return format2(l)

    case FMT3:
      return format3(l)

    default:
      return format1(l)
  }
}

/* Functions */
