package model

import (
	"errors"
	"time"
)

type Compra struct {
	Mercado string
	Data    time.Time
	Itens   []ItenDaCompra
}

type ItenDaCompra struct {
	Nome string
}

func NewCompra(mercado string, data time.Time, nomeDosItens []string) (*Compra, error) {
	var itens []ItenDaCompra

	if mercado == "" {
		return nil, errors.New("nome do mercado nao pode ser vazio")
	}

	if len(nomeDosItens) == 0 {
		return nil, errors.New("numero de itens nao pode ser vazio")
	}

	for _, nome := range nomeDosItens {
		itens = append(itens, ItenDaCompra{Nome: nome})
	}

	return &Compra{
		Mercado: mercado,
		Data:    time.Now(),
		Itens:   itens,
	}, nil

}
