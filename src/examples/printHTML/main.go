package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"strings"
)

func printHTMLheader(Title string) {
	fmt.Println("<!doctype html>")
	fmt.Println("<html lang=fr>")
	fmt.Println("<head>\n<meta charset=\"utf-8\">")
	fmt.Println("<title>", Title, "</title>")
	fmt.Println("</head>")
	fmt.Println("<body>")
}

func printHTMLfooter() {
	fmt.Println("</body></html>")
}

func getGravatarHash(email string) string {
	email = strings.TrimSpace(email)
	email = strings.ToLower(email)
	h := md5.New()
	io.WriteString(h, email)
	finalBytes := h.Sum(nil)
	return hex.EncodeToString(finalBytes)
}

func main() {
	gravatarHash := getGravatarHash("caleb@doxsey.net")

	printHTMLheader("Test")
	fmt.Println("<img src=\"http://www.gravatar.com/avatar/" + gravatarHash + "\"/>")
	printHTMLfooter()
}
