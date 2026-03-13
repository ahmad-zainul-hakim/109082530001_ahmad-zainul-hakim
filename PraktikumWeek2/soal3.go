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