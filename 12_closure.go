package main;
import "fmt";

func main(){
	number:=10;
	squared:=func ()(int){  //This is a closure
		return number*number;
	}
	fmt.Println(squared());

	greet:=func()(){
		fmt.Println("Hello World");
	}
	greet();
}