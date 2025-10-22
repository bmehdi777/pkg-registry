package api

import (
	"net/http"
	"pkg-registry/server/api/registry"
)

func NewApiMux() *http.ServeMux {
	apiMux := http.NewServeMux()

	health := Health{}
	apiMux.HandleFunc("/health", health.router)

	return apiMux
}

func NewApiRegistryMux() *http.ServeMux {
	apiRegistryMux := http.NewServeMux()

	npm := registry.NPM{}
	apiRegistryMux.HandleFunc("/npm", npm.RootRouter)
	apiRegistryMux.HandleFunc("/npm/{packageName}", npm.PackageRouter)
	apiRegistryMux.HandleFunc("npm/{packageName}/{version}", npm.PackageVersionRouter)

	return apiRegistryMux
}
