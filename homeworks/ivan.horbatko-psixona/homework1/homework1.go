/*
1. Описать обємні фігури: куля, конус, паралелепіпед використовуючи структури
2. реалізувать список/контейнер для фігур (додавать можна різні фігури). він повинен сортувати свій вміст
за обємом (для реалізації використовать інтерфейси)
*/
//package main
//
//import (
//	"fmt"
//	"math"
//)
//
//// Ball declaration (4 / 3 * math.Pi * r * r *r
//type Ball struct {
//	r float64
//}
//
//// Cone declaration (1/3*math.Pi*r*r*h)
//type Cone struct {
//	r, h float64
//}
//
//// Parallelepiped declaration (h * l * w)
//type Parallelepiped struct {
//	l, w, h float64
//}
//
//func (b *Ball) V() float64 {
//	return 4.0 / 3.0 * math.Pi * b.r * b.r * b.r
//}
//
//func (c *Cone) V() float64 {
//	return 1.0 / 3.0 * math.Pi * c.r * c.r * c.h
//}
//
//func (p *Parallelepiped) V() float64 {
//	return p.h * p.l * p.w
//}
//
//var (
//	m = make(map[string]interface{})
//	mList = make(map[string]float64)
//)
//
//type Messier interface {
//	V() float64
//}
//
//type List []Messier
//
//
//func (s List) Len() int {
//	return len(s)
//}
//func (s List) Swap(i, j int) {
//	s[i], s[j] = s[j], s[i]
//}
//func (s List) Less(i, j int) bool {
//	return s[i].V() < s[j].V()
//}
//
//func main() {
//	m["Ball"] = &Ball{r: 4}
//	m["Ball2"] = &Ball{r: 2}
//	m["Cone"] = &Cone{r: 5, h: 6}
//	m["Parallelepiped"] = &Parallelepiped{8, 3, 4.2}
//	m["Parallelepiped2"] = &Parallelepiped{6, 4, 4.8}
//	m["Cone2"] = &Cone{r: 6, h: 5}
//	for key, value := range m {
//		strKey := fmt.Sprintf("%v", key)
//		strValue, _ := fmt.Printf("%v", value)
//		mList[strKey] = float64(strValue)
//	}
//	fmt.Println(m)
//	fmt.Printf("%#v", mList)
//}

package main

import (
	"fmt"
	"math"
	"sort"
)

// Ball declaration (4 / 3 * math.Pi * r * r *r
type Ball struct {
	r float64
}

// Cone declaration (1/3*math.Pi*r*r*h)
type Cone struct {
	r, h float64
}

// Parallelepiped declaration (h * l * w)
type Parallelepiped struct {
	l, w, h float64
}

func (b *Ball) V() float64 {
	return 4.0 / 3.0 * math.Pi * b.r * b.r * b.r
}

func (c *Cone) V() float64 {
	return 1.0 / 3.0 * math.Pi * c.r * c.r * c.h
}

func (p *Parallelepiped) V() float64 {
	return p.h * p.l * p.w
}

var m = make(List, 0, 6)

type Messier interface {
	V() float64
}

type List []Messier

func receiver(shapes List) {
	for _, s := range shapes {
		fmt.Printf("Shape: %#v = %v\n", s, s.V())
	}
}

func (s List) Len() int {
	return len(s)
}
func (s List) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
func (s List) Less(i, j int) bool {
	return s[i].V() < s[j].V()
}

func main() {
	m = append(m, &Ball{10},
		&Cone{5, 8},
		&Parallelepiped{2, 3, 3},
		&Ball{6},
		&Cone{2.2, 8},
		&Parallelepiped{3.5, 2.2, 4.4})
	sort.Sort(m)
	receiver(m)
}
