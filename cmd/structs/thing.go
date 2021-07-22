package main

type Thing interface {
	GetThingy() int
}

type thing struct {
	thingy int
}

func NewThing(thingy int) Thing {
	t := &thing{
		thingy: thingy,
	}

	return t
}

func (t *thing) GetThingy() int {
	return t.thingy
}
