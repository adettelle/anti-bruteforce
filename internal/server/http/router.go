package internalhttp

import (
	"github.com/go-chi/chi/v5"
)

func NewRouter(bh *BucketHandler) chi.Router {
	r := chi.NewRouter()

	r.Get("/", mainPage)
	r.Get("/check/", bh.check)
	r.Get("/reset/", resetBucket)
	return r
}
