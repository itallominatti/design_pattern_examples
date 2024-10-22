package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Definindo a interface Veiculo
type Veiculo interface {
	BuscarCliente()
}

type CarroLuxo struct{}

func (c CarroLuxo) BuscarCliente() {
	fmt.Println("Carro de luxo está buscando o cliente...")
}

type CarroPopular struct{}

func (c CarroPopular) BuscarCliente() {
	fmt.Println("Carro popular está buscando o cliente...")
}

type Moto struct{}

func (m Moto) BuscarCliente() {
	fmt.Println("Moto está buscando o cliente...")
}

type VeiculoFactory struct {
	veiculo Veiculo
}

func (vf *VeiculoFactory) GetVeiculo(tipo string) Veiculo {
	switch tipo {
	case "luxo":
		return CarroLuxo{}
	case "popular":
		return CarroPopular{}
	case "moto":
		return Moto{}
	default:
		panic("Veículo não existe")
	}
}

func (vf *VeiculoFactory) BuscarCliente() {
	vf.veiculo.BuscarCliente()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	veiculosDisponiveis := []string{"luxo", "popular", "moto"}

	for i := 0; i < 10; i++ {
		tipo := veiculosDisponiveis[rand.Intn(len(veiculosDisponiveis))]
		factory := VeiculoFactory{}
		veiculo := factory.GetVeiculo(tipo)
		veiculo.BuscarCliente()
	}
}
