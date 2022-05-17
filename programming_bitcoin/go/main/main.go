package main


import (
	"blockchain/modulo/typedefine"
	"fmt"
)


func main(){
	var sample typedefine.Modulo;
	err := sample.Init(14, 19)
	if (err == nil){
		fmt.Println(sample.Num, "mod", sample.Order)
	}

	var sample2 typedefine.Modulo;
	err = sample2.Init(18, 19)
	if (err == nil){
		fmt.Println(sample2.Num, "mod", sample2.Order)
		err = sample.Add(sample2)
		fmt.Println("--- Add ---")
		if (err == nil){
			fmt.Println(sample.Num, "mod", sample.Order)
		}else{
			fmt.Println(err)
		}

		err = sample.Sub(sample2)
		fmt.Println("--- Sub ---")
		if (err == nil){
			fmt.Println(sample.Num, "mod", sample.Order)
		}else{
			fmt.Println(err)
		}

		err = sample.Mul(sample2)
		fmt.Println("--- Mul ---")
		if (err == nil){
			fmt.Println(sample.Num, "mod", sample.Order)
		}else{
			fmt.Println(err)
		}

		err = sample.Div(sample2)
		fmt.Println("--- Div ---")
		if (err == nil){
			fmt.Println(sample.Num, "mod", sample.Order)
		}else{
			fmt.Println(err)
		}

	}else{
		fmt.Println(err)
	}
}