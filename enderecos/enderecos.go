package enderecos

import "strings"

// TipoDeEnderecos verifica se o endereco tem o tipo valido e retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{
		"rua",
		"avenida",
		"estrada",
		"rodovia",
	}

	enderecoEmLetraMinuscula := strings.ToLower((endereco))
	primeiraPalavraDoEndereco := strings.Split(enderecoEmLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo Inválido"
}