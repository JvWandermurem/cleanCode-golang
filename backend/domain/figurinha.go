package domain

import (
	"time"
)

// Types para deixar salvo o Tipo e a Posição
type FigurinhaTipo string

const (
	TipoComum         FigurinhaTipo = "comum"
	TipoBrilhante     FigurinhaTipo = "brilhante"
	TipoLegendsOuro   FigurinhaTipo = "legends_ouro"
	TipoLegendsBronze FigurinhaTipo = "legends_bronze"
)

type FigurinhaPosicao string

const (
	PosicaoGoleiro      FigurinhaPosicao = "Goleiro"
	PosicaoZagueiro     FigurinhaPosicao = "Zagueiro"
	PosicaoMeioCampista FigurinhaPosicao = "Meio-campista"
	PosicaoAtacante     FigurinhaPosicao = "Atacante"
)

//Struct da Figurinha 
type Figurinha struct {
	ID        uint`gorm:"primaryKey" json:"id"`
	Numero    string           `gorm:"not null"   json:"numero"` 
	Tipo      FigurinhaTipo    `gorm:"default:comum" json:"tipo"`
	Posicao   FigurinhaPosicao `gorm:"not null"   json:"posicao"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
}

//Struct para criar uma nova Figurinha.
type CreateFigurinhaRequest struct {
	Numero  string `json:"numero" binding:"required"`
	Tipo    string `json:"tipo" binding:"required"`
	Posicao string `json:"posicao" binding:"required"`
}

// Atualizar Figurinha (ponteiro uau faz sentido)
type UpdateFigurinhaRequest struct {
	Numero  *string `json:"numero"`
	Tipo    *string `json:"tipo"`
	Posicao *string `json:"posicao"`
}
