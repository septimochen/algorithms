package graph

import "sort"

//need add edges to scc
func SCC(g Graph) (scc Graph) {
	scc = CreateGraphByType(g)
	//DFS and get forest
	dfsGraph, gT := DFS(g, nil), g.Transpose()
	//DFS of transpose in order of decreasing finish time
	dfsVertices := dfsGraph["dfsForest"].AllVertices()
	sort.Slice(dfsVertices, func(i, j int) bool {
		return dfsVertices[i].(*DFSElement).F > dfsVertices[j].(*DFSElement).F
	})
	dfsGraphOfT := DFS(gT, func(vertices []interface{}) {
		for i, v := range dfsVertices {
			vertices[i] = v.(*DFSElement).V
		}
	})
	//shrink all vertices, according to the root(disjoint-set)
	dfsForest := GetDFSComponent(dfsGraphOfT["dfsForest"])
	components := make(map[*DFSElement]Graph)
	for i := range dfsForest {
		components[i] = CreateGraphByType(g)
		for _,v := range dfsForest[i].AllVertices() {
			components[i].AddVertex(v.(*DFSElement).V)
			for _, e := range dfsForest[i].AllConnectedVertices(v) {
				components[i].AddEdge(Edge{e.(*DFSElement).V, v.(*DFSElement).V})
			}
			for _, e := range dfsGraphOfT["dfsForwardEdges"].AllConnectedVertices(v) {
				if e.(*DFSElement).FindRoot() == i {
					components[i].AddEdge(Edge{e.(*DFSElement).V, v.(*DFSElement).V})
				}
			}
			for _, e := range dfsGraphOfT["dfsBackEdges"].AllConnectedVertices(v) {
				if e.(*DFSElement).FindRoot() == i {
					components[i].AddEdge(Edge{e.(*DFSElement).V, v.(*DFSElement).V})
				}
			}
		}
	}
	//keep all cross edges
	for _,e := range dfsGraphOfT["dfsCrossEdges"].AllEdges() {
		scc.AddEdge(Edge{components[e.End.(*DFSElement).FindRoot()], components[e.Start.(*DFSElement).FindRoot()]})
	}

	return
}