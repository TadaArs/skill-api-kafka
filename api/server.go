package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"skill/database"
	router "skill/route"
	"skill/skill"
	"syscall"
	"time"
)

func main() {

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	db := database.Connect()
	defer db.Close()

	storage := skill.NewSkillStorage(db)
	pd := skill.NewProducer()
	skillHandler := skill.NewSkillHandler(storage, *pd)

	r := router.NewRouter(skillHandler)

	srv := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: r,
	}

	closeChan := make(chan struct{})

	go func() {
		<-ctx.Done()
		fmt.Println("shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				log.Println(err)
			}
		}

		close(closeChan)
	}()

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}

	<-closeChan

}
