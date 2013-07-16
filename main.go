package main

import (
	"net/http"
	"html/template"
"log"
"os"
)

var page = template.Must(template.New("T").Parse(`
<form action=/ method=post>
<textarea name=message>
</textarea>
<input type=submit value=Post>
</form>

<p>Your message was:</p>
<div id=dest></div>

<script>
var data = {{.}};

window.onload = function() {
	var d = document.getElementById('dest');
	d.textContent = data.UserInput;
};
</script>
`))

func main() {
	http.HandleFunc("/", handle)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	var data struct {
		SomeNumber int
		UserInput string
	}
	data.SomeNumber = 2
	data.UserInput = r.PostFormValue("message")
	w.Header().Set("Content-Type", "text/html")
	err := page.Execute(w, data)
	if err != nil {
		log.Fatal(err)
	}
}
