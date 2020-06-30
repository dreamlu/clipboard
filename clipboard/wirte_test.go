package clipboard

import "testing"

func TestWrite(t *testing.T) {
	b, _ := Read()
	t.Log(string(b))
	Write([]byte("new clip"))
	b, _ = Read()
	t.Log(string(b))
}
