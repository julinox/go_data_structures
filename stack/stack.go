/* Package */
package stack

/* Imports */
//import "fmt"

/* Glocals */

/* Types */
type Node struct {
  Next *Node
  Data interface{}
}

type Stack struct {
  Head *Node
  OutputFormat FormatType
}

/* Interface 'Stack' */
func (s *Stack) Push(data interface {}) (bool) {

  var node Node

  if (s == nil) {
    return false
  }

  node.Data = data
  if (s.Head == nil) {
    node.Next = nil
    s.Head = &node

  } else {
    node.Next = s.Head
    s.Head = &node
  }

  return true
}

func (s *Stack) Pop() (i interface {}) {

  if (s.Head == nil) {
    return nil
  }

  i = s.Head.Data
  s.Head = s.Head.Next
  return
}

/* Interface 'Stringer' */
func (s *Stack) String() (string) {

  switch (s.OutputFormat) {
    case FMT1:
      return format1(s)

    case FMT2:
      return format2(s)

    default:
      return format1(s)
  }
}

/* Functions */
