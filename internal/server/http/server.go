package internalhttp

import (
	"net/http"
	"time"

	"github.com/adettelle/anti-bruteforce/config"
	"github.com/adettelle/anti-bruteforce/internal/storage/memstorage"
)

type AttemptsRegistry interface {
	Register(login string, pwd string, ip string) bool
}

type Storager interface {
	CheckInWhiteList(ip string) bool
}

type Server struct {
	Srv *http.Server
	Cfg *config.Config
}

type BucketHandler struct {
	StoragerService AttemptsRegistry
}

func NewServer(cfg *config.Config) *Server {
	// rate := 10 * time.Millisecond
	// TODO как передать capacity, rate, added?
	capacityForLogin := 5 // 10
	capacityForPwd := 7   // 100
	capacityForIP := 10   // 1000
	memStorager := memstorage.New(capacityForLogin, capacityForPwd, capacityForIP)

	bucketHandler := BucketHandler{
		StoragerService: memStorager,
	}

	router := NewRouter(&bucketHandler)

	addr := "0.0.0.0:" + cfg.Port

	srv := http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	return &Server{Srv: &srv}
}

func mainPage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("start")) //nolint
}

// http://localhost:8080/check/?login=user1&&password=123&&ip=111.222.333.444
func (bh *BucketHandler) check(w http.ResponseWriter, r *http.Request) {
	// query-параметр
	login := r.URL.Query().Get("login")
	pwd := r.URL.Query().Get("password")
	ip := r.URL.Query().Get("ip")

	ok := bh.StoragerService.Register(login, pwd, ip) // CheckInWhiteList(ip)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	// w.Write([]byte(fmt.Sprintf("login = %s\n", login)))
	// w.Write([]byte(fmt.Sprintf("pwd = %s\n", pwd)))
	// w.Write([]byte(fmt.Sprintf("ip = %s\n", ip)))
}

func resetBucket(w http.ResponseWriter, r *http.Request) {

}
