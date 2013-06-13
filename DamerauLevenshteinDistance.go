package DamerauLevenshteinDistance

//import "fmt"

func Max(a, b int) (max int) {
	max = a
	if max < b {
		max = b
	}
	return
}

func Min(a, b int) (min int) {
	min = a
	if min > b {
		min = b
	}
	return
}

func Distance(source, target string) int {
	// Thanks to the Wikipedia gods: https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance
	if len(source) == 0 || len(target) == 0 {
		return Max(len(source), len(target))
	}

	var dict = make(map[rune]int, len(source)+len(target))
	var scores = make([][]int, len(source)+2)
	for i := range scores {
		scores[i] = make([]int, len(target)+2)
	}

	scores[0][0] = len(source) + len(target)

	for i, c := range source {
		scores[i+1][1] = i
		scores[i+1][0] = scores[0][0]
		dict[c] = 0
	}
	for i, c := range target {
		scores[1][i+1] = i
		scores[0][i+1] = scores[0][0]
		dict[c] = 0
	}

	for i, ci := range source {
		cio := ci
		cost := 0
		i = i + 1
		for j, cj := range target {
			//fmt.Printf("%v\n", scores)
			cjo := cj
			j = j + 1

			if cio == cjo {
				//fmt.Printf("1: %v:%v %v\n", i+1, j+1, scores[i][j])
				scores[i+1][j+1] = scores[i][j]
				cost = j
			} else {
				//fmt.Printf("2: %v:%v %v\n", i+1, j+1, Min(scores[i][j], Min(scores[i+1][j], scores[i][j+1])) + 1)
				scores[i+1][j+1] = Min(scores[i][j], Min(scores[i+1][j], scores[i][j+1])) + 1
			}

			//fmt.Printf("3: %v:%v %v | %v\n", i+1, j+1, scores[i+1][j+1], Min(scores[i+1][j+1], scores[dict[cjo]][cost] + (i - dict[cjo] - 1) + 1 + (j - cost - 1)))
			scores[i+1][j+1] = Min(scores[i+1][j+1], scores[dict[cjo]][cost] + (i - dict[cjo] - 1) + 1 + (j - cost - 1))
		}

		dict[cio] = i
	}

	return scores[len(source)+1][len(target)+1]
}
