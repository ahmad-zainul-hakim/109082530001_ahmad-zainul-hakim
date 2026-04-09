# <h1 align="center">Laporan Praktikum Modul 4</h1>
<p align="center">[Ahmad Zainul Hakim] - [109082530001]</p>

## Unguided 

### SOAL 1
#### soal1.go
  
```go
package main
import "fmt"

func main(){
	var a,b,c,d int 
	var Pac,Pbd,Cac, Cbd int
	fmt.Scan(&a,&b,&c,&d)
	permutasi(a,c,&Pac)
	kombinasi (a,c,&Cac)
	permutasi (b,d,&Pbd)
	kombinasi (b,d,&Cbd)
	fmt.Printf("%d %d \n%d %d", Pac,Cac,Pbd,Cbd )
}

func faktorial(x int, hasil *int){
	*hasil=1
	for i:=x;i>=1;i--{
		*hasil=*hasil*i
	}
}

func permutasi(n,r int, hasil*int){
	var faktorn, faktornr int
	faktorial(n,&faktorn)
	faktorial(n-r,&faktornr)
	*hasil=faktorn/faktornr
}

func kombinasi (n,r int, hasil*int){
	var faktorn,faktorr,faktornr int
	faktorial(n,&faktorn)
	faktorial(r,&faktorr)
	faktorial(n-r,&faktornr)
	*hasil= faktorn/(faktorr*faktornr)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/ahmad-zainul-hakim/109082530001_ahmad-zainul-hakim/blob/main/PraktikumWeek4/Output/soal1.png)
[penjelasan]
Program ini menampilkan cara menghitung permutasi dan kombinasi dengan menggunakan prosedur faktorial, Permutasian, dan kombinasi

###Soal 1
#### Soal2.go
```go
package main

import "fmt"

func main() {
	var nama1, nama2,win string
	var skornama1,skorsoal1,skornama2,skorsoal2 int
	fmt.Scan(&nama1)
	hitungSkor(nama1,&skorsoal1,&skornama1)
	// fmt.Print(skornama1,skorsoal1)
	fmt.Scan(&nama2)
	hitungSkor(nama2,&skorsoal2,&skornama2)

	if skorsoal1>skorsoal2{
		win=nama1
	}else if skorsoal1<skorsoal2{
		win=nama2
	}else if skorsoal1==skorsoal2{
		if skornama1>skornama2{
			win=nama1
		}else{
			win=nama2
		}
	}
	fmt.Print(win)
	switch win{
		case nama1:
			fmt.Print(" ",skornama1, " ",  skorsoal1)
		case nama2:
			fmt.Print(" ",skornama2, " ", skorsoal2)

	}
}
func hitungSkor(nama string, hasilskor, hasilsoal *int) {
	var skor int
	for i := 1; i <= 8; i++ {
		fmt.Scan(&skor)
		soal := 1
		if skor > 300 {
			skor = 0
			soal = 0
		}
		*hasilskor = *hasilskor + skor
		*hasilsoal = *hasilsoal + soal
	}
}

```

### Output Unguided
#### Output
![Screenshot Output Unguided 1_1](https://github.com/ahmad-zainul-hakim/109082530001_ahmad-zainul-hakim/blob/main/PraktikumWeek4/Output/soal2.png)
[penjelasan]
Program ini adalah program untuk menghitung skor dan menentukan pemenang. Pemenang dhitung berdasarkan berapa banyak soal yang diselesaikan. Jika jumlah soal yang diselesaikan sama, maka pmenang ditentukan berdasarkan pada siapa yang paling cepat menyelesaikan.