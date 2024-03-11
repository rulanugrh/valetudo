package algorithm

func CompressLZW(str string) []int {
	code := 256
  dict := make(map[string]int)
	for i := 0; i < 256; i++ {
		dict[string(rune(i))] = i
	}

	current := ""
	result := make([]int, 0)
	for _, c := range []byte(str) {
		phrase := current + string(c)
		if _, isTrue := dict[phrase]; isTrue {
			current = phrase
		} else {
			result = append(result, dict[current])
			dict[phrase] = code
			code++
			current = string(c)
		}
	}

	if current != "" {
		result = append(result, dict[current])
	}

	return result
}

func DecompresesLZW(data []int) string {
	code := 256
	dict := make(map[int]string)
	for i := 0; i < 256; i++ {
		dict[i] = string(rune(i))
	}

	current := string(rune(data[0]))
	result := current

	for _, element := range data[1:] {
		var word string
		if x, ok := dict[element]; ok {
			word = x
		} else if element == code {
			word = current + current[:1]
		} else {
			println("bad compressed")
		}

		result += word
		dict[code] = current + word[:1]
		code++
		current = word
	}

	return result
}
