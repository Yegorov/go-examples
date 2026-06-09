package lang

import (
	"fmt"
)

func Case(n int32) {
	switch {
	case n >= 10 && n <= 100:
		fmt.Println("10..100")
	case n > 100:
		fmt.Println("> 100")
	default:
		fmt.Print("Unknown")
	}
}

func Case2(val int32) int32 {
	var res int32
	switch val {
	case 10:
		res = val * val
	case 20, 30, 40, 50:
		res = val + 100
	}
	return res
}

func Case3(v any) string {
	switch val := v.(type) {
	case int8, int16, int32, int64:
		return fmt.Sprintf("int, val = %d", val)
	case float32, float64:
		return fmt.Sprintf("float, val = %f", val)
	default:
		return ""
	}
}

func Case4() {
	n := 10
	switch n {
	case 10:
		fmt.Print("Start")
		if n > 5 {
			break
		}
		fmt.Print("End")
	}
}

func Case5() {
	n := 10
	switch n {
	case 10:
		fmt.Print("Start")
		fallthrough
	case 20:
		fmt.Print("End")
	}
}

func Case6() {
OuterLoop:
	for i := 1; i <= 5; i++ {
		switch i {
		case 3:
			fmt.Println("3")
			break OuterLoop
		default:
			fmt.Println("i =", i)
		}
	}
	fmt.Print("End")
}
