package main;
import "fmt";

func main(){
	var i=15.5;
	var s="Hello";

	fmt.Printf("i:%v and s:%v\n",i,s);  //Gives value
	fmt.Printf("i:%#v and s:%#v\n",i,s);  //Prints values in Go syntax format
	fmt.Printf("i:%T and s:%T\n",i,s);  //Prints type of the variables
	fmt.Printf("%%\n")  //Prints the "%" sign

	var num=15;
	fmt.Printf("Value of num=%b\n",num)  //Base 2
	fmt.Printf("Value of num=%o\n",num)  //Base 8
	fmt.Printf("Value of num=%x\n",num)  //Base 16
}