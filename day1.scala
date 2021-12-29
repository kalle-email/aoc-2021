import scala.io.Source
object day1 {

  def part1(depths: List[Int]): Unit = {
    val count = depths.sliding(2).filter(p => p.head < p.last)
    println(s"part1: ${count.length}")
  }

  def part2(depths: List[Int]): Unit = {
    var count = 0
    for (t <- depths.sliding(4)) {
      if ((t.slice(0, 3).sum) < t.slice(1, 4).sum) {
        count += 1
      }
    }
    println(s"part2: ${count}")
  }

  def main(args: Array[String]): Unit = {
    val depths: List[Int] =
      Source.fromFile("day1.txt").getLines.map(_.toInt).toList

    part1(depths) // 1559
    part2(depths) // 1600
  }
}
