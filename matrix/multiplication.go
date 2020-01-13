package matrix

const N = 3

func Multiplication(a [][]int, b [][]int) [][N]int {
	c := make([][N]int, N)

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < N; k++ {
				c[i][j] += a[i][k] * b[k][j]
			}
		}
	}
	return c
}
