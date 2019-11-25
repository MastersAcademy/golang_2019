package main

import(
    "fmt"
    "sort"
    "math"
)

type Sphere struct {
    radius float64
}

type Parallelepiped struct {
    width, height, depth float64
}

type Cone struct {
    radius, height float64
}

type Figure interface {
    Volume() float64
}

type FigureContainer struct {
    figures []Figure
}

func (s *Sphere) Volume() float64 {
    return (4.0 / 3.0) * math.Pi * math.Pow(s.radius, 3)
}

func (p *Parallelepiped) Volume() float64 {
    return p.width * p.height * p.depth;
}

func (c *Cone) Volume() float64 {
    return math.Pi * math.Pow(c.radius, 2) * c.height / 3.0
}

// sorts figures in this container in ascending order of their volume
func (container *FigureContainer) Sort() {
    sort.Slice(container.figures, func (i, j int) bool {
        return container.figures[i].Volume() < container.figures[j].Volume()
    })
}

// Add a new figure to this container
func (container *FigureContainer) Add(figure Figure) {
    container.figures = append(container.figures, figure)
}

// get figures in this container
// This returns a copy that can be safely modified
func (container *FigureContainer) GetFigures() []Figure {
    out := make([]Figure, len(container.figures))
    copy(out, container.figures)
    return out
}

func main() {
    c := &FigureContainer{}
    c.Add(&Parallelepiped{1, 1, 2})
    c.Add(&Parallelepiped{1, 1, 1})
    c.Add(&Parallelepiped{1, 1, 3})
    c.Add(&Cone{1, 1})
    c.Add(&Sphere{1})
    c.Sort()
    for _, fig := range c.GetFigures() {
        fmt.Printf("Vol %f\n", fig.Volume())
    }
}
