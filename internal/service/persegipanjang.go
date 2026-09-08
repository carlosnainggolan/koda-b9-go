package service

func Area (p uint8, l uint8) uint8 {
	return p * l
}

func Keliling (p uint8, l uint8) uint8 {
	return 2 * (p + l)
}

func Gabungan (p uint8, l uint8) (a uint8, k uint8) {
	a = Area(p, l)
	k = Keliling(p, l)
	return a, k
}