# GoQuery - Dynamic Query Builder in Golang

GoQuery is a Go (Golang) package that provides a dynamic query builder for SQL statements. It allows you to easily create UPDATE and INSERT statements based on the provided column names, values, and checks. This package aims to simplify the process of generating SQL queries dynamically in your Go applications.

## Installation

To use GoQuery in your Go project, you can install it using `go get`:

```shell
go get github.com/your-username/goquery
```

## Usage

### Importing the Package

First, you need to import the GoQuery package in your Go code:

```go
import (
    "github.com/cetinboran/goquery"
    // Other necessary imports
)
```

### Creating a GoQuery Instance

You can create a GoQuery instance for a specific table by calling the `GoQueryInit` function:

```go
query := goquery.GoQueryInit("your_table_name")
```

### Setting Column Names and Values

You can set the column names and values for your query using the following methods:

```go
query.SetColmnNames([]string{"column1", "column2", "column3"})
query.SetValues([]interface{}{value1, value2, value3})
```

### Setting Checks

Checks are used to determine whether a specific column should be included in the query. You can set checks using the `SetChecks` method:

```go
query.SetChecks([]bool{true, false, true}) // Include column1 and column3 in the query
```

### Setting Unique Identifier

You can set a unique identifier and its corresponding value for the WHERE clause of your query:

```go
query.SetUniqueString("id")     // The column name for the unique identifier
query.SetUniqueValue(uniqueID)  // The value of the unique identifier
```

### Creating an UPDATE Query

To create an UPDATE query, you can use the `CreateUpdate` method:

```go
queryStr, args := query.CreateUpdate(true) // true for safe query with placeholders
```

To create an unsafe query without placeholders, set the argument to `false`.

### Creating an INSERT Query

To create an INSERT query, you can use the `CreateInsert` method:

```go
queryStr, args := query.CreateInsert(true) // true for safe query with placeholders
```

To create an unsafe query without placeholders, set the argument to `false`.

## Example

Here's a simple example of creating an UPDATE query:

```go
package main

import (
    "fmt"
    "github.com/cetinboran/goquery"
)

func main() {
    query := goquery.GoQueryInit("users")
    query.SetColmnNames([]string{"name", "email"})
    query.SetValues([]interface{}{"John Doe", "john.doe@example.com"})
    query.SetChecks([]bool{true, true})
    query.SetUniqueString("id")
    query.SetUniqueValue(1)

    queryStr, args := query.CreateUpdate(true)

    fmt.Println("Generated UPDATE query:")
    fmt.Println(queryStr)
    fmt.Println("Query arguments:")
    fmt.Println(args)
}
```

This will generate the following output:

```
Generated UPDATE query:
UPDATE users SET name = ?, email = ? WHERE id = ?
Query arguments:
[John Doe john.doe@example.com 1]
```

## Author

This GoQuery package was created by [cetinboran](https://github.com/cetinboran).
