package main

import (
	"fmt"
)

func main() {
	var arr [7]int
	arr[5] = 2
	arr[1] = 5
	fmt.Println(arr, " : ", len(arr), " ", cap(arr))

	slc := []string{"1", "2", "c", "d", "e", "f"}
	fmt.Println(slc, " : ", len(slc), " ", cap(slc))
	slc = append(slc, "d")
	slc = append(slc, "d")
	fmt.Println(slc, " : ", len(slc), " ", cap(slc))
	slc = append(slc, "d")
	slc = append(slc, "d")
	slc = append(slc, "d")
	slc = append(slc, "d")
	slc = append(slc, "d")
	fmt.Println(slc, " : ", len(slc), " ", cap(slc))
	newSlice := []string{"g", "k"}
	slc = append(slc, newSlice...)
	fmt.Println(slc, " : ", len(slc), " ", cap(slc))
	fmt.Println(" **************** PALINDROME    *********")

	str := "amor a roma"
	fmt.Println(isPalindrome(str))

	for i := range slc {
		fmt.Println(i)
	}
	/*Lista de Alumnos
	notas := []float64{}
	var (
		n    int
		nota float64
	)
	fmt.Println("Cantidad de notas ")
	fmt.Print("n: ")
	fmt.Scan(&n)
	for n > 0 {
		fmt.Printf("Nota %d: ", n)
		fmt.Scan(&nota)
		notas = append(notas, nota)
		n--
	}
	sum := 0.0
	for _, e := range notas {
		sum += e
	}
	fmt.Printf("primedio: %.3f", sum/float64(len(notas)))*/
	fmt.Println(fibo(15))

}
func fibo(n int) int {
	i, j := 0, 1
	for n-2 > 0 {
		aux := j
		j = i + j
		i = aux
		n--
		fmt.Printf(" %d : ", j)
	}
	fmt.Println("")
	return j
}
func isPalindrome(str string) (is bool) {
	l, r := 0, len(str)-1
	is = true
	for l < r {
		if str[l] != str[r] {
			is = false
			break
		}
		l++
		r--
	}
	return is
}
