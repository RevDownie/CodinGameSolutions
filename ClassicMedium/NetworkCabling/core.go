// RULES
// Given the x,y coordinates of buildings and the rule that a single cable must run from the x location of the most westerly to the most easterly house
// and a cable must run from the single main cable to each x,y location - calculate the total length of cable used

// INPUT
// 1. N - Number of x, y coords
// 2. N lines containing space separated x, y coords

// OUTPUT
// L - A single number representing the total length

// NOTES
// 0 < N <= 100,000
// Max value for L is 2^63 (i.e. 64-bit)
// -2^30 <= x/y <= 2^30
// Lines are only horizontal or vertical and not diagonal

// TO CONSIDER
// The main cable can run through buildings and remove the need for cables linking those buildings to the main line
// The length of the main cable is entirely determined by the distance between x0 and xN. The y position of the main line is the real problem
// The y pos should be an average of the y-coords (median? - it would run through at least one building) (do we prioritise running through multiple buildings)

package main

import "fmt"
import "os"
import "sort"

type Coord struct {
    x, y int64
}

/// CodinGame reserves stdout for returning answers and recommends
/// stderr for debug logging
///
func debugLog(msg interface{}) {
    fmt.Fprintln(os.Stderr, msg)
}

/// Runs the simulation. Coords as a flat list x1,y1, x2,y2, etc
///
func run(coordsFlat []int64) {

    // Convert flat list to paired coords
    coords := make([]Coord, len(coordsFlat)/2)
    for s, d := 0, 0; s<len(coordsFlat); s, d = s+2, d+1 {
        coords[d] = Coord{coordsFlat[s], coordsFlat[s+1]}
    }

    debugLog(coords)

    // Sort west to east to find length of main line
    sort.Slice(coords, func(a, b int) bool {
        return coords[a].x < coords[b].x
    })

    mainLineLength := coords[len(coords)-1].x - coords[0].x
    debugLog(fmt.Sprintf("Main Line Length: %d", mainLineLength))

    // Sort south to north to find the y median
    sort.Slice(coords, func(a, b int) bool {
        return coords[a].y < coords[b].y
    })

    mainLineY := coords[len(coords)/2].y
    debugLog(fmt.Sprintf("Main Line Y: %d", mainLineY))

    // Find the total length from each position to the main line (vertically)
    totalLineLength := mainLineLength

    for _,c := range coords {
        dy := c.y - mainLineY
        if dy < 0 {
            dy = -dy
        }
        totalLineLength += dy
    }

    debugLog(fmt.Sprintf("Total Line Length: %d", totalLineLength))

    // Send answer to CodinGame via stdout
    fmt.Println(totalLineLength)
}

/// Used by CodinGame to feed us the input via stdin
///
func main() {
    // Number of coords
    var n int
    fmt.Scan(&n)

    // Doubled as we store x and y individually
    n *= 2

    // Coords as a flat list x1,y1, x2,y2, etc
    coords := make([]int64, n)
    for i:=0; i<n; i++ {
        var c int64
        fmt.Scan(&c)
        coords[i] = c
    }

    // Test01 - Expected 4
    //coords = []int64{0,0,1,1,2,2}

    // Test02 - Expected 4
    //coords = []int64{1,2,0,0,2,2}

    // Test05 - Expected 18
    //coords = []int64{-5,-3,-9,2,3,-4}

    run(coords)
}