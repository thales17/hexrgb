package hexrgb

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

// ToRGB converts a string in hex format to an rgb string with commas between value
func ToRGB(h string) (string, error) {
	trimmedH := h
	if trimmedH[0] == '#' {
		trimmedH = h[1:]
	}
	if len(trimmedH) != 6 {
		return "", fmt.Errorf("Invalid hex string: %s", h)
	}
	hexR := trimmedH[0:2]
	hexG := trimmedH[2:4]
	hexB := trimmedH[4:6]

	decodedR, err := hex.DecodeString(hexR)
	if err != nil {
		return "", err
	}
	decodedG, err := hex.DecodeString(hexG)
	if err != nil {
		return "", err
	}
	decodedB, err := hex.DecodeString(hexB)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%d,%d,%d", uint8(decodedR[0]), uint8(decodedG[0]), uint8(decodedB[0])), nil
}

// ToHex converts a string in comma separated rgb form into a hex string
func ToHex(rgb string) (string, error) {
	comps := strings.Split(rgb, ",")
	if len(comps) != 3 {
		return "", fmt.Errorf("Invalid rgb string: %s", rgb)
	}
	r, err := strconv.Atoi(comps[0])
	if err != nil {
		return "", err
	}
	g, err := strconv.Atoi(comps[1])
	if err != nil {
		return "", err
	}
	b, err := strconv.Atoi(comps[2])
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("#%02X%02X%02X", r, g, b), nil
}
