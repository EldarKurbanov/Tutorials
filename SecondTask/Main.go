package SecondTask

import "fmt"

// Slice 1: К каждому элементу []int прибавить 1
func slice1(s *[]int) {
	for i := range *s {
		(*s)[i]++
	}
}

// Slice 2: Добавить в конец slice число 5
func slice2(s *[]int) {
	*s = append(*s, 5)
}

// Slice 3: Добавить в начало slice число 5
func slice3(addableNumber int, s *[]int) {
	*s = append(*s, 0)
	copy((*s)[1:], *s)
	(*s)[0] = addableNumber
}

func slice4(s *[]int) int {
	last := (*s)[len(*s) - 1]
	*s = (*s)[:len(*s) - 1]
	return last
}

// Slice 5: Взять первое число slice, вернуть его пользователю, а из slice этот элемент удалить
func slice5(s *[]int) int {
	first := (*s)[0]
	*s = (*s)[1:]
	return first
}

// Slice 6: Взять i-ое число slice, вернуть его пользователю, а из slice этот элемент удалить. Число i передаёт пользователь в функцию
func slice6(i int, s*[]int) int {
	element := (*s)[i]
	*s = append((*s)[:i], (*s)[i+1:]...)
	return element
}

func Main()  {
	fmt.Println("Slice 1: К каждому элементу []int прибавить 1:")
	mySlice := []int{1,2,3,4,5,6,7,8,9}
	slice1(&mySlice)
	fmt.Println(mySlice)
	fmt.Println()

	fmt.Println("Slice 2: Добавить в конец slice число 5:")
	slice2(&mySlice)
	fmt.Println(mySlice)
	fmt.Println()

	fmt.Println("Slice 3: Добавить в начало slice число 5:")
	slice3(5, &mySlice)
	fmt.Println(mySlice)
	fmt.Println()

	fmt.Println("Slice 4: Взять последнее число slice, вернуть его пользователю, а из slice этот элемент удалить:")
	fmt.Println(slice4(&mySlice))
	fmt.Println(mySlice)
	fmt.Println()

	fmt.Println("Slice 5: Взять первое число slice, вернуть его пользователю, а из slice этот элемент удалить:")
	fmt.Println(slice5(&mySlice))
	fmt.Println(mySlice)
	fmt.Println()

	fmt.Println("Slice 6: Взять i-ое число slice, вернуть его пользователю, а из slice этот элемент удалить. Число i передаёт пользователь в функцию:")
	fmt.Println(slice6(3, &mySlice))
	fmt.Println(mySlice)
	fmt.Println()
}