package repository

import (
	"github.com/JvWandermurem/cleanCode-golang/domain"
	"gorm.io/gorm"

)

//Interface do repository para acessor o DB
type FigurinhaRepository interface {
	Create(figurinha *domain.Figurinha) error
	FindAll(posicao, tipo string) ([]domain.Figurinha, error)
	FindByID(id uint) (*domain.Figurinha, error)
	Update(figurinha *domain.Figurinha) error
	Delete(id uint) error
}

// Interações com o Banco de dados com o ORM 
type figurinhaRepositoryImpl struct {
	db *gorm.DB
}

func NewFigurinhaRepository(db *gorm.DB) FigurinhaRepository {
	return &figurinhaRepositoryImpl{db: db}
}

// Funções Repository definidas no  /domain

func (r *figurinhaRepositoryImpl) Create(figurinha *domain.Figurinha) error {
	return r.db.Create(figurinha).Error
}

//func (Receiver-("Dono do método"))(Parâmetros de entrada)(variáveis de Saída) - estranhesas do golang =0
func (r *figurinhaRepositoryImpl) FindAll(posicao, tipo string) ([]domain.Figurinha, error) {
	var figurinhas []domain.Figurinha
	query := r.db.Model(&domain.Figurinha{})

	if posicao != "" {
		query = query.Where("posicao = ?", posicao)
	}
	if tipo != "" {
		query = query.Where("tipo = ?", tipo)
	}

	err := query.Find(&figurinhas).Error
	return figurinhas, err
}

func (r *figurinhaRepositoryImpl) FindByID(id uint) (*domain.Figurinha, error) {
	var figurinha domain.Figurinha
	err := r.db.First(&figurinha, id).Error
	if err != nil {
		return nil, err
	}
	return &figurinha, nil
}

func (r *figurinhaRepositoryImpl) Update( figurinha *domain.Figurinha) error {
	return r.db.Save(figurinha).Error
}

func (r *figurinhaRepositoryImpl) Delete(id uint) error {
	return r.db.Delete(&domain.Figurinha{}, id).Error
}