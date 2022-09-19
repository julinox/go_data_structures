/* Package */
package stack

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
func format1(s *Stack) (string) {

  var msg string

  i := 0
  it := s.Head

  for (it != nil) {
    //msg += fmt.Sprintf("%v | ", it.Data)
    msg += fmt.Sprintf("N%v: %v | ", i, it.Data)
    it = it.Next
    i++
  }
  return msg
}

func format2(s *Stack) (string) {

  var msg string

  i := 0
  it := s.Head

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
