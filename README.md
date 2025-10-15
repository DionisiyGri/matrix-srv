## Matrix API
| Endpoint | Description | Example Command |
|-----------|--------------|-----------------|
| **/echo** | Returns the matrix as uploaded | ```bash curl -F 'file=@/path/to/file.csv' http://localhost:8080/echo ``` |
| **/invert** | Transposes the matrix (swaps rows and columns) | ```bash curl -F 'file=@/path/to/file.csv' http://localhost:8080/invert ``` |
| **/flatten** | Returns all elements as a single comma-separated line | ```bash curl -F 'file=@/path/to/file.csv' http://localhost:8080/flatten ``` |
| **/sum** | Returns the sum of all matrix elements | ```bash curl -F 'file=@/path/to/file.csv' http://localhost:8080/sum ``` |
| **/multiply** | Returns the product of all matrix elements | ```bash curl -F 'file=@/path/to/file.csv' http://localhost:8080/multiply ``` |


### Run server
- Simply run: from projet directory ```go run main.go```
- Build executable ```go build -o matrix maing.go``` and run ```./matrix```
  
### Tests
```go test ./...``` - will execute all tests and output the results, including any error cases or expected output mismatches


### Notes
- The file must follow a valid matrix format with equal columns per row.
- Only integer values are supported.
- Invalid or missing files will return a 400 Bad Request.
