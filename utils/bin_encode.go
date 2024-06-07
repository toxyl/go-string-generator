package utils

import "strings"

// Binary Encoding function
func BinaryEncode(input string) string {
	var result strings.Builder
	for _, b := range []byte(input) {
		result.WriteString(byteToBinaryString(b))
	}
	return result.String()
}

// Helper function to convert a byte to a binary string
func byteToBinaryString(b byte) string {
	return strings.Join([]string{
		string('0' + (b>>7)&1),
		string('0' + (b>>6)&1),
		string('0' + (b>>5)&1),
		string('0' + (b>>4)&1),
		string('0' + (b>>3)&1),
		string('0' + (b>>2)&1),
		string('0' + (b>>1)&1),
		string('0' + b&1),
	}, "")
}
