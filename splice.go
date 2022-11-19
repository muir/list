package list

import (
	"fmt"
)

// ReplaceBeyond subsitutes replace into dest starting at start.
//
// Dest may be modified.
//
// Replace will panic if start is beyond the end of dest.
func Replace[E any](dest []E, start int, replace ...E) []E {
	if start > len(dest) {
		panic(fmt.Errorf("slice index out of bounds start %d > len %d", start, len(dest)))
	}
	return ReplaceBeyond[E](dest, start, replace...)
}

// ReplaceBeyond subsitutes replace into dest starting at start.  Start can
// be beyond the current end of dest.
//
// Dest may be modified.
func ReplaceBeyond[E any](dest []E, start int, replace ...E) []E {
	return SpliceBeyond[E](dest, start, start+len(replace), replace...)
}

// Splice removes replaces dest[start:end] with replace.  Replace
// can be larger or smaller and dest will grow or shrink as needed.
// Dest will be grown or shrunk as needed.
//
// End must be < len(dest) or else Splice panics.
//
// Dest may be modified.
func Splice[E any](dest []E, start int, end int, replace ...E) []E {
	if end > len(dest) {
		panic(fmt.Errorf("slice index out of bounds end %d > len %d", end, len(dest)))
	}
	return SpliceBeyond[E](dest, start, end, replace...)
}

// SpliceBeyond removes replaces dest[start:end] with replace.  Replace
// can be larger or smaller and dest will grow or shrink as needed.
// start or end can be beyond the current len or cap of dest and
// dest will be grown or shrunk as needed.
//
// End must be >= start or else SpliceBeyond panics.
//
// Dest may be modified.
func SpliceBeyond[E any](dest []E, start int, end int, replace ...E) []E {
	if end < start {
		panic(fmt.Errorf("invalid splice: end %d must be > start %d", end, start))
	}
	if start == len(dest) {
		// end is irrelevant
		return append(dest, replace...)
	}
	if start > len(dest) {
		// end is irrelevant
		if start+len(replace) > cap(dest) {
			// a copy is needed.  We might as well make it explicit
			n := make([]E, start+len(replace), ((start+len(replace))*3)/2)
			copy(n, dest)
			copy(n[start:], replace)
			return n
		}
		dest = dest[:start+len(replace)]
		copy(dest[start:], replace)
		return dest
	}
	if end > len(dest) {
		// any amount of end that is beyond the the current
		// slice doesn't matter
		end = len(dest)
	}
	max := len(dest) - (end - start) + len(replace)
	if max > cap(dest) {
		// not enough room, which means we're copying everything
		// anyway which makes things a bit simpler
		n := make([]E, max, (max*3)/2)
		copy(n, dest[:start])
		copy(n[start:], replace)
		if end < len(dest) {
			copy(n[start+len(replace):], dest[end:])
		}
		return n
	}
	if max > len(dest) {
		dest = dest[:max]
	}
	if end-start == len(replace) {
		copy(dest[start:end], replace)
		return dest
	}
	// replacement is a difference size from the part
	// being replaced.  We'll need to move things around
	//
	// The bit of dest that's after the replaced part
	// needs to move first.
	delta := len(replace) - (end - start)
	copy(dest[end+delta:], dest[end:])
	copy(dest[start:start+len(replace)], replace)
	if delta < 0 {
		dest = dest[:max]
	}
	return dest
}
