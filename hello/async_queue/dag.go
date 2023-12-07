package async_queue

// DAG 有向无环图
type DAG struct {
	Vertexes map[string]*Vertex
}

// Vertex 顶点
type Vertex struct {
	Key      string // 唯一标识符
	Value    interface{}
	Parents  map[string]*Vertex
	Children map[string]*Vertex
}

func NewDAG() *DAG {
	return &DAG{
		Vertexes: make(map[string]*Vertex),
	}
}

// AddVertex 添加顶点
func (dag *DAG) AddVertex(key string, value interface{}) {
	dag.Vertexes[key] = &Vertex{
		Key:      key,
		Value:    value,
		Parents:  make(map[string]*Vertex),
		Children: make(map[string]*Vertex),
	}
}

// AddEdge 添加边
func (dag *DAG) AddEdge(fromKey, toKey string) {
	from, ok1 := dag.Vertexes[fromKey]
	to, ok2 := dag.Vertexes[toKey]
	if ok1 && ok2 {
		from.Children[toKey] = to
		to.Parents[fromKey] = from
	}
}

// CanFinish 判断是否为有向无环图
func (dag *DAG) CanFinish() bool {
	return len(dag.BFS()) == len(dag.Vertexes)
}

// BFS 广度优先搜索，按访问顺序返回顶点列表
func (dag *DAG) BFS() []*Vertex {
	inDegrees := make(map[string]int)
	queue := make([]*Vertex, 0)
	result := make([]*Vertex, 0)

	for key, vertex := range dag.Vertexes {
		inDegrees[key] = len(vertex.Parents)
		if inDegrees[key] == 0 {
			queue = append(queue, vertex)
		}
	}

	for len(queue) > 0 {
		vertex := queue[0]
		queue = queue[1:]
		result = append(result, vertex)
		for _, ch := range vertex.Children {
			inDegrees[ch.Key]--
			if inDegrees[ch.Key] == 0 {
				queue = append(queue, ch)
			}
		}
	}

	return result
}
