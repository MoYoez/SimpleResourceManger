package main

import (
	endpoint "github.com/moyoez/starry-backend/endpoint"
	_ "github.com/moyoez/starry-backend/handler"
	_ "github.com/moyoez/starry-backend/sql"
)

// This Server Backend is for Starry RESMANAGER controller here, Restful API Design
// USE: GO-GIN + SQLITE + JWT USER AUTH TOKEN.

// BACKEND FINISHED.

func main() {
	// RUN GIN SERVICE.
	endpoint.RunGinServer()
}
