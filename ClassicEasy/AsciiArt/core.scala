import math._
import scala.util._

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
object Solution {

    def main(args: Array[String]) {
        
        val chars = List('A','B','C','D','E','F','G','H','I','J',
                         'K','L','M','N','O','P','Q','R','S','T',
                         'U','V','W','X','Y','Z','?')
        
        val l = readInt
        val h = readInt
        val t = readLine.toUpperCase
        val rows = for(i <- 0 until h) yield readLine
        
        def combineRowForChar(char : Char, rowIndex : Int) : String = {
            var index = chars.indexOf(char)
            if(index == -1) index = chars.indexOf('?')
            val cursorX = index * l
            rows(rowIndex).substring(cursorX, cursorX + l)
        }
        
        for(i <- 0 until h) {
            val combinedRow = t.map(combineRowForChar(_, i)).mkString
            println(combinedRow)
        }
    }
}