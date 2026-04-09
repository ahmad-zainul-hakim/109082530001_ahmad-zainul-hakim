package main
import "fmt"
func print(x int){
	if x==0 {
	return
		}else{
		print(x-1)
		bintang(x)
		fmt.Println("")
	}
}
func main() {
var n int
fmt.Scan(&n)
print(n)
}
func bintang(x int){
	if x==0{
		return
	}else{
		bintang(x-1)
	fmt.Print("*")
}
}