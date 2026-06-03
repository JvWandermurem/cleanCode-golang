package domain

import (
	"time"
)
//Structs da figurinha e struct para criar uma nova figurinha ou atualizar ela.
type Figurinha struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	Numero    string    `gorm:"not null" json:"numero"`
	Tipo      string    `json:"tipo"`
	Posicao   string    `json:"posicao"` 
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateFigurinhaRequest struct {
	Numero  string `json:"numero" binding:"required"`
	Tipo    string `json:"tipo" binding:"required"`
	Posicao string `json:"posicao" binding:"required"`
}

type UpdateFigurinhaRequest struct {
	Numero  *string `json:"numero"`
	Tipo    *string `json:"tipo"`
	Posicao *string `json:"posicao"`
}


//Funções do Gorm para a Camada Reposiory.
type FigurinhaRepository interface {
	Create(figurinha *Figurinha) error
	FindAll(posicao, tipo string) ([]Figurinha, error)
	FindByID(id int) (*Figurinha, error)
	Update(id int, figurinha *Figurinha) error
	Delete(id int) error
}

//Funções do Gorm para Camada de Service
type FigurinhaService interface {
	Create(req CreateFigurinhaRequest) (*Figurinha, error)
	FindAll(posicao, tipo string) ([]Figurinha, error)
	FindByID(id int) (*Figurinha, error)
	Update(id int, req UpdateFigurinhaRequest) (*Figurinha, error)
	Delete(id int) error
}

