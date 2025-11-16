package main

import (
	"fmt"
)

func main() {
	var (
		a int
		b int8
		c int16
		d int32
		e int64
		f uint
		g uint8
		h uint16
		i uint32
		j uint64
		k uintptr
		l float32
		m float64
		n complex64
		o complex128
		p bool
		q string
		r interface{}
		s []int
		t chan int
		u map[string]int
		v *int
		w func()
		x struct{}
		y struct {
			ya int
			yb bool
			yc rune
		}
	)

	fmt.Println("-----------------")
	fmt.Println("Zero values in golang")
	fmt.Println()

	fmt.Println("int:", a)
	fmt.Println("int8:", b)
	fmt.Println("int16:", c)
	fmt.Println("int32:", d)
	fmt.Println("int64:", e)
	fmt.Println("uint:", f)
	fmt.Println("uint8:", g)
	fmt.Println("uint16:", h)
	fmt.Println("uint32:", i)
	fmt.Println("uint64:", j)
	fmt.Println("uintptr:", k)
	fmt.Println("float32:", l)
	fmt.Println("float64:", m)
	fmt.Println("complex64:", n)
	fmt.Println("complex128:", o)
	fmt.Println("bool:", p)
	fmt.Println("string:", q)
	fmt.Println("interface:", r)
	fmt.Println("slice:", s)
	fmt.Println("channel:", t)
	fmt.Println("map:", u)
	fmt.Println("pointer:", v)
	fmt.Println("function:", w)
	fmt.Println("struct:", x)
	fmt.Println("struct:", y)

	fmt.Println()
	fmt.Println("End of zero values in golang")
	fmt.Println("-----------------")
}
