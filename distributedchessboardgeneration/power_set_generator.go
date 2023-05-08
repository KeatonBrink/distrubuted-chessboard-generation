package distributedchessboardgeneration

import (
	"log"
)

func DecodeBitBoard(cbInt int) Chessboard {
	var resetCB Chessboard
	resetCB.Reset()
	retCB := GenerateEmptyCB()
	for curP := 0; curP <= 7; curP++ {
		// Black Pawns
		if (1<<(curP+15))&cbInt > 0 {
			retCB.Board[6][curP] = resetCB.Board[6][curP]
		}
		// White Pawns
		if (1<<(curP+7))&cbInt > 0 {
			retCB.Board[1][curP] = resetCB.Board[1][curP]
		}
		// White Outer Row
		if curP < 4 {
			if (1<<curP)&cbInt > 0 {
				retCB.Board[0][curP] = resetCB.Board[0][curP]
			}
			// Black outer row
			if ((1<<(curP+23))&cbInt > 0) || curP == 4 {
				retCB.Board[7][curP] = resetCB.Board[7][curP]
			}
			// After index 4, some off by one errors occur, but the curP+1 should fix it
		} else if curP < 7 {
			if (1<<curP)&cbInt > 0 {
				retCB.Board[0][curP+1] = resetCB.Board[0][curP+1]
			}
			// Black outer row
			if (1<<(curP+23))&cbInt > 0 {
				retCB.Board[7][curP+1] = resetCB.Board[7][curP+1]
			}
		}
	}
	// Note, index 4 is king and is always printed, but is not i the bit board
	retCB.Board[0][4] = resetCB.Board[0][4]
	retCB.Board[7][4] = resetCB.Board[7][4]
	return retCB
}

/************************************************
Function: GetAll30PowerSet
Input:
Output: Powerset of 30 sent to output int channel

Example:
Powerset of 3,
	(1, 2, 4, 3, 5, 6, 7)
	(0b001, 0b010, 0b100, 0b011, 0b101, 0b110, 0b111)
************************************************/
// Algorithm for generating power set
func GetAll30PowerSet(outputChan chan<- int) {
	// Some constants for comparison
	// Bit Count is the number of bits in the that can be 1s
	const bitCount = 30
	// Upperbit is the overflow integer, anything generated bigger than this is bad
	const uppperBit = 1 << bitCount
	// Initialize the array of bits that will be modified
	var bitSet []int
	curSet := 0
	// While there are still sets not solved, ie the number is less than 1 <<30
	for curSet < uppperBit-1 {
		// If all zeros are at the end of the current set
		if (curSet<<len(bitSet))&(uppperBit-1) == 0 {
			// For loop to initialize next round
			for curIndex := 0; curIndex < len(bitSet); curIndex++ {
				bitSet[curIndex] = 1<<len(bitSet) - curIndex
			}
			// Append a new 1 bit
			bitSet = append(bitSet, 1)
			// Else the current bit count stays the same
		} else {
			for index, curBitVal := range bitSet {
				// If the bit can move without overflowing 30 bits
				if curBitVal < (1 << (bitCount - 1 - index)) {
					// Left shift the bit
					bitSet[index] = bitSet[index] << 1
					// Break out of for loop as no more bits need shifting
					break
					// Else left shift index twice past previous bit
				} else {
					bitSet[index] = bitSet[index+1] << 2
				}
			}
		}
		// Generate the set and send it to the channel
		curSet = bit2Int(bitSet)
		// Error checking
		if curSet >= uppperBit {
			log.Fatalf("Set value too big %v, Vs Upper limit %v", curSet, uppperBit)
		}
		// log.Printf("Current bit count %v and Value %v", count1bits(curSet), curSet)
		outputChan <- curSet
	}
}

// Counts the 1 bits in an integer
func count1bits(integer int) int {
	count := 0
	curBitInteger := 1 << 0
	for ; curBitInteger <= integer; curBitInteger = curBitInteger << 1 {
		if integer&curBitInteger > 0 {
			count++

		}
	}
	return count
}

func bit2Int(bitArray []int) int {
	retInt := 0
	for _, val := range bitArray {
		retInt = retInt | val
	}
	return retInt
}
