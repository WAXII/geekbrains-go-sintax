package main

import (
	"io"
	"net/http"
)

func hello(res http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	res.Header().Set("Content-Type", "text/html")
	io.WriteString(res,
		`<doctype html>
<html>
	<head>
    	<title>Hello, `+name+`!</title>
	</head>
	<body>
    	Hello, `+name+`!
	</body>
</html>`)
}

/**
Порт 80 у меня занят, так что ссылка http://localhost:8081/hello?name=World (В методичке ошибка в ссылке)
*/

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8081", nil)
}
