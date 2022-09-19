/* Package */
package hashtable

/* Imports */
import (
  "fmt"
)

/* Glocals */

/* Types */
type HashingAlgo func(interface {}, int) (int)

/* Interface */

/* Functions */
func Hashing1(key interface {}, capacity int) (int) {

  iStr, iBool := key.(string)

  if (iBool) {
    return sumStrBytes(iStr) % capacity
  }

  iInt, iBool := key.(int)

  if (iBool) {
    return sumStrBytes(fmt.Sprintf("%v", iInt)) % capacity
  }

  return -1
}

func sumStrBytes(str string) (int){

  /*
    String to bytes. Then sum
  */
  var sum int

  sum = 0
  iStr := []byte(str)

  for it := 0; it < len(iStr); it++ {
    sum += int(iStr[it])
  }
  return sum
}
