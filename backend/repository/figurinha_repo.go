package repository

import (
	"github.com/JvWandermurem/cleanCode-golang/domain"
	"gorm.io/gorm"

)

// Interações com o Banco de dados com o ORM

type figurinhaRepositoryImpl struct {
	db *gorm.DB
}

func NewFigurinhaRepository(db *gorm.DB) domain.FigurinhaRepository {
	return &figurinhaRepositoryImpl{db: db}
}

func (r *figurinhaRepositoryImpl) Create(figurinha *domain.Figurinha) error {
	return r.db.Create(figurinha).Error
}

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

func (r *figurinhaRepositoryImpl) FindByID(id int) (*domain.Figurinha, error) {
	var figurinha domain.Figurinha
	err := r.db.First(&figurinha, id).Error
	if err != nil {
		return nil, err
	}
	return &figurinha, nil
}

func (r *figurinhaRepositoryImpl) Update(id int, figurinha *domain.Figurinha) error {
	return r.db.Save(figurinha).Error
}

func (r *figurinhaRepositoryImpl) Delete(id int) error {
	return r.db.Delete(&domain.Figurinha{}, id).Error
}