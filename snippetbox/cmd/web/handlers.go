package main

import (
    "fmt"
    "html/template"
    "net/http"
    "strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
  if r.URL.Path != "/" {
    app.notFound(w) //use the notFound() helper
      return
  }

  files := []string{
    "./ui/html/base.tmpl",
    "./ui/html/partials/nav.tmpl",
    "./ui/html/pages/home.tmpl",
  }
  
  // Use the template.ParseFiles() function to read the template file into a
  // template set. If there's an error, we log the detailed error message and use
  // the http.Error() function to send a generic 500 Internal Server Error
  // response to the user.
  ts, err := template.ParseFiles(files...)
  if err != nil {
    app.serverError(w, err) // Use the serverError() helper
  }

  // We then use the Execute() method on the template set to write the
    // template content as the response body. The last parameter to Execute()
    // represents any dynamic data that we want to pass in, which for now we'll
    // leave as nil.
  err = ts.ExecuteTemplate(w, "base", nil)
  if err != nil {
    app.serverError(w, err) // Use the serverError() helper
  }
}

// Change the signature of the snippetView handler so it is defined as a method
// against *application.
func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
  id, err := strconv.Atoi(r.URL.Query().Get("id"))
  if err != nil || id < 1 {
    app.notFound(w) // Use the notFound() helper
    return
  }

  fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

// Change the signature of the snippetCreate handler so it is defined as a method
// against *application.
func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {
  if r.Method != http.MethodPost {
    w.Header().Set("Allow", http.MethodPost)
    app.clientError(w, http.StatusMethodNotAllowed)
    return
  }

  w.Write([]byte("Create a new snippet..."))
}
