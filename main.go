package main
import "fmt"
func main(){
   fmt.Println(fib(2))
}

func fib(a int)int {
 if a==0 || a==1{
   return 1
}
 return fib(a-1)+ fib(a-2)
}

func min(a,b int)int {
if a<b{
return a}
return b 
}
