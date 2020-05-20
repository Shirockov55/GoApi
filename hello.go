package main

import "fmt"

func main(){
	fmt.Println("Hello world")

	var web string = "hip " + "\nhop!"
	var len = len(web);
	fmt.Println(web, len)

	var num float64 = 6.456546456;
	fmt.Printf("%.2f \n", num)
	fmt.Printf("%T \n", num)

	if num < 5 {
		fmt.Println("мало")
	} else if (num > 5) || num == 5 {
		fmt.Println("много")
	}

	switch len {
	case 9:
		fmt.Println("в точку")
	}

	var i = 0
	for i <= 10 {
		//fmt.Println(i)
		i++
	}

	for i:= 0; i <= 5; i++ {
		fmt.Println(i)
	}

	var arr[3] int
	arr[0] = 45
	arr[1] = 55
	arr[2] = 65
	fmt.Println(arr[1])

	nums := [3]int { 1,2,4 }
	for i, value := range nums {
		fmt.Println(i, value)
	}

	// Типа ассоциативных массивов
	websites := make(map[string]int)
	websites["mySite"] = 56
	fmt.Println(websites["mySite"])

	// замыкание
	multi := func() int{
		websites["mySite"] *= 2
		return websites["mySite"]
	}

	// Откладывание до конца функции main
	defer fmt.Println(666)

	fmt.Println(multi())

	var r1 int
	var r2 int
	r1, r2 = summ(6, 8)
	fmt.Println(r1, r2)
	// Указатели
	var x int = 0
	pointer(&x)
	fmt.Println(x)

	bob:= Cats{ "Bob", 32 }
	fmt.Println(bob.name)
	fmt.Println(bob.test())
}

func summ(num_1 int, num_2 int) (int, int) {
	var res1 int
	var res2 int
	res1 = num_1 + num_2
	res2 = num_1 - num_2
	return res1, res2
}

// функция с указателем
func pointer(x *int){
	*x = 2
}

// Структуры
type Cats struct {
	name string
	age int
}

// Типа метода для структуры
func (cat *Cats) test() float64 {
	return float64(cat.age) * 2.6767
}
