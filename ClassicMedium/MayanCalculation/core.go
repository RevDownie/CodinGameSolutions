// NUMERICAL SYSTEM
// Uses base 20
// Each numeral is an LxH pattern consisting of the character set .-o
// There are 20 numbers ranging from 0-19
// A number is a vertical column of numerals where each numeral is multiplied by 20 to the power of position and added to the next one (top to bottom) => (X x 20^2) + (Y x 20^1) + (Z x 20^0). The top most number is the most significant
// Numeral patterns change with each test

// INPUT
// 1. The width L and height H of a numeral
// 2. H lines of 20xL characters representing the possible numerals
// 3. The number of lines S1 for the first number
// 4. S1 lines representing the first number - each line has L characters
// 5. The number of lines S2 for the second number
// 6. S2 lines representing the second number - each line has L characters
// 7. The operation to perform between the 2 numbers (*, /, +, -)

// OUTPUT
// The result of the operation in mayan H lines per section with each line having length L (i.e. the number is split across multiple lines)

// NOTES
// L and H between 0 and 100 (excl)
// S between 0 and 1000 (excl)
// Mayan numbers will not exceed 2 ^ 63 (64-bit signed)

package main

import "fmt"
import "os"
import "math"

/// CodinGame reserves stdout for returning answers and recommends
/// stderr for debug logging
///
func debugLog(msg interface{}) {
    fmt.Fprintln(os.Stderr, msg)
}

/// Convert the given integer number to a sequence of mayan numeral patterns. where each patter simply using the index
/// Requires that the 20 mayan numerals are in ascending order 0 - 19
///
func base10to20(toConvert int64, mayanNumerals []string) string {

    //toConvert
    return mayanNumerals[toConvert]
}

/// Convert the given mayan number to an integer. The mayan number comprises several mayan numerals in base 20 with the first being the most
/// significant. The number is converted using the formula (X0 x 20^N) + (X1 x 20^1) + (XN x 20^0). Each mayan numeral pattern is converted
/// to an integer simply using the index so requires that the 20 mayan numerals are in ascending order 0 - 19
///
func base20to10(toConvert []string, mayanNumerals []string) int64 {
    var total int64 = 0

    for in, n := range mayanNumerals {
        for ic, c := range toConvert {
            if n == c {
                power := len(toConvert) - ic - 1
                total += int64(in) * int64(math.Pow(20, float64(power)))
            }
        }
    }

    return total
}

/// CodinGame feeds us a number as a series of numerals ordered vertically. We get the number
/// in lines. We convert those lines into individual numerals and return them with the most significant first
/// e.g [ABC,DEF,GHI] => [ABCDEFGHI]
///
func parseSplitBase20Number(lines []string, numeralWidth int) []string {
    output := make([]string, len(lines)/numeralWidth)
    for lineIndex, l := range lines {
        numeralIndex := (lineIndex / numeralWidth)
        output[numeralIndex] = output[numeralIndex] + l
    }

    return output
}

/// Runs the simulation by taking 2 mayan numerals and performing the given arithmetic operation on them using
/// the given numerals for base10 conversion
///
func run(numeralWidth int, numeralsLines []string, aLines []string, bLines []string, operator string) {

    // Split the lines along the numeral boundaries into individual Mayan numeral strings - lowest first
    var numerals [20]string
    for lineIdx:=0; lineIdx<len(numeralsLines); lineIdx++ {
        for charIdx:=0; charIdx<20; charIdx++ {
            numeralIndex := (charIdx / numeralWidth)
            numerals[numeralIndex] = numerals[numeralIndex] + string(numeralsLines[lineIdx][charIdx])
        }
    }

    // Split the operands into individual Mayan numerals - most significant first
    a20 := parseSplitBase20Number(aLines, numeralWidth)
    b20 := parseSplitBase20Number(bLines, numeralWidth)

    // Convert from Mayan to base 10 ints for math operations
    var a10 = base20to10(a20, numerals[:])
    var b10 = base20to10(b20, numerals[:])

    var result10 int64 = 0
    switch operator {
    case "+":
        result10 = a10 + b10
    case "-":
        result10 = a10 - b10
    case "*":
        result10 = a10 * b10
    case "/":
        result10 = a10 / b10
    }

    // Convert back to base 20 and then to the Mayan numerals 
    result20 := base10to20(result10, numerals[:])
    
    debugLog(fmt.Sprintf("%s %s %s => %d %s %d = %d => %s", a20, operator, b20, a10, operator, b10, result10, result20))

    // Send the answer to CodinGame split into lines
    line := 0
    for i:=0; i<len(numeralsLines); i++ {

        fmt.Println(result20[line:line+numeralWidth])
        line += numeralWidth
    }
}


func main() {
    // L and H are the width and height of each numeral pattern e.g. 
    // .oo.
    // o..o
    // .oo.
    // ....
    //
    // Has a L and H of 4
    // var L, H int
    // fmt.Scan(&L, &H)
    
    // for i := 0; i < H; i++ 
    // {
    //     var numeral string
    //     fmt.Scan(&numeral)
    // }
    // var S1 int
    // fmt.Scan(&S1)
    
    // for i := 0; i < S1; i++ 
    // {
    //     var num1Line string
    //     fmt.Scan(&num1Line)
    // }
    // var S2 int
    // fmt.Scan(&S2)
    
    // for i := 0; i < S2; i++ 
    // {
    //     var num2Line string
    //     fmt.Scan(&num2Line)
    // }
    // var operation string
    // fmt.Scan(&operation)
    
    
    // // fmt.Fprintln(os.Stderr, "Debug messages...")
    // fmt.Println("result")// Write answer to stdout

    a := []string{"o...", "....", "....", "...."}
    b := []string{"o...", "....", "....", "...."}
    numerals := []string{".oo.o...oo..ooo.oooo....o...oo..ooo.oooo....o...oo..ooo.oooo....o...oo..ooo.oooo", "o..o................____________________________________________________________", ".oo.....................................________________________________________", "............................................................____________________"}
    run(4, numerals, a, b, "+")
}