package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/sync/errgroup"

	v1 "homework/internal/api/v1"
	"homework/internal/config"
	airportsService "homework/internal/service/airports"
	bookingService "homework/internal/service/booking"
	flightsService "homework/internal/service/flights"
	airportsStorage "homework/internal/storage/postgresql/airports"
	bookingStorage "homework/internal/storage/postgresql/booking"
	flightsStorage "homework/internal/storage/postgresql/flights"
	"homework/specs"
)

func main() {
	var (
		err         error
		ctx, cancel = signal.NotifyContext(
			context.Background(),
			syscall.SIGHUP,
			syscall.SIGINT,
			syscall.SIGTERM,
			syscall.SIGQUIT,
		)
	)
	defer cancel()

	cfg, err := config.InitConfig(os.Args)
	if err != nil {
		log.Fatal("get config: ", err.Error())
		return
	}

	dbpool, err := pgxpool.New(context.Background(), cfg.Db.Postgresql)
	if err != nil {
		log.Fatalf("Unable to create connection pool: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	// инициализация хранилищ
	airportsStorage := airportsStorage.NewAirportsStorage(dbpool)
	bookingStorage := bookingStorage.NewBookingStorage(dbpool)
	flightsStorage := flightsStorage.NewFlightsStorage(dbpool)

	// инициализация сервисов
	airportsService := airportsService.NewAirportsService(airportsStorage)
	bookingService := bookingService.NewBookingService(bookingStorage)
	flightsService := flightsService.NewFlightsService(flightsStorage)

	apiServer := v1.NewAPIServer(
		airportsService,
		bookingService,
		flightsService,
	)

	err = startHTTPServer(ctx, cfg, apiServer)
	if err != nil {
		log.Fatal("starting server: ", err.Error())
	}
}

func startHTTPServer(
	ctx context.Context,
	cfg *config.Config,
	apiServer specs.ServerInterface,
	middlewares ...specs.MiddlewareFunc,
) error {
	handler := specs.HandlerWithOptions(apiServer, specs.ChiServerOptions{
		BaseURL:     cfg.BasePath,
		Middlewares: middlewares,
	})

	router := chi.NewRouter()
	router.Handle("/*", handler)

	httpServer := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	group.Go(func() error {
		<-ctx.Done()
		return httpServer.Shutdown(ctx)
	})

	return group.Wait()
}
