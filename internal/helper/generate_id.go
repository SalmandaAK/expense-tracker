package helper

import (
	"maps"
	"slices"
)

// GenerateNumberId extracts Id which comparable with int type from a map.
// Then it finds a number before unused Id in a list of Id (which is in ascending sort).
// And then it generates Id by adding 1 into the number it found.
func GenerateNumberId[K ~int, V any](m map[K]V) int {
	idList := make([]int, 0, len(m))
	ids := maps.Keys(m)
	for id := range ids {
		idList = append(idList, int(id))
	}
	slices.Sort(idList)
	// lo and hi are the indices of lowest used Id and highest used Id respectively.
	lo, hi := 0, len(idList)-1

	// If the len(idList) == idList[hi] (highest Id), then there's no unused Id up to the highest used Id). So the next available Id is the highest used Id + 1.
	if len(idList) == idList[hi] {
		return idList[hi] + 1
	}

	// If len(idList) < idList[hi], there's any unused Id between lowest used Id, idList[0] and highest used Id (idList[len(idList) - 1])
	// We are going to find the first unused empty Id by using binary search algorithm until the range between hi and lo is only 1 (hi will always be greater than lo).
	for hi-lo != 1 {
		i := int(uint(lo+hi) >> 1)
		if len(idList[:i+1]) == idList[i] {
			lo = i
		} else {
			hi = i
		}
	}

	// The first unused Id will be positioned after lo, so the next available id will be idList[lo] + 1.
	return idList[lo] + 1
}
