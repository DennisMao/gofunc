/*
计数器-应用层
*/
package counter

type Counter struct {
	Uid     string `json:"uid"`
	Count   int64  `json:"count"`
	Expired string `json:"expired"`
}

func Add(uid string, number string) error {
	return CounterCache.Add(uid, number)
}

func Del(uid string) error {
	return CounterCache.Del(uid)
}

func Get(uid string) (error, string) {
	return CounterCache.Get(uid)
}

func ShowList() (error, []Counter) {
	return CounterCache.ShowList()
}
