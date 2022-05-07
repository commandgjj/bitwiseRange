package bitwiseRange

import (
	"fmt"
	"math"
)

var (
	bitmasks []uint32
)

type Range struct {
	Start  uint32
	End    uint32
}

type BitRange struct {
	Value uint32
	Mask  uint32
}

// bitmasks: 0x00000000, 0x00000001, 0x00000003, 0x00000007, ..., 0x7fffffff
func initialize() {
	var curMask uint32 = 0
	bitmasks = append(bitmasks, curMask)
	for i := 0; i < 31; i++ {
		curMask <<= 1
		curMask ++
		bitmasks = append(bitmasks, curMask)
	}
}

func GetBitwiseRanges(ranges []Range, isPort bool) []BitRange {
	bitRange := []BitRange{}
	if len(bitmasks) == 0 {
		initialize()
	}

	for _, r := range ranges {
		min := r.Start
		max := r.End

		// Entire field is wildcarded
		if isPort && min == 0 && max == math.MaxUint16 {
			return bitRange
		}

		if min == 0 && max == math.MaxUint32 {
			// Entire integer was used to wildcard -- it may be wrong
			return bitRange
		}

		length := len(bitmasks)
		port := min
		for port <= max {
			for i := 1; i < length; i++ {
				j := bitmasks[i]
				mask := ^j
				// Check if we have found the least sig. bits which if flipped
				// we wouldn't be in range. If so, that bit should be the last
				// unmasked one.
				if (port & mask) < min || (port | j) > max {
					bitRange = append(bitRange, BitRange{port, ^bitmasks[i - 1]})
					// set port to be above the added range
					port = (port | bitmasks[i - 1]) + 1
					break
				}
			}
		}
	}
	return bitRange
}
// Test
//func main() {
//	ranges := []Range{
//		{22, 40},
//		{1000, 1999},
//		{0, 65535},
//	}
//	bitRange := GetBitwiseRanges(ranges, true)
//	for _, rangeMask := range bitRange {
//		fmt.Printf("%d 0x%x\n", rangeMask.Value, rangeMask.Mask)
//	}
//}
