/* Package */
package heap

/* Imports */
import ()

/* Glocals */

/* Types */

/* Interface */
type Heap interface {
  HeapifyUp()
  HeapifyDown()
  HeapifyInsert(int)
  HeapifyRemove(int)
}

/* Functions */
