# <h1 align="center">Laporan Praktikum Modul 2 </h1>
<p align="center">[Ahmad Zainul Hakim] - [109082530001]</p>

## Modul 2 

### 1.Soal 1
#### soal1.go
  
```go
package main
import "fmt"
func main() {
var (
satu, dua, tiga string
temp string
)
fmt.Print("Masukan input string: ")
fmt.Scanln(&satu)
fmt.Print("Masukan input string: ")
fmt.Scanln(&dua)
fmt.Print("Masukan input string: ")
fmt.Scanln(&tiga)
fmt.Println("Output awal = " + satu + " " + dua + " " + tiga)
temp = satu
satu = dua
dua = tiga
tiga = temp
fmt.Println("Output akhir = " + satu + " " + dua + " " + tiga)
}
```
### Output soal1 :
[]
[]

### Soal 2 :
```go
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

	fmt.Printf("%d \n %d \n %d",fogog,gohof,hofog)

}
```

### Output Soal 2 :
[]
[]

### Soal 3 :
```go
package main
import "fmt"

func main(){
	var cx1, cy1, radius1 float64
	var cx2, cy2, radius2 float64
	var x, y float64
	fmt.Scanln(&cx1, &cy1, &radius1)
	fmt.Scanln(&cx2, &cy2, &radius2)
	fmt.Scanln(&x, &y)

	dalamlingkaran1:=didalam(cx1, cy1,radius1,x,y)
	dalamlingkaran2:=didalam(cx2, cy2,radius2,x,y) 

	if dalamlingkaran1 && dalamlingkaran2 == true{
		fmt.Println("Titik berada di dalam lingkaran 1 dan 2")
	}else if !dalamlingkaran1 && !dalamlingkaran2{
		fmt.Println("Titik berada di luar lingkaran 1 dan 2")
	}else if dalamlingkaran1{
		fmt.Println("Titik berada di dalam lingkaran 1")
	}else if dalamlingkaran2{
		fmt.Println("Titik berada di dalam lingkaran 2")
	}
}
func hitungJarak (a, b, c, d float64) float64{
	dx := a - c
	dy := b - d
	return dx*dx + dy*dy
}

func didalam (cx,cy,radius,x,y float64) bool {
	dx := hitungJarak(cx, cy, x, y)
	return dx <= radius*radius
}
```
### Output soal 3 :
[]
[]
