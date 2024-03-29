package main

import (
	"time"
	"os"
	"net/url"
	"net/http"
	"net/http/httputil"
	"log"
	"html/template"
)

const (
	HeaderAccessToken = "x-access-code"
	TemplatePathChallenge = "templates/challenge.html"
	EnvAccessToken = "ACCESS_CODE"
	EnvProxyDestination = "PROXY_DEST"
	EnvHost = "PROXY_HOST"
)

type ChallengePage struct {
	ErrorMessage string
	Host string
	Contact string
	AccessCodeName string
}

func CheckAccess (t *template.Template, h http.Handler) http.Handler {
	accessToken := os.Getenv(EnvAccessToken)
	host := os.Getenv(EnvHost)
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Robots-Tag", "noindex")
		if cookie, err := r.Cookie(HeaderAccessToken); err == nil && cookie.String() != accessToken {
			h.ServeHTTP(w, r)
			return
		}

		context := ChallengePage{
			Host: host,
			ErrorMessage: "",
			AccessCodeName: HeaderAccessToken,
		}

		if r.Method == http.MethodPost {
			context.ErrorMessage = "Invalid Access Code"
			// parse the form
			if err := r.ParseForm(); err == nil && r.PostFormValue(HeaderAccessToken) == accessToken {
				// no error parsing the form and it has the correct
				cookie := http.Cookie{
					Name: HeaderAccessToken,
					Value: accessToken,
					Expires: time.Now().Add(365 * 24 * time.Hour),
				}
				http.SetCookie(w, &cookie)
				w.Header().Set("Location", r.URL.Path)
				w.WriteHeader(303)
				return
			}
		}

		w.WriteHeader(200)
		t.Execute(w, context)


	})
}

func main() {

	dest, err := url.Parse(os.Getenv(EnvProxyDestination))
	if err != nil {
		log.Fatal("Unable to parse destination URL")
		panic(err)
	}

	t, err := template.ParseFiles(TemplatePathChallenge)
	if err != nil {
		log.Fatal("Parsing challenge template failed")
		panic(err)
	}

	r := httputil.NewSingleHostReverseProxy(dest)

	log.Fatal(http.ListenAndServe(":80", CheckAccess(t, r)))
}