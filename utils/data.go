package utils

func IsBinaryData(data []byte) bool {
	// Define a threshold of how many non-printable characters define binary data
	const threshold = 0.3 // 30% of data being non-printable indicates binary

	// Count non-printable characters
	var nonPrintableCount int
	for _, b := range data {
		// Ignore common printable characters
		if b == 0x09 || b == 0x0A || b == 0x0D { // tab, new line, carriage return
			continue
		}
		if b < 0x20 || b > 0x7E { // outside of printable ASCII range
			nonPrintableCount++
		}
	}

	// Calculate ratio of non-printable characters
	nonPrintableRatio := float64(nonPrintableCount) / float64(len(data))

	// Return true if non-printable ratio exceeds the threshold
	return nonPrintableRatio > threshold
}
