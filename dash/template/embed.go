package template

import (
	"embed"
	"github.com/rs/zerolog/log"
	"io/fs"
	"net/http"
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
