package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/emaaForlin/JasferInventoryClient/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var version string = "dev"

func main() {
	gin.SetMode(gin.ReleaseMode)
	l := log.New(os.Stdout, "JIClient: ", log.LstdFlags)
	l.Println(version)

	err := godotenv.Load()
	if err != nil {
		l.Printf("Error loading .env %q", err)
	}
	addressAndPath := os.Getenv("SRV_ADDR")
	apikey := os.Getenv("APIKEY")
	client := handlers.NewClient(l, addressAndPath, apikey)

	router := gin.Default()
	// Load pages and it assests
	router.LoadHTMLGlob("./templates/*.html")
	router.Static("assets/", "./templates/assets")
	router.Static("edit/assets", "./templates/assets")
	router.Static("config/assets", "./templates/assets")

	// Route the pages
	router.GET("/", client.Index)
	router.GET("/add", client.Add)
	router.GET("/config", client.ConfigPage)

	router.GET("/edit/:id", client.EditPage)
	router.GET("/delete/:id", client.DeleteProduct)

	router.POST("/add", client.AddPost)
	router.POST("/edit/:id", client.EditPost)
	router.POST("/delete/:id", client.DeleteProduct)
	router.POST("/config", client.Config)
	router.POST("/config/updnow", client.UpdNow)

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
		// service connections
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.Println("timeout of 5 seconds.")
	log.Println("Server exiting")
}
