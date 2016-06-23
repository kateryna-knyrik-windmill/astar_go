package main

import (
	"fmt"
)

type Scene struct {
	rows, columns int
	scene         [][]byte
}

func (scene *Scene) initScene(rows int, columns int) {
	scene.rows = rows
	scene.columns = columns

	scene.scene = make([][]byte, scene.rows)
	for i := 0; i < scene.rows; i++ {
		scene.scene[i] = make([]byte, scene.columns)
		for j := 0; j < scene.columns; j++ {
			if i == 0 || i == scene.rows-1 || j == 0 || j == scene.columns -1 {
				scene.scene[i][j] = '#'
			} else {
				scene.scene[i][j] = ' '
			}
		}
	}
}

func (scene *Scene) draw() {
	for i := 0; i < scene.rows; i++ {
		for j := 0; j < scene.columns; j++ {
			var color string
			switch scene.scene[i][j] {
			case '#':
				color = FgBlack
			case 'A':
				color = FgRed
			case 'B':
				color = FgBlue
			case '*':
				color = FgGreen
			}
			fmt.Printf("%s%c%s", color, scene.scene[i][j], Reset)
		}
		fmt.Printf("\n")
	}
}

func (scene *Scene) addWalls(num int) {
	for i := 0; i < num; i++ {
		origin := GetRandInt(2)
		length := GetRandInt(16) + 1
		row := GetRandInt(scene.rows)
		column := GetRandInt(scene.columns)
		switch origin {
		case 0:
			for i := 0; i < length; i++ {
				if column +i >= scene.columns {
					break
				}
				scene.scene[row][column +i] = '#'
			}

		case 1:
			for i := 0; i < length; i++ {
				if row+i >= scene.rows {
					break
				}
				scene.scene[row+i][column] = '#'
			}
		}
	}
}
