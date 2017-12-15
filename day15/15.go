package day15

func GenA(seed int) func() int {
	a := seed
	b := (a * 16807) % 2147483647
	return func() int {
		a, b = b, (b*16807)%2147483647
		return a
	}
}

func GenB(seed int) func() int {
	a := seed
	b := (a * 48271) % 2147483647
	return func() int {
		a, b = b, (b*48271)%2147483647
		return a
	}
}

func GenX(seed int) func() int {
	// gen x produces gen a numbers that are multiples of 4
	genA := GenA(seed)
	return func() int {
		a := genA()
		for a%4 != 0 {
			a = genA()
		}
		return a
	}
}

func GenY(seed int) func() int {
	// gen b produces gen b numbers that are multiples of 8
	genB := GenB(seed)
	return func() int {
		b := genB()
		for b%8 != 0 {
			b = genB()
		}
		return b
	}
}

func Solve1(seedA, seedB int) int {
	matches := 0
	genA := GenA(seedA)
	genB := GenB(seedB)
	for i := 0; i < 40000000; i++ {
		a, b := genA(), genB()
		al, bl := a&0xFFFF, b&0xFFFF
		if al == bl {
			matches++
		}
	}
	return matches
}

func Solve2(seedA, seedB int) int {
	matches := 0
	genX := GenX(seedA)
	genY := GenY(seedB)

	for i := 0; i < 5000000; i++ {
		a, b := genX(), genY()
		al, bl := a&0xFFFF, b&0xFFFF
		if al == bl {
			matches++
		}
	}

	return matches
}
