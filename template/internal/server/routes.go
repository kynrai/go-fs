package server

import (
	"module/placeholder/internal/server/assets"
	"module/placeholder/internal/server/handlers"
	"net/http"
)

func (s *Server) Routes() {
	// filserver route for assets
	assetMux := http.NewServeMux()
	assetMux.Handle("GET /*", http.StripPrefix("/assets/", http.FileServer(http.FS(assets.FS))))
	s.r.Handle("GET /assets/*", assetMux)

	// handlers for normal routes with all general middleware
	routesMux := http.NewServeMux()
	routesMux.Handle("GET /", handlers.PageIndex())
	routesHandler := s.middlewares(routesMux)

	s.r.Handle("GET /", routesHandler)

	s.srv.Handler = s.r
}
