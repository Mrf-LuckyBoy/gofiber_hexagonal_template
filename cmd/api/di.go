package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/Mrf-LuckyBoy/test-go/internal/adapters/cache"
	httpadapter "github.com/Mrf-LuckyBoy/test-go/internal/adapters/http/handlers"
	"github.com/Mrf-LuckyBoy/test-go/internal/adapters/repository/mariadb"
	"github.com/Mrf-LuckyBoy/test-go/internal/adapters/thirdparty"
	"github.com/Mrf-LuckyBoy/test-go/internal/core/domain"
	"github.com/Mrf-LuckyBoy/test-go/internal/service"
	"github.com/Mrf-LuckyBoy/test-go/pkg/config"
)

type Container struct {
	BookHandler *httpadapter.BookHandler
	AuthHandler *httpadapter.AuthHandler
	UserHandler *httpadapter.UserHandler
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
	db.AutoMigrate(
		&domain.Book{},
	)

	memCache := cache.NewRistrettoCache()

	// call repository, service, handler constructors
	bookRepo := mariadb.NewBookRepositoryMariaDB(db)
	bookService := service.NewBookService(bookRepo, memCache)
	userClient := thirdparty.NewUserAPIClient("https://6785e2a7f80b78923aa4afb7.mockapi.io/api/v1")
	userService := service.NewUserService(userClient, memCache)
	// handler
	bookHandler := httpadapter.NewBookHandler(bookService)
	authHandler := httpadapter.NewAuthHandler(cfg.JWTSecret)
	userHandler := httpadapter.NewUserHandler(userService)

	return &Container{BookHandler: bookHandler, AuthHandler: authHandler, UserHandler: userHandler}
}

// bookService
