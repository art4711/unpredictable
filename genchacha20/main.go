package main

import (
	"fmt"
	"os"
	"io"
	"flag"
)

func step(w io.Writer, a, b, c, rot int) {
	fmt.Fprintf(w, "\t\tx%.2d += x%.2d\n", a, b)
	fmt.Fprintf(w, "\t\tx = x%.2d ^ x%.2d\n", c, a)
	fmt.Fprintf(w, "\t\tx%.2d = (x << %d) | (x >> %d)\n", c, rot, 32 - rot)
}

func quarterround(w io.Writer, a, b, c, d int) {
	step(w, a, b, d, 16)
	step(w, c, d, b, 12)
	step(w, a, b, d, 8)
	step(w, c, d, b, 7)
	fmt.Fprintln(w, "")
}

func core(w io.Writer, rounds int) {
	fmt.Fprintln(w, "package unpredictable")
	fmt.Fprintln(w, "")	
	fmt.Fprintln(w, "func (s *stream) core(output *block) {")

	fmt.Fprintln(w, "\tvar (")
	for i := 0; i < 16; i++ {
		fmt.Fprintf(w, "\t\tx%.2d = s.state[%d]\n", i, i)
	}
	fmt.Fprintln(w, "\t)")

	fmt.Fprintln(w, "")

	fmt.Fprintln(w, "\tvar x uint32")

	fmt.Fprintln(w, "")

	fmt.Fprintf(w, "\tfor i := 0; i < %d; i += 2 {\n", rounds)
	quarterround(w, 0, 4, 8, 12)
	quarterround(w, 1, 5, 9, 13)
	quarterround(w, 2, 6, 10, 14)
	quarterround(w, 3, 7, 11, 15)
	quarterround(w, 0, 5, 10, 15)
	quarterround(w, 1, 6, 11, 12)
	quarterround(w, 2, 7, 8, 13)
	quarterround(w, 3, 4, 9, 14)
	fmt.Fprintln(w, "\t}")

	fmt.Fprintln(w, "")

	for i := 0; i < 16; i++ {
		fmt.Fprintf(w, "\toutput[%d] = s.state[%d] + x%.2d\n", i, i, i)
	}

	fmt.Fprintln(w, "}")
}

func main() {
	flag.Parse()
	o, err := os.OpenFile(flag.Arg(0), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	core(o, 20)
}