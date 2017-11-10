// Read inputs from System.in, Write outputs to use print.
// Your class name has to be Solution
import math._
import scala.util._

object Solution {
  def main(args: Array[String]) {
    val n = readInt
    val strengths = for(i <- 1 to n) yield readInt
    val strengthsOrdered = strengths.sorted
    val diffs = strengthsOrdered.sliding(2, 1).map(_.reduce((a, b) => Math.abs(b - a))).toList
    val min = diffs.min
    //println(strengthsOrdered)
    //println(diffs)
    println(min)
  }
}