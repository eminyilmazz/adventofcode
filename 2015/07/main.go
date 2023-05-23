package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Gate struct {
	Input1  string
	Input2  string
	Output  string
	Operand string
}

func main() {
	fmt.Println(partOne())
}

func partOne() uint16 {
	instructions := readInput()
	circuit := make(map[string]uint16)

	for {
		for _, instruction := range instructions {
			if signal, output := parseInstruction(instruction, &circuit); output != "" {
				circuit[output] = signal
			}
		}
		fmt.Println(circuit)
		if signal, ok := circuit["a"]; ok {
			return signal
		}
	}
}

func parseInstruction(instruction string, circuit *map[string]uint16) (uint16, string) {
	parts := strings.Split(instruction, " -> ")
	fmt.Println(parts)
	inputs, output := parts[0], parts[1]
	//if signal, err := strconv.ParseUint(output, 10, 16); err != nil {
	//	fmt.Println(err)
	//	return uint16(signal), output
	//}
	if strings.HasPrefix(inputs, "NOT ") {
		input := strings.TrimPrefix(inputs, "NOT ")
		return evaluateGate(Gate{Operand: "NOT", Input1: input, Input2: ""}, circuit), output
	}

	for _, gate := range []Gate{
		{Operand: "AND"}, {Operand: "OR"}, {Operand: "LSHIFT"}, {Operand: "RSHIFT"},
	} {
		if strings.Contains(inputs, gate.Operand) {
			gateParts := strings.Split(inputs, " "+gate.Operand+" ")
			gate.Input1, gate.Input2 = gateParts[0], gateParts[1]
			evaluated := evaluateGate(gate, circuit)
			return evaluated, output
		}
	}
	return (*circuit)[inputs], output
}

func evaluateGate(gate Gate, circuit *map[string]uint16) uint16 {
	input1, input2 := (*circuit)[gate.Input1], (*circuit)[gate.Input2]

	switch gate.Operand {
	case "AND":
		return input1 & input2
	case "OR":
		return input1 | input2
	case "LSHIFT":
		shift, _ := strconv.Atoi(gate.Input2)
		return input1 << uint(shift)
	case "RSHIFT":
		shift, _ := strconv.Atoi(gate.Input2)
		return input1 >> uint(shift)
	case "NOT":
		return ^input1 & 0xFFFF
	default:
		return 0
	}
}

func readInput() []string {
	input, err := os.Open("input.txt")
	var lines []string
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
