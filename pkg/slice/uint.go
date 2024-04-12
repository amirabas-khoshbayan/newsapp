package slice

func MapFromUintToUint64(l []uint) []uint64 {
	r := make([]uint64, len(l))
	for i := range l {
		r[i] = uint64(l[i])
	}

	return r
}

func MapFromUint64ToUint(l []uint64) []uint {
	r := make([]uint, len(l))
	for i := range l {
		r[i] = uint(l[i])
	}

	return r
}
