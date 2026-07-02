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