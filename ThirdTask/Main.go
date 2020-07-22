package ThirdTask

import (
	"fmt"
	"golang.org/x/tour/pic"
	"golang.org/x/tour/reader"
	"image"
	"image/color"
	"io"
	"math"
	"os"
	"sort"
	"strings"
)

// Раздел задания 1: https://tour.golang.org/ Раздел методы и интерфейсы

// Упражнение: Stringers
type IPAddr [4]byte

func stringers() {
	hosts := map[string]IPAddr {
		"loopback": {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

func (ipAddr IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ipAddr[0], ipAddr[1], ipAddr[2], ipAddr[3])
}

// Упражнение: ошибки
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func mySqrt(number float64) (float64, error) {
	if number < 0 {
		return 0, ErrNegativeSqrt(number)
	}
	const accuracy = 1e9
	var zOld float64
	z := 1.0
	for i := 0; int(zOld * accuracy) != int(z * accuracy); i++ {
		zOld = z
		z = z - (math.Pow(z, 2) - number) / (2 * z)
	}
	return z, nil
}

// Упражнение: Reader
type MyReader struct{
}

func (myReader MyReader) Read(b []byte) (int, error) {
	for i := range b {
		b[i] = 'A'
 	}
 	return len(b), nil
}

// Упражнение: rot13Reader
type rot13Reader struct {
	r io.Reader
}

func (reader rot13Reader) Read(b []byte) (int, error) {
	quantity, err := reader.r.Read(b)
	if err != nil {
		return 0, err
	}
	for i := 0; i < quantity; i++ {
		if b[i] >= 'a' && b[i] <= 'm' || b[i] >= 'A' && b[i] <= 'M' {
			b[i] = b[i] + 13
		} else if b[i] >= 'n' && b[i] <= 'z' || b[i] >= 'N' && b[i] <= 'Z' {
			b[i] = b[i] - 13
		}
	}

	return quantity, nil
}

// Упражнение: изображения
type Image struct{}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 100, 100)
}

func (img Image) At(x int, y int) color.Color {
	return color.RGBA{R: uint8((x + y) / 2), G: uint8(x + y), B: uint8(x - y), A: 1}
}

// Конец раздела задания 1

// Задание 2: Мапа с товарами. Написать методы добавления, удаления, изменения цены товара, изменения имени товара.
func AddProduct(products *map[string]float32, name string, price float32) {
	(*products)[name] = price
}

func RemoveProduct(products *map[string]float32, name string) {
	delete(*products, name)
}

func ChangePriceProduct(products *map[string]float32, name string, newPrice float32) {
	(*products)[name] = newPrice
}

func ChangeNameProduct(products *map[string]float32, oldName string, newName string) {
	price := (*products)[oldName]
	delete(*products, oldName)
	(*products)[newName] = price
}

// Задание 3: Пользователь даёт список товаров, программа должна по map с наименованиями товаров посчитать сумму заказа
func CalculateSumOrder(products *map[string]float32, productsInOrder []string) float32 {
	var sum float32
	for i := range productsInOrder {
		sum += (*products)[productsInOrder[i]]
	}
	return sum
}


// Задание 4: Сделать 1е, но у нас приходит несколько сотен таких списков заказов и мы хотим запоминать уже посчитанные заказы, чтобы если встречается такой же, то сразу говорить его цену без расчёта.
func CalculateSumOrderWithMemory() func (products *map[string]float32, productsInOrder []string) float32 {
	memory := make(map[string]float32)
	return func(products *map[string]float32, productsInOrder []string) float32 {
		var productsInOrderKey string
		for i := range productsInOrder {
			productsInOrderKey += productsInOrder[i]
		}
		if memory[productsInOrderKey] == 0 {
			fmt.Println("Идёт подсчёт суммы...")
			memory[productsInOrderKey] = CalculateSumOrder(products, productsInOrder)
		} else {
			fmt.Println("Сумма берётся из памяти...")
		}
		return memory[productsInOrderKey]
	}
}

// Задание 5: Сделать пользовательские аккаунты со счетом типа "вася: 300р, петя: 30000000р"
func makeTestAccounts(accounts *map[string]float32) {
	(*accounts)["Вася"] = 300
	(*accounts)["Петя"] = 30000000
}

// Есть map аккаунтов и счетов, как описано в 3. Надо вывести ее в отсортированном виде с сортировкой: по имени в алфавитном порядке, по имени в обратном порядке, по количеству денег по убыванию
// Заимствование из https://stackoverflow.com/a/18695740
type Pair struct {
	Key string
	Value float32
}

type PairList []Pair

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }
func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }
// Конец заимствования из https://stackoverflow.com/a/18695740

