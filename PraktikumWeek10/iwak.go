package main
import "fmt"
func main(){
	var iwak [1000]int 
	var x,y,total,totalberat int
	fmt.Print("Masukan banyak ikan yang akan dijual : ")
	fmt.Scanln(&x)
	fmt.Print("Masukan kapasitas ikan perwadah : ")
	fmt.Scanln(&y)
	fmt.Print("Berat ikan ikan yang akan dijual : ")
	for i:=0;i<x;i++{
		fmt.Scan(&iwak[i])
	}
	fmt.Print("Berat perwadah : ")
	for i:=0;i<x;i+=y{
		for j:=i;j<i+y && j<x;j++{
			total=total+iwak[j]
		}
		totalberat=totalberat+total
		fmt.Printf("%d ",total)
		total=0
	}
	wadah:=x/y
	if x%y > 0{
		wadah=wadah+1
	}
	ratarata:=float32(totalberat)/float32(wadah)
	fmt.Printf("\nRata rata berat perwadah : %.2f",ratarata)
}