package main;
import "fmt";

func main(){
	var arr1=[5]int{0,1,2,3,4};
	arr2:=[3]string{"hello","hi","world"};
	arr3:=[...]int{12,3,55,6,23,63};  //Size is automatically inferred

	fmt.Println(arr1,arr2,arr3);

	arr4:=[5]int{};
	fmt.Println(arr4);  //Printed with default values
	fmt.Println(arr1[2]);  //Accessing the elements of array

	fmt.Println(len(arr2));
}