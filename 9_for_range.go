package main;
import "fmt";

func main(){
	nums:=[]int{1,2,3,4,5};
	for i,val:=range nums{
		fmt.Println("Value:",val,"at index:",i);
	}

	for _,val:=range nums{  //We use '_' to ignore any of the return values
		fmt.Println(val);
	}
}