package main
import "fmt"

func main(){
	var a,b,c,d int 
	var Pac,Pbd,Cac, Cbd int
	fmt.Scan(&a,&b,&c,&d)
	permutasi(a,c,&Pac)
	kombinasi (a,c,&Cac)
	permutasi (b,d,&Pbd)
	kombinasi (b,d,&Cbd)
	fmt.Printf("%d %d \n%d %d", Pac,Cac,Pbd,Cbd )
}

func faktorial(x int, hasil *int){
	*hasil=1
	for i:=x;i>=1;i--{
		*hasil=*hasil*i
	}
}

func permutasi(n,r int, hasil*int){
	var faktorn, faktornr int
	faktorial(n,&faktorn)
	faktorial(n-r,&faktornr)
	*hasil=faktorn/faktornr
}

func kombinasi (n,r int, hasil*int){
	var faktorn,faktorr,faktornr int
	faktorial(n,&faktorn)
	faktorial(r,&faktorr)
	faktorial(n-r,&faktornr)
	*hasil= faktorn/(faktorr*faktornr)
}