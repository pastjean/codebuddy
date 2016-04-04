package pairs

import "math/rand"

// GeneratePairs takes a src array and returns an array of pairs.
// If the number of elements in the src array is not even the last pair is gonna
// have an additional element
func GeneratePairs(src []string) [][]string {
	pairs := make([][]string, len(src)/2)

	for i := 0; i < len(src)/2; i++ {
		pairs[i] = []string{src[i*2], src[i*2+1]}
	}

	// We're gonna generate a triple for the last group
	// if the array has a odd number of events
	if len(src)%2 == 1 {
		pairs[len(pairs)-1] = append(pairs[len(pairs)-1], src[len(src)-1])
	}

	return pairs
}

// Shuffle shuffles an array.
//
// src is the source array that is gonna be shuffled, a copy is made.
//
// seed is an integer to initialize a randomized source. The seed can be
// something like time.Now().UnixNano()
func Shuffle(src []string, seed int64) []string {
	r := rand.New(rand.NewSource(seed))

	dest := make([]string, len(src))
	perm := r.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}
	return dest
}
