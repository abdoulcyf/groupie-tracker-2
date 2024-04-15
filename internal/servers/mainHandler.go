package servers

import (
	"fmt"
	"main/internal/handlers"
	"net/http"
	"strings"
)

func MainHandler(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.URL.Path == "/":
		fmt.Println("/home-page")
		handlers.HomeHandler(w, r)
	case strings.HasPrefix(r.URL.Path, "/artist/"):
		fmt.Println("/artist/")
		handlers.ArtistHandler(w, r)
	// case "/ascii-art":
	// 	fmt.Println("/ascii-art")
	// 	handlers.AsciiToArt_handler(w, r)
	default:
		w.WriteHeader(404)
		handlers.ErrorPage(w, NotFound404)
	}
}
