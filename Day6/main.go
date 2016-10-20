package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", index)
	http.Handle("/image/", http.StripPrefix("./public/images", http.FileServer(http.Dir("public"))))
	http.ListenAndServe(":8080", nil)
}

func index(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("content-type", "text/html")
	io.WriteString(res, `<img src="/image/go.png" style="width: 125px;">`)
	http.Redirect(res, req, "/", 302)
}

// func image(res http.ResponseWriter, req *http.Request) {
// 	b, err := os.Open("./public/images/go.png")
// 	image, err := ioutil.ReadAll(b)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	res.Write(image)
// }
