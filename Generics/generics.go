package main

import "fmt"

type IntFloat interface{
	int | float64
	
}

func sumFloatIntMaps[k comparable , V IntFloat ] (m map[k]V) V{
	var sum V
	for _ , a := range m{
		sum += a
	}

	return sum
}

func sumFloatIntSlice[V IntFloat](s []V) V{
	var sum V
	for _ , a :=range s{
		sum +=a
	}

	return sum
}

func main(){

	mint := map[string]int{
		"first" : 2,
		"second" : 3,
	}
	// fmt.Println(mint)
	// sumFloatInt(mint)
	// fmt.Println(mint)
	mfloat := map[int]float64{
		1 : 2.2,
		2 : 3.2,
	}

	// fmt.Println(mfloat)
	// sumFloatInt(mfloat)
	// fmt.Println(mfloat)

	// fmt.Printf("Generic Sums: %v and %v\n",sumFloatInt[string, int](mint),sumFloatInt[string, float64](mfloat))
	fmt.Printf("Generic Sums of maps: %v and %v\n",sumFloatIntMaps(mint),sumFloatIntMaps(mfloat))

	
	sint := []int{2,3,4,56,7,8}
	sfloat := []float64{2.2,3.4,5.6,7.8,9.8}
	fmt.Printf("Generic Sums of slices: %v and %v\n",sumFloatIntSlice(sint),sumFloatIntSlice(sfloat))
	
}

// sumFloatInt[string, int](mint)
// key value isme string and int ki hai