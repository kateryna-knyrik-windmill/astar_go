package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Point struct {
	X, Y    int
	H, G, F int
	Parent  *Point
}

func (point Point) String() string {
	return "[" + strconv.Itoa(point.X) + ", " + strconv.Itoa(point.Y) + ", " + strconv.Itoa(point.F) + "]"
}

var random *rand.Rand

func init() {
	random = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func Clear() {
	fmt.Printf("\033[100B")
	for i := 0; i < 100; i++ {
		fmt.Printf("\033[1A")
		fmt.Printf("\033[K")
	}
}

func GetRandInt(limit int) int {
	return random.Intn(limit)
}
