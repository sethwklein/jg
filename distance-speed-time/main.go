package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func mainError() (err error) {
	rand.Seed(time.Now().UnixNano())

	// the questions tell you a speed
	// they give you coords to figure a distance
	// they give you start time
	// you're supposed to figure eta

	// if i pick a distance
	// a start time
	// and a speed

	// start time is in minutes
	// and seems to be rounded to 5
	// so 24 * 60 / 5
	// we would also like to avoid overlapping midnight, i suppose
	// but maybe not, so keep it simple

	start := rand.Intn(24*60/5) * 5
	fmt.Printf("%02d%02d hours\n", start/60, start%60)

	// speed is in knots
	// range maybe 5 to 20
	// whole numbers

	speed := rand.Intn(20-5+1) + 5
	fmt.Printf("%d knots\n", speed)

	// distance
	// distance is in nm
	// is hidden behind coordinates
	// too short doesn't make sense
	// because you'd just say, "almost there"
	// a knot is a nm/h

	distance := rand.Float32()*30 + 1
	fmt.Printf("%0.2f nm\n", distance)

	fmt.Print("press enter to see the answer")
	fmt.Scanln()

	// eta is start + distance / speed

	eta := (start + int(distance*60/float32(speed))) % (24 * 60)
	fmt.Printf("eta: %02d%02d hours\n", eta/60, eta%60)

	// lat and long
	// are %2°%02.1f'

	// https://www.movable-type.co.uk/scripts/latlong.html
	// the formula at the top of the page might be more complex than needed
	// scroll down a bit

	return nil
}

func mainCode() int {
	err := mainError()
	if err == nil {
		return 0
	}
	fmt.Fprintf(os.Stderr, "%v: Error: %v\n", filepath.Base(os.Args[0]), err)
	return 1
}

func main() {
	os.Exit(mainCode())
}
