package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// For each line, determine if it's a Negative (L) or (R) positive.
// if the newPointer would pass 100 or 0, add the remainder to the "other side" of the dial

func main() {

	dial := newDial()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("failed to open file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		// split by position
		direction := line[0:1]
		rotations := line[1:]

		rotationsParsedToInt, err := strconv.Atoi(rotations)
		if err != nil {
			log.Fatalf("failed converting string to int for rotations input")
		}

		dial.moveDial(direction, rotationsParsedToInt)

	}

	fmt.Printf("Pointer At Zero Counter: %d \n", dial.RotationLandOnZeroCount)
	fmt.Printf("Pointer Passed Zero Counter (clicks): %d \n", dial.FullRotationCount)
	fmt.Printf("Password: %d \n", dial.RotationLandOnZeroCount+dial.FullRotationCount)

	if err := scanner.Err(); err != nil {
		log.Fatalf("error during scanning: %v", err)
	}
}

type Dial struct {
	ceiling                 int
	floor                   int
	position                int
	RotationLandOnZeroCount int
	FullRotationCount       int
}

func newDial() *Dial {
	return &Dial{
		ceiling:                 99,
		floor:                   0,
		position:                50,
		RotationLandOnZeroCount: 0,
		FullRotationCount:       0,
	}
}

func (d *Dial) moveDial(direction string, rotations int) {

	fmt.Printf("Moving Dial... | Current position: %d | Direction to move: %s | Rotations %d \n", d.position, direction, rotations)
	additionalRotation := false
	switch direction {
	case "R":
		fullRotations := rotations / 100
		rotationsRemaining := rotations % 100

		newPosition := rotationsRemaining + d.position

		if newPosition > 99 {
			newPosition = (newPosition - d.ceiling) - 1
			additionalRotation = true

		}

		if newPosition == 0 {
			d.RotationLandOnZeroCount += 1
		}
		if additionalRotation && newPosition != 0 {
			fullRotations += 1
		}

		d.FullRotationCount += fullRotations
		d.position = newPosition

		fmt.Printf("FullRotations=%d \n", fullRotations)

	case "L":

		fullRotations := rotations / 100

		rotationsRemaining := rotations % 100

		newPosition := d.position - rotationsRemaining

		if newPosition < 0 {
			newPosition = d.ceiling - (rotationsRemaining - d.position) + 1
			additionalRotation = true
		}

		if newPosition == 0 {
			d.RotationLandOnZeroCount += 1
		}
		if additionalRotation && newPosition != 0 && d.position != 0 {
			fullRotations += 1
		}

		d.FullRotationCount += fullRotations
		d.position = newPosition
		fmt.Printf("FullRotations=%d \n", fullRotations)

	}

	fmt.Printf("Dial moved to position %d \n", d.position)

}
