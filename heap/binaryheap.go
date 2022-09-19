/* Package */
package heap

/* Imports */
import (
)

/* Glocals */

/* Types */
type BinaryHeap struct {
  Heap []int
  Last int
}

/* Interface */

func (bh *BinaryHeap) HeapifyInsert(i int) {

  if (bh == nil) {
    return
  }

  bh.Heap = append(bh.Heap, i)
  bh.Last++
  bh.HeapifyUp(bh.Last)
}

func (bh *BinaryHeap) HeapifyRemove(i int) {

  var pos int
  var found bool

  if (bh == nil) {
    return
  }

  for p, n := range bh.Heap {
    if (n == i) {
      pos = p
      found = true
      break
    }
  }

  if (!found) {
    return
  }

  bh.Heap[pos] = bh.Heap[bh.Last]
  bh.Heap = bh.Heap[:len(bh.Heap) - 1]
  bh.Last--
  bh.Heapify(pos)
}

func (bh *BinaryHeap) HeapifyUp(pos int) {

  var father int

  if (bh == nil) {
    return
  }

  for pos > 0 {
    if (pos % 2 != 0) {
      father = (pos - 1) / 2

    } else {
      father = (pos - 2) / 2
    }

    if (bh.Heap[father] > bh.Heap[pos]) {
      aux := bh.Heap[father]
      bh.Heap[father] = bh.Heap[pos]
      bh.Heap[pos] = aux
    } else {
      break
    }

    pos = father
  }
}

func (bh *BinaryHeap) HeapifyDown(pos int) {

  var candidate int

  if (bh == nil) {
    return
  }

  for true {
    lPos := 2 * pos + 1
    rPos := 2 * pos + 2
    if (lPos > bh.Last) {
      break
    }

    if (rPos > bh.Last) {
      candidate = lPos

    } else {
      if (bh.Heap[rPos] < bh.Heap[lPos]) {
        candidate = rPos

      } else {
        candidate = lPos
      }
    }

    if (bh.Heap[candidate] < bh.Heap[pos]) {
      aux := bh.Heap[pos]
      bh.Heap[pos] = bh.Heap[candidate]
      bh.Heap[candidate] = aux
      pos = candidate

    } else {
      break
    }
  }

}

func (bh *BinaryHeap) Heapify(pos int) {

  /*
    Checks whether to heapify up or down
  */

  var father int

  if (bh == nil) {
    return
  }

  if (pos % 2 != 0) {
    father = (pos - 1) / 2

  } else {
    father = (pos - 2) / 2
  }

  if (father < 0) {
    bh.HeapifyDown(pos)
    return
  }

  if (bh.Heap[father] < bh.Heap[pos]) {
    bh.HeapifyDown(pos)

  } else {
    bh.HeapifyUp(pos)
  }
}


/* Functions */
func BinaryHeapInit() (*BinaryHeap) {

  var bh BinaryHeap

  bh.Heap = make([]int, 0)
  bh.Last = -1
  return &bh
}
