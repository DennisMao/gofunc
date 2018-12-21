# gzip 
`compress/gzip` 是官方的gzip压缩解压库。支持压缩等级设置。

## 压缩 
```
func pack(rawData []byte) []byte {
```

支持设置压缩比
1最低,最快
9最高,最慢
```
func packLevel(rawData []byte, level int) ([]byte, error) {
```

## 解压
```
func unpack(rawData []byte) ([]byte, error) {
```

## 优化建议
由于gzip包每次使用都需要开辟一个buffer,因此建议使用的时候对buffer引入对象池,减少gc压力。官方包的压缩效率并不乐观,可使用`https://github.com/klauspost/compress`进行替换,实测性能有25%+提升。