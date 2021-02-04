package condition_test

import (
	"runtime"
	"testing"
)

func TestIfMultiSec(t *testing.T) {

	if a := 1 == 1; a {
		t.Log(a)
	}

	// if v, err := someFunc(); err == nil {
	// 	t.Log(v)
	// } else {
	// 	t.Log(err)
	// }
}

func TestSwitchMultiCase(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch i {
		case 0, 2:
			t.Log("Even")
		case 1, 3:
			t.Log("Odd")
		default:
			t.Log("it is not 0-3")
		}
	}

	switch os := runtime.GOOS; os {
	case "darwin":
		t.Log("OS x")
	case "linux":
		t.Log("linux")
	default:
		t.Logf("%s.", os)
	}
}

func TestSwitchCaseCondition(t *testing.T) {
	for i := 0; i < 5; i++ {
		switch {
		case i%2 == 0:
			t.Log("Even")
		case i%2 == 1:
			t.Log("Odd")
		default:
			t.Log("unkonw")
		}
	}
}
