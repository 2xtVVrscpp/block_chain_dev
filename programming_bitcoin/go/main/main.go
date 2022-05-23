package main


import (
	"blockchain/modulo/typedefine"
	"fmt"
)


func main(){
	var sample typedefine.Modulo;
	err := sample.Init(3, 19)
	if (err == nil){
		fmt.Println(sample.Num, "mod", sample.Order)
	}

	var sample2 typedefine.Modulo;
	err = sample2.Init(7, 19)
	if (err == nil){
		fmt.Println(sample2.Num, "mod", sample2.Order)
	}

	target := sample
	modulo := sample2
	err = target.Add(modulo)
	fmt.Println("--- Add ---")
	if (err == nil){
		fmt.Println(target.Num, "mod", target.Order)
	}else{
		fmt.Println(err)
	}

	target = sample
	modulo = sample2
	err = target.Sub(modulo)
	fmt.Println("--- Sub ---")
	if (err == nil){
		fmt.Println(target.Num, "mod", target.Order)
	}else{
		fmt.Println(err)
	}

	target = sample
	modulo = sample2
	err = target.Mul(modulo)
	fmt.Println("--- Mul ---")
	if (err == nil){
		fmt.Println(target.Num, "mod", target.Order)
	}else{
		fmt.Println(err)
	}

	target = sample
	modulo = sample2
	err = target.Div(modulo)
	fmt.Println("--- Div ---")
	if (err == nil){
		fmt.Println(target.Num, "mod", target.Order)
	}else{
		fmt.Println(err)
	}
}