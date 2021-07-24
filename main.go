package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/unrolled/render"
)

// EchoRSP response for /get
type EchoRSP struct {
	Agent   string      `json:"user-agent"`
	Addr    string      `json:"address"`
	Headers interface{} `json:"headers"`
}

// WebApp Main web app
type WebApp struct {
	addr   string
	r      *chi.Mux
	render *render.Render
}

// NewWA creation
func NewWA(addr string) *WebApp {
	return &WebApp{
		addr:   addr,
		r:      chi.NewRouter(),
		render: render.New(),
	}
}

// Run main runner
func (wa *WebApp) Run() {
	wa.r.Use(middleware.RequestID)
	wa.r.Use(middleware.RealIP)
	wa.r.Use(middleware.Recoverer)
	wa.r.Use(middleware.Logger)
	wa.r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome"))
	})
	wa.r.Get("/other", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("opther"))
	})
	wa.r.Route("/v1", func(r chi.Router) {
		r.Get("/get", wa.simpleEcho) // GET /articles/123
	})
	log.Println("Running web mode on: ", wa.addr)
	http.ListenAndServe(wa.addr, wa.r)
}

func (wa *WebApp) simpleEcho(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Header)
	wa.render.JSON(w, http.StatusOK, &EchoRSP{
		Agent:   r.UserAgent(),
		Addr:    r.RemoteAddr,
		Headers: r.Header,
	})

}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// commands
	webCmd := flag.NewFlagSet("web", flag.ExitOnError)

	// Params
	listen := webCmd.String("listen", ":5656", "Address to listen")

	flag.Parse()
	if len(os.Args) < 2 {
		fmt.Println("Command Error: worker or web is required")
	}

	switch os.Args[1] {
	case "web":
		err := webCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal("Error parsing args")
		}
		web := NewWA(*listen)
		web.Run()
	default:
		fmt.Printf("Please use 'web' command")
	}

}
