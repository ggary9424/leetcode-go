package leedcode_go

func champagneTower(poured int, query_row int, query_glass int) float64 {
	glasses := [101][101]float64{}
	glasses[0][0] = float64(poured)
	for i := 0; i <= query_row; i++ {
		for j := 0; j <= i; j++ {
			var dividedPoured float64
			dividedPoured = (glasses[i][j] - 1) / 2.0
			if dividedPoured > 0 {
				glasses[i+1][j] += dividedPoured
				glasses[i+1][j+1] += dividedPoured
			}
		}
	}

	if glasses[query_row][query_glass] > 1 {
		return 1
	} else {
		return glasses[query_row][query_glass]
	}
}
