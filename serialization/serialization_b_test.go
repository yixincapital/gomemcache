package serialization

import "testing"

type ColorGroup struct {
	Id     int
	Name   string
	Colors []string
}

var group = ColorGroup{
	Id:     1,
	Name:   "Reds",
	Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
}

func BenchmarkMarshalByJson(b *testing.B) {
	provider := &JsonProvider{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		provider.Marshal(group)
	}
}

func BenchmarkUnmarshalByJson(b *testing.B) {
	provider := &JsonProvider{}
	bytes, _ := provider.Marshal(group)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		provider.Unmarshal(bytes, &group)
	}
}

func BenchmarkMarshalByFFJson(b *testing.B) {
	provider := &FFJsonProvider{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		provider.Marshal(group)
	}
}

func BenchmarkUnmarshalByFFJson(b *testing.B) {
	provider := &FFJsonProvider{}
	bytes, _ := provider.Marshal(group)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		provider.Unmarshal(bytes, &group)
	}
}

func BenchmarkMarshalByMsgPack(b *testing.B) {
	provider := &MsgpackProvider{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		provider.Marshal(group)
	}
}

func BenchmarkUnmarshalByMsgPack(b *testing.B) {
	provider := &MsgpackProvider{}
	bytes, _ := provider.Marshal(group)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		provider.Unmarshal(bytes, &group)
	}
}
