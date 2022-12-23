package main;
import "fmt";

func main(){
	//Declared sinble variables
	var name string="This is a string";
	var num int=23;
	var value float32=23.54;
	fmt.Print(name,num,value);  //Prints without spaces
	fmt.Println(name,num,value);  //Prints with spaces	

	var a=23;  //Automatically inferred by compiler
	b:="Hello World";  //Automatically inferred by compiler. This operator cannot be used outside a function
	fmt.Println(a,b);

	//Multiple varialbes declaration
	
}