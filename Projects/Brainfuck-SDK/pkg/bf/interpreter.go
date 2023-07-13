package bf

import (
	"fmt"
	"io"
)

const DefaultMemorySize = 3e4

type Interpreter struct {
	output io.Writer
	input  io.Reader

	memory  [DefaultMemorySize]byte
	pointer int
}

func NewInterpreter(output io.Writer, input io.Reader) *Interpreter {
	return &Interpreter{
		output: output,
		input:  input,
	}
}

func (intrpr *Interpreter) Execute(program []byte) error {
	prog := tokenize(program)

	return intrpr.execute(prog)
}

type instruction int8

const (
	_ instruction = iota

	incPointer
	decPointer

	incByte
	decByte

	write
	read

	startLoop
	endLoop
)

type program []instruction

func tokenize(p []byte) program {
	prog := make(program, 0)

	for _, i := range p {
		switch i {
		case '>':
			prog = append(prog, incPointer)
		case '<':
			prog = append(prog, decPointer)
		case '+':
			prog = append(prog, incByte)
		case '-':
			prog = append(prog, decByte)
		case '.':
			prog = append(prog, write)
		case ',':
			prog = append(prog, read)
		case '[':
			prog = append(prog, startLoop)
		case ']':
			prog = append(prog, endLoop)
		}
	}

	return prog
}

func (intrpr *Interpreter) execute(prog program) error {
	for i := 0; i < len(prog); i++ {
		switch prog[i] {
		case incPointer:
			intrpr.pointer++
		case decPointer:
			intrpr.pointer--
		case incByte:
			intrpr.memory[intrpr.pointer]++
		case decByte:
			intrpr.memory[intrpr.pointer]--
		case write:
			if _, err := intrpr.output.Write([]byte{intrpr.memory[intrpr.pointer]}); err != nil {
				return err
			}
		case read:
			data := make([]byte, 0, 1)
			if _, err := intrpr.input.Read(data); err != nil {
				return err
			}

			intrpr.memory[intrpr.pointer] = data[0]
		case startLoop:
			if intrpr.memory[intrpr.pointer] == 0 {
				i = intrpr.searchEndLoop(i, prog)
			}
		case endLoop:
			i = intrpr.searchStartLoop(i, prog) - 1
		}
	}

	return nil
}

func (intrpr *Interpreter) searchStartLoop(i int, prog program) int {
	var stack int
	for i := i - 1; i != 0; i-- {
		instr := prog[i]

		if instr == endLoop {
			stack++
		}

		if instr == startLoop {
			stack--
		}

		if stack == -1 {
			return i
		}
	}

	panic(fmt.Sprint("invalid loop opened at ", i))

	return -1
}

func (intrpr *Interpreter) searchEndLoop(i int, prog program) int {
	var stack int
	for i := i + 1; i < len(prog); i++ {
		instr := prog[i]

		if instr == startLoop {
			stack++
		}

		if instr == endLoop {
			stack--
		}

		if stack == -1 {
			return i
		}
	}

	panic(fmt.Sprint("invalid loop closed at ", i))

	return -1
}
