package provider

import "fmt"
import "strconv"

func NewBarcode() *Barcode {
	return &Barcode{Provider{}}
}

func (barcode *Barcode) EAN8() string {
	code, _ := barcode.ean(8)
	return code
}

func (barcode *Barcode) EAN13() string {
	code, _ := barcode.ean(13)
	return code
}

// Generates EAN8 and EAN12 barcodes, depending on the parameter.
// Returns an error if lenght isn't either 8 or 13
func (barcode *Barcode) ean(length int) (string, error) {

	if length != 8 && length != 13 {
		return "", fmt.Errorf("ean: length must be 8 or 13")
	}

	var code []int

	for i := 0; i < length-1; i++ {
		code = append(code, barcode.Provider.Digit())
	}

	var weights []int

	if length == 8 {
		weights = []int{3, 1, 3, 1, 3, 1, 3}
	} else if length == 13 {
		weights = []int{1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3}
	}

	weighted_sum := 0

	zipped, _ := zip(code, weights)

	for _, e := range zipped {
		weighted_sum += e[0] * e[1]
	}

	var check_digit = (10 - weighted_sum%10) % 10

	code = append(code, check_digit)

	var str string
	for _, value := range code {
		str += strconv.Itoa(value)
	}

	return str, nil
}

// zip implements python zip function, that joins two arrays with the same length into an resulting tuples array.
// zip([]int{1,2,3}, []int{4,5,6}) yields [][2]int{[1,4], [2,5], [3,6]}
// Returns an error if the lengths of the parameters isn't the same
func zip(a, b []int) ([][2]int, error) {

	if len(a) != len(b) {
		return nil, fmt.Errorf("zip: arguments must be of same length")
	}

	r := make([][2]int, len(a), len(a))

	for i, e := range a {
		r[i] = [2]int{e, b[i]}
	}

	return r, nil
}
