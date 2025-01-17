package main //Type assertions

import "fmt"

func main() {
	var i interface{} = "HOLA senior" //создаем интерфейс с типом данных стринг
	s := i.(string)
	fmt.Println(s) //вывод "HOLA senior"

	t := i.(string)
	fmt.Println(t) //вывод "HOLA senior"

	k := i.(int)
	fmt.Println(k) //паника, потому что типы не соответствуют
}
