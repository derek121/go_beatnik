# Go Beatnik!
Implementation of an interpreter for the Beatnik esoteric programming language in Go.

## About

As described on Wikipedia [1]:

    Beatnik is a simple stack-based esoteric programming language, by Cliff L. Biffle. A beatnik program consists of any sequence of English words. Each word is assigned the score you would get for it in a Scrabble game. The value of the score determines what function is performed. Functions include pushing the score of the next word onto the stack, testing the stack and skipping forward or backward in the program and other stack operations.

## Example

From esolangs.org. A valid Beatnik program (found in `beatnik-esolang-example.txt`):

```
Hello, aunts! Around, around, swim!
```

The code contains trace output to monitor the execution- what operation is run, the state of the stack, etc. This program executes as follows:

```
Input a character, take its ASCII value, add 7 to it, and output the corresponding ASCII character. For example, if the input character is A, the output character will be H.
```

The `A` and the `H` may be seen in the output below.

```
$ ./beatnik beatnik-esolang-example.txt 
A
opInputCharPush: 8 65
Stack: [65]
opPush: 5 7
Stack: [65 7]
opPopTwoAddPush: 7 72
Stack: [72]
opPopOutput: 9 H
HStack: []
Stack at end: []
$
```

Further examples for each operation may be found in the `src` directory.

## References

[Wikipedia](https://en.wikipedia.org/wiki/Beatnik_\(programming_language\))

[esolangs.org](https://esolangs.org/wiki/Beatnik)

