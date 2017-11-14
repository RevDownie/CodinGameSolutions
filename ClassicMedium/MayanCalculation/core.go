// NUMERICAL SYSTEM
// Uses base 20
// Each numeral is an LxH pattern consisting of the character set ._o
// There are 20 numbers ranging from 0-19
// A number is a vertical column of numerals where the top most number is the most significant (i.e. the last numeral is the unit column, the one to the left is the 20 column, etc)
// Numeral patterns change with each test

// INPUT
// 1. The width L and height H of a numeral
// 2. H lines of 20xL characters representing the possible numerals
// 3. The number of lines S1 for the first number
// 4. S1 lines representing the first number - each line has L characters
// 5. The number of lines S2 for the second number
// 6. S2 lines representing the second number - each line has L characters
// 7. The operation to perform between the 2 numbers (*, /, +, _)

// OUTPUT
// The result of the operation in mayan H lines per section with each line having length L (i.e. the number is split across multiple lines)

// NOTES
// L and H between 0 and 100 (excl)
// S between 0 and 1000 (excl)
// Numbers will not exceed 2 ^ 63 (64-bit required)

package main

import "fmt"
import "os"
import "math"
import "errors"

/// CodinGame reserves stdout for returning answers and recommends
/// stderr for debug logging
///
func debugLog(msg interface{}) {
    fmt.Fprintln(os.Stderr, msg)
}

/// Convert the given decimal integer number to a sequence of mayan numeral patterns with the first numeral being the most significant
/// Requires that the 20 mayan numerals are in ascending order 0 - 19
///
func base10to20(toConvert int64, mayanNumerals []string) []string {

    //Handling zero as a special case as log10 0 is undefined but can also
    //be used to speed up any of the single unit cases
    if toConvert < 20 {
        return []string{mayanNumerals[toConvert]}
    }

    numDigits := int(math.Floor(math.Log10(math.Abs(float64(toConvert))))) + 1

    output := make([]string, numDigits)

    //We need to remove leading zeros. This will track where our first number is
    startIndex := -1

    //For example the value 21 is 1 unit and 1 20 so the mayan value returned will be 1,1
    //In order to carry the remainder we need to start from the most significant column
    remaining := toConvert
    for i:=0; i<numDigits; i++ {
        denom := int64(math.Pow(20, float64(numDigits - i - 1)))
        numInCol := remaining/denom
        output[i] = mayanNumerals[numInCol]
        if numInCol > 0 && startIndex < 0 {
            startIndex = i
        }
        remaining = remaining % denom
    }

    //Trim the start
    return output[startIndex:]
}

/// Convert the given mayan number to a decimal integer. The mayan number comprises several mayan numerals in base 20 with the first being the most
/// significant. The number is converted using the formula (X0 x 20^N) + (X1 x 20^1) + (XN x 20^0). Each mayan numeral pattern is converted
/// to an decimal value simply using the index so requires that the 20 mayan numerals are in ascending order 0 - 19
///
func base20to10(toConvert []string, mayanNumerals []string) (total int64, err error) {
    total = 0
    err = nil

    for ic, c := range toConvert {
        found := false
        for in, n := range mayanNumerals {
            if n == c {
                power := len(toConvert) - ic - 1
                total += int64(in) * int64(math.Pow(20, float64(power)))
                found = true
                break
            }
        }

        if found == false {
            err = errors.New("Cannot find numeral")
            return
        }
    }

    return
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
        for charIdx:=0; charIdx<len(numeralsLines[lineIdx]); charIdx++ {
            numeralIndex := (charIdx / numeralWidth)
            numerals[numeralIndex] = numerals[numeralIndex] + string(numeralsLines[lineIdx][charIdx])
        }
    }

    // Split the operand numbers into individual Mayan numerals - most significant first
    a20 := parseSplitBase20Number(aLines, numeralWidth)
    b20 := parseSplitBase20Number(bLines, numeralWidth)

    // Convert from Mayan to base 10 ints for math operations
    a10, err := base20to10(a20, numerals[:])
    if err != nil {
        panic(err)
    }
    b10, err := base20to10(b20, numerals[:])
    if err != nil {
        panic(err)
    }

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
    
    debugLog(numerals)
    debugLog(fmt.Sprintf("%s %s %s => %d %s %d = %d => %s", a20, operator, b20, a10, operator, b10, result10, result20))

    // Send the answer to CodinGame split into lines
    for i:=0; i<len(result20); i++ {
        line := 0
        for j:=0; j<len(result20[i])/numeralWidth; j++ {
            fmt.Println(result20[i][line:line+numeralWidth])
            line += numeralWidth
        }
    }
}

