package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	httpadapter "github.com/Mrf-LuckyBoy/test-go/internal/adapters/http"
	"github.com/Mrf-LuckyBoy/test-go/internal/adapters/repository/mariadb"
	"github.com/Mrf-LuckyBoy/test-go/internal/core/usecase"
	"github.com/Mrf-LuckyBoy/test-go/pkg/config"
	"github.com/Mrf-LuckyBoy/test-go/pkg/logger"
)

type Container struct {
	BookHandler *httpadapter.BookHandler
}

func BuildContainer(cfg *config.Config) *Container {
	// setup database connection
	// using mysql/mariadb with gorm
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.DBUser, cfg.DBPass, cfg.DBHost, cfg.DBPort, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(fmt.Errorf("failed to connect database: %w", err))
	}

	// call repository, service, handler constructors
	repo := mariadb.NewBookRepositoryMariaDB(db)
	if err := repo.AutoMigrate(); err != nil {
		logger.L.Println("migration failed:", err)
	}

	svc := usecase.NewBookService(repo)
	bh := httpadapter.NewBookHandler(svc)

	return &Container{BookHandler: bh}
}
