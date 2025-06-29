package enderecos

import (
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndreco(t *testing.T) {

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Rua Joao", "Rua"},
		{"Avenida adão", "Avenida"},
		{"Rodovia ABC", "Rodovia"},
		{"Preça ABC", "Tipo Inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("Tipo recebido é diferente do esperado! Esperava %s e recebeu %s",
				tipoDeEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}
}
