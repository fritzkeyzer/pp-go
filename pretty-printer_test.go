package pp_go

import (
	"testing"
)

func TestDump(t *testing.T) {

	type someDumbStruct struct {
		strVal string
		intVal int
		flVal  float64
	}

	s := someDumbStruct{
		strVal: "blah blah blah",
		intVal: 1337,
		flVal:  42.0,
	}

	Dump(&s)

	type anotherDumbStruct struct {
		intVal int
		vals   []int
		strVal []someDumbStruct
	}

	strVals := []someDumbStruct{
		{
			strVal: "a",
			intVal: 0,
			flVal:  0,
		},
		{
			strVal: "b",
			intVal: 1,
			flVal:  1,
		},
		{
			strVal: "c",
			intVal: 2,
			flVal:  2,
		},
	}

	a := anotherDumbStruct{
		intVal: 6969,
		vals:   []int{0, 1, 2, 3},
		strVal: strVals,
	}

	Dump(&a)

}
