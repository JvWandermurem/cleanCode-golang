// @title           API Figurinhas Copa 2026
// @version         1.0
// @description     Backend em Go utilizando Clean Code para gestão de figurinhas do álbum.
// @host            localhost:8080
// @BasePath        /api/v1
package main

import (
	"log"

	"github.com/JvWandermurem/cleanCode-golang/config"
	"github.com/JvWandermurem/cleanCode-golang/domain"
	"github.com/JvWandermurem/cleanCode-golang/handler"
	"github.com/JvWandermurem/cleanCode-golang/repository"
	"github.com/JvWandermurem/cleanCode-golang/router"
	"github.com/JvWandermurem/cleanCode-golang/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	configuracao := config.Load()

	db := connectDatabase(configuracao.DBPath)

	figurinhaRepo := repository.NewFigurinhaRepository(db)
	figurinhaSvc  := service.NewFigurinhaService(figurinhaRepo)
	figurinhaHdl  := handler.NewFigurinhaHandler(figurinhaSvc)

	r := router.Setup(figurinhaHdl)

	log.Printf("Servidor rodando em: http://localhost:%s", configuracao.Port)
	log.Printf("Documentação em: http://localhost:%s/swagger/index.html", configuracao.Port)

	if err := r.Run(":" + configuracao.Port); err != nil {
		log.Fatal("Erro ao inicializar o servidor: ", err)
	}
}

func connectDatabase(path string) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		log.Fatal("Não foi possível conectar ao banco SQLite: ", err)
	}

	if err := db.AutoMigrate(&domain.Figurinha{}); err != nil {
		log.Fatal("Erro ao executar a migração automática do banco: ", err)
	}

	return db
}