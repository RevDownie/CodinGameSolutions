import math._
import scala.util._

object Solution {

    case class Location(name : String, address : String, longitude : Double, latitude : Double)
    
    def distBetweenPoints(longA : Double, latA : Double, longB : Double, latB : Double) : Double = {
        val longRadA = longA * 0.0174532925
        val longRadB = longB * 0.0174532925
        val latRadA = latA * 0.0174532925
        val latRadB = latB * 0.0174532925
        
        val x = (longRadB - longRadA) * Math.cos((latRadA + latRadB)/2.0)
        val y = (latRadB - latRadA)
        val d = Math.sqrt(x * x + y * y) * 6371.0
        d
    }
    
    def main(args: Array[String]) {
        val personLong = readLine.replace(',', '.').toDouble
        val personLat = readLine.replace(',', '.').toDouble
        val count = readInt
        
        var closestDist : Double = 10000
        var closestName = ""
        
        for(i <- 0 until count) {
            var dist : Double  = 0
            var name  = ""
            val line = readLine
            line split ';' match {
                case Array(index, n, addr, phoneNum, lo, la) => {
                    val location = Location(n, addr, lo.replace(',', '.').toDouble, la.replace(',', '.').toDouble)
                    dist = distBetweenPoints(personLong, personLat, location.longitude, location.latitude)
                    name = location.name
                }
                case _ => {
                    val Array(place, pos) = line split ";;"
                    val Array(index, n, addr) = place split ";"
                    val Array(long, lat) = pos.split(";").map(_.replace(',', '.')).map(_.toDouble)
                    val location = Location(n, addr, long, lat)
                    dist = distBetweenPoints(personLong, personLat, location.longitude, location.latitude)
                    name = location.name
                }
            }

            if(dist <= closestDist) {
                closestDist = dist
                closestName = name
            }
        }

        println(closestName)
    }
}