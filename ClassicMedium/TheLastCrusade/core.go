// RULES
// Find the path from top to bottom of a grid of tiles where each tile has a specified entry and exit points.
// Gravity pulls the character downwards so the bottom exit is always selected where possible

// INPUT - GAME
// 2 digits, space separated, that describe the width and height of the map grid
// H lines representing the tilesets for a row of the map grid. Each tileset id is space separated
// 1 integer specifying the x coord of the exit cell (the y coord is always the bottom row) - Not required in Ep. 1

// INPUT - TURN
// 2 digits, space separated, that describe the current position of the character
// Single word (TOP, LEFT, RIGHT) that describes the entry

// OUTPUT - TURN
// A single line with 2 space separated ints for the x, y coords the character should move to

// NOTES
// Ep 1 - There is always a path and it is always continuous (no need for open and closed sets)
// This is a turn based game so there are multiple inputs and outputs
// Response time must be <= 150ms

// TILE SETS
// 0  - Blocker
// 1  - Enter L,T,R - Exit B,B,B
// 2  - Enter L,R - Exit R,L
// 3  - Enter T - Exit B
// 4  - Enter T,R - Exit L,B
// 5  - Enter L,T - Exit B,R
// 6  - Enter L,T,R - Exit R,N,L
// 7  - Enter T,R - Exit B,B
// 8  - Enter L,R - Exit B,B
// 9  - Enter L,T - Exit B,B
// 10 - Enter T - Exit L
// 11 - Enter T - Exit R
// 12 - Enter R - Exit B
// 13 - Enter L - Exit B

// TO CONSIDER
// Essentially we are mapping a tile id and a incoming direction to an outgoing direction. 
// Can generate a single unique key using a pairing function allowing us to do a single lookup

package main

import "fmt"
import "os"

const None = 0
const Left = 1
const Top = 2
const Right = 3
const Bottom = 4

/// Holds the map data
///
type Grid struct {
    width, height int
    cells []int
}

/// Generates a map thats key is a pairing of tileId and entry direction and thats value
/// is the exit direction
///
func generateTileMappings() map[int]int {
	tilemap := make(map[int]int)

	// For any values that don't exist 0 (None) will be returned

	tilemap[cantorPair(1, Left)] = Bottom
	tilemap[cantorPair(1, Top)] = Bottom
	tilemap[cantorPair(1, Right)] = Bottom

	tilemap[cantorPair(2, Left)] = Right
	tilemap[cantorPair(2, Right)] = Left

	tilemap[cantorPair(3, Top)] = Bottom

	tilemap[cantorPair(4, Top)] = Left
	tilemap[cantorPair(4, Right)] = Bottom

	tilemap[cantorPair(5, Left)] = Bottom
	tilemap[cantorPair(5, Top)] = Right

	tilemap[cantorPair(6, Left)] = Right
	tilemap[cantorPair(6, Top)] = None
	tilemap[cantorPair(6, Right)] = Left

	tilemap[cantorPair(7, Top)] = Bottom
	tilemap[cantorPair(7, Right)] = Bottom

	tilemap[cantorPair(8, Left)] = Bottom
	tilemap[cantorPair(8, Right)] = Bottom

	tilemap[cantorPair(9, Left)] = Bottom
	tilemap[cantorPair(9, Top)] = Bottom

	tilemap[cantorPair(10, Top)] = Left

	tilemap[cantorPair(11, Top)] = Right

	tilemap[cantorPair(12, Right)] = Bottom

	tilemap[cantorPair(13, Left)] = Bottom

	return tilemap
}

/// CodinGame reserves stdout for returning answers and recommends
/// stderr for debug logging
///
func debugLog(msg interface{}) {
    fmt.Fprintln(os.Stderr, msg)
}

/// Generate a unique number from the 2 given numbers
///
func cantorPair(a int, b int) int {
	return ((a + b) * (a + b + 1)) / 2 + b
}

/// Convert the direction string to a unique number
///
func parseDirection(dir string) int {
	switch dir {
		case "LEFT":
			return Left
		case "TOP":
			return Top
		case "RIGHT":
			return Right
	}

	// Blocked
	return None
}

/// Generate new x and y based on direction of movement
///
func moveInDirection(x int, y int, dirId int) (nx int, ny int) {
	switch dirId {
		case Left:
			nx = x - 1
			ny = y
		case Right:
			nx = x + 1
			ny = y
		case Bottom:
			nx = x
			ny = y + 1
		default:
			nx = x
			ny = y
	}

	return
}

/// Uses the current location, entry direction and tiles to decide what direction to head in and ultimately what the
/// next location is
///
func decide_next_location(x int, y int, dir string, grid *Grid, tilemappings map[int]int) (nx int, ny int, tileId int) {
	tileId = grid.cells[grid.width * y + x]
	incomingDirId := parseDirection(dir)
	key := cantorPair(tileId, incomingDirId)
	outgoingDirId := tilemappings[key]
	nx, ny = moveInDirection(x, y, outgoingDirId)
	return
}

/// Used by CodinGame to feed us the input via stdin
///
func main() {
	// Map of tileid and incoming direction to outgoing direction
	tilemappings := generateTileMappings()

	// Read width and height of the grid
    var w, h int
    fmt.Scan(&w, &h)
    debugLog(fmt.Sprintf("w %d, h %d", w, h))

    // Build the grid of tile ids
    cells := make([]int, w * h)
    for y:=0; y<h; y++ {
    	for x:=0; x<w; x++ {
    		var tileId int
    		fmt.Scan(&tileId)
    		cells[w * y + x] = tileId
    	}
    }

    //Not needed in this episode
    var exitX int
    fmt.Scan(&exitX)
    debugLog(fmt.Sprintf("Exit x %d", exitX))

    grid := Grid{w, h, cells}

    // Each turn we read the position and entry direction and output the next location
    for turn:=0; turn<100; turn++ { //Should be a forever loop but want my tests to terminate
    	var x, y int
    	fmt.Scan(&x, &y)

    	var dir string
    	fmt.Scan(&dir)

		nextX, nextY, tileId := decide_next_location(x, y, dir, &grid, tilemappings)
		debugLog(fmt.Sprintf("Sx %d, Sy %d. Dir %s. TileId %d. Nx %d, Ny %d", x, y, dir, tileId, nextX, nextY))

		// Send answer to CodinGame
		fmt.Printf("%d %d\n", nextX, nextY)
    }
}