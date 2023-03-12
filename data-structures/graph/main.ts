interface GraphNode<T> {
  value: T;
  neighbors: Set<GraphNode<T>>;
}

interface Graph<T> {
  vertices: GraphNode<T>[]; // set of vertices (nodes)
  hasNode(node: GraphNode<T>): boolean;
  addVertex(value: T): void; // Add a vertex to the graph
  addEdge(vertex1: GraphNode<T>, vertex2: GraphNode<T>): void; // Add an edge between two vertices
  removeVertex(value: T): void; // Remove a vertex from the graph
  removeEdge(vertex1: GraphNode<T>, vertex2: GraphNode<T>): void; // Remove an edge between two vertices
  getNeighbors(value: T): Set<T>; // Get the set of neighboring vertices for a given vertex
  size(): number; // Get the number of vertices in the graph
}
