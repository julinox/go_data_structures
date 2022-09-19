/* Package */
package linkedlist

/* Imports */
import (
  "fmt"
)

/* Glocals */

/* Types */
type FormatType uint8

const (
  FMT1 FormatType = iota
  FMT2
  FMT3
)

/* Interface */

/* Functions */
func format1(l *List) (string) {

  var msg string

  i := 0
  it := l.Head

  for (it != nil) {
    msg += fmt.Sprintf("N%v: %v | ", i, it.Data)
    it = it.Next
    i++
  }
  return msg
}

func format2(l *List) (string) {

  var msg string

  i := 0
  it := l.Head

  for (it != nil) {
    if (it.Next != nil) {
      msg += fmt.Sprintf("N%v: %v\n", i, it.Data)
    } else {
      msg += fmt.Sprintf("N%v: %v", i, it.Data)
    }

    it = it.Next
    i++
  }

  return msg
}

func format3(l *List) (string) {

  var msg string

  i := 0
  it := l.Head

  for (it != nil) {
    if (it.Next != nil) {
      msg += fmt.Sprintf("Node: %v | Data: %v\n", i, it.Data)
    } else {
      msg += fmt.Sprintf("Node: %v | Data: %v", i, it.Data)
    }

    it = it.Next
    i++
  }

  return msg
}
