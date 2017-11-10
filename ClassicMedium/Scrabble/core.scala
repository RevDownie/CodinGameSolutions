import math._
import scala.util._

object Solution {

    val letterScores = Map('e' -> 1, 'a' -> 1, 'i' -> 1, 'o' -> 1, 'n' -> 1, 'r' -> 1, 't' -> 1, 'l' -> 1, 's' -> 1, 'u' -> 1,
                           'd' -> 2, 'g' -> 2,
                           'b' -> 3, 'c' -> 3, 'm' -> 3, 'p' -> 3,
                           'f' -> 4, 'h' -> 4, 'v' -> 4, 'w' -> 4, 'y' -> 4,
                           'k' -> 5,
                           'j' -> 8, 'x' -> 8,
                           'q' -> 10, 'z' -> 10)
                               
    def generateUnallowedLetters(allowedLetters : String) : Seq[Char] = {
        val letters = "abcdefghijklmnopqrstuvwxyz"
  
        def loop(allowedLetters : Seq[Char], remainingLetters : String) : String = {
            allowedLetters match {
                case h :: t =>  loop(t, remainingLetters.replace(h.toString, ""))
                case _ => remainingLetters
            }
        }
        
        loop(allowedLetters.toList, letters).toList
    }
    
    def toScore(string : String) : Int = {
        string.map(letterScores.getOrElse(_, 0)).sum
    }
    
    def filterPossibles(words : Seq[String], allowedLetters : String) : Seq[String] = {
        val unallowedLetters = generateUnallowedLetters(allowedLetters)
        def containsUnallowed(word : String) : Boolean = (for(letter <- unallowedLetters) yield word.contains(letter)).contains(true)
        def hasEnoughLetters(word : String) : Boolean = (for(letter <- allowedLetters) yield word.count(_ == letter) <= allowedLetters.count(_ == letter)).contains(false) == false
        words.filterNot(_.length > allowedLetters.length).filterNot(containsUnallowed(_)).filter(hasEnoughLetters(_))
    }
        
    def main(args: Array[String]) {
        val n = readInt
        val words = for(i <- 0 until n) yield readLine
        val allowedLetters = readLine
        val possibles = filterPossibles(words, allowedLetters)
        //println(allowedLetters)
        //println(unallowedLetters)
        //println(words)
        //println(possibles)
        
        val scores = possibles.map(toScore(_))
        val max = scores.max
        val answer = possibles(scores.indexOf(max))
        println(answer)
    }
}