package algorithm


const (
	PatternSize int = 100
)

func SerchNext(stack string, needle string) int {
  ret := KMP(stack, needle)
  if len(ret) > 0 {
    return ret[len(ret)-1]
  }

  return -1
}

func SearchString(stack string, needle string) int {

	ret := KMP(stack, needle)
	if len(ret) > 0 {
		return ret[0]
	}

	return -1
}

func KMP(stack string, needle string) []int {
	next := preKMP(needle)
	i := 0
	j := 0
	m := len(needle)
	n := len(stack)

	x := []byte(needle)
	y := []byte(stack)
	var ret []int

	if m == 0 || n == 0 {
		return ret
	}

	if n < m {
		return ret
	}

	for j < n {
		for i > -1 && x[i] != y[j] {
			i = next[i]
		}
		i++
		j++

		if i >= m {
			ret = append(ret, j-i)
			i = next[i]
		}
	}

	return ret
}
func preMP(x string) [PatternSize]int {
	var i, j int
	lenght := len(x) - 1

	var mpNext [PatternSize]int
	i = 0
	j = -1
	mpNext[0] = -1

	for i < lenght {
		for j > -1 && x[i] != x[j] {
			j = mpNext[j]
		}
		i++
		j++
		mpNext[i] = j
	}

	return mpNext
}

func preKMP(x string) [PatternSize]int {
	var i, j int
	lenght := len(x) - 1

	var kmpNext [PatternSize]int
	i = 0
	j = -1
	kmpNext[0] = -1

	for i < lenght {
		for j > -1 && x[i] != x[j] {
			j = kmpNext[j]
		}
		i++
		j++

		if x[i] == x[j] {
			kmpNext[i] = kmpNext[j]
		} else {
			kmpNext[i] = j
		}
	}
	return kmpNext
}
