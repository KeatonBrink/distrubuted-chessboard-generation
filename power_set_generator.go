package distributedchessboardgeneration

func DecodeBitBoard(cbInt int) Chessboard {
	var retCB Chessboard
	if 1<<0&cbInt > 0 {
		retCB
	}
}

// Inefficient algorithm for generating all power sets
func GetAll30PowerSet(outputChan chan<- int) {
	for curPieceCount := 1; curPieceCount <= 30; curPieceCount++ {
		PowerSetGen(curPieceCount, outputChan)
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

// Grunt work of power set, given an integer
func PowerSetGen(nbits int, outputChan chan<- int) {
	for i := 1; i < (1 << 30); i++ {
		bitCount := count1bits(i)
		if bitCount == nbits {
			outputChan <- i
		}
	}
	if nbits == 30 {
		close(outputChan)
	}
}
