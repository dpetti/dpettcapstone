package main
import(
   "database/sql"
  _ "github.com/go-sql-driver/mysql"
//   "fmt"
   "log"
)
func main() {


    // Open up our database connection.
    // I've set up a database on my local machine using phpmyadmin.
    // The database is called testDb
    db, err := sql.Open("mysql", "dpetti:isc496@/test")

    // if there is an error opening the connection, handle it
    if err != nil {
        panic(err.Error())
    }

    // defer the close till after the main function has finished
    // executing
    defer db.Close()
    err = db.Ping()

        if err != nil {
            log.Fatalln(err)

        } else {
            log.Println("mysqld is WORKING!!!!")

        }


}
