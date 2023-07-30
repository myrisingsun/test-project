package main

import (
	"fmt"
	"reflect"
)

// выше импорт стандартных библиотек

func main() {
	message := "я тут ты там 0"
	fmt.Println(reflect.TypeOf(message))
	//reflect.TypeOf(message) - печать типа переменной message
}
