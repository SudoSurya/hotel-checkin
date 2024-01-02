package utils

import (
	"fmt"
	"strconv"
	"time"
)



func RandomString() string {
	return fmt.Sprintf("%d", time.Now().Nanosecond())
}

func Str_int(str string) int{
    val,err := strconv.Atoi(str)
    if err != nil {
        fmt.Println(err)
    }
    return val
}
