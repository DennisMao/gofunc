package strings

type String struct {
	arr  []byte
	size int
}

func New(str string) *String {
	return &String{[]byte(str), len(str)}
}

///////// Match algorithms ///////////
// Match sub string 'T' on a string 'S' by using common matching algorithm.
func IndexKmp(S, T *String) int {
	return -1
}

// Match sub string 'T' on a string 'S' by using KMP algorithm.
func IndexKmp(S, T *String) int {
	return -1
}
