package main

import (
	"html/template"
	"log"
	"net/http"
)

type Postre struct {
	Nombre      string
	Descripcion string
	Precio      float64
	Imagen      string
}

func main() {
	// Servir archivos estáticos (CSS, JS, imágenes, etc.)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Rutas
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/catalogo", catalogoHandler)
	http.HandleFunc("/contacto", contactoHandler)

	// Iniciar el servidor
	log.Println("Server started on: http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about", nil)
}

func catalogoHandler(w http.ResponseWriter, r *http.Request) {
	postres := []Postre{
		{"Brownies", "Chocolatosos y húmedos, puedes agregarle el topping que quieras.", 8, "/static/imagenes/cata/braunis.jpg"},
		{"Berlines", "Una deliciosa masa esponjosa y rellenos de crema pastelera o dulce de leche.", 3.50, "/static/imagenes/cata/berlines.jpg"},
		{"Mini-donas", "Esponjosas y puedes escoger el color que quieras para su decoración.", 2, "/static/imagenes/cata/donas.jpg"},
		{"Volcán de chocolate", "Bizcocho de chocolate cubierto de chocolate (estilo torta Matilda).", 50, "/static/imagenes/cata/volcan.jpg"},
		{"Tres leche", "Delicioso bizcocho de vainilla, acompañado de 3 leches, con dulce de leche y crema chantilly.", 10, "/static/imagenes/cata/tres leche.jpg"},
		{"Torta red velvet", "La torta terciopelo rojo acompañada con buttercream de queso crema y un toque de dulce de leche.", 30, "/static/imagenes/cata/red.jpg"},
	}
	renderTemplate(w, "catalogo", postres)
}

func contactoHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "contacto", nil)
}

func renderTemplate(w http.ResponseWriter, tmpl string, data interface{}) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl + ".html")
	err := parsedTemplate.Execute(w, data)
	if err != nil {
		log.Printf("Error occurred while executing the template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
