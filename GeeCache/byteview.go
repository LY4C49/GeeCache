package GeeCache

type ByteView struct {
	bv []byte
}

// 实现了被缓存对象具有的Len()方法
func (v ByteView) Len() int {
	return len(v.bv)
}

func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.bv)
}

func (v ByteView) String() string {
	return string(v.bv)
}

func cloneBytes(b []byte) []byte {
	new_bv := make([]byte, len(b))
	copy(new_bv, b)
	return new_bv
}
