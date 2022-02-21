package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/emaaForlin/JasferInventoryClient/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	l := log.New(os.Stdout, "JISoftware-prototype: ", log.LstdFlags)

	client := handlers.NewClient(l, "http://localhost:9090/products")

	router := gin.Default()

	// Load pages and it assests
	router.LoadHTMLGlob("./templates/*.html")
	//router.LoadHTMLFiles("./templates/index.html", "./templates/add.html", "./templates/update.html")
	router.Static("assets/", "./templates/assets")
	router.Static("edit/assets", "./templates/assets")

	// Route the pages
	router.GET("/", client.Index)
	router.GET("/add", client.Add)
	router.GET("/edit/:id", client.EditGet)
	router.GET("/delete/:id", client.DeleteProduct)

	router.POST("/add", client.AddPost)
	router.POST("/edit/:id", client.EditPost)
	router.POST("/delete/:id", client.DeleteProduct)

	// all the stuff needed to start serving the page are down here
	// setting up http server
	s := &http.Server{
		Addr:         ":9091",           // configure the bind address
		Handler:      router,            // set the default handler
		ErrorLog:     l,                 // set logger for the server
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
		ReadTimeout:  5 * time.Second,   // max time to read requests from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
	}

	go func() {
		err := s.ListenAndServe()
		l.Printf("Server listening on %s", s.Addr)

		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