/// Used by CodinGame to feed us the input via stdin
///
func main() {
    // L and H are the width (length) and height of each numeral pattern e.g. 
    // .oo.
    // o..o
    // .oo.
    // ....
    //
    // Has a L and H of 4
    var l, h int
    fmt.Scan(&l, &h)

    // The list of Mayan numerals are sent horizontally by scan lines (top to bottom)
    numerals := make([]string, h)
    for i:=0; i<h; i++ {
        var line string
        fmt.Scan(&line)
        numerals[i] = line
    }

    // The height of the first operand in lines
    var s1 int
    fmt.Scan(&s1)

    // The number is then scanned in lines top to bottom
    a := make([]string, s1)
    for i:=0; i<s1; i++ {
        var line string
        fmt.Scan(&line)
        a[i] = line
    }

    // The height of the second operand in lines
    var s2 int
    fmt.Scan(&s2)

    // The number is then scanned in lines top to bottom
    b := make([]string, s2)
    for i:=0; i<s2; i++ {
        var line string
        fmt.Scan(&line)
        b[i] = line
    }

    // The operation
    var operation string
    fmt.Scan(&operation)

    //TEST 01 - Simple addition
    // a = []string{"o...", "....", "....", "...."}
    // b = []string{"o...", "....", "....", "...."}
    // numerals = []string{".oo.o...oo..ooo.oooo....o...oo..ooo.oooo....o...oo..ooo.oooo....o...oo..ooo.oooo", "o..o................____________________________________________________________", ".oo.....................................________________________________________", "............................................................____________________"}
    // operation = "+"
    // l = 4

    //TEST 02 - Carry addition
    // a = []string{"ooo.", "____", "____", "____"}
    // b = []string{"ooo.", "....", "....", "...."}
    // numerals = []string{".oo.o...oo..ooo.oooo....o...oo..ooo.oooo....o...oo..ooo.oooo....o...oo..ooo.oooo", "o..o................____________________________________________________________", ".oo.....................................________________________________________", "............................................................____________________"}
    // operation = "+"
    // l = 4

    //TEST 05 - Subtraction
    // a = []string{"o...","____","....","....","ooo.","....","....","...."}
    // b = []string{"oo..","____","....","...."}
    // numerals = []string{".oo.o...oo..ooo.oooo....o...oo..ooo.oooo....o...oo..ooo.oooo....o...oo..ooo.oooo", "o..o................____________________________________________________________", ".oo.....................................________________________________________", "............................................................____________________"}
    // operation = "-"
    // l = 4

    //TEST 05 - Great Multiplication
    // a = []string{"o...","....","....","....","....","____","____","....","oo..","____","____","____","....","____","....","...."}
    // b = []string{"oooo","....","....","....","ooo.","____","____","____","oo..","____","____","....","....","____","____","....","oo..","____","____","...."}
    // numerals = []string{".oo.o...oo..ooo.oooo....o...oo..ooo.oooo....o...oo..ooo.oooo....o...oo..ooo.oooo", "o..o................____________________________________________________________", ".oo.....................................________________________________________", "............................................................____________________"}
    // operation = "*"
    // l = 4

    //TEST 09 - Zero
    // a = []string{"....","____","____","____","ooo.","____","....","....","oo..","____","____","....","o...","____","____","____"}
    // b = []string{".oo.","o..o",".oo.","...."}
    // numerals = []string{".oo.o...oo..ooo.oooo....o...oo..ooo.oooo....o...oo..ooo.oooo....o...oo..ooo.oooo", "o..o................____________________________________________________________", ".oo.....................................________________________________________", "............................................................____________________"}
    // operation = "*"
    // l = 4

    run(l, numerals, a, b, operation)
}