package daily

func rotate(matrix [][]int) {
	N := len(matrix)
	for i := 0; i < N/2; i++ {
		for j := i; j < N-i-1; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[N-j-1][i]
			// fmt.Println(i, j, matrix)
			matrix[N-j-1][i] = matrix[N-i-1][N-j-1]
			// fmt.Println(i, j, matrix)
			matrix[N-i-1][N-j-1] = matrix[j][N-i-1]
			// fmt.Println(i, j, matrix)
			matrix[j][N-i-1] = tmp
			// fmt.Printf("[%d,%d] <- [%d,%d]<- [%d,%d]<- [%d,%d]", N-j-1, i, N-i-1, N-j-1, j, N-i-1, i, j)
			// fmt.Println(i, j, matrix)
		}
	}
}
