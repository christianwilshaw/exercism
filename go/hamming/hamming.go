package hamming

import "errors"

//Distance returns the hamming distance between 2 strings
func Distance(a, b string) (int, error) {
	var alength = len(a)
	var blength = len(b)
	var hammingDistance = 0

	if alength != blength {
		return -1, errors.New("strand dont match")
	}

	for i := 0; i < alength; i++ {
		if a[i] != b[i] {
			hammingDistance++
		}
	}

	return hammingDistance, nil
}
