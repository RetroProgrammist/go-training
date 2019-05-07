package helper

import (
	"fmt"
)

/*
	Инициализация при импорте пакета
	* Может быть несколько
*/
func init() { 
	fmt.Println("init 1")
}
func init() { 
	fmt.Println("init 2")
}
func init() { 
	fmt.Println("init 3")
}

/*
	Экспортируемые функции должны начинаться с Заглавной буквы 
*/
func Help() {
	notExportFunc()
}