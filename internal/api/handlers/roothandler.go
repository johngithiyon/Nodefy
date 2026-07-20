package handlers

import (
	"net/http"
	"strings"
)

func Roothandler(w http.ResponseWriter, r *http.Request) {

	host := r.Host

	// Remove port
	host = strings.Split(host, ":")[0]

	if host == "nodefy.in" {
		Renderhomepage(w, r)
		return
	}

	if strings.HasSuffix(host, ".workspace.nodefy.in") {
		WorkspaceHanlder(w, r)
		return
	}

	if strings.HasSuffix(host, ".nodefy.in") {
		Deployproxy(w, r)
		return
	}

	http.NotFound(w, r)
}