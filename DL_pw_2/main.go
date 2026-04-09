package main

import (
	"fmt"
	"strings"
)

type CustomNum struct {
	value []uint64
}

func NewCustomNum() *CustomNum {
	return &CustomNum{}
}

func (customNum *CustomNum) SetHex(hexString string) {
	n := (len(hexString) + 15) / 16 // split into chunks of 16 characters
	data := make([]uint64, n)

	for i := 0; i < n; i++ {
		start := len(hexString) - (i+1)*16
		if start < 0 {
			start = 0
		}
		chunk := hexString[start : len(hexString)-i*16]
		data[i] = parseHexToUint64(chunk)
	}

	customNum.value = data
}

func parseHexToUint64(hexString string) uint64 {
	var result uint64
	multiplier := uint64(1)

	for i := len(hexString) - 1; i >= 0; i-- {
		val := hexString[i]
		if val >= '0' && val <= '9' {
			result += uint64(val-'0') * multiplier
		} else if val >= 'A' && val <= 'F' {
			result += uint64(val-'A'+10) * multiplier
		} else if val >= 'a' && val <= 'f' { // Also handle lowercase hex digits
			result += uint64(val-'a'+10) * multiplier
		}
		multiplier *= 16
	}

	return result
}

func (customNum *CustomNum) GetHex() string {
	hexChars := "0123456789ABCDEF"
	var result string

	for _, num := range customNum.value {
		for i := 0; i < 16; i++ {
			result = string(hexChars[num&0xF]) + result
			num >>= 4
		}
	}

	for len(result) > 1 && result[0] == '0' {
		result = result[1:]
	}

	return result
}

func GetHex(input []uint64) string {
	hexChars := "0123456789ABCDEF"
	var result string

	for _, num := range input {
		for i := 0; i < 16; i++ {
			result = string(hexChars[num&0xF]) + result
			num >>= 4
		}
	}

	// Removing leading zeros
	for len(result) > 1 && result[0] == '0' {
		result = result[1:]
	}

	return result
}

func (customNum *CustomNum) Inv() {
	result := make([]uint64, len(customNum.value))
	//since every elements stores 64 bits
	for i, val := range customNum.value {
		result[i] = ^val
	}

	customNum.value = result
}

func toBinaryString(val uint64) string {
	binStr := ""
	for i := 0; i < 64; i++ {
		if val%2 == 1 {
			binStr = "1" + binStr
		} else {
			binStr = "0" + binStr
		}
		val /= 2
	}
	return binStr
}

func fromBinaryString(binStr string) uint64 {
	val := uint64(0)
	factor := uint64(1)
	for i := len(binStr) - 1; i >= 0; i-- {
		if binStr[i] == '1' {
			val += factor
		}
		factor *= 2
	}
	return val
}

func (customNum *CustomNum) XOR(otherCustomNum *CustomNum) *CustomNum {
	maxLen := len(customNum.value)
	if len(otherCustomNum.value) > maxLen {
		maxLen = len(otherCustomNum.value)
	}

	a := make([]uint64, maxLen)
	b := make([]uint64, maxLen)
	copy(a[maxLen-len(customNum.value):], customNum.value)
	copy(b[maxLen-len(otherCustomNum.value):], otherCustomNum.value)

	result := make([]uint64, maxLen)

	for i := 0; i < maxLen; i++ {
		binA := toBinaryString(a[i])
		binB := toBinaryString(b[i])
		xorResult := ""

		for j := 0; j < 64; j++ {
			if binA[j] != binB[j] {
				xorResult += "1"
			} else {
				xorResult += "0"
			}
		}

		result[i] = fromBinaryString(xorResult)
	}

	return &CustomNum{result}
}

func (customNum *CustomNum) OR(otherCustomNum *CustomNum) *CustomNum {
	maxLen := len(customNum.value)
	if len(otherCustomNum.value) > maxLen {
		maxLen = len(otherCustomNum.value)
	}

	a := make([]uint64, maxLen)
	b := make([]uint64, maxLen)
	copy(a[maxLen-len(customNum.value):], customNum.value)
	copy(b[maxLen-len(otherCustomNum.value):], otherCustomNum.value)

	result := make([]uint64, maxLen)

	for i := 0; i < maxLen; i++ {
		binA := toBinaryString(a[i])
		binB := toBinaryString(b[i])
		orResult := ""

		for j := 0; j < 64; j++ {
			if binA[j] == '1' || binB[j] == '1' {
				orResult += "1"
			} else {
				orResult += "0"
			}
		}

		result[i] = fromBinaryString(orResult)
	}

	return &CustomNum{result}
}

