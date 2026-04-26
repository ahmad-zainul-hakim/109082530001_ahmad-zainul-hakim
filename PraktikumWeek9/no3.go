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