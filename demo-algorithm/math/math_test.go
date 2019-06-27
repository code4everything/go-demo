package math

import (
	"fmt"
	"testing"
)

func TestConvertToTitle(test *testing.T) {
	fmt.Println("A\t", convertToTitle(1))
	fmt.Println("AB\t", convertToTitle(28))
	fmt.Println("ZY\t", convertToTitle(701))
	fmt.Println("AZ\t", convertToTitle(52))
}
