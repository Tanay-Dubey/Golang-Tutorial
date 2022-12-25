package main;
import "fmt";

func operations(a int,b int)(int,int,int,int){
	sum:=a+b;
	diff:=a-b;
	pro:=a*b;
	div:=a/b;
	return sum,diff,pro,div;  //Can return multiple values in golang
}

func custom(a int,b ...int){  //'...' is ellipsis operator used to pass any number of arguments. But this operator should be the last parameter
	fmt.Println("Your number is:",a);
	sum:=0;
	for _,val:=range b{
		sum+=val;
	}
	fmt.Println("Total sum is:",sum);  
}

func main(){
	a,b,c,d:=operations(6,3);
	fmt.Println(a,b,c,d);
	custom(1);
	custom(1,2,3,4);
}