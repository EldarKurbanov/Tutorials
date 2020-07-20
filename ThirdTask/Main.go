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

	fmt.Println()

}