package api

import (
	"adams549659584/go-proxy-bingai/api/helper"
	"adams549659584/go-proxy-bingai/common"
	"net/http"
	"strings"
	"os"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, common.PROXY_WEB_PAGE_PATH, http.StatusFound)
		return
	}
	if strings.HasPrefix(r.URL.Path, "/turing") {
		if !helper.CheckAuth(r) {
			helper.UnauthorizedResult(w)
			return
		}
	}
	bingURL := os.Getenv("BING_PROXY_DM")
	if bingURL == "" {
		bingURL = common.BING_URL
	}
	url, _ := url.Parse(bingURL)
	common.NewSingleHostReverseProxy(url).ServeHTTP(w, r)
}

