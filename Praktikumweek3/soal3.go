package main
import "fmt"

func main(){
	var cx1, cy1, radius1 float64
	var cx2, cy2, radius2 float64
	var x, y float64
	fmt.Scanln(&cx1, &cy1, &radius1)
	fmt.Scanln(&cx2, &cy2, &radius2)
	fmt.Scanln(&x, &y)

	dalamlingkaran1:=didalam(cx1, cy1,radius1,x,y)
	dalamlingkaran2:=didalam(cx2, cy2,radius2,x,y) 

	if dalamlingkaran1 && dalamlingkaran2 == true{
		fmt.Println("Titik berada di dalam lingkaran 1 dan 2")
	}else if !dalamlingkaran1 && !dalamlingkaran2{
		fmt.Println("Titik berada di luar lingkaran 1 dan 2")
	}else if dalamlingkaran1{
		fmt.Println("Titik berada di dalam lingkaran 1")
	}else if dalamlingkaran2{
		fmt.Println("Titik berada di dalam lingkaran 2")
	}
}
func hitungJarak (a, b, c, d float64) float64{
	dx := a - c
	dy := b - d
	return dx*dx + dy*dy
}

func didalam (cx,cy,radius,x,y float64) bool {
	dx := hitungJarak(cx, cy, x, y)
	return dx <= radius*radius
}