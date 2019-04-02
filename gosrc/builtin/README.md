# Go底层 make与new区别
在$GOROOT/go/src/builtin.go中,声明了Go语言编译的关键字说明,里面的关健字在编译阶段会被替换为内建函数。
```
// The make built-in function allocates and initializes an object of type
// slice, map, or chan (only). Like new, the first argument is a type, not a
// value. Unlike new, make's return type is the same as the type of its
// argument, not a pointer to it. The specification of the result depends on
// the type:
//	Slice: The size specifies the length. The capacity of the slice is
//	equal to its length. A second integer argument may be provided to
//	specify a different capacity; it must be no smaller than the
//	length, so make([]int, 0, 10) allocates a slice of length 0 and
//	capacity 10.
//	Map: An initial allocation is made according to the size but the
//	resulting map has length 0. The size may be omitted, in which case
//	a small starting size is allocated.
//	Channel: The channel's buffer is initialized with the specified
//	buffer capacity. If zero, or the size is omitted, the channel is
//	unbuffered.
func make(Type, size IntegerType) Type

// The new built-in function allocates memory. The first argument is a type,
// not a value, and the value returned is a pointer to a newly
// allocated zero value of that type.
func new(Type) *Type
```

## 输出结果
Slice切片
```
===newSlice          
Pointer:&[]string(nil) 
Object:[]string(nil) 
DataStructure:&main.slice{array:(unsafe.Pointer)(nil), len:0, cap:0}
===newSliceAppended  
Pointer:&[]string{"1"} 
Object:[]string{"1"} 
DataStructure:&main.slice{array:(unsafe.Pointer)(0xc0000521c0), len:1, cap:1}
===makeSlice         
Pointer:&[]string{} 
Object:[]string{} 
DataStructure:&main.slice{array:(unsafe.Pointer)(0x57a1c8), len:0, cap:0}
===makeSliceAppended 
Pointer:&[]string{"1"} 
Object:[]string{"1"} 
DataStructure:&main.slice{array:(unsafe.Pointer)(0xc0000521f0), len:1, cap:1}
```
我们可以看到,new的切片中`array:(unsafe.Pointer)(nil)`array变量并未分配地址,而make出来的切片`(unsafe.Pointer)(0x57a1c8)`已成功被分配了一个零地址(由于我们初始化时候并未指定长度,默认为0)

Map字典
```
===newMap      
Pointer:&map[string]interface {}(nil)
Object:map[string]interface {}(nil) 
DataStructure:&main.hmap{count:0, flags:0x0, B:0x0, hash0:0x0, buckets:(unsafe.Pointer)(nil), oldbuckets:(unsafe.Pointer)(nil), nevacuate:0x0, overflow:(*[2]*[]*main.bmap)(nil)}
===makeMap     
Pointer:&map[string]interface {}{} 
Object:map[string]interface {}{} 
DataStructure:&main.hmap{count:824634196592, flags:0x0, B:0x0, hash0:0x0, buckets:(unsafe.Pointer)(nil), oldbuckets:(unsafe.Pointer)(nil), nevacuate:0x0, overflow:(*[2]*[]*main.bmap)(nil)}
```

Channel通道
```
===newChannel      
Pointer:(*chan int)(0xc000086030) 
Object:(chan int)(nil) 
DataStructure:&main.hchan{qcount:0x0, dataqsiz:0x0, buf:(unsafe.Pointer)(nil), elemsize:0x0, closed:0x0, elemtype:(unsafe.Pointer)(nil), sendx:0x0, recvx:0x0, recvq:struct { first unsafe.Pointer; last unsafe.Pointer }{first:(unsafe.Pointer)(nil), last:(unsafe.Pointer)(nil)}, sendq:struct { first unsafe.Pointer; last unsafe.Pointer }{first:(unsafe.Pointer)(nil), last:(unsafe.Pointer)(nil)}, lock:struct { key uintptr }{key:0x0}}
===makeChannel     
Pointer:(*chan int)(0xc000086038) 
Object:(chan int)(0xc000050060) 
DataStructure:&main.hchan{qcount:0xc000050060, dataqsiz:0x0, buf:(unsafe.Pointer)(nil), elemsize:0x0, closed:0x0, elemtype:(unsafe.Pointer)(nil), sendx:0x0, recvx:0x0, recvq:struct { first unsafe.Pointer; last unsafe.Pointer }{first:(unsafe.Pointer)(nil), last:(unsafe.Pointer)(nil)}, sendq:struct { first unsafe.Pointer; last unsafe.Pointer }{first:(unsafe.Pointer)(nil), last:(unsafe.Pointer)(nil)}, lock:struct { key uintptr }{key:0x0}}
```

TODO：
+ 1.printf %#v无法正确解析unsaft.Pointer
+ 2.builtin append等相关操作的二次初始化
+ 3.new map和chan 操作会引起panic的原因