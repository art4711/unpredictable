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

func add(w *wr, a, b int) {
	w.F("%s += %s\n", x(a), x(b))
}

func xor(w *wr, a, b int) {
	w.F("%s ^= %s\n", x(a), x(b))
}

func rot(w *wr, a, r int) {
	w.F("%s = (%s << %d) | (%s >> %d)\n", x(a), x(a), r, x(a), 32 - r)
}


type m struct {
	a int
	b int
	c int
	d int
}

func r(w *wr, x []m) {
	for _, u := range x {
		add(w, u.a, u.b)
		xor(w, u.d, u.a)
		rot(w, u.d, 16)
		add(w, u.c, u.d)
		xor(w, u.b, u.c)
		rot(w, u.b, 12)
	}
	for _, u := range x {
		add(w, u.a, u.b)
		xor(w, u.d, u.a)
		rot(w, u.d, 8)
		add(w, u.c, u.d)
		xor(w, u.b, u.c)
		rot(w, u.b, 7)
	}
}

func core(w *wr, rounds int) {
	w.Ln("func (s *stream) core(output *block) {")

	w.indent++

	for i := 0; i < 16; i++ {
		w.F("%s := s[%d]\n", x(i), i)
	}

	w.Nl()

	w.F("for i := %d; i > 0; i -= 2 {\n", rounds)
	w.indent++
	r(w, []m{ { 0, 4, 8, 12 }, { 1, 5, 9, 13 }, { 2, 6, 10, 14 }, { 3, 7, 11, 15 } })
	r(w, []m{ { 0, 5, 10, 15 }, { 1, 6, 11, 12 }, { 2, 7, 8, 13 }, { 3, 4, 9, 14 } })
	w.indent--
	w.Ln("}")

	w.Nl()

	for i := 0; i < 16; i++ {
		w.F("output[%d] = s[%d] + %s\n", i, i, x(i))
	}

	w.Ln("s[12]++")
	w.Ln("if s[12] == 0 {")
	w.indent++
	w.Ln("s[13]++")
	w.indent--
	w.Ln("}")
	w.Nl()

	w.indent--

	w.Ln("}")
}

func core_slice(w *wr, rounds int) {
	w.Ln("func (s *stream) core_slice(output []block) {")

	w.indent++

	w.Ln("for b := range output {")
	w.indent++

	for i := 0; i < 16; i++ {
		w.F("%s := s[%d]\n", x(i), i)
	}

	w.Nl()

	w.F("for i := %d; i > 0; i -= 2 {\n", rounds)
	w.indent++
	r(w, []m{ { 0, 4, 8, 12 }, { 1, 5, 9, 13 }, { 2, 6, 10, 14 }, { 3, 7, 11, 15 } })
	r(w, []m{ { 0, 5, 10, 15 }, { 1, 6, 11, 12 }, { 2, 7, 8, 13 }, { 3, 4, 9, 14 } })
	w.indent--
	w.Ln("}")

	w.Nl()

	for i := 0; i < 16; i++ {
		w.F("output[b][%d] = s[%d] + %s\n", i, i, x(i))
	}

	w.Ln("s[12]++")
	w.Ln("if s[12] == 0 {")
	w.indent++
	w.Ln("s[13]++")
	w.indent--
	w.Ln("}")
	w.Nl()

	w.indent--
	w.Ln("}")

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

	w.Ln("package unpredictable")
	w.Nl()

	core(w, 20)
	w.Nl()
	core_slice(w, 20)
}