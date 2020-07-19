package SecondTask

import (
	"fmt"
	"io/ioutil"
	"math"
	"math/rand"
	"sort"
	"strconv"
	"strings"
	"time"
)

// Постоянные
const MIN int = 1
const MAX int = 20

// Вспомогательные структуры данных
type void struct{}
var member void

// Раздел вспомогательных функций
func randomInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func randomSlice() []int {
	mySlice := make([]int, randomInt(MIN, MAX))
	for i := range mySlice {
		mySlice[i] = randomInt(1, 20)
	}
	return mySlice
}

func randomBigSlice() []int {
	mySlice := make([]int, randomInt(math.MaxInt32 / 10000, math.MaxInt32 / 1000))
	for i := range mySlice {
		mySlice[i] = randomInt(0, 500)
	}
	return mySlice
}

// Конец раздела вспомогательных функций

// Код заданий
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

// Slice 4: Взять последнее число slice, вернуть его пользователю, а из slice этот элемент удалить:
func slice4(s *[]int) int {
	last := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return last
}

// Slice 5: Взять первое число slice, вернуть его пользователю, а из slice этот элемент удалить
func slice5(s *[]int) int {
	first := (*s)[0]
	*s = (*s)[1:]
	return first
}

// Slice 6: Взять i-ое число slice, вернуть его пользователю, а из slice этот элемент удалить. Число i передаёт пользователь в функцию
func slice6(i int, s *[]int) int {
	element := (*s)[i]
	*s = append((*s)[:i], (*s)[i+1:]...)
	return element
}

// Slice 7: Объединить два slice и вернуть новый со всеми элементами первого и второго
func slice7(s1 *[]int, s2 *[]int) []int {
	return append(*s1, *s2...)
}

// Slice 8: Из первого slice удалить все числа, которые есть во втором
func slice8(s1 *[]int, s2 *[]int) {
	for i := range *s2 {
		currentNumber := (*s2)[i]
		k := len(*s1)
		for j := 0; j < k; j++ {
			if (*s1)[j] == currentNumber {
				*s1 = append((*s1)[:j], (*s1)[j+1:]...)
				k = len(*s1) // или k--
			}
		}
	}
}
// Slice 9: Сдвинуть все элементы slice на 1 влево. Нулевой становится последним, первый - нулевым, последний - предпоследним
func slice9(s *[]int) {
	sResult := make([]int, len(*s))
	// Нулевой элемент становится последним
	sResult[len(*s) - 1] = (*s)[0]
	// Остальные записываем со сдвигом
	for i := 0; i < len(*s) - 1; i++ {
		sResult[i] = (*s)[i + 1]
	}
	*s = sResult
}

// Slice 10: Тоже, но сдвиг на заданное пользователем i
func slice10(s *[]int, i int)  {
	for j := 0; j < i; j++ {
		slice9(s)
	}
}

// Slice 11: Тоже, что 9, но сдвиг вправо
func slice11(s *[]int) {
	sResult := make([]int, len(*s))
	// Последний элемент становится первым
	sResult[0] = (*s)[len(*s) - 1]
	// Остальные записываем со сдвигом
	for i := 1; i < len(*s); i++ {
		sResult[i] = (*s)[i - 1]
	}
	*s = sResult
}

// Slice 12: Тоже, но сдвиг на i
func slice12(s *[]int, i int)  {
	for j := 0; j < i; j++ {
		slice11(s)
	}
}

// Slice 13: Вернуть пользователю копию пераднного slice
func slice13(s *[]int) []int {
	s2 := make([]int, len(*s))
	copy(s2, *s)
	return s2
}

// Slice 14: В slice поменять все четные с ближайшими нечетными индексами. 0 и 1, 2 и 3, 4 и 5...
func slice14(s *[]int) {
	for i := 1; i < len(*s); i += 2 {
		temp := (*s)[i]
		(*s)[i] = (*s)[i-1]
		(*s)[i - 1] = temp
	}
}

