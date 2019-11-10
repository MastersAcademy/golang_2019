/*
1. Описать обємні фігури: куля, конус, паралелепіпед використовуючи структури
2. реалізувать список/контейнер для фігур (додавать можна різні фігури). він повинен сортувати свій вміст
за обємом (для реалізації використовать інтерфейси)
*/
package main

import (
	"fmt"
	"math"
	"sort"
)

// Ball declaration (4 / 3 * math.Pi * r * r *r
type Ball struct {
	ballName string
	r        float64
}

// Cone declaration (1/3*math.Pi*r*r*h)
type Cone struct {
	coneName string
	r, h     float64
}

// Parallelepiped declaration (h * l * w)
type Parallelepiped struct {
	parallelepipedName string
	l, w, h            float64
}

var (
	m = make(map[string]float64)
)

func (b *Ball) v() float64 {
	return 4.0 / 3.0 * math.Pi * b.r * b.r * b.r
}

func (c *Cone) v() float64 {
	return 1.0 / 3.0 * math.Pi * c.r * c.r * c.h
}

func (p *Parallelepiped) v() float64 {
	return p.h * p.l * p.w
}

func main() {
	ball1 := Ball{"ball1", 4.6}
	ball2 := Ball{"ball2", 5.7}
	cone1 := Cone{"cone1", 5.0, 6.2}
	cone2 := Cone{"cone2", 3.3, 8.8}
	parallelepiped1 := Parallelepiped{"parallelepiped1", 6.3, 9.2, 15}
	parallelepiped2 := Parallelepiped{"parallelepiped2", 8, 3, 4.4}
	m[ball1.ballName] = ball1.v()
	m[ball2.ballName] = ball2.v()
	m[cone1.coneName] = cone1.v()
	m[cone2.coneName] = cone2.v()
	m[parallelepiped1.parallelepipedName] = parallelepiped1.v()
	m[parallelepiped2.parallelepipedName] = parallelepiped2.v()
	fmt.Println(m)
	type sSort struct {
		Key   string
		Value float64
	}

	var ss []sSort
	for k, v := range m {
		ss = append(ss, sSort{k, v})
		//fmt.Println(ss)
	}

	sort.Slice(ss, func(i, j int) bool {
		//fmt.Println(ss[i].Value, ss[j].Value, i , j)
		return ss[i].Value > ss[j].Value
	})

	for _, sSort := range ss {
		fmt.Printf("%s, %v\n", sSort.Key, sSort.Value)
	}
}
