package leetcode

func AssertTrue(b bool) {
	AssertFalse(!b)
}

func AssertFalse(b bool) {
	if b {
		panic("assert error")
	}
}

func AssertEquals(a, b interface{}) {
	AssertTrue(a == b)
}
