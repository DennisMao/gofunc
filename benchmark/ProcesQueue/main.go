package main

// 待我处理消息公共结构
type TodoList struct {
	Sid        string `json:"sid"`         //待处理的单号
	Status     string `json"status"`       //当前状态
	NoticeTime string `json:"notice_time"` //发布时间
}

func main() {

}
