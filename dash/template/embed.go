package template

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/rs/zerolog/log"
)

//go:generate yarn
//go:generate yarn build
//go:embed all:dist
var content embed.FS

func Dist() http.FileSystem {
	dist, err := fs.Sub(content, "dist")
	if err != nil {
		log.Panic().Err(err).Msg("failed to find dist")
	}
	return http.FS(dist)
}
