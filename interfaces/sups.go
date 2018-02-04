/*
From the previous worker/supervisor lab, create a Printer interface to print out info
on worker/supervisor. Define a function print that takes in a Printer and displays
the correct information on either types.
Make sure to test out your new function!
*/
package methods

import (
	"fmt"
	"strings"
)

type (
	printer interface {
		info() string
	}

	worker struct {
		id   int
		name string
	}

	supervisor struct {
		worker
		level int
	}
)

func (w *worker) info() string {
	return fmt.Sprintf("*[%d] %s", w.id, w.name)
}

func (w *worker) setName(newOne string) {
	w.name = newOne
}

func (s *supervisor) info() string {
	return fmt.Sprintf("%s (%d)", s.worker.info(), s.level)
}

func (s *supervisor) setLevel(newOne int) {
	s.level = newOne
}

func main() {
	w := worker{id: 100, name: "Fred"}
	s := supervisor{worker: worker{200, "Frank"}, level: 10}
	fmt.Println(w.info(), s.info())

	w.setName("Freddy")
	s.setName("Franky")
	s.setLevel(11)
	fmt.Println(strings.Join(print(&w, &s), "\n"))
}

func print(ps ...printer) []string {
	res := make([]string, len(ps))
	for i, p := range ps {
		res[i] = p.info()
	}
	return res
}
