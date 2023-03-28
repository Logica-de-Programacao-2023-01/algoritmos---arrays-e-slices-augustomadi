package main

func main() {
	lista := [3]int{4, 2, 1}
	soma := 0

	for i := 0; i < len(lista); i++ {
		soma += lista[i]
	}
	print(soma)
}
