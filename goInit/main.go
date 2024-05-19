package main

import (
	"fmt"
	"goInit/model"
	"time"
)

func main() {
	endereco := model.Endereco{
		Rua:    "mauro",
		Numero: 35,
		Cidade: "SP",
	}

	pessoa := model.Pessoa{
		Nome:            "Gabriel",
		Endereco:        endereco,
		DataDeNacimento: time.Date(1997, 5, 12, 0, 0, 0, 0, time.Local),
	}

	//idade := model.CalculaIdade(pessoa)
	pessoa.CalculaIdade()

	fmt.Println("My infos:", pessoa)

	// Using Automovel
	createMoto := model.Automovel{
		Ano:    2024,
		Placa:  "XPTO",
		Modelo: "CG",
	}

	moto := model.Moto{
		Automovel:   createMoto,
		Cilindradas: 20,
	}

	fmt.Println("I am Using a veiculo", moto)

	// Minhas compras
	var nomeDosItens []string
	nomeDosItens = append(nomeDosItens, "rice")
	nomeDosItens = append(nomeDosItens, "bean")
	nomeDosItens = append(nomeDosItens, "meat")

	compra, err := model.NewCompra("Real", time.Now(), nomeDosItens)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(compra)
	}

}
