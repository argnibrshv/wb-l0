package app

import (
	"context"
	"database/sql"
	"net/http"
	"sub/api"
	"sub/api/middleware"
	"sub/internals/app/cache"
	"sub/internals/app/db"
	"sub/internals/app/handlers"
	"sub/internals/app/processors"
	"sub/internals/app/subscribers"
	"sub/internals/cfg"
	"time"

	stan "github.com/nats-io/stan.go"
	log "github.com/sirupsen/logrus"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type Server struct {
	ctx    context.Context
	sc     stan.Conn
	sub    stan.Subscription
	srv    *http.Server
	db     *sql.DB
	cache  *cache.Cache
	config cfg.Cfg
}

func NewServer(ctx context.Context, config cfg.Cfg) *Server {
	server := new(Server)
	server.ctx = ctx
	server.config = config
	return server
}

func (server *Server) Serve() {
	log.Println("Starting server")

	var err error
	server.db, err = sql.Open("pgx", server.config.GetDBURL())
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Database connected")

	server.sc, err = stan.Connect("test-cluster", "test-sub")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Nats-streaming connected")

	server.cache = cache.NewCache()
	ordersStorage := db.NewOrdersStorage(server.db)
	subscriber := subscribers.NewSubscriber(&server.sc, ordersStorage, server.cache)

	server.sub, err = subscriber.Subscribe()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Subscribed to the Nats-streaming channel")

	ordersProcessor := processors.NewOrdersProcessor(ordersStorage, server.cache)
	ordersHandler := handlers.NewOrdersHandler(ordersProcessor)

	routes := api.CreateRoutes(ordersHandler)
	routes.Use(middleware.RequestLog)

	server.srv = &http.Server{
		Addr:    ":" + server.config.Port,
		Handler: routes,
	}

	log.Println("Server started on port:", server.config.Port)

	err = server.srv.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		log.Fatalln(err)
	}

	log.Printf("server exited properly")
}

func (server *Server) ShutDown() {
	log.Printf("server stopped")
	ctxShutDown, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	server.db.Close()
	server.sc.Close()
	defer func() {
		cancel()
	}()
	if err := server.srv.Shutdown(ctxShutDown); err != nil {
		log.Fatal("server shutdown failed:", err)
	}
}
