package main
import "fmt"

func permutasi(n, r int) int {
hasil:= faktorial(n)/faktorial(n-r)
return hasil
}

func faktorial(n int) int {
	hasil:=1
	if n==0{
		hasil=1
	}else if n>0{
		for n=n;n>0;n--{
			hasil=hasil*n
		}
	}
	return hasil
}

func kombinasi (n, r int) int {
	hasil:=faktorial(n)/(faktorial(r)*faktorial(n-r)) 
	return hasil
}

func main() {
var a,b,c,d int
fmt.Scan(&a,&b,&c,&d)
permutasiac:= (permutasi(a,c))
kombinasiac:= (kombinasi(a,c))
permutasibd:= (permutasi(b,d))
kombinasibd:= (kombinasi(b,d))
fmt.Printf(" %d %d\n%d %d", permutasiac, kombinasiac, permutasibd, kombinasibd)
}