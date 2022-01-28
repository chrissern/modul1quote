package main

import (
	"fmt"

	"github.com/chrissern/modul1quote/quotes"
)

func main() {
	fmt.Println(quotes.GetGo())
	fmt.Println(quotes.GetGlass())
	fmt.Println(quotes.GetOpt())
	fmt.Println(quotes.GetHello())
}
