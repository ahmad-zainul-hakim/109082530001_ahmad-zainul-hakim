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