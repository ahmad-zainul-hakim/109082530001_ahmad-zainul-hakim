# <h1 align="center">Laporan Praktikum Modul 1 - ... </h1>
<p align="center">[Ahmad Zainul Hakim] - [109082530001]</p>

## Unguided 

### 1. SOAL 1
#### soal1.go
  
```go
	package main
	import "fmt"

	func main() {
		var n int
		fmt.Scan(&n)
		for i:=0;i<=n;i++{
			fmt.Print(fibonacci(i)," ")
		}
	}
	func fibonacci(n int ) int{
		if n==0{
			return 0
		}else if n==1{
			return 1
		}else{
		return fibonacci(n-1)+fibonacci(n-2)
		}
		}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/ahmad-zainul-hakim/109082530001_ahmad-zainul-hakim/blob/main/Praktikumweeek1/Output/Example.png)
[penjelasan]ini adalah contoh bro!

### 1. SOAL 1
#### soal1.go
  
```go
package main
import "fmt"
func print(x int){
	if x==0 {
	return
		}else{
		print(x-1)
		bintang(x)
		fmt.Println("")
	}
}
func main() {
var n int
fmt.Scan(&n)
print(n)
}
func bintang(x int){
	if x==0{
		return
	}else{
		bintang(x-1)
	fmt.Print("*")
}
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/ahmad-zainul-hakim/109082530001_ahmad-zainul-hakim/blob/main/Praktikumweeek1/Output/Example.png)
[penjelasan]ini adalah contoh bro!

### 1. SOAL 1
#### soal1.go
  
```go
package main
import "fmt"

func main(){
	var n int 
	fmt.Scan(&n)
	faktor(n, 1)
}
func faktor(n , i int){

	if i==n+1{
		return
	}else if n%i == 0{
		fmt.Print(i," ")
	}
	faktor(n, i+1)
}
```
### Output Unguided :

##### Output 
![Screenshot Output Unguided 1_1](https://github.com/ahmad-zainul-hakim/109082530001_ahmad-zainul-hakim/blob/main/Praktikumweeek1/Output/Example.png)
[penjelasan]ini adalah contoh bro!

