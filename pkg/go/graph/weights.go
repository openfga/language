package graph

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

// WeightMap is a map of where the key is a type (e.g. folder, user) and the value is the weight/complexity to reach that type.
type WeightMap map[string]int

func (wt WeightMap) String() string {
	var sb strings.Builder

	// Extract keys and sort them
	keys := make([]string, 0, len(wt))
	for k := range wt {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		formatV := strconv.Itoa(wt[k])
		if wt[k] == math.MaxInt {
			formatV = "+∞"
		}
		sb.WriteString(fmt.Sprintf("%v=%s,", k, formatV))
	}
	formattedWeights := sb.String()
	if len(formattedWeights) > 0 {
		formattedWeights = formattedWeights[:len(formattedWeights)-1]
	}

	return fmt.Sprintf("weights:[%v]", formattedWeights)
}

func Intersection(maps ...WeightMap) (WeightMap, error) {
	if len(maps) == 0 {
		return nil, fmt.Errorf("%w: no maps given to compute intersection", ErrBuildingGraph)
	}

	// Copy the keys from the first map to avoid mutating the map itself later on
	intersectionMap := make(WeightMap)
	for key, value := range maps[0] {
		intersectionMap[key] = value
	}

	// For each subsequent map, retain only the keys that exist in the current intersection map
	for i := 1; i < len(maps); i++ {
		currentMap := maps[i]
		for key := range intersectionMap {
			if _, exists := currentMap[key]; !exists {
				delete(intersectionMap, key)
			}
		}
	}

	return intersectionMap, nil
}
