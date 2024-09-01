package main

import (
	"fmt"

	"github.com/solarlune/dngn"
)

var GameMap *dngn.Layout
var intGameMap []int
func InitMap(wallVal rune, doorVal rune, splitCount int, minRoomSize int) []int {
	ops := &dngn.BSPOptions{WallValue: wallVal, DoorValue: doorVal, SplitCount: splitCount, MinimumRoomSize: minRoomSize}
	
	GameMap := dngn.NewLayout(30,30)
	GameMap.GenerateBSP(*ops)
	flatGameMapData := flatten(GameMap.Data)
	for _, r := range flatGameMapData {
		if r == '1' {
			intGameMap = append(intGameMap, 1)
			} else {
				intGameMap = append(intGameMap, 2)
			}
		}
	// fmt.Printf("This: %+v", flatGameMapData)
	fmt.Print(intGameMap)
	fmt.Print(len(intGameMap))
	return intGameMap
}

func flatten[T any](slices [][]T) []T{
	var flat_slice []T
	for _, s := range slices {
		flat_slice = append(flat_slice, s...)
	}
	return flat_slice
}