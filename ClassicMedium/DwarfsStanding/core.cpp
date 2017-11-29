// DFS, Tree (directed), Recursion

// OBJECTIVE
// Find the longest chain of a tree

// NOTES
// Tree not graph
// No cyclic or bidirectional dependencies
// Number of edges is < 10,000

// INPUT
// N int representing the number of links/edges
// N lines consisting of 2, space separated, ints describing a link between 2 nodes
// Node Ids are less that 10,000

// OUTPUT
// Single int describing the longest chain of the tree

// TO CONSIDER
// Just use a depth first search and track the longest chain.
// The input doesn't always appear to be in order - i.e. the root isn't first - so we have to identify the root (a root has no parents)
// Node ids are not contiguous
// As Node ids are constrained to 10,000 would it be quicker using an array rather than a map
// Usually I would create the graph using pointers but as the input comes in as numbers might be better just using the numbers as look-up
// The input of test03 suggests that there can be more than one tree/root (nods 8 and 9 don't relate to anything)

#include <iostream>
#include <fstream>
#include <vector>
#include <array>
#include <algorithm>
#include <sstream>

//Toggle between local tests and CodinGame submission
#define RUN_LOCAL 1

const int k_maxNumNodes = 10000;
const int k_numTests = 4;

///
struct NodeData
{
	bool m_isRoot = true;
	std::vector<int> m_children;
};

/// Starting at the given node, recurse through the tree depth first until we have exhausted all
/// chains. We keep track of the longest chain. The tree is represented as an adjacency list for
/// each node index.
///
/// NOTE: The chain counts is the number of nodes - not links
///
int FindLongestChain(int nodeIndex, int longestChainLen, const std::array<NodeData, k_maxNumNodes>& tree)
{
	int prevLongestChainLen = longestChainLen;
	for(int childIdx : tree[nodeIndex].m_children)
	{
		int chainLen = FindLongestChain(childIdx, prevLongestChainLen + 1, tree);
		longestChainLen = std::max(longestChainLen, chainLen);
	}

	return longestChainLen;
}

/// Run a single test reading from stdin and calculating the longest chain
/// 
void Run()
{
	std::array<NodeData, k_maxNumNodes> tree;

	int numEdges;
	std::cin >> numEdges;
	std::cin.ignore();

	for(int i=0; i<numEdges; ++i)
	{
		int n1, n2;
		std::cin >> n1 >> n2;
		std::cin.ignore();

		tree[n1].m_children.push_back(n2);
		//n2 is the child of another node and cannot be a root
		tree[n2].m_isRoot = false;
	}

	//Find all the roots and then find the longest chain in each tree.
	//Pick the longest of the longest
	int longestChain = 0;
	for(int i=0; i<tree.size(); ++i)
	{
		if(tree[i].m_isRoot == true)
		{
			int chain = FindLongestChain(i, 1, tree);
			longestChain = std::max(chain, longestChain);
		}
	}

	//Output answer to CodinGame via stdout
	std::cout << longestChain << std::endl;

#if RUN_LOCAL
	//For the tests I've appended the expected output to the end of the test case
	int expectedOutput;
	std::cin >> expectedOutput;
	std::cin.ignore();

	//Debug content goes to error as stdout is reserved for the answer
	std::cerr << (longestChain == expectedOutput ? "SUCCESS" : "FAIL") << std::endl;
#endif
}


/// Read input from CodinGame (or tests) via stdin
///
int main()
{
#if RUN_LOCAL
	for(int i=1; i<=k_numTests; ++i)
	{
		std::ostringstream testPath;
		testPath << "Tests/Test" << i << ".txt";
		std::ifstream test(testPath.str());
		// Redirect cin to read from our test files
		std::cin.rdbuf(test.rdbuf());
		Run();
	}
#else
	Run();
#endif

	return 0;
}