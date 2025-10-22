package registry

import "net/http"

type NPM struct {}

func (n *NPM) RootRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		n.get(w, r)
	default: 
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Send empty object for root registry
func (n *NPM) get(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{}"))
	w.WriteHeader(http.StatusOK)
}

func (n *NPM) PackageRouter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
