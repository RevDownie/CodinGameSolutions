# OBJECTIVE
# Find the central node of a tree that has the shortest maximum distance to all the leafs of the tree and return that shortest distance

# INPUT
# N - Number of links
# Next N lines are 2 space separated ints specifying a link between 2 nodes.

# OUTPUT
# Single int representing the minimal amount of steps required to proagate to every node on the graph

# NOTES
# No cyclic links e.g. a => b => c => a
# A propagation step is counted as traversing to a new depth in the tree and not a step for each link (i.e. propagation is concurrent)
# Graphs can be big (up to 150,000 nodes)
# Looking at the test data the nodes are not always numbered from zero or contiguously

# TO CONSIDER
# We should always take the max distance from our epicentre to the tips as the answer as this will represent the minimum time required
# Do we calculate the max distance for each node and track the shortest or is there an algorithm that will allow us to pinpoint the epicentre node immediately?
# The blurb for the problem specifies memoization which prob means that a straightforward search approach will be too slow and we need to cache the max dists. No mention as to what the time limit is - forum posts suggests that timeout is ~5000ms
# We don't have to search the whole tree everytime we just have to search until we exceed the current max
# We could always thread the search and calculate node distances in parallel
# Is there a heuristic that will give us the best node to start with that will give us more early outs during traversal?

import sys
from collections import defaultdict

# Logging to stdout is reserved for returning the answer. To debug log
# we must log to stderr
#
def debug_log(msg):
	sys.stderr.write(str(msg) + '\n')

# Perform a depth-first search from the given node and store the distance from each node to the furthest leaf node. Will stop if
# the max distance matches or exceeds the given threshold
# 
def find_longest_distance_to_leaf(start_node, links, max_dist_thresh):
	to_visit = list()
	visited = set()
	depth = list()

	to_visit.append(start_node)
	visited.add(start_node)
	depth.append(0)

	max_dist = 0

	while len(to_visit) > 0:
		node = to_visit.pop()
		node_links = links[node]
		current_depth = depth.pop()

		max_dist = max(max_dist, current_depth)
		if max_dist >= max_dist_thresh:
			return max_dist_thresh
		# debug_log(str.format("Visiting node {} dist is {}", node, current_depth))

		for linked_node in node_links:
			if linked_node not in visited:
				visited.add(linked_node)
				to_visit.append(linked_node)
				depth.append(current_depth + 1)
	
	# debug_log("")
	return max_dist


# Links are adjacency lists where each node key of the dictionary contains the other linked node indices
# We traverse the graph and calculate the max distance from each node to the the edge of the graph
# and then find the shortest of those. 
# 
# We optimise by 
# 	* Stop traversing if we exceed the current max
# 
def run(links, expected_output = None):
	min_max_dist = sys.maxint
	for n in links.keys():
		min_max_dist = find_longest_distance_to_leaf(n, links, min_max_dist)

	if expected_output != None:
		print("SUCCESS" if min_max_dist == expected_output else str.format("FAIL expected {}", expected_output))

	# Output the answer to CodinGame
	print(min_max_dist)

# Used by the tests to convert lines of linked nodes into adjacency lists
# 
def parse_links(unparsed_links):
	links = defaultdict(list)
	for nodes in unparsed_links:
		n1, n2 = [int(node) for node in nodes.split()]
		links[n1].append(n2)
		links[n2].append(n1)

	return links

# Used when running in CodinGame to read input from stdinput
#
if __name__ == "__main__":
	num_links = int(raw_input())

	links = defaultdict(list)
	for i in xrange(num_links):
		n1, n2 = [int(j) for j in raw_input().split()]
		links[n1].append(n2)
		links[n2].append(n1)

	run(links)
