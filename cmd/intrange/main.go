package main

import (
	"golang.org/x/tools/go/analysis/unitchecker"

	"github.com/ckaznocha/intrange"
)

func main() { unitchecker.Main(intrange.Analyzer) }
