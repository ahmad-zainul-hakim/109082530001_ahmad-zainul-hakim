# <h1 align="center">Laporan Praktikum Modul 10</h1>
<p align="center">[Ahmad Zainul Hakim] - [109082530001]</p>

## Modul 10 

### 1. [Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak kelinci yang akan dijual. Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual. Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan terbesar.]

#### soal1.go
```go
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
```
##### Output soal 1
![Screenshot Output soal 1](https://github.com/ahmad-zainul-hakim/109082530001_ahmad-zainul-hakim/blob/main/PraktikumWeek10/iwak.png)
[penjelasan]

### 2. [Sebuah program digunakan untuk menentukan tarif ikan yang akan dijual ke pasar. Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat ikan yang akan dijual. Masukan terdiri dari dua baris, yang mana baris pertama terdiri dari dua bilangan bulat x dan y. Bilangan x menyatakan banyaknya ikan yang akan dijual, sedangkan y adalah banyaknya ikan yang akan dimasukan ke dalam wadah. Baris kedua terdiri dari sejumlah x bilangan riil yang menyatakan banyaknya ikan yang akan dijual. Keluaran terdiri dari dua baris. Baris pertama adalah kumpulan bilangan riil yang menyatakan total berat ikan di setiap wadah (jumlah wadah tergantung pada nilai x dan y, urutan ikan yang dimasukan ke dalam wadah sesuai urutan pada masukan baris ke-2). Baris kedua adalah sebuah bilangan riil yang menyatakan berat rata-rata ikan di setiap wadah.]

#### soal2.go
```go
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
```

##### Output soal 2
![Screenshot Output soal 2](https://github.com/ahmad-zainul-hakim/109082530001_ahmad-zainul-hakim/blob/main/PraktikumWeek10/kelinci.png)
[penjelasan]

### 3 [Pos Pelayanan Terpadu (posyandu) sebagai tempat pelayanan kesehatan perlu mencatat data berat balita (dalam kg). Petugas akan memasukkan data tersebut ke dalam array. Dari data yang diperoleh akan dicari berat balita terkecil, terbesar, dan reratanya.]
```go
package main
import "fmt"
type arrBalita [100]float64
func main() {
	var data arrBalita
	var n int
	var min, max float64
	fmt.Print("Masukan banyak data berat balita : ")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Printf("Masukan berat balita ke-%d: ", i+1)
		fmt.Scan(&data[i])
	}
	hitungMinMax(data, n, &min, &max)
	rata2 := ratarataberat(data, n)
	fmt.Printf("Berat balita minimum: %.2f kg\n", min)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", max)
	fmt.Printf("Rerata berat balita: %.2f kg\n", rata2)
} 
func hitungMinMax(berat arrBalita, n int, bMin, bMax *float64) {
	*bMin = berat[0]
	*bMax = berat[0]
	for i := 1; i < n; i++ {
		if berat[i] < *bMin {
		*bMin = berat[i]
		}
		if berat[i] > *bMax {
			*bMax = berat[i]
		}
	}
}
func ratarataberat(arrBerat arrBalita, n int) float64 {
	var total float64 = 0

	for i := 0; i < n; i++ {
		total += arrBerat[i]
	}
	return total / float64(n)
}
```
##### Output soal 3
![Screenshot Output soal 3](https://github.com/ahmad-zainul-hakim/109082530001_ahmad-zainul-hakim/blob/main/PraktikumWeek10/posyandu.png)
[penjelasan] 

