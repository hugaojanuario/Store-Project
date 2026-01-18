package models

import (
	"LojaDoHugo/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosNoBanco() []Produto {

	db := db.ConectarComBanco()

	selectDosProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	pdt := Produto{}
	produtos := []Produto{}

	for selectDosProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDosProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		pdt.Id = id
		pdt.Nome = nome
		pdt.Descricao = descricao
		pdt.Preco = preco
		pdt.Quantidade = quantidade

		produtos = append(produtos, pdt)
	}

	defer db.Close()
	return produtos
}

func CriarNovoProduto(nome, descricao string, preco float64, quantidade int) {
	db := db.ConectarComBanco()

	insereNoBanco, err := db.Prepare("insert into produtos (nome, descricao, preco, quantidade) values ($1,$2,$3,$4)")
	if err != nil {
		panic(err.Error())
	}

	insereNoBanco.Exec(nome, descricao, preco, quantidade)
}

func DeletePorId(id string) {
	db := db.ConectarComBanco()

	deletaProduto, err := db.Prepare("delete from produtos where id= $1")
	if err != nil {
		panic(err.Error)
	}

	deletaProduto.Exec(id)
	db.Close()
}
