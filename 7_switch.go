package main;
import "fmt";

func main(){
	var num int;
	fmt.Println("Enter a number between 1 to 5");
	fmt.Scanln(&num);
	switch num{  //break is implicit in golang switch
	case 1:
		fmt.Println("The number is 1");
	
	case 2:
		fmt.Println("The number is 2");

	case 3:
		fmt.Println("The number is 3");

	case 4:
		fmt.Println("The number is 4");

	case 5:
		fmt.Println("The number is 5");

	default:
		fmt.Println("The number is not in the range");
	}
	//fallthrough keyword is used to run next case block without checking its condition
}