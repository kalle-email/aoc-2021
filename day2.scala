import scala.io.Source

object day2 {

  def part1(tuples: List[(String, Int)]): Unit = {
    var depth = 0
    var horizontal = 0
    tuples.foreach {
      case ("forward", amount) => horizontal += amount
      case ("down", amount)    => depth += amount
      case ("up", amount)      => depth -= amount
      case (_, _)              => throw new Exception("should not happen")
    }
    println(s"${depth * horizontal}")
  }

  def part2(tuples: List[(String, Int)]): Unit = {
    var depth = 0
    var horizontal = 0
    var aim = 0
    tuples.foreach {
      case ("forward", amount) => {
        horizontal += amount
        depth += (aim * amount)
      }
      case ("down", amount) => aim += amount
      case ("up", amount)   => aim -= amount
      case (_, _)           => throw new Exception("should not happen")
    }
    println(s"${depth * horizontal}")
  }

  def main(args: Array[String]): Unit = {
    val input: List[String] =
      Source.fromFile("day2.txt").getLines.toList

    // val CoolWay= input
    //   .flatMap(_.split(" ").toList)
    //   .grouped(2)
    //   .collect { case List(a, b) => (a.toString, b.toInt) }
    //   .toList

    // split into (direction, amount)
    val commandTuples = for (l <- input) yield {
      val split = l.split(" ")
      val tuple = (split(0).toString, split(1).toInt)
      tuple
    }

    part1(commandTuples)
    part2(commandTuples)
  }
}
