package main

import "fmt"

func main(){

	interable := 1

	for interable <= 100 {

		if interable%3 == 0 && interable%5 == 0{
			fmt.Println("Pin Pan")
			interable ++
			continue
		}else if interable%5 == 0{
			fmt.Println("Pan")
			interable ++
			continue
		}else if interable%3 == 0{
			fmt.Println("Pin")
			interable ++
			continue
		}

		fmt.Println(interable)
		interable ++

	}

}