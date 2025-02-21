package slicesandarray

func Sum(v []int) (result int) {
	for _, number := range v {
		result += number
	}
	return result
}

func SumAll(numberToSums ...[]int) (sums []int) {

	for _, v := range numberToSums {
		sums = append(sums, Sum(v))
	}
	return
}

func SumAllTails(numberToSumTail ...[]int) (tails []int) {
	for _, v := range numberToSumTail {
		if len(v) == 0 {
			tails = append(tails, 0)
		} else {
			tail := v[1:]
			tails = append(tails, Sum(tail))
		}
	}
	return
}
