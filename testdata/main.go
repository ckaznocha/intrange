package main

import "testing"

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

	var b *testing.B
	for i := 0; i < b.N; i++ {
	}

	for i := 0; b.N >= i; i++ {
	}

	var n int
	for i := 0; i < n; i++ {
		n--
	}

	for i := 0; i < n; i++ {
		n++
	}

	// Example from https://github.com/ckaznocha/intrange/issues/12
	var what string
	for i := 0; i < len(what); i++ {
		if what[i] == 'v' && i+1 < len(what) && what[i+1] >= '0' && what[i+1] <= '9' {
			what = what[:i] + what[i+1:]
		}
	}

	for i := 0; i < len(what); i++ { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	var t struct{ n int }
	for i := 0; i < t.n; i++ { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; i < t.n; i++ {
		t.n++
	}

	var s []int
	for i := 0; i < len(s); i++ { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; i < len(s); i++ {
		s = append(s, 4)
	}

	var m map[int]int
	for i := 0; i < len(m); i++ { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; i < len(m); i++ {
		m[4] = 4
	}

	var t2 struct{ m map[int]int }
	for i := 0; i < len(t2.m); i++ { // want `for loop can be changed to use an integer range \(Go 1\.22\+\)`
	}

	for i := 0; i < len(t2.m); i++ {
		t2.m[4] = 4
	}
}
