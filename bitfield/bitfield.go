package bitfield

type Bitfield []byte

// Has Piece tells if a bitfield ahs a paricular index set
func (bf Bitfield) HasPiece(index int) bool{
	byteIndex := index/8
	offset := index%8
	if byteIndex < 0 || byteIndex >= len(bf){
		return false
	}
	return bf[byteIndex] >> (7-offset)&1 !=0
}

// SetPiece sets a bit in the bitfield
func (bf Bitfield) SetPiece(index int) {
	byteIndex := index / 8
	offset := index % 8

	// silently discard invalid bounded index
	if byteIndex < 0 || byteIndex >= len(bf) {
		return 
	}
	bf[byteIndex] |= 1 << (7 - offset)
}