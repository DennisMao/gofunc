package strings

type String struct {
	arr  []byte
	size int
}

func New(str string) *String {
	return &String{[]byte(str), len(str)}
}

///////// Match algorithms ///////////
// indexCommon can find the first position of sub string 'T'
// on a string 'S' by using common matching algorithm.
func indexCommon(S, T *String) int {
	lenS := S.size
	lenT := T.size

	if lenS < lenT {
		return -1
	}

	for i := 0; i < lenS-lenT; i++ {
		j := 0
		for ; j < lenT; j++ {
			if S.arr[i+j] != T.arr[j] {
				break
			}
		}

		if j == lenT {
			return i
		}

	}

	return -1
}

// indexKmp can find the first position of sub string 'T'
// on a string 'S' by using KMP algorithm.
func indexKmp(S, T *String) int {
	lenS := S.size
	lenT := T.size

	if lenS < lenT {
		return -1
	}

	for i := 0; i < lenS-lenT; {
		j := 0
		for ; j < lenT; j++ {
			if S.arr[i+j] != T.arr[j] {
				break
			}
		}

		if j == lenT {
			return i
		}

		if j > 0 {
			i += j
		} else {
			i++
		}

	}

	return -1
}
