package main

import (
	"fmt"
)

func setBit(value int64, i int) int64 {
	return value | (1 << i)
}

func clearBit(value int64, i int) int64 {
	return value & ^(1 << i)
}

func isBitSet(value int64, i int) bool {
	return (value & (1 << i)) != 0
}

func main() {
	var value int64
	var i int
	var bitValue int

	fmt.Print("Enter a number ")
	if _, err := fmt.Scan(&value); err != nil {
		fmt.Println("Incorrect input")
		return
	}

	fmt.Print("Enter the bit position")
	if _, err := fmt.Scan(&i); err != nil || i < 0 || i > 63 {
		fmt.Println("Incorrect bit position")
		return
	}

	fmt.Print("Enter the bit value ")
	if _, err := fmt.Scan(&bitValue); err != nil || (bitValue != 0 && bitValue != 1) {
		fmt.Println("Incorrect bit value")
		return
	}

	// Установка бита
	if bitValue == 1 {
		value = setBit(value, i)
		fmt.Printf("bit %d to 1, value %d\n", i, value)
	} else {
		value = clearBit(value, i)
		fmt.Printf("bit %d to 0, value %d\n", i, value)
	}

	// Проверка состояния i-го бита
	if isBitSet(value, i) {
		fmt.Printf("Bit %d to 1\n", i)
	} else {
		fmt.Printf("Bit %d to 0\n", i)
	}
}
