package main
import "fmt"
func main(){
	// numbers := [...]int{21,23,27,30,33}

	// for index := range numbers{
	// 	fmt.Println(numbers[index])


	// }

	// map

	// subjectMarks := map[string]float32{"Java":80 , "Python": 81 , "GoLang" : 85}
	// fmt.Println("Marks Obtained")

	// //use range to iterate through key-value pair
	// for subject := range subjectMarks{
	// 	// fmt.Println(subject , ":" , subjectMarks[subject])
	// 	fmt.Println(subject)
	
	// }



	// break and continue

	for i:= 1;i<=3;i++{
		for j:=1 ; j<=3;j++{
			if i==3{
				continue
			}

			fmt.Println("i= " , i , "j=",j)
		}
	}
}