package main

import "fmt"

type CurrentAccount struct {
	holder        string
	agencyNumber  int
	accountNumber int
	balance       float64
}

func main() {
	accountJohn := CurrentAccount{"John", 589, 123456, 125.50}
	accountElon := CurrentAccount{"Elon", 678, 192345, 10234.50}

	fmt.Println(accountJohn)
	fmt.Println(accountElon)

}
