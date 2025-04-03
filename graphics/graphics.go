package graphics

const (
	WIDTH  = 64
	HEIGHT = 32
)

type Graphics struct {
	Screen [WIDTH * HEIGHT]uint32
}

func (g *Graphics) Clear() {
	for i := range g.Screen {
		g.Screen[i] = 0
	}
}
