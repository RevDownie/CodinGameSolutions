import math._
import scala.util._

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/
object Player {

    def clamp(value : Int, min : Int, max : Int) : Int = Math.max(Math.min(value, max), min)

    def seqToTuple[A](seq : Seq[A]) : (A, A) = {
        (seq(0), seq(1))
    }
    def main(args: Array[String]) {
    
        val n = readInt // the number of points used to draw the surface of Mars.
        
        val samplePoints = (for(i <- 0 until n) yield for(i <- readLine split " ") yield i.toInt) map(seqToTuple(_))
        val flats = samplePoints.groupBy(_._2).values
        val longestFlat = flats.maxBy(_.length)

        val flatMinX = longestFlat(0)._1
        val flatMaxX = longestFlat.last._1
        val flatY = longestFlat(0)._2

        // game loop
        while(true) {
            // hs: the horizontal speed (in m/s), can be negative.
            // vs: the vertical speed (in m/s), can be negative.
            // f: the quantity of remaining fuel in liters.
            // r: the rotation angle in degrees (-90 to 90).
            // p: the thrust power (0 to 4).
            val Array(x, y, hs, vs, f, r, p) = for(i <- readLine split " ") yield i.toInt
            
            //Find the direction to reach the flat ground. Rotate towards that vector
            //and thrust with min of 4 until we reach the flat ground and then rotate
            //to vertical
            var angle = 0
            var thrust = 3
            if(x > flatMaxX || x < flatMinX) {
                val closestFlatX = Math.min(Math.abs(flatMinX - x), Math.abs(flatMaxX - x))
                angle = clamp((Math.atan2(flatY - y, closestFlatX - x) * 57.2957795).toInt, -90, 90);
                thrust = 4
            }
            else {
                if(Math.abs(vs) >= 40) thrust = 4
            }

            println("%d %d".format(angle, thrust))
        }
    }
}