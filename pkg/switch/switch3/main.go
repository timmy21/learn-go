package main

import (
	"fmt"
	"strconv"
)

// 类型 switch 用于进行一系列的类型断言
func ToInt(i any) (int, error) {
	switch v := i.(type) {
	case int8:
		return int(v), nil
	case uint8:
		return int(v), nil
	case int16:
		return int(v), nil
	case uint16:
		return int(v), nil
	case int32:
		return int(v), nil
	case uint32:
		return int(v), nil
	case int64:
		return int(v), nil
	case uint64:
		return int(v), nil
	case string:
		n, err := strconv.ParseInt(v, 10, 0)
		if err != nil {
			return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
		}
		return int(n), nil
	default:
		return 0, fmt.Errorf("unable to cast %#v of type %T to int", i, i)
	}
}

func main() {
	fmt.Println(ToInt("1"))
}