// Slice 15: Упорядочить slice в порядке: прямом, обратном, лексикографическом
func slice15(s *[]int) ([]int, []int, []int) {
	sliceDirectOrder := slice13(s)
	sliceReverseOrder := slice13(s)
	sliceLexicographicOrder := slice13(s)
	sliceLexicographicOrderStrings := make([]string, len(*s))

	sort.Sort(sort.IntSlice(sliceDirectOrder))
	sort.Sort(sort.Reverse(sort.IntSlice(sliceReverseOrder)))
	// Преобразуем массив чисел в массив строк
	for i := range sliceLexicographicOrder {
		sliceLexicographicOrderStrings[i] = strconv.Itoa(sliceLexicographicOrder[i])
	}
	// Сортируем
	sort.Strings(sliceLexicographicOrderStrings)
	// И обратно
	for i := range sliceLexicographicOrderStrings {
		sliceLexicographicOrder[i], _ = strconv.Atoi(sliceLexicographicOrderStrings[i])
	}
	return sliceDirectOrder, sliceReverseOrder, sliceLexicographicOrder
}

// Map 1: Есть текст, надо посчитать сколько раз каждое слова встречается
func map1(text string) map[string]int {
	text = strings.ToLower(text)
	words := strings.Fields(text)
	resultMap := make(map[string]int)
	for i := range words {
		resultMap[words[i]]++
	}
	return resultMap
}

// Map 2: Есть очень большой массив или slice целых чисел, надо сказать какие числа в нем упоминаются хоть по разу
func map2(s []int) []int {
	set := make(map[int]void)
	for i := range s {
		set[s[i]] = member
	}
	result := make([]int, len(set))
	i := 0
	for number, _ := range set {
		result[i] = number
		i++
	}
	return result
}

// Map 3: Есть два больших массива чисел, надо найти, какие числа упоминаются в обоих
func map3(s1 []int, s2 []int) []int {
	set1 := map2(s1)
	set2 := map2(s2)
	set3 := make(map[int]void)
	for number, _ := range set1 {
		set3[number] = member
	}
	for number, _ := range set2 {
		set3[number] = member
	}
	result := make([]int, len(set3))
	i := 0
	for number, _ := range set3 {
		result[i] = number
		i++
	}
	return result
}

// Map 4: Сделать Фибоначчи с мемоизацией
func map4() func(n int) int {
	m := make(map[int]int)
	past := -1
	current := 1
	i := 0
	return func(n int) int {
		if m[n] == 0 {
			// Число не было подсчитано. Считаем
			fmt.Println("Идёт подсчёт для n =", n)
			for i <= n {
				newN := past + current
				past = current
				current = newN
				m[i] = newN
				i++
			}
		}
		return m[n]
	}
}


