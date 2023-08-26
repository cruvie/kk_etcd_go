package test

import (
	"html/template"
	"net/http"
	"os"
	"testing"
)

func TestName(t *testing.T) {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/tmpl", tmpl)
	server.ListenAndServe()
}

type Person struct {
	Name string
	Age  int
}

func tmpl(w http.ResponseWriter, r *http.Request) {
	t1, err := template.ParseFiles("test.html")
	if err != nil {
		panic(err)
	}
	t1.Execute(w, "hello world")
}

func Test1(t *testing.T) {
	p := Person{"longshuai", 23}
	tmpl, _ := template.New("test").Parse("Name: {{.Name}}, Age: {{.Age}}")
	_ = tmpl.Execute(os.Stdout, p)
}
