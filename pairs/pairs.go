package pairs

import "math/rand"

// GeneratePairs takes a src array and returns an array of pairs
// If the number of elements in the src array is not even one pair is gonna
// have an additional element
func GeneratePairs(src []string) [][]string {
	pairs := make([][]string, len(src)/2)

	for i := 0; i < len(src)/2; i++ {
		pairs[i] = []string{src[i*2], src[i*2+1]}
	}
	// We're gonna generate a triple for the last pair
	if len(src)%2 == 1 {
		pairs[len(pairs)-1] = append(pairs[len(pairs)-1], src[len(src)-1])
	}

	return pairs
}

// Shuffle an array with a seed
func Shuffle(src []string, seed int64) []string {
	r := rand.New(rand.NewSource(seed))

	dest := make([]string, len(src))
	perm := r.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}
	return dest
}
