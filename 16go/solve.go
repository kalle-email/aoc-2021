package main

import (
	"fmt"
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
	hex := "005532447836402684AC7AB3801A800021F0961146B1007A1147C89440294D005C12D2A7BC992D3F4E50C72CDF29EECFD0ACD5CC016962099194002CE31C5D3005F401296CAF4B656A46B2DE5588015C913D8653A3A001B9C3C93D7AC672F4FF78C136532E6E0007FCDFA975A3004B002E69EC4FD2D32CDF3FFDDAF01C91FCA7B41700263818025A00B48DEF3DFB89D26C3281A200F4C5AF57582527BC1890042DE00B4B324DBA4FAFCE473EF7CC0802B59DA28580212B3BD99A78C8004EC300761DC128EE40086C4F8E50F0C01882D0FE29900A01C01C2C96F38FCBB3E18C96F38FCBB3E1BCC57E2AA0154EDEC45096712A64A2520C6401A9E80213D98562653D98562612A06C0143CB03C529B5D9FD87CBA64F88CA439EC5BB299718023800D3CE7A935F9EA884F5EFAE9E10079125AF39E80212330F93EC7DAD7A9D5C4002A24A806A0062019B6600730173640575A0147C60070011FCA005000F7080385800CBEE006800A30C023520077A401840004BAC00D7A001FB31AAD10CC016923DA00686769E019DA780D0022394854167C2A56FB75200D33801F696D5B922F98B68B64E02460054CAE900949401BB80021D0562344E00042A16C6B8253000600B78020200E44386B068401E8391661C4E14B804D3B6B27CFE98E73BCF55B65762C402768803F09620419100661EC2A8CE0008741A83917CC024970D9E718DD341640259D80200008444D8F713C401D88310E2EC9F20F3330E059009118019A8803F12A0FC6E1006E3744183D27312200D4AC01693F5A131C93F5A131C970D6008867379CD3221289B13D402492EE377917CACEDB3695AD61C939C7C10082597E3740E857396499EA31980293F4FD206B40123CEE27CFB64D5E57B9ACC7F993D9495444001C998E66B50896B0B90050D34DF3295289128E73070E00A4E7A389224323005E801049351952694C000"

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
func parsePacket(fullNum []string, ptr *int) packet {

	// package type & versions.
	b := strings.Join(fullNum[*ptr:*ptr+3], "")
	pVer, _ := strconv.ParseInt(b, 2, 0)
	*ptr += 3
	pType, _ := strconv.ParseInt(strings.Join(fullNum[*ptr:*ptr+3], ""), 2, 0)
	*ptr += 3

	packet := packet{version: int(pVer), typeCode: int(pType), subPackets: []packet{}}
	if pType == 4 {
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

// func parseOperatorNum(fullNum []string, ptr int)            {}
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
