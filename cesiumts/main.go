package main

import (
	"io"
	"net/http"
	"os"
	"path"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	ctsHandlers "github.com/geo-data/cesium-terrain-server/handlers"
	"github.com/geo-data/cesium-terrain-server/stores/fs"
	"github.com/gorilla/mux"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func main() {
	tilesetRoot := getEnv("TILESET_ROOT", ".")
	baseTerrainUrl := getEnv("BASE_TERRAIN_URL", "/tilesets")

	// Workaround to tilesetRoot not being current working directory: Change CWD to tilesetRoot.
	err := os.Chdir(path.Join(tilesetRoot))
	if err != nil {
		panic(err)
	}

	store := fs.New(tilesetRoot)

	r := mux.NewRouter()
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "PONG")
	})
	r.HandleFunc(baseTerrainUrl+"/{tileset}/layer.json", ctsHandlers.LayerHandler(store))
	r.HandleFunc(baseTerrainUrl+"/{tileset}/{z:[0-9]+}/{x:[0-9]+}/{y:[0-9]+}.terrain", ctsHandlers.TerrainHandler(store))

	handler := ctsHandlers.AddCorsHeader(r)
	http.Handle("/", handler)

	lambda.Start(httpadapter.New(http.DefaultServeMux).ProxyWithContext)
}
