package main

import (
	"strconv"
	"fmt")

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.

// func (i IPAddr) String() string{
// 	return fmt.Sprintf(" %v . %v .%v .%v", i[0],i[1],i[2],i[3])
// }


func (ip IPAddr) String() string{
	 output  := ""
	for i , v:= range ip{
		if i == len(ip) -1{
			output += strconv.Itoa((int(v))) // cause v is in bytes uint8 so we need to first change it to string then to string
		}else{
			output +=  strconv.Itoa((int(v))) + "."
		}
	}
	return output
}

// func strconv(i int) {
// 	panic("unimplemented")
// }
func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
