package main

import (
	"io"
	"net/http"

	"github.com/muly/golang/authentication/oauth/fb"
	"github.com/muly/golang/authentication/oauth/github"
	"github.com/muly/golang/authentication/oauth/gitlab"
	"github.com/muly/golang/authentication/oauth/google"
	"github.com/muly/golang/authentication/oauth/linkedin"
)

const loginhtml = `<!DOCTYPE html>
<html>
<head></head>
<body>
<p><a href="/facebooklogin">LOGIN WITH FACEBOOK </a></p>
<p><a href="/githublogin">LOGIN WITH GITHUB</a></p>
<p><a href="/gitlablogin">LOGIN WITH GITLAB</a></p>
<p><a href="/googlelogin">LOGIN WITH GOOGLE</a></p>
<a href="/linkedinlogin">LOGIN WITH LINKEDIN</a>
</body>
</html>
`

func main() {
	http.HandleFunc("/", handleindex)
	http.HandleFunc("/githublogin", github.HandleLogin)
	http.HandleFunc("/gitlablogin", gitlab.HandleLogin)
	http.HandleFunc("/facebooklogin", fb.HandleLogin)
	http.HandleFunc("/linkedinlogin", linkedin.HandleLogin)
	http.HandleFunc("/googlelogin", google.HandleLogin)
	http.ListenAndServe(":8081", nil)
}

func handleindex(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, loginhtml)
}
