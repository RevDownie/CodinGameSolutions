import math._
import scala.util._
import scala.collection.immutable.HashMap

object Solution {

    def main(args: Array[String]) {
        
        var mimeMap = new HashMap[String, String]
        
        val n = readInt // Number of elements which make up the association table.
        val q = readInt // Number Q of file names to be analyzed.
        
        for(i <- 0 until n) {
            // ext: file extension
            // mt: MIME type.
            val Array(ext, mt) = readLine split " "
            mimeMap = mimeMap + (ext.toLowerCase -> mt)
        }
        val extensionPattern = """(.*)[.]([^.]*)""".r

        for(i <- 0 until q) {
            readLine match { 
                case extensionPattern(name, ext) => println(mimeMap.getOrElse(ext.toLowerCase, "UNKNOWN"))
                case _ => println("UNKNOWN")
            }
        }
    }
}