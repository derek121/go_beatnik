package main

// https://en.wikipedia.org/wiki/Beatnik_(programming_language)
// https://esolangs.org/wiki/Beatnik

// --- ABCDEFGHIJKLMNOPQRSTUVWXYZ ---
//  1: A   E   I  L NO  RSTU
//  2:    D  G
//  3:  BC         M  P
//  4:      F H             VW Y
//  5:           K
//  6:
//  7:
//  8:          J             X
//  9:
// 10:                 Q        Z
// --- ABCDEFGHIJKLMNOPQRSTUVWXYZ ---

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

var wordVals = map[rune]int{
	'A': 1, 'a': 1,
	'B': 3, 'b': 3,
	'C': 3, 'c': 3,
	'D': 2, 'd': 2,
	'E': 1, 'e': 1,
	'F': 4, 'f': 4,
	'G': 2, 'g': 2,
	'H': 4, 'h': 4,
	'I': 1, 'i': 1,
	'J': 8, 'j': 8,
	'K': 5, 'k': 5,
	'L': 1, 'l': 1,
	'M': 3, 'm': 3,
	'N': 1, 'n': 1,
	'O': 1, 'o': 1,
	'P': 3, 'p': 3,
	'Q': 10, 'q': 10,
	'R': 1, 'r': 1,
	'S': 1, 's': 1,
	'T': 1, 't': 1,
	'U': 1, 'u': 1,
	'V': 4, 'v': 4,
	'W': 4, 'w': 4,
	'X': 8, 'x': 8,
	'Y': 4, 'y': 4,
	'Z': 10, 'z': 10,
}

const opMock1 = 1
const opMock2 = 2
const opMock3 = 3
const opMock4 = 4

const opPush = 5
const opPopDiscard = 6
const opPopTwoAddPush = 7
const opInputCharPush = 8
const opPopOutput = 9
const opPopTwoSubtractPush = 10
const opPopTwoSwapPush = 11
const opPopPushTwice = 12
const opPopSkipAheadIfZero = 13
const opPopSkipAheadIfNotZero = 14
const opPopSkipBackIfZero = 15
const opPopSkipBackIfNotZero = 16
const opStop = 17

func main() {
	var words []string
	words = readProgWords()
	// fmt.Printf("%d %v\n", len(words), words)

	var scores []int
	scores = getScores(words)
	//fmt.Printf("%d %T %v\n", len(scores), scores[0], scores)

	run(scores)
}

