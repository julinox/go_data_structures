/* Package */
package graphos

/* Imports */
import (
  "os"
  "fmt"
)

/* Glocals */

/* Types */
type GraphList struct {
  Max int
  Height int
  Flags byte
  Edges map[int][]Edge
  Vertices map[int]*Vertex
}

/* Interface 'Grapho' */
func (gl *GraphList) GraphHeight() (int) {

  if (gl == nil) {
    return -1
  }

  return gl.Height
}

func (gl *GraphList) GraphFlags() (byte) {

  if (gl == nil) {
    return 0
  }

  return gl.Flags
}

func (gl *GraphList) GraphFlagSet(flag byte) {

  if (gl == nil) {
    return
  }

  gl.Flags |= flag
}
func (gl *GraphList) VertexAdd(id, depth int, tag string, info interface {}) (bool) {

  var newVertex Vertex

  if (gl == nil) {
    return false
  }

  if (gl.Vertices[id] != nil) {
    return false
  }

  newVertex.ID = id
  newVertex.Depth = depth
  newVertex.Tag = tag
  newVertex.Info = info
  newVertex.InDegree = 0
  gl.Vertices[id] = &newVertex
  return true
}

func (gl *GraphList) VertexRemove(v int) (bool) {

  if (gl == nil || gl.Vertices[v] == nil) {
    return false
  }

  for _, p := range gl.Edges[v] {
    gl.EdgeRemove(v, p.Peer, p.Weight)
  }

  if (gl.Flags & GRAPH_DIRECTED == GRAPH_DIRECTED) {
    for i, edges := range gl.Edges {
      for j, p := range edges {
        if (p.Peer == v) {
          gl.Edges[i][j] = gl.Edges[i][len(gl.Edges[i]) - 1]
          gl.Edges[i] = gl.Edges[i][:len(gl.Edges[i]) - 1]
        }
      }
    }
  }

  delete(gl.Vertices, v)
  return true
}

func (gl *GraphList) VertexAddDepth(id, depth int) {

  if (gl == nil || gl.Vertices[id] == nil) {
    return
  }

  gl.Vertices[id].Depth = depth
  if (depth > gl.Height) {
    gl.Height = depth
  }
}

func (gl *GraphList) VertexAddTag(id int, tag string) {

  if (gl == nil || gl.Vertices[id] == nil) {
    return
  }

  gl.Vertices[id].Tag = tag
}

func (gl *GraphList) VertexAddInfo(id int, info interface {}) {

  if (gl == nil || gl.Vertices[id] == nil) {
    return
  }

  gl.Vertices[id].Info = info
}

func (gl *GraphList) VertexTag(v int) (string) {

  if (gl == nil || gl.Vertices[v] == nil) {
    return ""
  }

  return gl.Vertices[v].Tag
}

func (gl *GraphList) VertexInfo(v int) (interface {}) {

  if (gl == nil || gl.Vertices[v] == nil) {
    return nil
  }

  return gl.Vertices[v].Info
}

func (gl *GraphList) VertexMax() (int) {

  if (gl == nil) {
    return -1
  }

  return gl.Max
}

func (gl *GraphList) VertexList() (*[]int) {

  var vertices []int

  if (gl == nil) {
    return &([]int{})
  }

  for v := range gl.Vertices {
    vertices = append(vertices, v)
  }

  return &vertices
}

func (gl *GraphList) VertexNeighbours(v int) (*[]int) {

  var neighbours []int

  if (gl == nil || v < 0) {
    return &([]int{})
  }

  for _, v := range gl.Edges[v] {
    neighbours = append(neighbours, v.Peer)
  }

  return &neighbours
}

func (gl *GraphList) VertexEdges(v int) (*[]Edge) {

  var neighbours []Edge

  if (gl == nil || v < 0) {
    return &([]Edge{})
  }

  for _, e := range gl.Edges[v] {
    neighbours = append(neighbours, e)
  }

  return &neighbours
}

func (gl *GraphList) VertexInDegree(v int) (int) {

  if (gl == nil || gl.Vertices[v] == nil) {
    return -1
  }

  return gl.Vertices[v].InDegree
}

//##################### EDGE #####################
func (gl *GraphList) EdgeAdd(v1, v2, weight int) (bool) {

  /*
    Vertices 'v1' and 'v2' are added if necessary
  */

  if (gl == nil) {
    return false
  }

  if (gl.Vertices[v1] == nil) {
    gl.VertexAdd(v1, -1, "", nil)
  }

  if (gl.Vertices[v2] == nil) {
    gl.VertexAdd(v2, -1, "", nil)
  }

  if !(gl._edgeAdd(v1, v2, weight)) {
    return false
  }

  if !(gl.Flags & GRAPH_DIRECTED == GRAPH_DIRECTED) {
    gl._edgeAdd(v2, v1, weight)

  }

  return true
}

func (gl *GraphList) _edgeAdd(v1, v2, weight int) (bool) {

  if (gl == nil) {
    return false
  }

  if (gl.Edges == nil) {
    // Double check
    gl.Edges = make(map[int][]Edge)
  }

  // Check if edge already exist
  for i, v := range gl.Edges[v1] {
    if (v.Peer == v2) {
      // Update edge weight
      if !(gl.Flags & GRAPH_MULTIEDGE == GRAPH_MULTIEDGE) {
        gl.Edges[v1][i].Weight = weight
        return true
      }
    }
  }

  if (v1 > gl.Max) {
    gl.Max = v1
  }

  if (v2 > gl.Max) {
    gl.Max = v2
  }

  gl.Edges[v1] = append(gl.Edges[v1], Edge{v2, weight})
  gl.Vertices[v2].InDegree += 1
  return true
}

