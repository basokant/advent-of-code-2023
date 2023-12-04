package util

type Graph[T any] struct {
	edges    map[int][]Vertex[T]
	vertices []Vertex[T]
	length   int
}

type Vertex[T any] struct {
	data T
	id   int
}

var nextId int

func (g *Graph[T]) addVertex(data T) {
	vertex := Vertex[T]{
		data: data,
		id:   nextId,
	}
	g.vertices = append(g.vertices, vertex)
	nextId += 1
}

func (g *Graph[T]) addEdge(a Vertex[T], b Vertex[T]) error {
	g.edges[a.id] = append(g.edges[a.id], b)
	g.edges[b.id] = append(g.edges[b.id], a)

	return nil
}
