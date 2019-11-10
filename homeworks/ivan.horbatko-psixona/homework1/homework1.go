/*
1. Описать обємні фігури: куля, конус, паралелепіпед використовуючи структури
2. реалізувать список/контейнер для фігур (додавать можна різні фігури). він повинен сортувати свій вміст
за обємом (для реалізації використовать інтерфейси)
*/
package main

import (
	"fmt"
	"math"
)

// Ball declaration (4 / 3 * math.Pi * r * r *r
type Ball struct {
	//ballName string
	r float64
}

// Cone declaration (1/3*math.Pi*r*r*h)
type Cone struct {
	//coneName string
	r, h float64
}

// Parallelepiped declaration (h * l * w)
type Parallelepiped struct {
	//parallelepipedName string
	l, w, h float64
}

var (
	m = make(map[string]float64)
)

func (b *Ball) v() float64 {
	return 4 / 3 * math.Pi * b.r * b.r * b.r
}

func (c *Cone) v() float64 {
	return 1 / 3 * math.Pi * c.r * c.r * c.h
}

func (p *Parallelepiped) v() float64 {
	return p.h * p.l * p.w
}

func main() {
	ball1 := Ball{1}
	ball2 := Ball{1.5}
	cone1 := Cone{5.0, 6.2}
	parallelepiped1 := Parallelepiped{6, 9, 15}
	m["ball1"] = ball1.v()
	m["ball2"] = ball2.v()
	m["cone1"] = cone1.v()
	m["Parallelepiped1"] = parallelepiped1.v()
	fmt.Println(m)
}
