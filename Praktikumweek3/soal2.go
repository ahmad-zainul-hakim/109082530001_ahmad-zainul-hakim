package main
import "fmt"

func f( x int) int{
 hasil:= x*x
 return hasil
}

func g (x int) int{
	hasil := x-2
	return hasil
}

func h (x int) int{
	hasil := x+1
	return hasil
}

func main (){
	var x1, x2, x3 int
	fmt.Scan(&x1,&x2,&x3) 
	fogog:= f(g(h(x1))) 
	gohof:= g(h(f(x2)))
	hofog:= h(f(g(x3)))

	fmt.Printf(" %d \n %d \n %d",fogog,gohof,hofog)

}