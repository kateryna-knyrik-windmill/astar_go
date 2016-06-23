package main

import (
	"fmt"
	"math"
	"os"
)

var originPoint, destinationPoint Point
var openList, closeList, path []Point

/**
	set start point
 */
func setOriginPoint(scene *Scene) {
	originPoint = Point{GetRandInt(scene.rows-2) + 1, GetRandInt(scene.columns -2) + 1, 0, 0, 0, nil}
	if scene.scene[originPoint.X][originPoint.Y] == ' ' {
		scene.scene[originPoint.X][originPoint.Y] = 'A'
	} else {
		setOriginPoint(scene)
	}
}

/**
 	Set the destination point
  */
func setDestinationPoint(scene *Scene) {
	destinationPoint = Point{GetRandInt(scene.rows-2) + 1, GetRandInt(scene.columns -2) + 1, 0, 0, 0, nil}

	if scene.scene[destinationPoint.X][destinationPoint.Y] == ' ' {
		scene.scene[destinationPoint.X][destinationPoint.Y] = 'B'
	} else {
		setDestinationPoint(scene)
	}
}

/**
 	Init origin, destination. Put the origin point into the openlist by the way
  */
func initAstar(scene *Scene) {
	setOriginPoint(scene)
	setDestinationPoint(scene)
	openList = append(openList, originPoint)
}

func findPath(scene *Scene) {
	current := getFMin()
	addToCloseList(current, scene)
	walkable := getWalkable(current, scene)
	for _, point := range walkable {
		addToOpenList(point)
	}
}

func getFMin() Point {
	if len(openList) == 0 {
		fmt.Println("No possible path")
		os.Exit(-1)
	}
	index := 0
	for i, point := range openList {
		if (i > 0) && (point.F <= openList[index].F) {
			index = i
		}
	}
	return openList[index]
}

func getWalkable(point Point, scene *Scene) []Point {
	var around []Point
	row, column := point.X, point.Y
	left := scene.scene[row][column -1]
	up := scene.scene[row-1][column]
	right := scene.scene[row][column +1]
	down := scene.scene[row+1][column]
	leftup := scene.scene[row-1][column -1]
	rightup := scene.scene[row-1][column +1]
	leftdown := scene.scene[row+1][column -1]
	rightdown := scene.scene[row+1][column +1]
	if (left == ' ') || (left == 'B') {
		around = append(around, Point{row, column - 1, 0, 0, 0, &point})
	}
	if (leftup == ' ') || (leftup == 'B') {
		around = append(around, Point{row - 1, column - 1, 0, 0, 0, &point})
	}
	if (up == ' ') || (up == 'B') {
		around = append(around, Point{row - 1, column, 0, 0, 0, &point})
	}
	if (rightup == ' ') || (rightup == 'B') {
		around = append(around, Point{row - 1, column + 1, 0, 0, 0, &point})
	}
	if (right == ' ') || (right == 'B') {
		around = append(around, Point{row, column + 1, 0, 0, 0, &point})
	}
	if (rightdown == ' ') || (rightdown == 'B') {
		around = append(around, Point{row + 1, column + 1, 0, 0, 0, &point})
	}
	if (down == ' ') || (down == 'B') {
		around = append(around, Point{row + 1, column, 0, 0, 0, &point})
	}
	if (leftdown == ' ') || (leftdown == 'B') {
		around = append(around, Point{row + 1, column - 1, 0, 0, 0, &point})
	}
	return around
}

func addToOpenList(point Point) {
	updateWeight(&point)
	if checkExist(point, closeList) {
		return
	}
	if !checkExist(point, openList) {
		openList = append(openList, point)
	} else {
		if openList[findPoint(point, openList)].F > point.F { //New path found
			openList[findPoint(point, openList)].Parent = point.Parent
		}
	}
}

// Update G, H, F of the point
func updateWeight(point *Point) {
	if checkRelativePos(*point) == 1 {
		point.G = point.Parent.G + 10
	} else {
		point.G = point.Parent.G + 14
	}
	absx := (int)(math.Abs((float64)(destinationPoint.X - point.X)))
	absy := (int)(math.Abs((float64)(destinationPoint.Y - point.Y)))
	point.H = (absx + absy) * 10
	point.F = point.G + point.H
}

func removeFromOpenList(point Point) {
	index := findPoint(point, openList)
	if index == -1 {
		os.Exit(0)
	}
	openList = append(openList[:index], openList[index+1:]...)
}

func addToCloseList(point Point, scene *Scene) {
	removeFromOpenList(point)
	if (point.X == destinationPoint.X) && (point.Y == destinationPoint.Y) {
		generatePath(point, scene)
		scene.draw()
		os.Exit(1)
	}
	if scene.scene[point.X][point.Y] != 'A' {
		scene.scene[point.X][point.Y] = '·'
	}
	closeList = append(closeList, point)
}

func checkExist(chekExistsPoint Point, pointsArray []Point) bool {
	for _, point := range pointsArray {
		if chekExistsPoint.X == point.X && chekExistsPoint.Y == point.Y {
			return true
		}
	}
	return false
}

func findPoint(findPoint Point, pointsArray []Point) int {
	for index, point := range pointsArray {
		if findPoint.X == point.X && findPoint.Y == point.Y {
			return index
		}
	}

	return -1
}

func checkRelativePos(point Point) int {
	parent := point.Parent
	horizontal := (int)(math.Abs((float64)(point.X - parent.X)))
	vertical := (int)(math.Abs((float64)(point.Y - parent.Y)))
	return horizontal + vertical
}

func generatePath(point Point, scene *Scene) {
	if (scene.scene[point.X][point.Y] != 'A') && (scene.scene[point.X][point.Y] != 'B') {
		scene.scene[point.X][point.Y] = '*'
	}
	if point.Parent != nil {
		generatePath(*(point.Parent), scene)
	}
}
