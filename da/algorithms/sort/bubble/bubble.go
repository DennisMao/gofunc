package quick

func SortBubble(data *[]string) {
	d := *data
	lenD := len(d)
	if lenD < 2 {
		return
	}

	i := 0
	for ; i < lenD; i++ {
		for j := i; j < lenD; j++ {
			if d[i] > d[j] { // 升降序控制
				d[i], d[j] = d[j], d[i]
			}
		}
	}
}
