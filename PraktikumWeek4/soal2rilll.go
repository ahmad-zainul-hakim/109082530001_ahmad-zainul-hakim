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
