package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/cody-s-lee/advent-of-code-2021/internal/get_input"
)

func main() {
	body := get_input.Body("https://adventofcode.com/2021/day/16/input")
	body = strings.TrimSpace(body)

	bits := make([]bool, len(body)*4)

	for i, r := range body {
		subbits := bitmap[r]
		for j := 0; j < 4; j++ {
			bits[i*4+j] = subbits[j]
		}
	}

	root, _ := ParsePacket(bits)
	fmt.Printf("%v\n", root)

	totalVersion := 0

	queue := []Packet{root}

	for len(queue) > 0 {
		p := queue[0]
		queue = queue[1:]
		totalVersion += p.header().Version
		switch v := p.(type) {
		case *OperatorPacket:
			queue = append(queue, v.SubPackets...)
		default:
		}
	}

	result := root.execute()

	fmt.Printf("Total version sum: %d\n", totalVersion)
	fmt.Printf("Result: %d\n", result)
}

type Packet interface {
	header() *PacketHeader
	setParent(Packet)
	execute() int
}

type PacketHeader struct {
	Version    int
	ID         int
	PacketType PacketType
	Parent     Packet
}

type PacketType int

// ParsePacket parses the next packet, including any subpackets, from the input bits
// returning the packet discovered and the remainder unused bits
func ParsePacket(bits []bool) (Packet, []bool) {
	header, bits := ParseHeader(bits)

	var packet Packet

	switch header.ID {
	case 4: // literal
		var value int
		value, bits = ParseLiteralValue(bits)

		packet = &LiteralPacket{
			PacketHeader: header,
			Value:        value,
		}
	default: // operator
		var mode bool
		var subPackets []Packet
		mode, subPackets, bits = ParseOperatorBody(bits)

		packet = &OperatorPacket{
			PacketHeader: header,
			Mode:         mode,
			SubPackets:   subPackets,
		}

		for _, subPacket := range subPackets {
			subPacket.setParent(packet)
		}
	}

	return packet, bits
}

func ParseLiteralValue(bits []bool) (int, []bool) {
	valueBits := make([]bool, 0)
	lastGroup := false

	for !lastGroup {
		lastGroup = !bits[0]
		group := bits[1:5]
		bits = bits[5:]

		valueBits = append(valueBits, group...)
	}

	return toInt(valueBits), bits
}

func ParseOperatorBody(bits []bool) (bool, []Packet, []bool) {
	mode := bits[0]
	subPackets := make([]Packet, 0)
	remainderIndex := 0

	switch mode {
	case false: // Mode 0 next 15 bits are length of bits of subpackets
		nSubBits := toInt(bits[1:16])
		bits = bits[16:]
		subBits := bits[:nSubBits]
		bits = bits[nSubBits:]
		for len(subBits) > 0 {
			var subPacket Packet
			subPacket, subBits = ParsePacket(subBits)
			subPackets = append(subPackets, subPacket)
		}
	case true: // Mode 1 next 11 bits are number of subpackets
		nSubPackets := toInt(bits[1:12])
		bits = bits[12:]
		for i := 0; i < nSubPackets; i++ {
			var subPacket Packet
			subPacket, bits = ParsePacket(bits)
			subPackets = append(subPackets, subPacket)
		}
	}

	remainder := bits[remainderIndex:]
	return mode, subPackets, remainder
}

func ParseHeader(bits []bool) (*PacketHeader, []bool) {
	version := bits[0:3]
	id := bits[3:6]
	remainder := bits[6:]

	packet := &PacketHeader{
		Version: toInt(version),
		ID:      toInt(id),
	}

	return packet, remainder
}

func toInt(bits []bool) int {
	value := 0

	for _, bit := range bits {
		value = 2 * value
		if bit {
			value++
		}
	}

	return value
}

type LiteralPacket struct {
	*PacketHeader
	Value int
}

func (lp LiteralPacket) header() *PacketHeader {
	return lp.PacketHeader
}

func (lp LiteralPacket) setParent(parent Packet) {
	lp.header().Parent = parent
}

func (lp LiteralPacket) execute() int {
	return lp.Value
}

type OperatorPacket struct {
	*PacketHeader
	Mode       bool
	SubPackets []Packet
}

func (op OperatorPacket) header() *PacketHeader {
	return op.PacketHeader
}

func (op OperatorPacket) setParent(parent Packet) {
	op.header().Parent = parent
}

func (op OperatorPacket) execute() int {
	switch op.ID {
	case 0: // sum
		result := 0
		for _, sp := range op.SubPackets {
			result += sp.execute()
		}
		return result
	case 1: // product
		result := 1
		for _, sp := range op.SubPackets {
			result *= sp.execute()
		}
		return result
	case 2: // min
		result := math.MaxInt
		for _, sp := range op.SubPackets {
			val := sp.execute()
			if val < result {
				result = val
			}
		}
		return result
	case 3: // max
		result := 0
		for _, sp := range op.SubPackets {
			val := sp.execute()
			if val > result {
				result = val
			}
		}
		return result
	case 5: // greater than
		a := op.SubPackets[0].execute()
		b := op.SubPackets[1].execute()

		result := 0
		if a > b {
			result = 1
		}
		return result
	case 6: // less than
		a := op.SubPackets[0].execute()
		b := op.SubPackets[1].execute()

		result := 0
		if a < b {
			result = 1
		}
		return result
	case 7: // equal to
		a := op.SubPackets[0].execute()
		b := op.SubPackets[1].execute()

		result := 0
		if a == b {
			result = 1
		}
		return result
	}

	return 0
}

var bitmap = map[rune][]bool{
	'0': {false, false, false, false},
	'1': {false, false, false, true},
	'2': {false, false, true, false},
	'3': {false, false, true, true},
	'4': {false, true, false, false},
	'5': {false, true, false, true},
	'6': {false, true, true, false},
	'7': {false, true, true, true},
	'8': {true, false, false, false},
	'9': {true, false, false, true},
	'A': {true, false, true, false},
	'B': {true, false, true, true},
	'C': {true, true, false, false},
	'D': {true, true, false, true},
	'E': {true, true, true, false},
	'F': {true, true, true, true},
}
