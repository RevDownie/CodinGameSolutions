// OBJECTIVE
// Find the bottom and right marked neighbours for marked nodes. We check right and down until we reach a node or the end of the grid

// INPUT
// int representing the width of the grid
// int representing the height of the grid
// Height number of lines containing a string with width characters where '.' is empty and '0' is marked.

// OUTPUT
// A line for each marked cell as a string in the format "cellX cellY rX rY bX bY"
// If there is no neighbour the coords should be -1 -1

#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include <memory>

//Toggle between local tests and CodinGame submission
#define RUN_LOCAL 1

const int k_numTests = 4;
const char k_mark = '0';

/// For stepping along the grid it is easier to split the index into
/// x and y coords
///
struct Coord
{
	int x, y;

	Coord(int in_x, int in_y)
	: x(in_x), y(in_y)
	{}
};

/// Data about the grid and accessors to the data
///
class Grid
{
public:
	Grid(std::unique_ptr<char[]> cells, int w, int h)
	: m_cells(std::move(cells)), m_width(w), m_height(h), m_area(w * h)
	{}

	int GetWidth() const { return m_width; }

	int GetHeight() const { return m_height; }

	int GetArea() const { return m_area; }

	char TryGetCellContent(const Coord& coord) const 
	{ 
		if(coord.x < 0 || coord.x >= m_width || coord.y < 0 || coord.y >= m_height)
			return -1;

		return m_cells[coord.y * m_width + coord.x];
	}

	char TryGetCellContent(int index) const 
	{ 
		if(index < 0 || index >= m_width * m_height)
			return -1;

		return m_cells[index];
	}

private:
	std::unique_ptr<char[]> m_cells;
	int m_width;
	int m_height;
	int m_area;
};

/// Read the grid from stdin into the grid data structure
///
Grid ReadGrid()
{
	int width;
	std::cin >> width;
	std::cin.ignore();

	int height;
	std::cin >> height;
	std::cin.ignore();

	int area = width * height;

	// Build the grid
	std::unique_ptr<char[]> cells(new char[area]);

	for(int y=0; y<height; ++y)
	{
		std::string row;
		std::getline(std::cin, row);

		for(int x=0; x<row.length(); ++x)
		{
			cells[y * width + x] = row[x];
		}
	}

	return Grid(std::move(cells), width, height);
}

/// Search in the given direction until we hit a marked node
/// or the end of the grid. If we hit the end of the grid
/// return -1, -1
///
Coord TryFindMarkedNeighbour(Coord coord, int dirX, int dirY, const Grid& grid)
{
	char content = -1;

	do
	{
		coord.x += dirX;
		coord.y += dirY;

		content = grid.TryGetCellContent(coord);

		if(content == k_mark)
		{
			return coord;
		}
	}
	while(content != -1);

	return Coord(-1, -1);
}

/// Run a single test reading from stdin
/// 
void Run()
{
	Grid grid = ReadGrid();

	bool success = true;

	//Find each marked node - find its neigbours and output
	for(int i=0; i<grid.GetArea(); ++i)
	{
		if(grid.TryGetCellContent(i) == k_mark)
		{
			Coord current(i % grid.GetWidth(), i / grid.GetWidth());
			Coord right(TryFindMarkedNeighbour(current, 1, 0, grid));
			Coord bottom(TryFindMarkedNeighbour(current, 0, 1, grid));

			std::ostringstream outputLine;
			outputLine << current.x << " " << current.y << " " << right.x << " " << right.y << " " << bottom.x << " " << bottom.y;

#if RUN_LOCAL
			//For the tests I've appended the expected output lines to the end of the test case
			std::string expectedOutputLine;
			std::getline(std::cin, expectedOutputLine);

			if(outputLine.str() != expectedOutputLine)
			{
				success = false;
			}
#endif

			//Output answer to CodinGame via stdout
			std::cout << outputLine.str() << std::endl;
		}
	}

#if RUN_LOCAL
			std::cout << (success ? "SUCCESS" : "FAIL") << std::endl << std::endl;
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