package main

import (
    "database/sql"
    "flag"
    "log"
    "net/http"
    "os"

    _ "github.com/go-sql-driver/mysql"
)

// Define an application struct to hold the application-wide dependencies for the
// web application. For now we'll only include fields for the two custom loggers, but
// we'll add more to it as the build progresses.
type application struct {
  errorLog *log.Logger
  infoLog *log.Logger
}

func main() {
  // Define a new command-line flag with the name 'addr', a default value of ":4000"
  // and some short help text explaining what the flag controls. The value of the
  // flag will be stored in the addr variable at runtime. 
  addr := flag.String("addr", ":4000", "HTTP network address")
  // Define a new command-line flag for the MySQL DSN string
  dsn := flag.String("dsn", "web:tydeptrai@/snippetbox?parseTime=true", "MySQL data source name")

  // Importantly, we use the flag.Parse() function to parse the command-line flag.
  // This reads in the command-line flag value and assigns it to the addr
  // variable. You need to call this *before* you use the addr variable
  // otherwise it will always contain the default value of ":4000". If any errors are
  // encountered during parsing the application will be terminated.
  flag.Parse()

  infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
  errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Lshortfile)

  db, err := openDB(*dsn)
  if err != nil {
    errorLog.Fatal(err)
  }

  // We also defer a call to db.Close(), so that the connection pool is closed
  // before the main() function exits.
  defer db.Close()

  app := &application{
    errorLog: errorLog,
    infoLog: infoLog,
  }

  srv := &http.Server{
    Addr: *addr,
    ErrorLog: errorLog,
    // Call the new app.routes() method to get the servemux containing our routes.
    Handler: app.routes(),
  }

  infoLog.Printf("Starting server on %s", *addr)
  err = srv.ListenAndServe()
  errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
  db, err := sql.Open("mysql", dsn)
  if err != nil {
    return nil, err
  }
  if err = db.Ping(); err != nil {
    return nil, err
  }

  return db, nil
}