func (customNum *CustomNum) AND(otherCustomNum *CustomNum) *CustomNum {
	maxLen := len(customNum.value)
	if len(otherCustomNum.value) > maxLen {
		maxLen = len(otherCustomNum.value)
	}

	a := make([]uint64, maxLen)
	b := make([]uint64, maxLen)
	copy(a[maxLen-len(customNum.value):], customNum.value)
	copy(b[maxLen-len(otherCustomNum.value):], otherCustomNum.value)

	result := make([]uint64, maxLen)

	for i := 0; i < maxLen; i++ {
		binA := toBinaryString(a[i])
		binB := toBinaryString(b[i])
		andResult := ""

		for j := 0; j < 64; j++ {
			if binA[j] == '1' && binB[j] == '1' {
				andResult += "1"
			} else {
				andResult += "0"
			}
		}

		result[i] = fromBinaryString(andResult)
	}

	return &CustomNum{result}
}

func (customNum *CustomNum) ShiftL(n uint) {
	resultData := make([]uint64, len(customNum.value))

	uint64ToBinary := func(val uint64) string {
		s := ""
		for i := 0; i < 64; i++ {
			if val%2 == 1 {
				s = "1" + s
			} else {
				s = "0" + s
			}
			val /= 2
		}
		return s
	}
	binaryToUint64 := func(s string) uint64 {
		val := uint64(0)
		multiplier := uint64(1)
		for i := len(s) - 1; i >= 0; i-- {
			if s[i] == '1' {
				val += multiplier
			}
			multiplier *= 2
		}
		return val
	}

	combinedBinary := ""
	for _, chunk := range customNum.value {
		combinedBinary += uint64ToBinary(chunk)
	}

	combinedBinary += strings.Repeat("0", int(n))

	for i := 0; i < len(resultData); i++ {
		startIndex := 64 * i
		endIndex := startIndex + 64
		if endIndex <= len(combinedBinary) {
			resultData[i] = binaryToUint64(combinedBinary[startIndex:endIndex])
		}
	}

	customNum.value = resultData
}

func main() {
	//shiftL
	aShiftL := NewCustomNum()
	aShiftL.SetHex("403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c")
	aShiftL.ShiftL(4)
	if aShiftL.GetHex() == "403DB8AD88A3932A0B7E8189AED9EEFFB8121DFAC05C3512FDB396DD73F6331C" {
		fmt.Printf("shift left works correct\n")
	} else {
		fmt.Printf("shift left works incorrect, result is: %s\n", aShiftL.GetHex())
	}
	
	//AND
	aAND := NewCustomNum()
	bAND := NewCustomNum()
	aAND.SetHex("51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4")
	bAND.SetHex("403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c")
	cAND := aAND.AND(bAND)
	andedHex := cAND.GetHex()
	if andedHex == "403D208400A113220340808088D16A1B10121078400C1002748196DD62460204" {
		fmt.Printf("bitwise AND works correct\n")
	} else {
		fmt.Printf("bitwise AND works incorrect, result is: %s\n", andedHex)
	}

	//OR
	a1 := NewCustomNum()
	b1 := NewCustomNum()
	a1.SetHex("51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4")
	b1.SetHex("403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c")
	c1 := a1.OR(b1)
	oredHex := c1.GetHex()
	if oredHex == "51BFF8AD9CAFD72EABFFBFC9BEFFFFFFFCFFBFFAFFDD779AFDF3D7FDF7F73FBC" {
		fmt.Printf("bitwise OR works correct\n")
	} else {
		fmt.Printf("bitwise OR works incorrect, result is: %s\n", oredHex)
	}

	//XOR
	a := NewCustomNum()
	b := NewCustomNum()
	a.SetHex("51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4")
	b.SetHex("403db8ad88a3932a0b7e8189aed9eeffb8121dfac05c3512fdb396dd73f6331c")
	c := a.XOR(b)
	xoredHex := c.GetHex()
	if xoredHex == "1182D8299C0EC40CA8BF3F49362E95E4ECEDAF82BFD167988972412095B13DB8" {
		fmt.Printf("XOR works correct\n")
	} else {
		fmt.Printf("XOR works incorrect, result is: %s\n", xoredHex)
	}

	//INV
	cN := NewCustomNum()
	cN.SetHex("A")
	cN.Inv()
	res := cN.GetHex()
	if res == "FFFFFFFFFFFFFFF5" {
		fmt.Printf("bitwise NOT works correct\n")
	} else {
		fmt.Printf("bitwise NOT works incorrect, result is: %s\n", res)
	}

	//SET and GET
	tests := []string{
		"ABCDEF",
		"1234567890",
		"FEDCBA",
		"1AF",
		"51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4",
		"70983d692f648185febe6d6fa607630ae68649f7e6fc45b94680096c06e4fadb",
	}

	for _, test := range tests {
		customNum := NewCustomNum()
		customNum.SetHex(test)
		if customNum.GetHex() != strings.ToUpper(test) {
			fmt.Printf("Failed for %s. Got %s\n", test, customNum.GetHex())
		} else {
			fmt.Printf("Passed for %s\n", test)
		}
	}
}
