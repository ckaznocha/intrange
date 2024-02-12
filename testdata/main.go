package main

func main() {
	for i := 2; i < 10; i++ {
	}

	for i := 0; i < 10; i += 2 {
	}

	for i := 0; i < 10; i++ {
		i += 1
	}

	for i := 0; i < 10; i++ {
		i++
	}

	for i := 0; i < 10; i++ {
		i = i + 1
	}

	for i := 0; i < 10; i++ { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := uint32(0); i < 10; i++ { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0x0; i < 10; i++ { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; i < 10; i += 1 { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; i < 10; i += 0x1 { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; i < 10; i = i + 1 { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; i < 10; i = i + 0x1 { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; i < 10; i = 1 + i { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; i < 10; i = 0x1 + i { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; 10 > i; i++ { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0x0; 10 > i; i++ { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; 10 > i; i += 1 { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; 10 > i; i += 0x1 { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; 10 > i; i = i + 1 { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; 10 > i; i = i + 0x1 { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; 10 > i; i = 1 + i { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; 10 > i; i = 0x1 + i { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	const x = 10

	for i := 2; i < x; i++ {
	}

	for i := 0; i < x; i += 2 {
	}

	for i := 0; i < x; i++ { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := uint32(0); i < uint32(x); i++ { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0x0; i < x; i++ { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; i < x; i += 1 { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; i < x; i += 0x1 { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; i < x; i = i + 1 { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; i < x; i = i + 0x1 { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; i < x; i = 1 + i { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; i < x; i = 0x1 + i { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; x > i; i++ { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; x > i; i += 1 { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; x > i; i += 0x1 { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; x > i; i = i + 1 { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; x > i; i = i + 0x1 { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; x > i; i = 1 + i { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; x > i; i = 0x1 + i { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}
}
