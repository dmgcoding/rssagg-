package main

import (
	"fmt"
	"os"
	"log"
	"github.com/joho/godotenv"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"net/http"
)

func main(){
	godotenv.Load()
	 
	portString := os.Getenv("PORT")

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*","http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"*"},
		AllowCredentials: true,
	}))

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz",handlerReadiness)
	v1Router.Get("/error",handleError)

	router.Mount("/v1",v1Router)
	
	fmt.Println("server starting at port:",portString)

	srv := &http.Server{
		Handler: router,
		Addr: ":"+portString,
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}