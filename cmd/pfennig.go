package cmd

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	log "github.com/sirupsen/logrus"
	flag "github.com/spf13/pflag"

	"github.com/embik/pfennig/pkg/data"
    "github.com/embik/pfennig/pkg/router"
    mw "github.com/embik/pfennig/pkg/middleware"
)

func StartPfennig() {
    log.Println("Starting pfennig")

    var ip net.IP
	var port int
	var dbPath string
	var wait time.Duration

	flag.IPVar(&ip, "bind-ip", net.IPv4(127, 0, 0, 1), "ip address to bind the server to")
	flag.IntVar(&port, "bind-port", 8080, "port to bind the server to")
	flag.DurationVar(&wait, "graceful-timeout", time.Second*15, "the duration for which the server gracefully waits for existing connections to close")
	flag.StringVar(&dbPath, "db-path", "pfennig.db", "location for sqlite database file")
	flag.Parse()

    log.Info("initalizing database")
	err := data.InitDB(dbPath)
	if err != nil {
		panic(err)
	}
	defer data.CloseDB()

    log.Info("creating dummy data")
	data.CreateDummyData()

    log.Info("starting api server")
    r := router.GetRouter()
	srv := &http.Server{
		Handler: mw.RequestLogging(r),
		Addr: fmt.Sprintf("%v:%v", ip, port),
	}

	go func() {
		log.Fatal(srv.ListenAndServe())
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	srv.Shutdown(ctx)

	log.Println("Shutting down")

	os.Exit(0)
}

