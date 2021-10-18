package iteracao

import "testing"

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a")
	}
}

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a")
	esperado := "aaaaa"

	if esperado != repeticoes {
		t.Errorf("esperado '%s', repeticoes '%s'", esperado, repeticoes)
	}
}
