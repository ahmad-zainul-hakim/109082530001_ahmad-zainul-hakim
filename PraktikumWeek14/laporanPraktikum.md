# <h1 align="center">Laporan Praktikum Modul 14 - Selection sort </h1>
<p align="center">[Ahmad Zainul Hakim] - [109082530001]</p>

## Unguided 
#### soal1.go
```go
package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)
		rumah := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}
		for j := 0; j < m-1; j++ {
			min := j
			for k := j + 1; k < m; k++ {
				if rumah[k] < rumah[min] {
					min = k
				}
			}
			rumah[j], rumah[min] = rumah[min], rumah[j]
		}
		for j := 0; j < m; j++ {
			if j > 0 {
				fmt.Print(" ")
			}
			fmt.Print(rumah[j])
		}
		fmt.Println()
	}
}
```
### Output Unguided :
##### Output 
![![Screenshot Output soal 1](https://github.com/ahmad-zainul-hakim/109082530001_ahmad-zainul-hakim/blob/main/PraktikumWeek14/OutputSoal1.png)]
[Penjelasan]
Dari soal diminta ubntuk bikin program yang dapat mengurutkan nomor rumah kerabat secara ascending (naik). Dengan menggunakan selection short program akan mencari nomor berapa saja yang paling kecil kemudian dipindahkan ke yang paling kiri, setelah itu lanjut ke sisa nomornya. Sisa nomor itu dicari yang paling kecil kemudian dipindah ke sebelah kanan nomor kecil yang telah dipindahkan itu.
## Unguided 
#### soal2.go
```go
package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)
		rumah := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}
		for j := 0; j < m-1; j++ {
			min := j
			for k := j + 1; k < m; k++ {
				if rumah[k] < rumah[min] {
					min = k
				}
			}
			rumah[j], rumah[min] = rumah[min], rumah[j]
		}
		first := true
		for j := 0; j < m; j++ {
			if rumah[j]%2 != 0 {
				if !first {
					fmt.Print(" ")
				}
				fmt.Print(rumah[j])
				first = false
			}
		}
		for j := m - 1; j >= 0; j-- {
			if rumah[j]%2 == 0 {
				if !first {
					fmt.Print(" ")
				}
				fmt.Print(rumah[j])
				first = false
			}
		}
		fmt.Println()
	}
}
```
### Output Unguided :
##### Output 
![![Screenshot Output soal 2](https://github.com/ahmad-zainul-hakim/109082530001_ahmad-zainul-hakim/blob/main/PraktikumWeek14/OutputSoal2.png)]
[Penjelasan]
Yang nomor 2 ini sebenarnya mirip mirip dengan nomor 1, cuman ada aturan tambahan. Jadi biar efisien, Hercules lebih baik melewati rumah ganjil dulu yang ada dikiri urut dari nomor terkecil, kemudian pas udah mentok ujung dia pindah ke kanan ke rumah bernomor genap yang ada disebelah. Program ini memanfaatkan kondisi "IF rumah%2" untuk membantu mengetahui rumah mana yang ganjil dan rumah mana yang genap, kemudian nanti di sorting ganjil dari kecil ke besar terus genap dari besar ke kecil.

## Unguided 
#### soal3.go
```go
package main
import "fmt"
func main() {
	var data []int
	for {
		var x int
		fmt.Scan(&x)
		if x == -5313 {
			break
		}
		if x == 0 {
			for i := 1; i < len(data); i++ {
				key := data[i]
				j := i - 1

				for j >= 0 && data[j] > key {
					data[j+1] = data[j]
					j--
				}
				data[j+1] = key
			}
			n := len(data)
			if n%2 == 1 {
				fmt.Println(data[n/2])
			} else {
				fmt.Println((data[n/2-1] + data[n/2]) / 2)
			}
		} else {
			data = append(data, x)
		}
	}
}
```
### Output Unguided :
##### Output 
![![Screenshot Output soal 3](https://github.com/ahmad-zainul-hakim/109082530001_ahmad-zainul-hakim/blob/main/PraktikumWeek14/OutputSoal3.png)]
[Penjelasan]
Di Soal ini kita disuruh untuk membuat program median, ada 2 aturan angka, yaitu angka 0 dan angka -5303, yang mana 0 itu untuk menghitung median dari angka angka disebelah 0, dan angka -5303 sebagai signal bahwasanya input berhenti. Pertama tama kita harus urutkan dulu angka angka yang telah diinputkan, kemudian diurutkan dari yang terkecil ke yang terbesar, baru setelah itu dicari mediannya. 