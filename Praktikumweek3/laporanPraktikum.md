# <h1 align="center">Laporan Praktikum Modul 3 </h1>
<p align="center">[Ahmad Zainul Hakim] - [109082530001]</p>

## Modul 3

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
![Screenshot Output soal 1](https://github.com/ahmad-zainul-hakim/109082530001_ahmad-zainul-hakim/blob/main/Praktikumweek3/output/soal1.png )
Ini adalah program untuk menghitung permutasian dan kombinasi. Untuk dapat menghitung hasil permutasian dan kombinasi, aku pakai fungsi faktorial, fungsi kombinasi dan fungsi permutasian.agar perhitungan buat permutasian dan kombinasi tidak ribet maka dibuatlah fungsi faktorial
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
![Screenshot Output soal 2](https://github.com/ahmad-zainul-hakim/109082530001_ahmad-zainul-hakim/blob/main/Praktikumweek3/output/soal2.png)
singkatnya kita terjemahin setiap fungsi dari bahasa matematika ke bahasa pemrograman. Terus kita gunakan pemahaman matematika untuk mengatur fungsi nya, contoh FoGoH bakal jadi F(G(H(x))). dengan merubah struktur nya maka akan berhasil dan ketemu hasilnya

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
![Screenshot Output soal 3](https://github.com/ahmad-zainul-hakim/109082530001_ahmad-zainul-hakim/blob/main/Praktikumweek3/output/soal3.png)
jadi program ini untuk menentukan apakah suatu titik ini masuk dalam lingkaran. Bagi yang bingung mungkin kujelaskan saja. Ibaratkan ada sebuah lingkaran di sebuah titik, titik itu adalah titik lingkarannya, notasi nya adalah "cx,cy" yang artinya center x & center y. Nah besar dari lingkaran itu dapat kita lihat dari radiusnya. kemudian katakanlah ada 2 lingkaran besar yang membuat suatu wilayah dimana wilayah tersebut berada termasuk dalam lingkaran 1 dan lingkaran 2. Nah dari situ akan terbagi jadi 4 wilayah. 1 wilayah yang masuk dengan 2 lingkaran. 1 wilayah dengan lingkaran 1, 1 wilayah dengan lingkaran 2 dan satu sisanya ga ikut wilayah manapun. Untuk menentukan suatu titik termasuk dalam wilayah apa. Kita menggunakan rumus (a-c)^2 + (b-d)^2 <= radius^2. Logikanya jika melebihi nilai r^2 maka titik itu ada diluar wilayah lingkaran tersebut. Dan dengan algoritma seperti itu program ini diciptakan.