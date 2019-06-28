package math

import (
	"testing"

	. "../leetcode"
)

func TestFindNthDigit(t *testing.T) {
	AssertEquals(findNthDigit(10), 1)
	AssertEquals(findNthDigit(11), 0)
	AssertEquals(findNthDigit(12), 1)
	AssertEquals(findNthDigit(13), 1)
	AssertEquals(findNthDigit(14), 1)
	AssertEquals(findNthDigit(15), 2)
	AssertEquals(findNthDigit(8898998), 0)
	AssertEquals(findNthDigit(872012), 5)
	AssertEquals(findNthDigit(19), 4)
	AssertEquals(findNthDigit(99), 4)
	AssertEquals(findNthDigit(17), 3)
	AssertEquals(findNthDigit(999), 9)
}

func TestIsPowerOfFour(t *testing.T) {
	AssertFalse(isPowerOfFour(3))
	AssertTrue(isPowerOfFour(1 << 2))
}

func TestConvertToTitle(test *testing.T) {
	AssertEquals(convertToTitle(1), "A")
	AssertEquals(convertToTitle(28), "AB")
	AssertEquals(convertToTitle(701), "ZY")
	AssertEquals(convertToTitle(52), "AZ")
}
