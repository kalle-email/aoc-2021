import scala.io.Source
object day3 {

  def part1(binNums: List[List[Int]]): Unit = {
    val cutOff = binNums.length / 2
    var gamma = "" // use stringBuilder to be cool ..
    var epsilon = ""
    // idk how to zip more than 2 collections in scala ..?
    // just 2 for loops : )

    for (i <- 0 to (binNums(0).length) - 1) { // loop over all 'positions'
      var zippedNum = 0
      for (currNum <- binNums) { // loop over all binary nums.
        zippedNum += currNum(i)
      }
      if (zippedNum > cutOff) {
        gamma += "1"
        epsilon += "0"
      } else {
        gamma += "0"
        epsilon += "1"
      }
    }
    println(
      s"Part1 ${Integer.parseInt(gamma, 2) * Integer.parseInt(epsilon, 2)}"
    )
  }

  // this is so ugly, way better to just have 2 separate loops for oxy/co2 imo
  def part2_helper(
      _binNums: List[List[Int]],
      compareFunc: (Int, Int) => Boolean
  ): Int = {

    var binNums = _binNums
    var cutOff = binNums.length / 2

    for (i <- 0 to (binNums(0).length) - 1) { // loop over all 'positions'
      var numOfOnes = 0
      for (currNum <- binNums) { // loop over all binary nums.
        numOfOnes += currNum(i)
      }

      var numOfZeroes = binNums.length - numOfOnes

      if (compareFunc(numOfOnes, numOfZeroes)) {
        binNums = binNums.filter(arr => arr(i) == 1)
      } else {
        binNums = binNums.filter(arr => arr(i) == 0)
      }
      cutOff = binNums.length / 2

      if (binNums.length == 1) {
        return Integer.parseInt(binNums(0).mkString(""), 2)
      }
    }
    return 0 // should not happen
  }

  def part2(binNums: List[List[Int]]): Unit = {

    val moreThan: (Int, Int) => Boolean = (x, y) => x >= y
    val lessThan: (Int, Int) => Boolean = (x, y) => x < y // ugh

    val oxy = part2_helper(binNums, moreThan)
    val co2 = part2_helper(binNums, lessThan)
    println(s"Part2 ${oxy * co2}")
  }

  def main(args: Array[String]): Unit = {
    val input: List[String] =
      Source.fromFile("../resources/day3.txt").getLines.toList
    val binNums = for (l <- input) yield {
      (l.map(_.asDigit)).toList
    }

    part1(binNums)
    part2(binNums)
  }
}
