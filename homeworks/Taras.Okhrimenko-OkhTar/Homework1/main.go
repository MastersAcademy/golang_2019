package main

import (
	"fmt"
	"math"
	"sort"
)

// Figures interface for volume
type Figures interface {
	volume() float64
}

// Sphere struct create
type Sphere struct {
	radius float64
}

func (s *Sphere) volume() float64 {
	return (4.0 * math.Pi * math.Pow(s.radius, 3)) / 3
}

// Сone struct create
type Сone struct {
	radius, height float64
}

func (c *Сone) volume() float64 {
	return (math.Pi * math.Pow(c.radius, 2) * c.height) / 3
}

// Parallelepiped struct create
type Parallelepiped struct {
	a, b, height float64
}

func (p *Parallelepiped) volume() float64 {
	return p.a * p.b * p.height
}

// ArrFig type create
type ArrFig []Figures

func (a ArrFig) Len() int           { return len(a) }
func (a ArrFig) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ArrFig) Less(i, j int) bool { return a[i].volume() < a[j].volume() }

// OutRes output array with figure volume
func OutRes(arr ArrFig) {
	for _, el := range arr {
		fmt.Printf("%.4v \t||| %v (%T)\n", el.volume(), el, el)
	}
	fmt.Printf("\n")
}

func main() {
	sph1 := &Sphere{9.3}
	sph2 := &Sphere{3.4}
	cone1 := &Сone{7.3, 2.2}
	cone2 := &Сone{3, 8.1}
	paral1 := &Parallelepiped{5, 3.1, 3.3}
	paral2 := &Parallelepiped{3.9, 3.7, 4}
	paral3 := &Parallelepiped{0.4, 6.4, 5}

	arr := make(ArrFig, 0, 7)

	arr = append(arr, sph1, sph2, cone1, cone2, paral1, paral2, paral3)

	fmt.Println("Input data (before sort):")
	OutRes(arr)

	sort.Sort(arr)

	fmt.Println("Sorted figure volume values:")
	OutRes(arr)
}
