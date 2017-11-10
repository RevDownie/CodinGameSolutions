import math._
import scala.util._

object Player {

    case class Bounds(min : (Int, Int), max : (Int, Int))
    
    def getMidpoint(a : Int, b : Int) : Int = {
        Math.floor((b - a)/2.0f).toInt + a
    }
    
    def getNextBounds(dir : String, pos : (Int, Int), bounds : Bounds) : Bounds = {
        dir match {
            case "U" => Bounds(bounds.min, (bounds.max._1, pos._2)) 
            case "UR" => Bounds((pos._1, bounds.min._2), (bounds.max._1, pos._2)) 
            case "R" => Bounds((pos._1, bounds.min._2), bounds.max) 
            case "DR" => Bounds(pos, bounds.max) 
            case "D" => Bounds((bounds.min._1, pos._2), bounds.max)  
            case "DL" => Bounds((bounds.min._1, pos._2), (pos._1, bounds.max._2)) 
            case "L" => Bounds(bounds.min, (pos._1, bounds.max._2))  
            case "UL" => Bounds(bounds.min, pos) 
        }
    }
    
    def getNextPos(dir : String, pos : (Int, Int), bounds : Bounds) : (Int, Int) = {
        dir match {
            case "U" => (pos._1, getMidpoint(bounds.min._2, pos._2)) 
            case "UR" => (getMidpoint(pos._1, bounds.max._1), getMidpoint(bounds.min._2, pos._2))
            case "R" => (getMidpoint(pos._1, bounds.max._1), pos._2)
            case "DR" => (getMidpoint(pos._1, bounds.max._1), getMidpoint(pos._2, bounds.max._2)) 
            case "D" => (pos._1, getMidpoint(pos._2, bounds.max._2)) 
            case "DL" => (getMidpoint(bounds.min._1, pos._1), getMidpoint(pos._2, bounds.max._2)) 
            case "L" => (getMidpoint(bounds.min._1, pos._1), pos._2)  
            case "UL" => (getMidpoint(bounds.min._1, pos._1), getMidpoint(bounds.min._2, pos._2))  
        }
    }
    
    def main(args: Array[String]) {
        val Array(w, h) = for(i <- readLine split " ") yield i.toInt
        val n = readInt
        val Array(x0, y0) = for(i <- readLine split " ") yield i.toInt
        var nextPos = (x0, y0)
        var nextBounds = Bounds((0, 0), (w, h))
        
        // game loop
        while(true) {
            val bombDir = readLine
        
            //Find the mid point between the edge and the current pos in the given direction
            val pos = nextPos
            val bounds = nextBounds
            nextPos = getNextPos(bombDir, pos, bounds)
            nextBounds = getNextBounds(bombDir, pos, bounds)
            println("%d %d".format(nextPos._1, nextPos._2))
        }
    }
}