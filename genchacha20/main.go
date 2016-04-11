package main

import (
	"fmt"
	"os"
	"io"
	"flag"
	"strings"
)

type wr struct {
	w io.Writer
	indent int
}

func (w *wr) Ln(o ...interface{}) {
	fmt.Fprint(w.w, strings.Repeat("\t", w.indent))
	fmt.Fprintln(w.w, o...)
}

func (w *wr) F(f string, o ...interface{}) {
	fmt.Fprintf(w.w, strings.Repeat("\t", w.indent) + f, o...)
}

func (w *wr) Nl() {
	fmt.Fprintln(w.w, "")
}

func x(i int) string {
	return fmt.Sprintf("x%.2d", i)
}

func step(w *wr, a, b, c, rot int) {
	w.F("%s += %s\n", x(a), x(b))
	w.F("%s ^= %s\n", x(c), x(a))
	w.F("%s = (%s << %d) | (%s >> %d)\n", x(c), x(c), rot, x(c), 32 - rot)
}

func quarterround(w *wr, a, b, c, d int) {
	step(w, a, b, d, 16)
	step(w, c, d, b, 12)
	step(w, a, b, d, 8)
	step(w, c, d, b, 7)
	w.Nl()
}

func core(w *wr, rounds int) {
	w.Ln("package unpredictable")
	w.Nl()
	w.Ln("func (s *stream) core(output *block) {")

	w.indent++

	for i := 0; i < 16; i++ {
		w.F("%s := s.state[%d]\n", x(i), i)
	}

	w.Nl()

	w.F("for i := 0; i < %d; i += 2 {\n", rounds)
	w.indent++
	quarterround(w, 0, 4, 8, 12)
	quarterround(w, 1, 5, 9, 13)
	quarterround(w, 2, 6, 10, 14)
	quarterround(w, 3, 7, 11, 15)
	quarterround(w, 0, 5, 10, 15)
	quarterround(w, 1, 6, 11, 12)
	quarterround(w, 2, 7, 8, 13)
	quarterround(w, 3, 4, 9, 14)
	w.indent--
	w.Ln("}")

	w.Nl()

	for i := 0; i < 16; i++ {
		w.F("output[%d] = s.state[%d] + %s\n", i, i, x(i))
	}

	w.Ln("s.state[12]++")
	w.Ln("if s.state[12] == 0 {")
	w.indent++
	w.Ln("s.state[13]++")
	w.indent--
	w.Ln("}")
	w.Nl()

	w.indent--

	w.Ln("}")
}

func main() {
	flag.Parse()
	o, err := os.OpenFile(flag.Arg(0), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	w := &wr{ o, 0 }
	core(w, 20)
}