package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

/*
Esto mi querido amigo es un closure, o lo mas parecido, y en cristiano significa
que hay una funcion anonima dentro de otra que, tiene acceso a variable de scope de la funcion contenedora
... muy loco, no
*/
func CheckAuth() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := false
			fmt.Println("Checking authentication . . .")
			if flag {
				//Aca esto es si cumple con el middleware
				f(w, r)
			} else {
				return
			}
		}
	}
}

func Logging() Middleware {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			defer func() {
				log.Println(r.URL.Path, time.Since(start))
			}()
			//Esto es importante, es como el next() en javascript
			f(w, r)
		}
	}
}