func (gl *GraphList) EdgeRemove(v1, v2, weight int) (bool) {

  if (gl == nil) {
    return false
  }

  if !(gl._edgeRemove(v1,v2, weight)) {
    return false
  }

  if !(gl.Flags & GRAPH_DIRECTED == GRAPH_DIRECTED) {
    gl._edgeRemove(v2, v1, weight)
  }

  return true
}

func (gl *GraphList) _edgeRemove(v1, v2, weight int) (bool) {

  var i int

  if (gl == nil) {
    return false
  }

  for i = 0; i < len(gl.Edges[v1]); i++ {
    if (gl.Edges[v1][i].Peer == v2) {
      if (gl.Flags & GRAPH_MULTIEDGE == GRAPH_MULTIEDGE) {
        if (gl.Edges[v1][i].Weight == weight) { break }

      } else { break }
    }
  }

  gl.Edges[v1][i] = gl.Edges[v1][len(gl.Edges[v1]) - 1]
  gl.Edges[v1] = gl.Edges[v1][:len(gl.Edges[v1]) - 1]
  if (len(gl.Edges[v1]) < 1) {
    delete(gl.Edges, v1)
  }

  gl.Vertices[v2].InDegree -= 1
  return true
}

func (gl *GraphList) EdgeWeight(v1, v2 int) (int) {

  if (gl == nil) {
    return 0
  }

  edges := gl.Edges[v1]
  for _, e := range edges {
    if (e.Peer == v2) {
      return e.Weight
    }
  }

  return 0
}

//##################### PRINT #####################
func (gl *GraphList) PrintFlags() (string) {

  var ret string

  if (gl == nil) {
    return ""
  }

  ret = "Flags: "
  if (gl.Flags & GRAPH_DIRECTED == GRAPH_DIRECTED) {
    ret += "DIRECTED |"

  } else {
    ret += "UNDIRECTED | "
  }

  if (gl.Flags & GRAPH_MULTIEDGE == GRAPH_MULTIEDGE) {
    ret += "MULTI-EDGE"

  } else {
    ret += "SINGLE-EDGE"
  }

  if (gl.Flags & GRAPH_HAS_CYCLE == GRAPH_HAS_CYCLE) {
    ret += " | CYCLE"
  }

  return ret
}

func (gl *GraphList) ExportAsGviz(fileName string) (error) {

  /*
    Export as 'graphviz' format
  */

  var err2 error = nil
  var arrow string = "--"

  if (fileName == "") {
    return nil
  }

  f, err1 := os.Create(fileName)
  if (err1 != nil) {
    return err1
  }

  if (gl.Flags & GRAPH_DIRECTED == GRAPH_DIRECTED) {
    arrow = "->"
    _, err2 = fmt.Fprintf(f, "digraph {\n")

  } else if (gl.Flags & GRAPH_MULTIEDGE == GRAPH_MULTIEDGE) {
    _, err2 = fmt.Fprintf(f, "graph {\n")

  } else {
    _, err2 = fmt.Fprintf(f, "strict graph {\n")
  }

  defer f.Close()
  if (err2 != nil) {
    return err2
  }

  for _, v1 := range *(gl.VertexList()) {
    l1 := fmt.Sprintf("%v", v1)
    aux := len(*(gl.VertexNeighbours(v1))) - 1
    if (gl.VertexTag(v1) != "") {
      l1 = fmt.Sprintf("%v_%v", gl.VertexTag(v1), v1)
    }

    fmt.Fprintf(f, "  %v %v {", l1, arrow)
    for i, v2 := range *(gl.VertexNeighbours(v1)) {
      l2 := fmt.Sprintf("%v", v2)
      space := " "
      if (i == aux) {
        space = ""
      }

      if (gl.VertexTag(v2) != "") {
        l2 = fmt.Sprintf("%v_%v", gl.VertexTag(v2), v2)
      }
      _, err3 := fmt.Fprintf(f, "%v%v", l2, space)
      if (err3 != nil) {
        return err3
      }
    }

    fmt.Fprintf(f, "}\n")
  }

  fmt.Fprintf(f, "}\n")
  return nil
}

/* Interface 'stringer' */
func (gl *GraphList) String() (string) {

  var ret string

  ret += fmt.Sprintf("(TAG/DEPTH/{INFO}/INDEGREE)\n")
  for i, _ := range gl.Vertices {
    ret += fmt.Sprintf("'%v'\t(%v/%v): [", i, gl.Vertices[i],  gl.Vertices[i].InDegree)
    for j, e := range gl.Edges[i] {
      if (j < len(gl.Edges[i]) - 1) {
        ret += fmt.Sprintf("%v, ", e)

      } else {
        ret += fmt.Sprintf("%v", e)
      }
    }
    ret += fmt.Sprintf("]\n")
  }

  ret += fmt.Sprintf("\nGraph-Height: %v\n", gl.GraphHeight())
  ret += gl.PrintFlags()
  return ret
}

/* Functions */
func InitGraphList() (*GraphList) {

  var graph GraphList

  graph.Max = -1
  graph.Height = -1
  graph.Flags = 0
  graph.Edges = make(map[int][]Edge)
  graph.Vertices = make(map[int]*Vertex)
  return &graph
}
