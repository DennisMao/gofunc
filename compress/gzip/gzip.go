package gzip

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

// 压缩
func pack(rawData []byte) []byte {

	buf := bytes.Buffer{}
	w := gzip.NewWriter(&buf)

	w.Write(rawData)
	w.Close()

	return buf.Bytes()
}

// 压缩
// 支持设置等级
func packLevel(rawData []byte, level int) ([]byte, error) {
	// 设置压缩等级 1最快9最慢
	// 此处可能导致CPU占率过高,使用对象池和内存共享优化
	buf := bytes.Buffer{}
	w, err := gzip.NewWriterLevel(&buf, level)
	if err != nil {
		return nil, err
	}

	w.Write(rawData)
	w.Close()

	return buf.Bytes(), nil
}

// 解压
func unpack(rawData []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewReader(rawData))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	b, _ := ioutil.ReadAll(r)
	return b, nil
}
