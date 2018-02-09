package main

// common structure
type TodoList struct {
	Sid        string `json:"sid"`
	Status     string `json"status"`
	NoticeTime string `json:"notice_time"`
}

// common array
type TodoListArray []TodoList

func (c TodoListArray) Len() int {
	return len(c)
}
func (c TodoListArray) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c TodoListArray) Less(i, j int) bool {
	return c[i].NoticeTime > c[j].NoticeTime
}
