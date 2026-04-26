# <h1 align="center">Laporan Praktikum Modul 9 - Array </h1>
<p align="center">[Ahmad Zainul Hakim] - [109082530001]</p>

### 1. SOAL 1
#### no1.go
  
```go
 package main
import "fmt"

	type titik struct {
	x, y int
}
type lingkaran struct {
	pusat titik
	r int
}

func main() {
	var lingkaran1, lingkaran2 lingkaran
	var t titik
	fmt.Scan(&lingkaran1.pusat.x, &lingkaran1.pusat.y, &lingkaran1.r)
	fmt.Scan(&lingkaran2.pusat.x, &lingkaran2.pusat.y, &lingkaran2.r)
	fmt.Scan(&t.x, &t.y)
	diLingkaran1 := dalamlingkaran(lingkaran1, t)
	diLingkaran2 := dalamlingkaran(lingkaran2, t)
	
	if diLingkaran1 && diLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

func dalamlingkaran(l lingkaran, t titik) bool {
	dx := t.x - l.pusat.x
	dy := t.y - l.pusat.y
	return dx*dx+dy*dy <= l.r*l.r
}
```
### Output Modul 9 no 1 :

##### Output no1.go
![Screenshot Output Modul 9 no 1](https://github.com/ahmad-zainul-hakim/109082530001_ahmad-zainul-hakim/blob/main/PraktikumWeek9/Output/no1.png)
Ini adalah Program untuk mencari suatu titik diantara 2 lingkaran. Anggap saja ada 2 lingkaran yaitu lingkaran A dan lingkaran B, kita memasukan koordinat titik pusat dan radius titik A dan B. Dan dari kedua lingkaran tersebut kita memasukan koordinat suatu titik, dan program ini akan menampilkan ada didalam lingkaran mana saja titik itu berada.

### 2. SOAL 2
#### no2.go
  
```go
package main
import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("masukan n : ")
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Print("Isi Array : ")
	fmt.Println(arr)

	fmt.Print("\nIndeks ganjil : ")
	for i := 0; i < len(arr); i++ {
		if i%2 == 1 {
			fmt.Print(arr[i], " ")
		}
	}

	fmt.Print("\nIndeks genap : ")
	for i := 0; i < len(arr); i++ {
		if i%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}

	var x int
	fmt.Print("\nMasukan X : ")
	fmt.Scan(&x)

	fmt.Print("\nIndeks kelipatan x :")
	if x != 0 {
		for i := 0; i < len(arr); i++ {
			if i%x == 0 {
				fmt.Print(arr[i], " ")
			}
		}
	} else {
		fmt.Println("x tidak boleh 0")
	}

	var hapus int
	fmt.Print("\nMasukan indeks yang ingin dihapus : ")
	fmt.Scan(&hapus)

	if hapus < 0 || hapus >= len(arr) {
		fmt.Println("Indeks tidak valid")
		return
	}

	temp := arr[:hapus]
	for i := hapus + 1; i < len(arr); i++ {
		temp = append(temp, arr[i])
	}
	arr = temp
	fmt.Println("\nArray setelah dihapus:", arr)
	
	jumlahratarata:=0
	for i:=0;i<len(arr);i++{
		jumlahratarata=jumlahratarata+arr[i]
	}
	var hasilratarata float32
	hasilratarata = float32(jumlahratarata)/float32(len(arr))
	fmt.Printf("rata rata = %.2f \n",hasilratarata)
	fmt.Print(float32(jumlahratarata),float32(len(arr)))

		var jumlah float64
	for i := 0; i < len(arr); i++ {
		selisih := float64(arr[i]) - float64(hasilratarata)
		jumlah += selisih * selisih
	}
	stdDev := math.Sqrt(jumlah / float64(len(arr)))
	fmt.Println("Standar deviasi: ", stdDev)

	var cari int
	fmt.Print("Masukkan bilangan yang ingin dicari frekuensinya: ")
	fmt.Scan(&cari)
	frekuensi := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == cari {
			frekuensi++
		}
	}
	fmt.Println("Frekuensi", cari, "=", frekuensi)
}
```
### Output Modul 9 no 2 :

##### Output no2.go
![Screenshot Output Modul 9 no 2](https://github.com/ahmad-zainul-hakim/109082530001_ahmad-zainul-hakim/blob/main/PraktikumWeek9/Output/no2.png)
Program diatas adalah program untuk Membuat array dengan panjang yang ditentukan oleh input user (n). Dan dari program itu kita bisa mengetahui nilai indeks ganjil dan genapnya apa saja, serta kia dapat menghapus salah satu indeks tersebut. Dan setelah kita menghapus salah 1 indeks tersebut, kita dapat mencari rata rata nya dan nilai standard deviasi nya. Selain itu juga, kita dapat mengetahui sebanyak apa nilai dalam suatu indeks yang muncul. 

### 3. SOAL 3
#### no3.go
```go
	package main
	import "fmt"

	func main() {
		var klubA,klubB string
		var skorA,skorB int
		var pemenang []string
		fmt.Print("Klub A : ")
		fmt.Scanln(&klubA)
		fmt.Print("Klub B : ")
		fmt.Scanln(&klubB)
		i:=1
		for {
			fmt.Printf("Pertandingan %d : ",i)
			fmt.Scanln(&skorA,&skorB)
			if skorA<0||skorB<0{
				break
			}else if skorA>skorB{
				pemenang = append(pemenang, klubA)
			}else if skorB>skorA{
				pemenang = append(pemenang, klubB)
			}else if skorA==skorB{
				pemenang = append(pemenang, "Draw")
			}
			i++
		}
		for j:=0;j<len(pemenang);j++{
			fmt.Printf("Hasil %d : %s\n",j+1,pemenang[j])
		}
		fmt.Println("Pertandingan selesai")
	}
```
### Output Modul 9 no 3 :
![Screenshot Output Modul 9 no 3](https://github.com/ahmad-zainul-hakim/109082530001_ahmad-zainul-hakim/blob/main/PraktikumWeek9/Output/no3.png)
Program diatas adalah sebuah program yang dapat merekap banyaknya pertandingan antara 2 team. Semua skor pertandingan yang telah dimainkan diinputkan dan direkap siapa saja pemenang nya. 

### 4. SOAL 4
#### no4.go
  
```go
```
### Output Modul 9 no 4 :
![Screenshot Output Modul 9 no 4](https://github.com/ahmad-zainul-hakim/109082530001_ahmad-zainul-hakim/blob/main/PraktikumWeek9/Output/no4.png)