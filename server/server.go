package server

import (
	"fmt"
	"net/http"
	"pkg-registry/server/api"
)

const ADDR = "localhost"
const PORT = "8081"

func StartServer() error {
	topMux := http.NewServeMux()
	apiMux := api.NewApiMux()
	registryMux := api.NewApiRegistryMux()

	topMux.Handle("/api/", http.StripPrefix("/api", apiMux))
	topMux.Handle("/registry/", http.StripPrefix("/registry", registryMux))

	fmt.Println("Server running on http://" + ADDR + ":" + PORT)
	err := http.ListenAndServe(ADDR+":"+PORT, topMux)
	return err
}
