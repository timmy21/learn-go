package main

import (
	"fmt"
	"strconv"
)

// 类型 switch 用于进行一系列的类型断言
func ToInt(i any) (int, error) {
	switch v := i.(type) {
	case int8:
		// v 类似为 int8
		return int(v), nil
	case uint8:
		// v 类似为 uint8
		return int(v), nil
	case int16:
		// v 类似为 int16
		return int(v), nil
	case uint16:
		// v 类似为 uint16
		return int(v), nil
	case int32:
		// v 类似为 int32
		return int(v), nil
	case uint32:
		// v 类似为 uint32
		return int(v), nil
	case int64:
		// v 类似为 int64
		return int(v), nil
	case uint64:
		// v 类似为 uint64
		return int(v), nil
	case string:
		// v 类型为 string
		n, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
		}
		return int(n), nil
	default:
		// v 类型为 interface{}
		return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
	}
}

func main() {
	fmt.Println(ToInt("1"))
}
