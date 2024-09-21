package main

import (
	"fmt"

	"github.com/solarlune/dngn"
)

var GameMap *dngn.Layout
var intGameMap []int

func InitMap(wallVal rune, doorVal rune, splitCount int, minRoomSize int) []int {
	ops := &dngn.BSPOptions{WallValue: wallVal, DoorValue: doorVal, SplitCount: splitCount, MinimumRoomSize: minRoomSize}

	GameMap := dngn.NewLayout(30, 30)
	GameMap.GenerateBSP(*ops)
	flatGameMapData := flatten(GameMap.Data)
	fmt.Print(flatGameMapData)
	for _, r := range flatGameMapData {
		switch r {
		case 32, 51:
			intGameMap = append(intGameMap, 0)
		case 48:
			intGameMap = append(intGameMap, 2)
		default:
			fmt.Printf("Unknown Tile value: %v", r)
		}
	}
	// fmt.Printf("This: %+v", flatGameMapData)
	// fmt.Print(intGameMap)
	// fmt.Print(len(intGameMap))
	return intGameMap
}

func InitOLayer(mapSize int, baseLayer []int) []int {
	oLayer := make([]int, mapSize)
	//TODO: make this dynamic, actually check for walls
	for i := range oLayer {
		if i == 20 {
			oLayer[i] = 1
			continue
		}
		if i == 21 {
			oLayer[i] = 3
			continue
		}
		oLayer[i] = 0
	}
	fmt.Printf("%+v", oLayer)
	return oLayer
}

func flatten[T any](slices [][]T) []T {
	var flat_slice []T
	for _, s := range slices {
		flat_slice = append(flat_slice, s...)
	}
	return flat_slice
}
