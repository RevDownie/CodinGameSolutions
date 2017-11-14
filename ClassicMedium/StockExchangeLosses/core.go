// RULES
// Given a graph going from T0 to TN find the largest loss made

// INPUT
// 1. Number of data points
// 2. Line containing space separated data points 

// OUTPUT
// The loss as a negative number of 0 if no loss

// NOTES
// Max number of data points is 100,000
// Max value for a data point is 2^31 (32-bit)
// Values are always greater than zero

package main

import "fmt"
import "os"

/// CodinGame reserves stdout for returning answers and recommends
/// stderr for debug logging
///
func debugLog(msg interface{}) {
    fmt.Fprintln(os.Stderr, msg)
}

/// Runs the simulation
///
func run(dataPoints []int) {
    debugLog(dataPoints)

    runStartVal := 0
    maxLoss := 0

    for _, d := range dataPoints {
        if d >= runStartVal {
            //Start of a new run
            runStartVal = d
            debugLog(fmt.Sprintf("Run Start: %d", runStartVal))
        } else {
            currentLoss := d - runStartVal
            debugLog(fmt.Sprintf("Current Loss: %d", currentLoss))

            if currentLoss < maxLoss {
                maxLoss = currentLoss
                debugLog(fmt.Sprintf("Max Loss: %d", maxLoss))
            }
        }
    }

    // Send answer to CodinGame via stdout
    fmt.Println(maxLoss)
}

/// Used by CodinGame to feed us the input via stdin
///
func main() {
    // Number of data points
    var n int
    fmt.Scan(&n)

    // Data points
    dataPoints := make([]int, n)
    for i:=0; i<n; i++ {
        var dp int
        fmt.Scan(&dp)
        dataPoints[i] = dp
    }

    //Test01
    //dataPoints = []int{3,2,4,2,1,5} //-3 expected

    //Test02
    dataPoints = []int{5,3,4,2,3,1} //-4 expected

    run(dataPoints)
}