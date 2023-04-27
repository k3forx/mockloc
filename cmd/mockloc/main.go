package main

import (
	"github.com/k3forx/mockloc"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(mockloc.Analyzer) }
