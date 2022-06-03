package main

func main() {
	server := NewServer(":3003")
	server.Handle("GET", "/", HandleRoot)
	//Es importante a los middle llamarlos como funcion, estamos esperando un handler de el
	server.Handle("POST", "/api", server.AddMiddleware(HandleApi, CheckAuth(), Logging()))
	server.Handle("POST", "/create", HandlePostReq)
	server.Handle("POST", "/user", HandleRightJson)
	server.Listen()
}

/*
Si están trabajando con Go en una empresa, puede que su proyecto tenga dependencias privadas de la misma organización.

Para poder acceder a esos módulos privados tendrán que agregar una variable de entorno a su entorno local, de esa forma go clonará los proyectos utilizando git, en vez de fallar cuando intenta buscarlos a través del proxy de módulos de Go.

Agregar la variable de entorno:
export GOPRIVATE="github.com/[el usuario dueño del módulo]"

# Si usas Go 1.13+ podés ejecutar:
go env -w GOPRIVATE="github.com/[el usuario dueño del módulo]"
Una vez que setees GOPRIVATE, Go va a usar git para clonar los repositorios que coincidan con la variable de entorno. Por defecto, Go clona módulos usando http pero para poder autenticarte con Github y poder clonar repositorios privados vas a necesitar ssh. Para forzar a git a usar ssh en Github, necesitas agregar lo siguiente en tu ~/.gitconfig:

[url "ssh://git@github.com/"]
    insteadOf = https://github.com/
Y voilá! Ya podés sincronizar las dependencias privadas de tu proyecto!
*/
