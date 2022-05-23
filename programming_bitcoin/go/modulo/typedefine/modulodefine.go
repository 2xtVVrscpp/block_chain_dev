package typedefine

import (
	"errors"
	//"math"
	//"fmt"
)

type Modulo struct{
	Order int
	Num int
}

func minus_to_plus(num, order int) int{
	for(num < 0){
		num = num + order
	}

	return num
}

func (target *Modulo) Init(num int, order int) error{
	if (order > 0){
		target.Order = order
		if(num < 0){
			num = minus_to_plus(num, target.Order)
		}
		target.Num = num % target.Order
		return nil
	}else{
		// nothing
		return errors.New("order is not correct")
	}
}

func (target *Modulo) Copy() error{
	var modulo Modulo;
	return modulo.Init(target.Num, target.Order)
}

func Copy(target Modulo) error{
	var modulo Modulo;
	return modulo.Init(target.Num, target.Order)
}

func (target *Modulo) Add(modulo Modulo) error {
	if (target.Order == modulo.Order){
		num := target.Num + modulo.Num
		if(num < 0){
			num = minus_to_plus(num, modulo.Order)
		}
		target.Num = num % target.Order
		return nil
	}else{
			return errors.New("order doesn't match")
	}
}

func (target *Modulo) Sub(modulo Modulo) error {
	if (target.Order == modulo.Order){
		num := target.Num - modulo.Num
		if(num < 0){
			num = minus_to_plus(num, modulo.Order)
		}
		target.Num = num % target.Order
		return nil
	}else{
			return errors.New("order doesn't match")
	}
}

func (target *Modulo) Mul(modulo Modulo) error {
	if (target.Order == modulo.Order){
		num := target.Num * modulo.Num
		if(num < 0){
			num = minus_to_plus(num, modulo.Order)
		}
		target.Num = num % target.Order
		return nil
	}else{
			return errors.New("order doesn't match")
	}
}

func (target *Modulo) Div(modulo Modulo) error {
	if (target.Order == modulo.Order){
		if(Pow(&modulo, -1) == nil){
			return target.Mul(modulo)
		}else{
			return errors.New("unknown error")
		}
	}else{
			return errors.New("order doesn't match")
	}
}

func Pow(target *Modulo, exponent int) error {
	for(exponent < 0){
		exponent = exponent + (target.Order - 1)
	}

	num := target.Num
	exponent -= 1

	for ;exponent > 0;exponent -= 1 {
		target.Num = (target.Num * num) % target.Order
	}

	return nil
}