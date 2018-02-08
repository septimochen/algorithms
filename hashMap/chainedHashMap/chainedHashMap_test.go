package chainedHashMap

import (
	"testing"
	"algorithms/hashMap"
	"reflect"
	"fmt"
)

func Test_ChainedHashMap(t *testing.T) {
	cmap := New(4)
	hashMap.BasicTestHashMap(t,cmap)
}

func Test_ChainedHashMapResize(t *testing.T) {
	cmap := new(ChainedHashMap)
	cmap.Init(4)
	hashMap.TestHashMapResize(t, cmap)
	if !reflect.DeepEqual(cmap.Cap, uint32(4+hashMap.DEFALUTCAP)) {
		t.Log(fmt.Sprintf("expect:", uint32(4+hashMap.DEFALUTCAP)) + fmt.Sprintf("but get:", cmap.Cap))
		t.Fail()
	}
}

func Test_ChainedHashMapDelete(t *testing.T) {
	cmap := new(ChainedHashMap)
	cmap.Init(4)
	hashMap.TestHashMapDelete(t, cmap)
	if !reflect.DeepEqual(cmap.Count, uint32(0)) {
		t.Log(fmt.Sprintf("expect:", 0) + fmt.Sprintf("but get:", cmap.Count))
		t.Fail()
	}
}

func BenchmarkChainedHashMap_HashInsert(b *testing.B) {
	hashMap.BenchmarkHashMapInsert(b,New(128))
}

func BenchmarkChainedHashMap_HashInsertDelete(b *testing.B) {
	hashMap.BenchmarkHashMapInsertDelete(b,New(128))
}

func BenchmarkChainedHashMap_HashGet(b *testing.B) {
	hashMap.BenchmarkHashMapGet(b,New(128))
}