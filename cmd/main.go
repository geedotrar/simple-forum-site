package main

import (
	"log"

	"github.com/geedotrar/simple-forum-site/internal/configs"
	"github.com/geedotrar/simple-forum-site/internal/handlers/memberships"
	"github.com/geedotrar/simple-forum-site/internal/handlers/posts"
	membershipRepo "github.com/geedotrar/simple-forum-site/internal/repository/memberships"
	postRepo "github.com/geedotrar/simple-forum-site/internal/repository/posts"
	membershipSvc "github.com/geedotrar/simple-forum-site/internal/service/memberships"
	postSvc "github.com/geedotrar/simple-forum-site/internal/service/posts"
	"github.com/geedotrar/simple-forum-site/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	var (
		cfg *configs.Config
	)

	err := configs.Init(
		configs.WithConfigFolder(
			[]string{"./internal/configs/"},
		),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal("Gagal inisiasi config", err)
	}
	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal inisiasi database", err)
	}

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	membershipRepo := membershipRepo.NewRepository(db)
	postRepo := postRepo.NewRepository(db)

	membershipService := membershipSvc.NewService(cfg, membershipRepo)
	postService := postSvc.NewService(cfg, postRepo)

	membershipHandler := memberships.NewHandler(r, membershipService)
	membershipHandler.RegisterRoute()

	postHandler := posts.NewHandler(r, postService)
	postHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
