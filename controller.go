package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"path"

	"github.com/samuel/go-zookeeper/zk"
)

// Controller handles HTTP requests.
type Controller struct {
	zkConn *zk.Conn
}

func (c *Controller) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := cleanPath(req.URL.Path)
	switch req.Method {
	case "GET":
		log.Printf("retrieving node %q", path)
		contents, _, err := c.zkConn.Get(path)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(contents)
		return
	case "OPTIONS":
		// TODO: Implement CORs correctly
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "DNT,X-CustomHeader,Keep-Alive,User-Agent,X-Requested-With,If-Modified-Since,Cache-Control,Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}
}

func cleanPath(str string) string {
	if str == "" || str == "/" {
		return str
	}
	return path.Clean(str)
}
