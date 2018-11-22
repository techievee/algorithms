package main

import (
	"fmt"
)

func main() {

	//Get the matrix
	n := 10

	//Create a slice to store the value
	var a = make([][]string, n)

	for i := range a {
		a[i] = make([]string, n)
	}

	//Assign the matrix value by x+y
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			//We initialize the matrix with the corresponding, position values
			a[i][j] = string(i+48) + string(j+48)

		}
	}

	//Print the matrix to confirm
	printMatrix(a, n)

	//Rotate the matrix
	rotateMatrix(a, n)

	//Print the rotated matrix
	printMatrix(a, n)
}

func rotateMatrix(a [][]string, n int) {

	//Rotate the outside layer first
	for layer := 0; layer < n/2; layer++ {
		i1 := layer
		//Rotate each element of the layer
		//Every layer will have (n-1 - layer) count elements
		for j := layer; j < (n - 1 - layer); j++ {
			j1 := j
			//printMatrix(a, n)

			temp := a[i1][j1]

			//Everytime this code executes four times, ie four side of the matrix rotation is performed
			//top, right, bottom,left
			//Formula applied , y->x, x-> (n-1-layer)%n
			for true {
				temp1 := a[j1][(n-1-i1)%(n)]
				a[j1][(n-1-i1)%(n)] = temp
				temp = temp1

				old := j1
				j1 = (n - 1 - i1) % (n)
				i1 = old

				if i1 == layer && j1 == j {
					break
				}

				//ALternate solution, perform fourtimes rotation manually
				//Refer Gayle solution
			}

		}
	}

}

func printMatrix(a [][]string, n int) {

	fmt.Println("\n\nMATRIX")
	//Assign the matrix value by x+y
	for i := 0; i < n; i++ {
		fmt.Printf("\n")
		for j := 0; j < n; j++ {
			fmt.Printf("\t%s", a[i][j])

		}
	}
}
