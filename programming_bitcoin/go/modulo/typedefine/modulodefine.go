package typedefine

import (
	"errors"
//	"math"
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

func (modulo *Modulo) Init(num int, order int) error{
	if (order > 0){
		modulo.Order = order
		if(num < 0){
			num = minus_to_plus(num, modulo.Order)
		}
		modulo.Num = num % modulo.Order
		return nil
	}else{
		// nothing
		return errors.New("order is not correct")
	}
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
		num := target.Num / modulo.Num
		if(num < 0){
			num = minus_to_plus(num, modulo.Order)
		}
		target.Num = num % target.Order
		return nil
	}else{
			return errors.New("order doesn't match")
	}
}