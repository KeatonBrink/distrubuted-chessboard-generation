package distributedchessboardgeneration

func DecodeBitBoard(cbInt int) Chessboard {
	var resetCB Chessboard
	resetCB.Reset()
	var retCB Chessboard
	retCB.HasCastled = false
	retCB.Board = make([][]ChessPiece, 8)
	for i := range retCB.Board {
		retCB.Board[i] = make([]ChessPiece, 8)
	}
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

// Inefficient algorithm for generating all power sets
func GetAll30PowerSet(outputChan chan<- int) {
	for curPieceCount := 1; curPieceCount <= 30; curPieceCount++ {
		PowerSetGen(curPieceCount, outputChan)
	}
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