// Код, запускающий функции выше
func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func Main() {
	mySlice := randomSlice()
	fmt.Println("Случайно созданный срез:", mySlice)
	fmt.Println()

	fmt.Println("Slice 1: К каждому элементу []int прибавить 1:")
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
	fmt.Println("Последнее число из среза:",  slice4(&mySlice))
	fmt.Println("Получившийся срез:", mySlice)
	fmt.Println()

	fmt.Println("Slice 5: Взять первое число slice, вернуть его пользователю, а из slice этот элемент удалить:")
	fmt.Println("Первое число из среза:", slice5(&mySlice))
	fmt.Println("Получившийся срез:", mySlice)
	fmt.Println()

	fmt.Println("Slice 6: Взять i-ое число slice, вернуть его пользователю, а из slice этот элемент удалить. Число i передаёт пользователь в функцию:")
	randomNumber := randomInt(0, len(mySlice))
	fmt.Println("Случайное i-ое число:", randomNumber, "Удалён элемент из среза:", slice6(randomNumber, &mySlice))
	fmt.Println("Получившийся срез:", mySlice)
	fmt.Println()

	mySlice2 := randomSlice()
	fmt.Println("Второй срез:", mySlice2)
	fmt.Println()

	fmt.Println("Slice 7: Объединить два slice и вернуть новый со всеми элементами первого и второго:")
	fmt.Println("Первый срез:", mySlice, "Второй срез:", mySlice2)
	mySlice = slice7(&mySlice, &mySlice2)
	fmt.Println("Получившийся первый срез: ", mySlice)
	fmt.Println()

	fmt.Println("Slice 8: Из первого slice удалить все числа, которые есть во втором:")
	fmt.Println("Первый срез:", mySlice, " Второй срез:", mySlice2)
	slice8(&mySlice, &mySlice2)
	fmt.Println("Получившийся первый срез:", mySlice)
	fmt.Println()

	// Иногда так получается, что первый срез остаётся совсем пустым
	if len(mySlice) == 0 {
		mySlice = randomSlice()
	}

	fmt.Println("Slice 9: Сдвинуть все элементы slice на 1 влево. Нулевой становится последним, первый - нулевым, последний - предпоследним:")
	fmt.Println("Начальный срез:", mySlice2)
	slice9(&mySlice2)
	fmt.Println("Сдвинутый влево срез:", mySlice2)
	fmt.Println()

	fmt.Println("Slice 10: Тоже, но сдвиг на заданное пользователем i:")
	randomNumber = randomInt(MIN, MAX)
	fmt.Println("Заданный i:", randomNumber)
	fmt.Println("Начальный срез:", mySlice2)
	slice10(&mySlice2, randomNumber)
	fmt.Println("Сдвинутый влево", randomNumber, "раз срез:", mySlice2)
	fmt.Println()

	fmt.Println("Slice 11: Тоже, что 9, но сдвиг вправо:")
	fmt.Println("Начальный срез:", mySlice2)
	slice11(&mySlice2)
	fmt.Println("Сдвинутый вправо срез:", mySlice2)
	fmt.Println()

	fmt.Println("Slice 12: Тоже, но сдвиг на i:")
	randomNumber = randomInt(MIN, MAX)
	fmt.Println("Заданный i:", randomNumber)
	fmt.Println("Начальный срез:", mySlice2)
	slice12(&mySlice2, randomNumber)
	fmt.Println("Сдвинутый вправо", randomNumber, "раз срез:", mySlice2)
	fmt.Println()

	fmt.Println("Slice 13: Вернуть пользователю копию переданного slice:")
	fmt.Println("Исходный срез:", mySlice)
	sliceCopy := slice13(&mySlice)
	fmt.Println("Полученная копия:", sliceCopy)
	randomNumber = randomInt(0, len(sliceCopy))
	randomNumber2 := randomInt(MIN, MAX)
	fmt.Println("Будем менять", randomNumber, "элемент среза на число", randomNumber2)
	sliceCopy[randomNumber] = randomNumber2
	fmt.Println("Изменённая копия:", sliceCopy, "Исходный срез:", mySlice)
	fmt.Println()

	fmt.Println("Slice 14: В slice поменять все четные с ближайшими нечетными индексами. 0 и 1, 2 и 3, 4 и 5...:")
	fmt.Println("Исходный срез:", mySlice2)
	slice14(&mySlice2)
	fmt.Println("Полученный срез:", mySlice2)
	fmt.Println()

	fmt.Println("Slice 15: Упорядочить slice в порядке: прямом, обратном, лексикографическом:")
	fmt.Println("Исходный срез:", mySlice2)
	sliceDirectOrder, sliceReverseOrder, sliceLexicographicOrder := slice15(&mySlice2)
	fmt.Println("Сортировка в прямом порядке:", sliceDirectOrder)
	fmt.Println("Сортировка в обратном порядке:", sliceReverseOrder)
	fmt.Println("Сортировка в лексикографическом порядке:", sliceLexicographicOrder)
	fmt.Println()

	fmt.Println("Map 1: Есть текст, надо посчитать сколько раз каждое слово встречается:")
	fmt.Println("Текст для проверки находится в файле Text.txt")
	fmt.Println("Для сокращения результата отображаются только слова с количеством повторений больше 5000")
	text, err := ioutil.ReadFile("SecondTask/Text.txt")
	if err != nil {
		fmt.Println("Файл Text.txt не открывается!")
	}
	myMap := map1(string(text))
	for key, value := range myMap {
		if value < 5000 {
			continue
		}
		fmt.Println("Слово", key, "встречается", value, "раз.")
	}
	fmt.Println()

	bigSlice := randomBigSlice()
	fmt.Println("Map 2: Есть очень большой массив или slice целых чисел, надо сказать какие числа в нем упоминаются хоть по разу:")
	//fmt.Println("Созданный случайно срез:")
	//fmt.Println(bigSlice)
	setSlice := map2(bigSlice)
	fmt.Println("Числа, упоминающиеся в нём хоть по разу:")
	fmt.Println(setSlice)
	fmt.Println()

	bigSlice2 := randomBigSlice()
	fmt.Println("Map 3: Есть два больших массива чисел, надо найти, какие числа упоминаются в обоих:")
	setSlice = map3(bigSlice, bigSlice2)
	fmt.Println("Числа, упоминающиеся в обоих массивах:")
	fmt.Println(setSlice)
	fmt.Println()

	fmt.Println("Map 4: Сделать Фибоначчи с мемоизацией:")
	fib := map4()
	fmt.Println("N = 10:", fib(10))
	fmt.Println("N = 6:", fib(6))
}
