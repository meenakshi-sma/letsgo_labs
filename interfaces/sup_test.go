package methods

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWorkerInfo(t *testing.T) {
	e := worker{id: 100, name: "Fred"}

	assert.Equal(t, "*[100] Fred", e.info())
}

func TestWorkerSetName(t *testing.T) {
	e := worker{id: 100, name: "Fred"}
	e.setName("Freddy")

	assert.Equal(t, "*[100] Freddy", e.info())
}

func TestSupInfo(t *testing.T) {
	e := supervisor{
		worker: worker{id: 200, name: "Frank"},
		level:  10,
	}

	assert.Equal(t, "*[200] Frank (10)", e.info())
}

func TestSupSetName(t *testing.T) {
	e := supervisor{
		worker: worker{id: 200, name: "Frank"},
		level:  10,
	}
	e.setName("Franky")

	assert.Equal(t, "*[200] Franky (10)", e.info())
}

func TestSupSetLevel(t *testing.T) {
	e := supervisor{
		worker: worker{id: 200, name: "Frank"},
		level:  10,
	}
	e.setLevel(11)

	assert.Equal(t, "*[200] Frank (11)", e.info())
}

func TestWorkerPrint(t *testing.T) {
	e := worker{id: 100, name: "Fred"}

	assert.Equal(t, []string{"*[100] Fred"}, print(&e))
}

func TestSupPrint(t *testing.T) {
	e := supervisor{
		worker: worker{id: 200, name: "Frank"},
		level:  10,
	}

	assert.Equal(t, []string{"*[200] Frank (10)"}, print(&e))
}
