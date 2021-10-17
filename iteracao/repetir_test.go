package iteracao

import "testing"

func TestRepetir(t *testing.T) {
	repeticoes := Repetir("a")
	esperado := "aaaaa"

	if esperado != repeticoes {
		t.Errorf("esperado '%s', repeticoes '%s'", esperado, repeticoes)
	}
}
