package main
import "fmt"
func main(){
	var array [1000]float32
	var n int
	fmt.Print("Masukan berapa jumlah anak kelinci yang mau ditimbang : ")
	fmt.Scan(&n)
	fmt.Print("Data berat kelinci : ")
	for i:=0;i<n;i++{
		fmt.Scan(&array[i])
		fmt.Print(" ")
	}
	min:= array[0]
	max:= array[0]
	for i:=1;i<n;i++{
		if array[i]<min{
			min=array[i]
		}
		if array[i]>max{
			max=array[i]
		}
	}
	fmt.Printf("\nBerat kelinci terkecil : %.2f\nBerat kelinci terbesar adalah : %.2f ",min,max)
}