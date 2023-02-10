package main

import "fmt"

type ninjaStar struct{
	owner string
}

type ninjaWeapon interface{
	attack()
}

type ninjaSword struct{
	owner string
}

func (ninjaStar) attack(){
	fmt.Println("Throwing Ninja Star")
}

// func (ninjaStar) pickUp(){
// 	fmt.Println("Picking up Ninja Star")
// }


func (ninjaSword) attack(){
	fmt.Println("Swinging ninja sword")

}

func main(){
	fmt.Println("Diving deep into assertion...\n")
	//  TODO ->maine bnaya interface {ninjaWeapon} type ka variable weapons
	//TODO -> DUCK TYPING 
	// us interface type se main initialise kr rha hai struct ko {ninjastar}

	weapons := []ninjaWeapon{
		ninjaStar{owner: "Onkar"},
		ninjaSword{owner: "Ojha"},
	}
	//TODO ab us weapon mein main for loop chala rha and attack ko access kr rha
	// for accesing the attack method
	for _, v := range weapons{
		v.attack()

		nr , ok := v.(ninjaStar)//type assertion

		if ok{
			// ns.pickUp()

			fmt.Println(nr.owner)
		}
		// ns , ok :=v.(ninjaSword)//type assertion

		// if ok{
		// 	// ns.pickUp()

		// 	fmt.Println(ns.owner)
		// }

		
		
	}

	fmt.Println(weapons)
}