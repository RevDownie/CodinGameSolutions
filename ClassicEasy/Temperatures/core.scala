import math._
import scala.util._

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
object Solution {

    def findClosestToZero(a : Int, b :Int) = {
        if(Math.abs(a) > Math.abs(b)) b 
        else if(Math.abs(a) < Math.abs(b)) a 
        else if(a > b) a
        else b
    }
    
    def main(args: Array[String]) {
        val n = readInt // the number of temperatures to analyse
        val temps = readLine // the N temperatures expressed as integers ranging from -273 to 5526
        
        // Write an action using println
        // To debug: Console.err.println("Debug messages...")
        //-5 -4 -2 12 -40 4 2 18 11 5
        if(n == 0) {
            print(0)
        }
        else {
            val closest = temps.split(' ').map(_.toInt).reduce((a, b) => findClosestToZero(a, b))
            println(closest)
        }
    }
}