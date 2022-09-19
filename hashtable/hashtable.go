/*
  This is a HASH table implementation.
  Colision handling method: 'Separate chaining' (using linked list)
  Insertion: Make new insertion head of the list

  Time complexity:
    - Worst case: O(N)
*/

/*
  OPS:
  - Insert, Search, Delete (HASH Table)
  - Insert, Search, Delete (bucket[linked list])
*/

/* Package */
package hashtable

/* Imports */
import (
  "fmt"
  "datastructures/linkedlist"
)

/* Glocals */
var Cap int = 5 //128

/* Types */
type KeyVal struct {
  Key interface {}
  Value interface {}
}

type HashTable struct {
  Algorithm HashingAlgo
  Capacity int
  hashes []*linkedlist.List
}

/* Interface */
func (ht *HashTable) Insert(kv KeyVal) (bool) {

  var pos int

  if (ht == nil) {
    return true
  }

  pos = ht.Algorithm(kv.Key, ht.Capacity)

  if (pos < 0) {
    return false
  }

  if (ht.hashes[pos] == nil) {
    ht.hashes[pos] = &linkedlist.List{}
  }

  ht.hashes[pos].Insert(kv)
  return true
}

func (ht *HashTable) Delete(key interface {}) (bool) {

  var keyPos, keyHashPos int

  if (ht == nil) {
    return true
  }

  keyHashPos = ht.Algorithm(key, ht.Capacity)
  _, keyPos = ht.Search(key)

  if (keyPos < 0) {
    return false
  }

  ht.hashes[keyHashPos].Delete(uint(keyPos))
  //fmt.Printf("HashPos: %v || Delete '%v' AT '%v'\n", keyHashPos, kv, keyPos)
  return true
}

func (ht *HashTable) Search(key interface {}) (KeyVal, int) {

  var keyPos, keyHashPos int

  if (ht == nil) {
    return KeyVal{}, -1
  }

  keyPos = 0
  keyHashPos = ht.Algorithm(key, ht.Capacity)

  for it := ht.hashes[keyHashPos].Head; it != nil; it = it.Next {
    t,_ := it.Data.(KeyVal)

    if (t.Key == key) {
      return t, keyPos
    }
    keyPos++
  }
  return KeyVal{}, -1
}

/* Functions */
func InitHashTable(capacity int, fn HashingAlgo) *HashTable {

  var ht HashTable

  if (capacity < 1) {
    return nil
  }

  ht.Algorithm = fn
  ht.Capacity = capacity
  ht.hashes = make([]*linkedlist.List, capacity)
  return &ht
}

func PrintAll(ht *HashTable) {

  if (ht == nil) {
    return
  }

  for it := 0; it < len(ht.hashes); it++ {
    fmt.Println(ht.hashes[it])
  }
}
