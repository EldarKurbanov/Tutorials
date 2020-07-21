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


func Main() {
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

}