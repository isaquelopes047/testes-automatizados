package main

import (
	"fmt"
	"testes-automatizados/enderecos"
)

func main() {
	TipoDeEnderecos := enderecos.TipoDeEndereco("avenida Paulista")
	fmt.Println(TipoDeEnderecos)
}
