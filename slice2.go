/*max значение в массиве*/
package main
import "fmt"
func main(){

		array := [5]int{}
		var a int
		for i:=0; i < 5; i++{
			fmt.Scan(&a)
			array[i] = a
		}

		var max int
		max = array[0]
		for _, elem := range array {
			if elem > max {
				max = elem
			}
		}

		fmt.Println(max)
}