# <h1 align="center">Laporan Praktikum Modul 10</h1>
<p align="center">[Ahmad Zainul Hakim] - [109082530001]</p>

## Modul 10 

### 1. [Sebuah program digunakan untuk mendata berat anak kelinci yang akan dijual ke pasar.
Program ini menggunakan array dengan kapasitas 1000 untuk menampung data berat anak
kelinci yang akan dijual.
Masukan terdiri dari sekumpulan bilangan, yang mana bilangan pertama adalah bilangan
bulat N yang menyatakan banyaknya anak kelinci yang akan ditimbang beratnya. Selanjutnya
N bilangan riil berikutnya adalah berat dari anak kelinci yang akan dijual.
Keluaran terdiri dari dua buah bilangan riil yang menyatakan berat kelinci terkecil dan
terbesar.]
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
### Output modul 10 :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/ahmad-zainul-hakim/109082530001_ahmad-zainul-hakim/blob/main/Praktikumweeek1/Output/Example.png)
[penjelasan]ini adalah contoh bro!
```go
package main
import "fmt"

func main(){
	var i, a int
}
```
### Output modul 10 :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/ahmad-zainul-hakim/109082530001_ahmad-zainul-hakim/blob/main/Praktikumweeek1/Output/Example.png)
[penjelasan]ini adalah contoh bro!

