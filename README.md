# GoQuery - Dynamic Query Builder in Golang

GoQuery is a library designed to assist in creating SQL queries in the Go programming language. This library simplifies database operations by leveraging struct structures.

## Usage

To use GoQuery, follow these steps:

1. First, include GoQuery in your project:

   ```go
   import "github.com/cetinboran/goquery"
   ```

2. Initialize the `GoQuery` object and set the required parameters:

   ```go
   // Define the database table name.
   query := goquery.GoQueryInit("your_table_name")

   // Set the struct type.
   yourStruct := YourStructType{}
   query.SetStruct(yourStruct)

   // Set query safety and specify the unique column.
   query.SetChecks([]bool{true, true, true}) // Mark your fields for query safety, for example.
   query.SetUnique("unique_column_name", uniqueValue)
   ```

3. Create SQL queries:
    - Creating an update query:
        ```go
        updateQuery, args := query.CreateUpdate(true) // Use safe query (true)
        ```

        or

        ```go
        updateQuery, args := query.CreateUpdate(false) // Without using safe query
        ```

        - Creating an insert query:

        ```go
        insertQuery, args := query.CreateInsert(true) // Use safe query (true)
        ```

        or

        ```go
        insertQuery, args := query.CreateInsert(false) // Without using safe query
        ```
        
    - detais:  If you set the `safe` parameter to `false`, the argument interface array will indeed be empty, and the values will be directly embedded into the string. In this case, the SQL query will not use parameterized placeholders.
    - However, if you set `safe` to `true`, the `args` interface array will be populated, and the SQL query string will contain placeholders, typically represented as `?`. This allows you to prepare the query in advance and then later supply the values, increasing security by protecting against SQL injection attacks.
    - So, in summary:
        - `safe` set to `false`: Values are directly embedded into the string, and `args` will be empty.
        - `safe` set to `true`: Values are parameterized with placeholders, and `args` will contain the corresponding values, enhancing security by using prepared statements.
        - Here's an example of how this affects the query generation:

    ```go
    // When safe is false:
    query := goquery.GoQueryInit("users")
    // ... Set other parameters
    insertQuery, args := query.CreateInsert(false)
    // insertQuery might look like: "INSERT INTO users (column1, column2) VALUES (value1, value2)"

    // When safe is true:
    query := goquery.GoQueryInit("users")
    // ... Set other parameters
    insertQuery, args := query.CreateInsert(true)
    // insertQuery might look like: "INSERT INTO users (column1, column2) VALUES (?, ?)"
    // args will contain the actual values to be used in place of the placeholders
    ```
    - Using the `safe` parameter with `true` is a recommended practice for building secure SQL queries to prevent SQL injection.

    

4. Use SQL queries:

   ```go
   // Example of executing the query in your database (this is just an example and may vary depending on the database used).
   _, err := db.Exec(updateQuery, args...)
   if err != nil {
       log.Fatal(err)
   }
   ```

## Example Usage

Here is an example of how to create an SQL query using GoQuery:

```go
package main

import (
	"fmt"
	"log"
	"github.com/cetinboran/goquery"
)

type User struct {
	ID       int    `column:"id"`
	Username string `column:"username"`
	Email    string `column:"email"`
}

func main() {
	// Initialize the GoQuery object and set the required parameters.
	query := goquery.GoQueryInit("users")
	user := User{ID: 1, Username: "john_doe", Email: "john@example.com"}
	query.SetStruct(user)
	query.SetChecks([]bool{true, true, true})
	query.SetUnique("id", user.ID)

	// Create an update query.
	updateQuery, args := query.CreateUpdate(true)

	// Display the generated query.
	fmt.Println("Update Query:", updateQuery)
	fmt.Println("Query Arguments:", args)

	// Execute the query in the database (this is just an example and may vary depending on the database used).
	_, err := db.Exec(updateQuery, args...)
	if err != nil {
		log.Fatal(err)
	}
}
```

## Contribution

If you'd like to contribute to this project, please fork it and submit a pull request. Feel free to report any bugs or issues on the GitHub issues page as well.

## License

This project is licensed under the MIT License. For more information, please see the [LICENSE](LICENSE) file.

---

This README file provides essential information to get started with the GoQuery library. Don't forget to create documentation and guides specific to your project for more details.