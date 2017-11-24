// OBJECTIVE
// Prevent the agent from exiting the cluster through the gateway by severing links between nodes

// INPUT - GAME
// N - Total number of nodes
// L - The number of links between nodes
// E - The number of exit gateways
// Next L lines - 2 ints indicating a link between those indexed nodes
// Next E lines - 1 int representing the index of the gateway

// INPUT - TURN
// 1 int which is the node index of the agent

// OUTPUT - TURN
// A single line with 2 space separated ints describing the nodes whose joining link should be cut

// NOTES
// Not all nodes start with links
// Reponse time must be <= 150ms
// Nodes cannot be connected to more than one gateway

// TO CONSIDER
// We can create a graph of all possible paths from the agent to the gateways and cut the link nearest the agent on the shortest route to a gateway
// Should use a breadth-first search as we will come across the shortest path first
// What's the best way to track the path taken? Just to find the shortest path first and then run again with the target and start swapped?

package main

import "fmt"
import "os"
import "container/list"


/// CodinGame reserves stdout for returning answers and recommends
/// stderr for debug logging
///
func debugLog(msg interface{}) {
    fmt.Fprintln(os.Stderr, msg)
}

///
func contains(vals []int, val int) bool {
	for _,v := range vals {
		if v == val {
			return true
		}
	}

	return false
}

///
func indexOf(vals []int, val int) int {
	for i,v := range vals {
		if v == val {
			return i
		}
	}

	return -1
}

/// Given the starting node and the links between each node search the tree breadth-first
/// until we come across the first gateway. Return the node indices of the path taken to reach that gateway
/// Returns nil if no exit found
///
func findShortestPathToGateway(startNodeIdx int, links [][]int, exits []int) *list.List {
	toVisitQueue := list.New()
	visitedFrom := make([]int, len(links))
	for i := range visitedFrom {
		visitedFrom[i] = -1
	}

	visitedFrom[startNodeIdx] = startNodeIdx
	toVisitQueue.PushFront(startNodeIdx)

    for {
    	queueEl := toVisitQueue.Back()
    	if queueEl == nil {
    		//Exhausted all nodes and didn't find anything
    		break
    	}

    	nodeIdx := queueEl.Value.(int)
    	toVisitQueue.Remove(queueEl)

	    for _,n := range links[nodeIdx] {
	    	//When we sever a link we set it to -1
	    	if n >= 0 {
	    		if contains(exits, n) == true {
	    			//Found an exit build the path by going back through the parent links
	    			path := list.New()
	    			path.PushFront(n)
	    			
	    			for {
	    				path.PushFront(nodeIdx)
	    				nodeIdx = visitedFrom[nodeIdx]
	    				if nodeIdx == visitedFrom[nodeIdx] {
	    					break
						}
	    			}
	    			return path
	    		}

	    		if visitedFrom[n] < 0 {
		    		visitedFrom[n] = nodeIdx
			    	toVisitQueue.PushFront(n)
			    }
			}
    	}
	}

	// No path
	return nil
}

/// Print the contents of the list each on a new line
///
func debugPrintList(l *list.List) {
	el := l.Front()
	for {
		if el == nil {
			break
		}
		debugLog(el.Value)

		el = el.Next()
	}
}

/// Used by CodinGame to feed us the input via stdin
///
func main() {

	//Read the initial immutable game data
	var numNodes, numLinks, numExits int
	fmt.Scan(&numNodes, &numLinks, &numExits)

	//Each node index contains a list of all nodes it links to
    links := make([][]int, numNodes)
    exits := make([]int, numExits)

	//Read the links between node indexes
	for i:=0; i<numLinks; i++ {
		var n1, n2 int;
		fmt.Scan(&n1, &n2)

		links[n1] = append(links[n1], n2)
		links[n2] = append(links[n2], n1)
	}

	//Read the node index of each exit gateway
	for i:=0; i<numExits; i++ {
		var n int;
		fmt.Scan(&n)

		exits[i] = n
	}

	//Each turn we read the node index of the agent and output the link to sever in order to trap it
    for {
    	var agentNodeIndex int;
		fmt.Scan(&agentNodeIndex)

	    //Find the path to the nearest gateway exit from the agent location and sever the first link on that path
	    //We sever the first link as there is an additional goal to limit the amount of moves that the agent can make
    	path := findShortestPathToGateway(agentNodeIndex, links, exits)

    	debugLog("Shortest Path:")
    	debugPrintList(path)

    	firstEl := path.Front()
    	n1 := firstEl.Value.(int)
    	n2 := firstEl.Next().Value.(int)
    	debugLog(fmt.Sprintf("Cutting link between %d => %d", n1, n2))

    	//Don't forget to update our local state to reflect the new graph
    	n2Idx := indexOf(links[n1], n2)
    	links[n1][n2Idx] = -1

    	//Tell CodinGame what link we severed
    	fmt.Printf("%d %d\n", n1, n2)
    }
}