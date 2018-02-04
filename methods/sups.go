/*
- Define two user defined types *worker* and *supervisor*.
- Both types must have an ID(int) and a name(string).
- A supervisor must have a level(int) to indicate her supervisor level type ie 1, 2, 3, etc...
- Worker support 2 functions:
  1. **info** to print out the worker ID
  2. **setName** to update the worker name
- Similarly a supervisor must support info and setName, but **info** must print out the supervisor level.
- Additionally users of the system must be able to update the supervisor level.
- Make sure to write tests to ensure both types conform to the specification.

Expectations:

worker.info()             // => *[100] Fred
sup.info()                // => *[200] Frank (10)
worker.setName("Freddie") // => *[100] Freddie
sup.setLevel(11)          // => *[200] Frank (11)
sup.setName("Franky")     // => *[200] Franky (11)
*/
package methods

import (
	"fmt"
)

type (
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
	fmt.Println(w.info(), s.info())
}