func sortMap(sortableMap *map[string]float32) {
	keysProducts := make([]string, 0)
	for k, _ := range *sortableMap {
		keysProducts = append(keysProducts, k)
	}
	sort.Strings(keysProducts)
	fmt.Println("По имени в алфавитном порядке:")
	for i := range keysProducts {
		fmt.Println(keysProducts[i], (*sortableMap)[keysProducts[i]])
	}
	fmt.Println("По имени в обратном порядке:")
	sort.Sort(sort.Reverse(sort.StringSlice(keysProducts)))
	for i := range keysProducts {
		fmt.Println(keysProducts[i], (*sortableMap)[keysProducts[i]])
	}
	pairList := make(PairList, len(*sortableMap))
	i := 0
	for key, value := range *sortableMap {
		pairList[i] = Pair{key, value}
		i++
	}
	sort.Sort(sort.Reverse(pairList))
	fmt.Println("По количеству денег по убыванию:")
	fmt.Println(pairList)
}


func Main() {
	defer fmt.Println("Конец программы!")
	fmt.Println("Задание 1: https://tour.golang.org/ Раздел методы и интерфейсы:")
	fmt.Println()

	fmt.Println("Упражнение: Stringers:")
	stringers()
	fmt.Println()

	fmt.Println("Упражнение: ошибки:")
	sqrt, err := mySqrt(2)
	fmt.Printf("Корень из 2: %v; Err = %v.\n", sqrt, err)
	sqrt, err = mySqrt(-2)
	fmt.Printf("Корень из -2: %v; Err = %v.\n", sqrt, err)
	fmt.Println()

	fmt.Println("Упражнение: Reader:")
	reader.Validate(MyReader{})
	fmt.Println()

	fmt.Println("Упражнение: rot13Reader:")
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	_, _ = io.Copy(os.Stdout, &r)
	fmt.Println()

	fmt.Println("Упражнение: изображения:")
	m := Image{}
	pic.ShowImage(m)
	fmt.Println()

	fmt.Println("---------------------------------------------------------") // Конец первого задания

	products := make(map[string]float32)
	fmt.Println("Задание 2: Мапа с товарами. Написать методы добавления, удаления, изменения цены товара, изменения имени товара:")
	fmt.Println("Добавим немного товаров...")
	AddProduct(&products, "Кукуруза вареная в/у, 450 г", 59.99)
	AddProduct(&products, "Баклажаны грунтовые, 1 кг", 49.89)
	AddProduct(&products, "Арбуз Чёрный Принц, 1 кг", 49.99)
	AddProduct(&products, "Фасоль стручковая резаная зам, 1 кг Импорт", 103.19)
	fmt.Println(products)
	fmt.Println("Удалим фасоль...")
	RemoveProduct(&products, "Фасоль стручковая резаная зам, 1 кг Импорт")
	fmt.Println(products)
	fmt.Println("Изменим цену на кукурузу...")
	ChangePriceProduct(&products, "Кукуруза вареная в/у, 450 г", 59.98)
	fmt.Println(products)
	fmt.Println("Сделаем название арбуза более политкорректным...")
	ChangeNameProduct(&products, "Арбуз Чёрный Принц, 1 кг", "Арбуз Тёмный Принц, 1 кг")
	fmt.Println(products)
	fmt.Println()

	fmt.Println("Задание 3: Пользователь даёт список товаров, программа должна по map с наименованиями товаров посчитать сумму заказа:")
	fmt.Println("Я хочу купить две кукурузы и баклажан")
	myOrder := []string{"Кукуруза вареная в/у, 450 г", "Кукуруза вареная в/у, 450 г", "Баклажаны грунтовые, 1 кг"}
	fmt.Println("Итого придётся потратить:", CalculateSumOrder(&products, myOrder))
	fmt.Println()

	fmt.Println("Задание 4: Сделать 1е, но у нас приходит несколько сотен таких списков заказов и мы хотим запоминать уже посчитанные заказы, чтобы если встречается такой же, то сразу говорить его цену без расчёта:")
	myOrder2 := []string{"Арбуз Тёмный Принц, 1 кг"}
	f := CalculateSumOrderWithMemory()
	fmt.Println("Заказываю две кукурузы и баклажан!")
	fmt.Println("С вас:", f(&products, myOrder))
	fmt.Println("Заказываю арбуз!")
	fmt.Println("С вас:", f(&products, myOrder2))
	fmt.Println("Заказываю две кукузы и баклажан!")
	fmt.Println("С вас:", f(&products, myOrder))
	fmt.Println()

	accounts := make(map[string]float32)
	fmt.Println("Задание 5: Сделать пользовательские аккаунты со счетом типа \"вася: 300р, петя: 30000000р\":")
	makeTestAccounts(&accounts)
	fmt.Println(accounts)
	fmt.Println()

	fmt.Println("Задание 6: Есть map аккаунтов и счетов, как описано в 3. Надо вывести ее в отсортированном виде с сортировкой: по имени в алфавитном порядке, по имени в обратном порядке, по количеству денег по убыванию:")
	fmt.Println("Сортируем товары:")
	sortMap(&products)
	fmt.Println()
	fmt.Println("Сортируем аккаунты:")
	sortMap(&accounts)
	fmt.Println()
}