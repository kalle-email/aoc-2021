package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type packet struct {
	version    int
	typeCode   int
	num        int
	subPackets []packet
}

func main() {

	input, _ := ioutil.ReadFile("day16.txt") // just one big line
	hex := string(input)

	var sb strings.Builder
	for _, c := range strings.Split(hex, "") {
		decimal, _ := strconv.ParseInt(c, 16, 0)
		curr := fmt.Sprintf("%04b", decimal)
		sb.WriteString(curr)
	}
	binaryNum := strings.Split(sb.String(), "")
	ptr := 0
	packet := parsePacket(binaryNum, &ptr)
	total1 := sumP1(packet)
	fmt.Printf("TOTAL P1: %d\n", total1)

	total2 := sumP2(packet)
	fmt.Printf("TOTAL P2: %d\n", total2)
}

func parsePacket(fullNum []string, ptr *int) packet {
	packetVersion, _ := strconv.ParseInt(strings.Join(fullNum[*ptr:*ptr+3], ""), 2, 0)
	*ptr += 3
	packetTypeCode, _ := strconv.ParseInt(strings.Join(fullNum[*ptr:*ptr+3], ""), 2, 0)
	*ptr += 3

	packet := packet{version: int(packetVersion), typeCode: int(packetTypeCode), subPackets: []packet{}}
	if packetTypeCode == 4 {
		num := parseLiteral(fullNum, ptr)
		packet.num = num
	} else {
		_subPackets := parseOperator(fullNum, ptr)
		packet.subPackets = append(packet.subPackets, _subPackets...)
	}
	return packet
}

func parseOperator(fullNum []string, ptr *int) []packet {
	newPackets := []packet{}
	switch fullNum[*ptr] {
	case "0":
		*ptr++
		lenBinary := strings.Join(fullNum[*ptr:*ptr+15], "")
		lenInt, _ := strconv.ParseInt(lenBinary, 2, 0)
		*ptr += 15
		packets := parseOperatorLength(fullNum, ptr, int(lenInt))
		newPackets = append(newPackets, packets...)
	case "1":
		*ptr++
		numBinary := strings.Join(fullNum[*ptr:*ptr+11], "")
		numInt, _ := strconv.ParseInt(numBinary, 2, 0)
		*ptr += 11
		packets := parseOperatorNum(fullNum, ptr, int(numInt))

		newPackets = append(newPackets, packets...)
	}
	return newPackets
}

func parseOperatorLength(fullNum []string, ptr *int, length int) []packet {
	subPackets := []packet{}

	bitsRead := 0
	for bitsRead < length {
		ptrBefore := *ptr
		subPackets = append(subPackets, parsePacket(fullNum, ptr))
		bitsRead += *ptr - ptrBefore
	}
	return subPackets
}

func parseOperatorNum(fullNum []string, ptr *int, num int) []packet {
	subPackets := []packet{}
	for i := 0; i < num; i++ {
		subPackets = append(subPackets, parsePacket(fullNum, ptr))
	}
	return subPackets
}

func parseLiteral(fullNum []string, ptr *int) int {

	cont := true
	numsBin := []string{}
	for cont {
		flag := fullNum[*ptr]
		if flag == "0" { // last num
			cont = false
		}
		*ptr++
		numsBin = append(numsBin, strings.Join(fullNum[*ptr:*ptr+4], ""))
		*ptr += 4
	}
	var sb strings.Builder
	for _, n := range numsBin {
		sb.WriteString(n)
	}

	numInt, _ := strconv.ParseInt(sb.String(), 2, 0)
	return int(numInt)
}

func sumP1(p packet) int {
	total := 0
	for _, subPacket := range p.subPackets {
		total += sumP1(subPacket)
	}
	return total + p.version
}

func sumP2(p packet) int { // this one is a beast but idk how to make it prettier in golang
	total := 0
	switch p.typeCode {
	case 0:
		for _, subPacket := range p.subPackets {
			total += sumP2(subPacket)
		}
	case 1:
		total = 1
		for _, subPacket := range p.subPackets {
			total *= sumP2(subPacket)
		}
	case 2:
		smallest := 999999999999999
		for _, subPacket := range p.subPackets {
			curr := sumP2(subPacket)
			if curr < smallest {
				smallest = curr
			}
		}
		total = smallest
	case 3:
		largest := -1
		for _, subPacket := range p.subPackets {
			curr := sumP2(subPacket)
			if curr > largest {
				largest = curr
			}
		}
		total = largest
	case 4:
		total = p.num
	case 5:
		total = 0
		firstValSet := false
		var firstVal int
		for _, subPacket := range p.subPackets {
			if !firstValSet {
				firstVal = sumP2(subPacket)
				firstValSet = true
			} else {
				if firstVal > sumP2(subPacket) {
					total = 1
				}
			}
		}
	case 6:
		total = 0
		firstValSet := false
		var firstVal int
		for _, subPacket := range p.subPackets {
			if !firstValSet {
				firstVal = sumP2(subPacket)
				firstValSet = true
			} else {
				if firstVal < sumP2(subPacket) {
					total = 1
				}
			}
		}
	case 7:
		total = 0
		firstValSet := false
		var firstVal int
		for _, subPacket := range p.subPackets {
			if !firstValSet {
				firstVal = sumP2(subPacket)
				firstValSet = true
			} else {
				if firstVal == sumP2(subPacket) {
					total = 1
				}
			}
		}
	}
	return total
}
