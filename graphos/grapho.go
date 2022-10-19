/*
  'Grapho' inteface.
  Vertex info can be stored using the 'info' (interface {}) parameter) either
  at creation time (VertexAdd()) or update (VertexAddInfo())
*/

/* Package */
package graphos

/* Imports */
import (
  "fmt"
)

/* Glocals */
const (
  GRAPH_DIRECTED byte   = 1 << 0
  GRAPH_HAS_CYCLE byte  = 1 << 1
  GRAPH_MULTIEDGE byte  = 1 << 2
)

/* Types */
type Vertex struct {
  ID int
  Depth int
  Tag string
  InDegree int
  Info interface {}
}

type Edge struct {
  Peer int
  Weight int
}

/* Interface 'Grapho'*/
type Grapho interface {
  GraphHeight() (int)
  GraphFlags() (byte)
  GraphFlagSet(byte)
  VertexAdd(int, int, string, interface{}) (bool)
  VertexRemove(int) (bool)
  VertexAddDepth(int, int)
  VertexAddTag(int, string)
  VertexAddInfo(int, interface {})
  VertexTag(int) (string)
  VertexInfo(int) (interface {})
  VertexMax() (int)
  VertexList() (*[]int)
  VertexNeighbours(int) (*[]int)
  VertexEdges(int) (*[]Edge)
  VertexInDegree(int) (int)
  EdgeAdd(int, int, int) (bool)
  EdgeRemove(int, int, int) (bool)
  EdgeWeight(int, int) (int)
  ExportAsGviz(string) ()
  PrintFlags() (string)
}

/* Interface 'Stringer' */
func (v *Vertex) String() (string) {

  if (v.Tag == "") {
    return fmt.Sprintf("NIL/%v/%v", v.Depth, v.Info)
  }
  return fmt.Sprintf("%v/%v/%v", v.Tag, v.Depth, v.Info)
}

/* Functions */
