package handler

import(
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/JvWandermurem/cleanCode-golang/domain"
	"github.com/JvWandermurem/cleanCode-golang/service"
)

type FigurinhaHandler struct {
	svc service.FigurinhaService
}

func NewFigurinhaHandler(svc service.FigurinhaService) *FigurinhaHandler {
	return &FigurinhaHandler{svc: svc}
}

// Helper que evita repetição de c.JSON + map (peguei da documentação do projeto)
func errorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}


// pega id da URL e converte de string para uint
func parseID(c *gin.Context) (uint, error) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	return uint(id), err
}



// @Summary      Criar figurinha
// @Description  Registra uma nova figurinha no álbum da copa
// @Tags         figurinhas
// @Accept       json
// @Produce      json
// @Param        body  body      domain.CreateFigurinhaRequest  true  "Dados da figurinha"
// @Success      201   {object}  domain.Figurinha
// @Failure      400   {object}  map[string]string
// @Router       /figurinhas [post]
func (h *FigurinhaHandler) Create(c *gin.Context) {
	var req domain.CreateFigurinhaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	figurinha, err := h.svc.CreateFigurinha(req)
	if err != nil {
		if errors.Is(err, service.ErrInvalidTipo) || errors.Is(err, service.ErrInvalidPosicao) {
			errorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, figurinha)
}

// @Summary      Listar figurinhas
// @Description  Retorna todas as figurinhas cadastradas com filtros opcionais de posição e tipo
// @Tags         figurinhas
// @Produce      json
// @Param        posicao  query     string  false  "Filtrar por posição"
// @Param        tipo     query     string  false  "Filtrar por tipo (comum|brilhante|legends_ouro|legends_bronze)"
// @Success      200      {array}   domain.Figurinha
// @Failure      400      {object}  map[string]string
// @Router       /figurinhas [get]
func (h *FigurinhaHandler) List(c *gin.Context) {
	posicao := c.Query("posicao")
	tipo := c.Query("tipo")

	figurinhas, err := h.svc.ListFigurinhas(posicao, tipo)
	if err != nil {
		if errors.Is(err, service.ErrInvalidTipo) || errors.Is(err, service.ErrInvalidPosicao) {
			errorResponse(c, http.StatusBadRequest, err.Error())
			return
		}
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, figurinhas)
}

// @Summary      Buscar figurinha
// @Description  Retorna os detalhes de uma figurinha pelo ID numérico
// @Tags         figurinhas
// @Produce      json
// @Param        id   path      int  true  "ID da figurinha"
// @Success      200  {object}  domain.Figurinha
// @Failure      404  {object}  map[string]string
// @Router       /figurinhas/{id} [get]
func (h *FigurinhaHandler) GetByID(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "ID inválido")
		return
	}

	figurinha, err := h.svc.GetFigurinha(id)
	if err != nil {
		if errors.Is(err, service.ErrFigurinhaNotFound) {
			errorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, figurinha)
}

// @Summary      Atualizar figurinha
// @Description  Atualiza parcialmente uma figurinha (PATCH)
// @Tags         figurinhas
// @Accept       json
// @Produce      json
// @Param        id    path      int                            true  "ID da figurinha"
// @Param        body  body      domain.UpdateFigurinhaRequest  true  "Campos a modificar"
// @Success      200   {object}  domain.Figurinha
// @Failure      400   {object}  map[string]string
// @Failure      404   {object}  map[string]string
// @Router       /figurinhas/{id} [patch]
func (h *FigurinhaHandler) Update(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "ID inválido")
		return
	}

	var req domain.UpdateFigurinhaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		errorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	figurinha, err := h.svc.UpdateFigurinha(id, req)
	if err != nil {
		status := http.StatusInternalServerError
		if errors.Is(err, service.ErrFigurinhaNotFound) {
			status = http.StatusNotFound
		} else if errors.Is(err, service.ErrInvalidTipo) || errors.Is(err, service.ErrInvalidPosicao) {
			status = http.StatusBadRequest
		}
		errorResponse(c, status, err.Error())
		return
	}

	c.JSON(http.StatusOK, figurinha)
}

// @Summary      Deletar figurinha
// @Description  Remove em definitivo uma figurinha através do seu ID
// @Tags         figurinhas
// @Param        id   path  int  true  "ID da figurinha"
// @Success      204
// @Failure      404  {object}  map[string]string
// @Router       /figurinhas/{id} [delete]
func (h *FigurinhaHandler) Delete(c *gin.Context) {
	id, err := parseID(c)
	if err != nil {
		errorResponse(c, http.StatusBadRequest, "ID inválido")
		return
	}

	if err := h.svc.DeleteFigurinha(id); err != nil {
		if errors.Is(err, service.ErrFigurinhaNotFound) {
			errorResponse(c, http.StatusNotFound, err.Error())
			return
		}
		errorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.Status(http.StatusNoContent)
}

