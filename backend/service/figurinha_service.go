package service

import(
	"errors"
	"github.com/JvWandermurem/cleanCode-golang/domain"
	"github.com/JvWandermurem/cleanCode-golang/repository"
	"gorm.io/gorm"

)

//Erros padronizados
var (
	ErrFigurinhaNotFound = errors.New("figurinha não encontrada")
	ErrInvalidTipo       = errors.New("tipo inválido, deve ser 'comum', 'brilhante', 'legends_ouro' ou 'legends_bronze'")
	ErrInvalidPosicao    = errors.New("posição inválida, deve ser'Goleiro', 'Zagueiro', 'Meio-campista' ou 'Atacante'")
)

type FigurinhaService interface {
	CreateFigurinha(req domain.CreateFigurinhaRequest) (*domain.Figurinha, error)
	ListFigurinhas(posicao, tipo string) ([]domain.Figurinha, error)
	GetFigurinha(id uint) (*domain.Figurinha, error)
	UpdateFigurinha(id uint, req domain.UpdateFigurinhaRequest) (*domain.Figurinha, error)
	DeleteFigurinha(id uint) error
}

type figurinhaServiceImpl struct {
	repo repository.FigurinhaRepository
}

func NewFigurinhaService(repo repository.FigurinhaRepository) FigurinhaService {
	return &figurinhaServiceImpl{repo: repo}
}

// Funçções da interface
func (s *figurinhaServiceImpl) CreateFigurinha(req domain.CreateFigurinhaRequest) (*domain.Figurinha, error) {
	if !isValidTipo(req.Tipo) {
		return nil, ErrInvalidTipo
	}
	if !isValidPosicao(req.Posicao) {
		return nil, ErrInvalidPosicao
	}

	figurinha := &domain.Figurinha{
		Numero:  req.Numero,
		Tipo:    domain.FigurinhaTipo(req.Tipo),
		Posicao: domain.FigurinhaPosicao(req.Posicao),
	}

	if err := s.repo.Create(figurinha); err != nil {
		return nil, err
	}
	return figurinha, nil
}

func (s *figurinhaServiceImpl) ListFigurinhas(posicao, tipo string) ([]domain.Figurinha, error) {
	if tipo != "" && !isValidTipo(tipo) {
		return nil, ErrInvalidTipo
	}
	if posicao != "" && !isValidPosicao(posicao) {
		return nil, ErrInvalidPosicao
	}
	return s.repo.FindAll(posicao, tipo)
}

func (s *figurinhaServiceImpl) GetFigurinha(id uint) (*domain.Figurinha, error) {
	figurinha, err := s.repo.FindByID(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, ErrFigurinhaNotFound
	}
	return figurinha, err
}

func (s *figurinhaServiceImpl) UpdateFigurinha(id uint, req domain.UpdateFigurinhaRequest) (*domain.Figurinha, error) {
	figurinha, err := s.GetFigurinha(id)
	if err != nil {
		return nil, err
	}

	if req.Numero != nil {
		figurinha.Numero = *req.Numero
	}
	if req.Tipo != nil {
		if !isValidTipo(string(*req.Tipo)) {
			return nil, ErrInvalidTipo
		}
		figurinha.Tipo = domain.FigurinhaTipo(*req.Tipo)
	}
	if req.Posicao != nil {
		if !isValidPosicao(string(*req.Posicao)) {
			return nil, ErrInvalidPosicao
		}
		figurinha.Posicao = domain.FigurinhaPosicao(*req.Posicao)
	}

	if err := s.repo.Update(figurinha); err != nil {
		return nil, err
	}
	return figurinha, nil
}

func (s *figurinhaServiceImpl) DeleteFigurinha(id uint) error {
	if _, err := s.GetFigurinha(id); err != nil {
		return err
	}
	return s.repo.Delete(id)
}

//Helpers para deixar o código de validação mais limpo, seria bom depois colocar todos esses helpers e erros em uma pasta Utils talvez
func isValidTipo(tipo string) bool {
	return tipo == string(domain.TipoComum) || tipo == string(domain.TipoBrilhante) ||tipo == string(domain.TipoLegendsOuro) ||tipo == string(domain.TipoLegendsBronze)
}

func isValidPosicao(posicao string) bool {
	return posicao == string(domain.PosicaoGoleiro) ||posicao == string(domain.PosicaoZagueiro) ||posicao == string(domain.PosicaoMeioCampista) ||posicao == string(domain.PosicaoAtacante)
}