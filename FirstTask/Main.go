package FirstTask

import (
	"fmt"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/wc"
	"math"
	"strings"
)

//Упражнение: циклы и функции
func mySqrt(number float64) float64 {
	const accuracy = 1e9
	var zOld float64
	z := 1.0
	for i := 0; int(zOld * accuracy) != int(z * accuracy); i++ {
		zOld = z
		z = z - (math.Pow(z, 2) - number) / (2 * z)
	}
	return z
}

//Упражнение: срезы
func Pic(dx, dy int) [][]uint8 {
	picture := make([][]uint8, dy)
	for i := range picture {
		picture[i] = make([]uint8, dx)
		for j := range picture[i] {
			picture[i][j] = uint8((i + j) / 2)
		}
	}
	return picture
}

//Упражнение: карты
func WordCount(s string) map[string]int {
	array := strings.Fields(s)
	m := make(map[string]int)
	for i := range array {
		m[array[i]]++
	}
	return m
}

//Упражнение: замыкание Фибоначчи
func fibonacci() func() int {
	past, current := -1, 1
	return func() int {
		newN := past + current
		past = current
		current = newN
		return newN
	}
}

func Main() {
	//Упражнение: циклы и функции
	fmt.Println(math.Sqrt(2))
	fmt.Println(mySqrt(2))
	//Упражнение: срезы
	pic.Show(Pic)
	//Упражнение: карты
	wc.Test(WordCount)
	//Упражнение: замыкание Фибоначчи
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}