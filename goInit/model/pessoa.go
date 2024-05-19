package model

import "time"

type Pessoa struct {
	Nome            string
	Endereco        Endereco
	DataDeNacimento time.Time
	Idade           int
}

// Method
func (p *Pessoa) CalculaIdade() {
	anoNascimento := p.DataDeNacimento.Year()
	anoAtual := time.Now().Year()

	p.Idade = anoAtual - anoNascimento
}

// func CalculaIdade(p Pessoa) int {
// 	anoNascimento := p.DataDeNacimento.Year()
// 	anoAtual := time.Now().Year()

// 	return anoAtual - anoNascimento

// }
