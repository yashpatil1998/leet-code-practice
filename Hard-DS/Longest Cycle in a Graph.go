// https://leetcode.com/problems/longest-cycle-in-a-graph/description/

// added comments using ai

package HardDs

// longestCycle finds the length of the longest cycle in a directed graph
// represented by an edges array. edges[i] = j means a directed edge from node i to node j.
// If edges[i] = -1, node i has no outgoing edge.
func longestCycle(edges []int) int {
	// Build the directed graph from the edges array.
	// The graph is represented using an adjacency list where the key is the node
	// and the value is a slice of its outgoing neighbors.
	graph := make(map[int][]int)
	// We need to know the total number of nodes, which is the length of the edges array.
	numNodes := len(edges)

	// Initialize the graph map keys for all potential nodes (0 to numNodes-1),
	// even if they have no outgoing edges (-1). This ensures our main loop
	// iterates through all possible nodes.
	for i := 0; i < numNodes; i++ {
		graph[i] = []int{} // Initialize with an empty slice of neighbors
	}

	// Populate the adjacency list based on the edges array.
	for i, e := range edges {
		if e != -1 { // If there is an outgoing edge
			// Add a directed edge from node i to node e.
			graph[i] = append(graph[i], e)
			// Ensure the target node 'e' also exists as a key in the map,
			// even if it has no outgoing edges itself.
			if _, ok := graph[e]; !ok {
				graph[e] = []int{}
			}
		}
	}

	// visitedAndDepth map serves a dual purpose for DFS:
	// - If visitedAndDepth[node] == 0: Node is unvisited.
	// - If visitedAndDepth[node] > 0: Node is currently being visited in the *current* DFS path,
	//                                and the value is the step count/depth at which it was
	//                                first entered in this path (starting from 1).
	// - If visitedAndDepth[node] == -1: Node has been fully explored (its subtree has been
	//                                 visited, and no cycle was found *starting from a node
	//                                 in its subtree* that goes back into that subtree).
	//                                 This prevents redundant re-visiting of fully explored components.
	visitedAndDepth := make(map[int]int)

	// Initialize visitedAndDepth for all nodes.
	for i := 0; i < numNodes; i++ {
		visitedAndDepth[i] = 0 // Initially, all nodes are unvisited.
	}

	// maxCycleLength stores the length of the longest cycle found across all DFS traversals.
	maxCycleLength := -1

	// Iterate through all nodes in the graph.
	// This is necessary to handle disconnected components.
	for node := range graph {
		// If the node hasn't been visited yet (state 0), start a new DFS traversal from it.
		if visitedAndDepth[node] == 0 {
			// Start DFS from this node at step 1. Pass a pointer to maxCycleLength
			// so the recursive calls can update the global maximum.
			dfs(node, graph, visitedAndDepth, 1, &maxCycleLength)
		}
		// If visitedAndDepth[node] is > 0, it means this node is currently in the recursion
		// stack of a DFS started from an earlier node. We don't need to start a new DFS.
		// If visitedAndDepth[node] is -1, this node and its reachable subtree have been
		// fully explored in a previous DFS. We skip it.
	}

	// Return the maximum cycle length found. If no cycle was found, it remains -1.
	return maxCycleLength
}

// dfs performs a Depth-First Search traversal to find cycles and their lengths.
// - currentNode: The node currently being visited.
// - graph: The adjacency list representation of the directed graph.
// - visitedAndDepth: Map to track node states (0: unvisited, >0: visiting with step, -1: fully visited).
// - currentStep: The current step count/depth in the ongoing DFS path.
// - maxCycleLength: A pointer to the variable storing the maximum cycle length found so far.
func dfs(currentNode int, graph map[int][]int, visitedAndDepth map[int]int, currentStep int, maxCycleLength *int) {
	// Base Case 1: If the node has been fully explored (-1), return immediately.
	// This prevents re-processing nodes in components already checked.
	if visitedAndDepth[currentNode] == -1 {
		return
	}

	// Base Case 2: If the node is currently being visited (> 0), a cycle is detected.
	// The value in visitedAndDepth[currentNode] is the step at which we first
	// entered this node in the current path.
	if visitedAndDepth[currentNode] > 0 {
		// Calculate the length of the cycle.
		// currentStep is the step when we *re-entered* the node.
		// visitedAndDepth[currentNode] is the step when we *first entered* the node.
		cycleLength := currentStep - visitedAndDepth[currentNode]

		// Update the maximum cycle length found so far if the current cycle is longer.
		if cycleLength > *maxCycleLength {
			*maxCycleLength = cycleLength
		}
		// Return after finding and processing the cycle.
		return
	}

	// Mark the current node as being visited in the current path by storing the current step.
	visitedAndDepth[currentNode] = currentStep

	// Explore the neighbor(s) of the current node.
	// In this problem's structure (edges[i] = j), each node has at most one outgoing edge.
	if neighbors, ok := graph[currentNode]; ok && len(neighbors) > 0 {
		neighbor := neighbors[0] // Get the single outgoing neighbor

		// If the neighbor is a valid node (not -1).
		if neighbor != -1 {
			// Recursively call DFS on the neighbor, incrementing the step count.
			dfs(neighbor, graph, visitedAndDepth, currentStep+1, maxCycleLength)
		}
	}

	// After exploring the path starting from currentNode (either reached a visited node,
	// a -1 edge, or the end of a path), mark the currentNode as fully explored (-1).
	// This signifies that the subtree/path starting from this node has been processed
	// and won't lead to a new cycle by re-entering it from a different path.
	visitedAndDepth[currentNode] = -1
}
