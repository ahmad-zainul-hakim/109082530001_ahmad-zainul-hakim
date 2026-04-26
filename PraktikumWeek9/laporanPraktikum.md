# <h1 align="center">Laporan Praktikum Modul 9 - Array </h1>
<p align="center">[Ahmad Zainul Hakim] - [109082530001]</p>

### 1. SOAL 1
#### soal1.go
  
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
