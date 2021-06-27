package main

import (
	"bufio"
	"bytes"
	"os"
	"testing"
)

func Benchmark_run(b *testing.B) {
	bs, err := os.ReadFile("samples/D.in")
	if err != nil {
		b.Fail()
	}
	r := bufio.NewReader(bytes.NewBuffer(bs))
	input := readInput(r)

	for i := 0; i < b.N; i++ {
		run(input)
	}
}
