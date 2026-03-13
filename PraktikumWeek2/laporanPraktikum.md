# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Ahmad Zainul Hakim] - [109082530001]</p>

## Unguided 

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
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/ahmad-zainul-hakim/109082530001_ahmad-zainul-hakim/blob/main/PraktikumWeek2/output/soal1.png)
[penjelasan] 
Jadi program ini menggunakan prinsip yang sama ketika kita memindahkan 3 bola berwarna kewadah yang berbeda dengan bantuan 1 wadah kosong.Variabel "satu" ,"dua", "tiga" berperan sebagai wadah tujuan yang menjadi tempat akhir menyimpan bola. Dan variable "temp" sebagai wadah kosong tempat sementara untuk menyimpan bola.Endingnya posisi bola berubah, dan wadah kosong tetap menjadi wadah kosong


