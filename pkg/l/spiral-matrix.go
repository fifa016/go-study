/*
 * @Descripttion:
 * @Author: jzh
 * @Date: 2022-06-14 22:49:08
 */
package l

func SpiralOrder(matrix [][]int) []int {
	res := []int{}
	flag := make([][]bool, len(matrix))
	for i := 0; i < len(matrix); i++ {
		flag[i] = make([]bool, len(matrix[0]))
	}

	i, j := 0, 0
	count := 0

	for count < len(matrix)*len(matrix[0]) {
		//right
		for j < len(matrix[0]) && !flag[i][j] {
			res = append(res, matrix[i][j])
			flag[i][j] = true
			count++
			j++
		}
		j--
		i++

		//down
		for i < len(matrix) && !flag[i][j] {
			res = append(res, matrix[i][j])
			flag[i][j] = true
			count++
			i++
		}
		i--
		j--

		//left
		for j >= 0 && !flag[i][j] {
			res = append(res, matrix[i][j])
			flag[i][j] = true
			count++
			j--
		}
		j++
		i--

		//up
		for i >= 0 && !flag[i][j] {
			res = append(res, matrix[i][j])
			flag[i][j] = true
			count++
			i--
		}
		i++
		j++

	}

	return res
}
