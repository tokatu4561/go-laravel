module myapp

go 1.17

replace github.com/tsawler/celeritas => ../celeritas

require github.com/tsawler/celeritas v0.0.0

require (
	github.com/go-chi/chi/v5 v5.0.7 // indirect
	github.com/joho/godotenv v1.4.0 // indirect
)
