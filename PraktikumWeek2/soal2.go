package main
import "fmt"
func main() {
	var m,k,h,u string 
	var percobaanbenar int
	var Berhasil bool
	for i:=1;i<=5;i++{
		fmt.Printf("Percobaan %d : ",i)
		fmt.Scan(&m,&k,&h,&u)
		if m=="merah" && k=="kuning" && h=="hijau" && u=="ungu"{
			percobaanbenar=percobaanbenar+1
		}
	}
	if percobaanbenar==5{
		Berhasil=true
	}else{
		Berhasil=false
	}	
	fmt.Printf("Berhasil = %t ",Berhasil)
}