package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)
		rumah := make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&rumah[j])
		}
		for j := 0; j < m-1; j++ {
			min := j
			for k := j + 1; k < m; k++ {
				if rumah[k] < rumah[min] {
					min = k
				}
			}
			rumah[j], rumah[min] = rumah[min], rumah[j]
		}
		first := true
		for j := 0; j < m; j++ {
			if rumah[j]%2 != 0 {
				if !first {
					fmt.Print(" ")
				}
				fmt.Print(rumah[j])
				first = false
			}
		}
		for j := m - 1; j >= 0; j-- {
			if rumah[j]%2 == 0 {
				if !first {
					fmt.Print(" ")
				}
				fmt.Print(rumah[j])
				first = false
			}
		}
		fmt.Println()
	}
}