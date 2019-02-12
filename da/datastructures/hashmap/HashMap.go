// Name: Binary search tree
package HashMap

const (
	CompareEqual = 1 << iota
	CompareLess
	CompareMore
	LoadFactor        = 6.5
	DefaultCapability = 1024
	DefaultB          = 10
	BucketSize        = 8
)

type HashNode struct {
	key   string
	value interface{}
	next  *HashNode // 链地址法
}

type HashMap struct {
	bucket []HashNode
	size   int
	seed   int
	b      uint8 // 底层顺序存储长度的阶数 log2(size)
}

func New(capabilty int) *HashMap {
	if capabilty == 0 {
		capabilty = DefaultCapability
	}
	return makeHashmap(capabilty)
}

// makeHashmap can create a HashMap storage,malloc memory
// and return the pointer to the HashMap.
//
// The key storage is an array with set length.And the array will be
// separated logically into several buckets (length = 8).
func makeHashmap(capabilty int) *HashMap {

	// Find a suitable B for the bucket size
	b := uint8(0)
	for ; float32(capabilty) > LoadFactor*float32(uintptr(1)<<b); b++ {
	}

	h := HashMap{
		bucket: make([]HashNode, 2<<b, 2<<b),
		size:   0,
		seed:   2 << b,
	}

	return &h
}

// 哈希函数
// 取余
func hashFunc(key string, seed int) int {

	keyNum := int(0)

	for _, k := range []byte(key) {
		keyNum += int(k)
	}

	return keyNum % seed
}

func (this *HashMap) Insert(k string, v interface{}) {
	idx, isFind := _searchWithHot(this, k)
	if isFind {
		return
	}

	if idx == -1 {
		return
	}

	if this.size > this.seed {
		return
	}

	this.bucket[idx].key = k
	this.bucket[idx].value = v
	this.size++
}

func (this *HashMap) Delete(k string) {
	idx, isFind := _searchWithHot(this, k)
	if !isFind {
		return
	}

	if idx == -1 {
		return
	}

	this.bucket[idx].key = ""
	this.bucket[idx].value = nil
	this.size--
}

func (this *HashMap) Search(k string) (interface{}, bool) {
	idx := _search(this, k)
	if idx == -1 {
		return nil, false
	}

	return this.bucket[idx].value, true
}

// Find the index from the bucket for a key
func _search(m *HashMap, k string) int {
	idx := hashFunc(k, m.seed)
	if idx > m.seed {
		return -1
	}

	if m.bucket[idx].key == k {
		return idx
	}

	return -1
}

// Find the index from the bucket for a key
func _searchWithHot(m *HashMap, k string) (int, bool) {
	idx := hashFunc(k, m.seed)
	if idx > m.seed {
		return -1, false
	}

	if m.bucket[idx].key == k {
		return idx, true
	}

	return idx, false
}
