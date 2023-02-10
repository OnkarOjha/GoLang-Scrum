package main

import "fmt"

type MotorVehicle interface{
	Mileage() float64
}

type BMW struct{
	distance float64
	fuel float64
	averagespeed float64
}

type Audi struct{
	distance float64
	fuel float64
}
// mileage function defined for both the struct type
func (b BMW) Mileage() float64{
	return b.distance/b.fuel
}

func (a Audi) Mileage() float64{
	return a.distance/a.fuel
}
// to calculate total mileage
func totalMileage(m []MotorVehicle){
	tm := 0.0
	for _,v := range m{
		tm = tm + v.Mileage()
	}

	fmt.Printf("Totoal mileage per month %f km/l",tm)
}


func main(){
	fmt.Println("working with interface")

	b1 := BMW{
		distance: 167.9,
		fuel: 36,
		averagespeed: 56,
	}

	a1:= Audi{
		distance: 123.9,
		fuel: 45,
		//averagespeed: 67,
	}
	person := []MotorVehicle{b1, a1}
	totalMileage(person)
	// totalMileage()
}