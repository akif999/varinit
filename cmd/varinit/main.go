package main

import (
	"github.com/akif999/varinit"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(varinit.Analyzer) }
