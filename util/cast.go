package util

import (
	"fmt"
	"strconv"
)

func ToInt(arg interface{}) int {
	var res int

	switch arg.(type) {
	case string:
		var err error

		res, err = strconv.Atoi(arg.(string))

		if err != nil {
			panic("error converting string to integer" + err.Error())
		}
	default:
		panic(fmt.Sprintf("unhandled type for"))
	}

	return res
}
