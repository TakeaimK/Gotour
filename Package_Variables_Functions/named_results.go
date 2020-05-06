package main

import "fmt"

func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}

func money(total int)(a,b,c int){
	a = total/1000
	total %= 1000
	b = total/100
	total %= 100
	c = total/10
	return
}

func main() {
	fmt.Println(split(17))

	var temp int
	fmt.Print("만원 이하의 돈을 입력하세요")
	//a,b,c := money(fmt.Scanln(&temp)) => error(too many arguments)
	fmt.Scanln(&temp)
	a,b,c := money(temp)
	//fmt.Print("천원 : %d장, 백원 : %d개, 십원 : %d개 입니다.", a,b,c)
	fmt.Printf("천원 : %d장, 백원 : %d개, 십원 : %d개 입니다.", a,b,c)
}
