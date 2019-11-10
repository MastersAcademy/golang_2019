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
	r float64
}

//  Cone declaration (1/3*math.Pi*r*r*h)
type Cone struct {
	r, h float64
}

// Parallelepiped declaration (h * l * w)
type Parallelepiped struct {
	l, w, h float64
}

var (
	b Ball
	c Cone
	p Parallelepiped
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