func run(prog []int) {
	var i int
	var stack []int
	var top, next int

	// TODO: less or greater than the consts

FOR:
	for i < len(prog) {
		switch prog[i] {
		case opMock1:
			// a
			fmt.Printf("opMock1: %d\n", opMock1)
			fmt.Println("Beatnik? HA!!")

		case opMock2:
			fmt.Printf("opMock2: %d\n", opMock2)
			fmt.Println("Beatnik? HA!!")

		case opMock3:
			fmt.Printf("opMock3: %d\n", opMock3)
			fmt.Println("Beatnik? HA!!")

		case opMock4:
			fmt.Printf("opMock4: %d\n", opMock4)
			fmt.Println("Beatnik? HA!!")

		case opPush:
			// Push the next word's value onto the stack.
			// k zzy k zzya
			fmt.Printf("opPush: %d %d\n", opPush, prog[i+1])
			stack = push(stack, prog[i+1])
			fmt.Printf("Stack: %v\n", stack)
			i++ // To go past the value to the push command

		case opPopDiscard:
			// Pop a number from the stack and discard it.
			// k zzy k zzya ka
			fmt.Printf("opPopDiscard: %d\n", opPopDiscard)
			stack, _ = pop(stack)
			fmt.Printf("Stack: %v\n", stack)

		case opPopTwoAddPush:
			// Pop two numbers, add them, and push the result.
			// k zzy k zzya kaa
			stack, top = pop(stack)
			stack, next = pop(stack)
			fmt.Printf("opPopTwoAddPush: %d %d\n", opPopTwoAddPush, top+next)
			stack = push(stack, top+next)
			fmt.Printf("Stack: %v\n", stack)

		case opInputCharPush:
			// Input a character and push its value.
			// kaaa
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Split(bufio.ScanBytes)
			scanner.Scan()
			input := scanner.Text()
			//fmt.Printf("Scanned... %T\n", input[0])
			//fmt.Println(input)
			fmt.Printf("opInputCharPush: %d %d\n", opInputCharPush, int(input[0]))
			stack = push(stack, int(input[0]))
			fmt.Printf("Stack: %v\n", stack)

		case opPopOutput:
			// Pop a number and output it as a character.
			// k zzzzzzzzzkaa kaaaa
			stack, top = pop(stack)
			fmt.Printf("opPopOutput: %d %c\n", opPopOutput, top)
			fmt.Printf("%c", top)
			fmt.Printf("Stack: %v\n", stack)

		case opPopTwoSubtractPush:
			// Pop two numbers, subtract the first one popped from the second
			// one popped, and push the result.
			// k zzy k zzya z
			stack, top = pop(stack)
			stack, next = pop(stack)
			fmt.Printf("opPopTwoSubtractPush: %d %d\n", opPopTwoSubtractPush, next-top)
			stack = push(stack, next-top)
			fmt.Printf("Stack: %v\n", stack)

		case opPopTwoSwapPush:
			// Pop two numbers, swap them, and push them back.
			// k zzy k zzya za
			stack, top = pop(stack)
			stack, next = pop(stack)
			fmt.Printf("opPopTwoSwapPush: %d %d %d\n", opPopTwoSwapPush, next, top)
			stack = push(stack, top)
			stack = push(stack, next)
			fmt.Printf("Stack: %v\n", stack)

		case opPopPushTwice:
			// Pop a number and push it twice.
			// k zzy k zzya zaa
			stack, top = pop(stack)
			fmt.Printf("opPopPushTwice: %d %d\n", opPopPushTwice, top)
			stack = push(stack, top)
			stack = push(stack, top)
			fmt.Printf("Stack: %v\n", stack)

		case opPopSkipAheadIfZero:
			// Pop a number and skip ahead n (actually n+1) words if the number is zero.
			// zaaa
			fmt.Printf("opPopSkipAheadIfZero. NO-OP\n")
			log.Fatalf("Unsupported operation- opPopSkipAheadIfZero. Zero cannot appear on the stack.")

		case opPopSkipAheadIfNotZero:
			// Pop a number and skip ahead n (actually n+1) words if the number isn't zero.
			// k a zaaaa k aa k aaa
			stack, top = pop(stack)
			fmt.Printf("opPopSkipAheadIfNotZero: %d %d\n", opPopSkipAheadIfNotZero, top)
			fmt.Printf("Stack: %v\n", stack)
			i += top

		case opPopSkipBackIfZero:
			// Pop a number and skip back n words if the number is zero.
			// zaaaaa
			fmt.Printf("opPopSkipBackIfZero. NO-OP\n")
			log.Fatalf("Unsupported operation- opPopSkipBackIfZero. Zero cannot appear on the stack.")

		case opPopSkipBackIfNotZero:
			// Pop a number and skip back n words if the number isn't zero.
			// Prog: Push 1. Skip ahead if not 0 (is 1; skips over Stop). Stop.
			//       Push 3. Skip back if not 0 (to Stop)
			// k a     zaaaa         zaaaaaaa k aa    zaaaaaa
			// Push 1. Skip ahead 1. Stop.    Push 2. Skip back
			stack, top = pop(stack)
			fmt.Printf("opPopSkipBackIfNotZero: %d %d\n", opPopSkipBackIfNotZero, top)
			fmt.Printf("Stack: %v\n", stack)
			i = (i - top - 1) // -1 to undo the +1 at end of switch

		case opStop:
			fmt.Printf("opStop: %d\n", opStop)
			break FOR

		default:
			if prog[i] > 23 {
				fmt.Printf("Beatnik applause, man... (%d)\n", prog[i])
			} else {
				// log.Fatalf("Illegal score: %d", prog[i])
				fmt.Printf("Unsupported score: %d\n", prog[i])
			}
		} // case

		i++
	} // for

	fmt.Printf("Stack at end: %v\n", stack)
} // func

func push(stack []int, val int) []int {
	return append(stack, val)
}

func pop(stack []int) ([]int, int) {
	return stack[:len(stack)-1], stack[len(stack)-1]
}

func getScores(words []string) []int {
	var scores []int

	for _, word := range words {
		// TODO: type of word[0] is unit8, but type of letter below is int32. ???
		// fmt.Println("Word: " + word)

		var score int
		for _, letter := range word {
			score += wordVals[letter]
			//fmt.Printf("Letter: %v %d\n", letter, wordVals[letter])
		}
		// fmt.Printf("Score: %d\n", score)
		scores = append(scores, score)
	}

	return scores
}

func readProgWords() []string {
	var all []string

	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Split(bufio.ScanWords)

	if len(os.Args) != 2 {
		log.Fatalf("usage: %s <filename>", os.Args[0])
	}

	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatalf("Error opening file %s: %v", os.Args[1], err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		w := scanner.Text()
		all = append(all, w)
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Scan file error: %v", err)
	}

	return all
}
