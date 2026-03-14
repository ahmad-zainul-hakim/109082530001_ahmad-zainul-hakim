# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Ahmad Zainul Hakim] - [109082530001]</p>

## Soal1 

### 1. [Soal Latihan Modul 2A]
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

##### Output 
![Screenshot Output soal 1](https://github.com/ahmad-zainul-hakim/109082530001_ahmad-zainul-hakim/blob/main/PraktikumWeek2/output/soal1.png)
[penjelasan] 
Jadi program ini menggunakan prinsip yang sama ketika kita memindahkan 3 bola berwarna kewadah yang berbeda dengan bantuan 1 wadah kosong.Variabel "satu" ,"dua", "tiga" berperan sebagai wadah tujuan yang menjadi tempat akhir menyimpan bola. Dan variable "temp" sebagai wadah kosong tempat sementara untuk menyimpan bola.Endingnya posisi bola berubah, dan wadah kosong tetap menjadi wadah kosong

### Soal2 :

#### soal2.go
  
```go
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
```
### Output Soal2 :

##### Output soal2 
![Screenshot Output soal 2](https://github.com/ahmad-zainul-hakim/109082530001_ahmad-zainul-hakim/blob/main/PraktikumWeek2/output/soal2.png)
[penjelasan] 
Dengan memanfaatkan sistem perulangan, kita bisa membuat program yang akan menentukan apakah 5 percobaan tersebut berhasil atau tidak. Caranya dengan membuat perulangan yang berhenti di iterasi ke 5. Kemudian untuk setiap input (percobaan) benar, variabel "percobaan benar" akan bertambah 1 point. Dengan adanya logika "if percobaanbenar==5" algoritma akan memberikan output berupa true atau false, tergantung dari kelima inputan yang dimasukan.

### Soal3 :

#### Soal3.go

```go
package main

import "fmt"
func main() {
	var berat,kg,sisaGram,biaya,pengiriman,total int
	fmt.Print("Berat parsel (gram) : ")
	fmt.Scan(&berat)
	kg=berat/1000
	sisaGram = berat%1000
	biaya=(berat/1000)*10000
	if berat>10000{
		pengiriman=sisaGram*5
	}else if sisaGram>=500{
		pengiriman=sisaGram*5
	}else if sisaGram<500{
		pengiriman=sisaGram*15
	}
	total=biaya+pengiriman
	if berat>10000{
		total=biaya
	}
	fmt.Printf("Detail berat : %d kg + %d gram \nDetail biaya : Rp.%d + Rp.%d \nTotal biaya : Rp.%d",kg,sisaGram,biaya,pengiriman,total)
	}
```
### Output soal3 :
![Screenshot Output soal 2](https://github.com/ahmad-zainul-hakim/109082530001_ahmad-zainul-hakim/blob/main/PraktikumWeek2/output/soal3.png)

[Penjelasan]
Program ini dapat menghitung total biaya pengiriman parsel berdasarkan berat parsel.Dengan harga 10k per 1 kg program ini menghitung total biayanya. Tapi sebelum menghitung,ada beberapa aturan 
1.Jika sisa berat tidak kurang dari 500
gram, maka tambahan biaya kirim hanya Rp. 5,- per gram
2.jika sisa berat kurang dari 500
gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram
3.Sisa berat (yang
kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.

Dengan memanfaatkan logika if, else if dan else. Program dapat dibuat. ketika ada input masuk, algoritma akan mengecek apakah dia bernilai benar untuk "if" pertama. jika tidak maka akan dilanjut ke "else if" dan apabila tidak benar, dia akan melanjutkan ke "else". Jika input benar di salah satu kondisi (if / else if/else) maka algoritma akan menampilkan output sesuai yang dimana nilai benar itu berhenti.