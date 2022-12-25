package main;
import (
	"fmt"
	"strconv"
);

func main(){
	var a int=10;
	var b float32=23.7;
	var c string="101";
	var d string="23.36";

	fmt.Println(float32(a));
	fmt.Println(int(b));
	fmt.Println(strconv.ParseInt(c,2,64));  //The second argument is the base of the value present in the string
	fmt.Println(strconv.ParseFloat(d,64));
}

//The above functions also return another value called error to store the value of errors;