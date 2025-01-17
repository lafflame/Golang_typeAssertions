package main //Type assertions

import "fmt"

func typeAssertion() {
	var i interface{} = "HOLA senior" //создаем интерфейс с типом данных стринг
	s := i.(string)
	fmt.Println(s) //вывод "HOLA senior"

	t := i.(string)
	fmt.Println(t) //вывод "HOLA senior"

	k := i.(int)
	fmt.Println(k) //паника: типы не соответствуют
}

func typeSwitch(i interface{}) {
	switch v := i.(type) {
	case string:
		fmt.Println("type: string ->", v)
	case int:
		fmt.Println("type: int ->", v)
	case bool:
		fmt.Println("type: bool ->", v, "\n")
	}
}
func main() {
	fmt.Println("Пример Type Switch")
	typeSwitch(12)
	typeSwitch("como estas - это 'как дела' на испанском")
	typeSwitch(true)

	fmt.Println("Пример Type Assertion")
	typeAssertion()
}
