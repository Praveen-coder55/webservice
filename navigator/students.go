package handlers

import (
	"log"
	"net/http"

	"github.com/Praveen-coder55/webservice.git/data"
)

type Student struct {
	l *log.Logger
}

func NewStudents(l *log.Logger) *Student {
	return &Student{l}
}

func (q *Student) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// ls := data.GetStudents()
	// err := ls.ToJSON(rw)

	// if err != nil {
	// 	http.Error(rw, "Unable to find json", http.StatusInternalServerError)
	// }

	if r.Method == http.MethodGet {
		q.getStudents(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)

}

func (q *Student) getStudents(rw http.ResponseWriter, r *http.Request) {
	ls := data.GetStudents()
	err := ls.ToJSON(rw)

	if err != nil {
		http.Error(rw, "Unable to find json", http.StatusInternalServerError)
	}

}
