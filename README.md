# streetmarketAPI

- To start the API, execute "go build" and "./feira-api"
- To use the script to read a CSV file and insert in database, remove the comment block at line 12 in main.go
- To make a request, use "Postman" or another software which supports REST calls

The response statuses gonna be:
- 200 - OK
- 400 - Bad request (if some required field is empty)
- 500 - Internal server error (generic error)

- Its possible to check all logs on "smlogs.csv" at logs folder

- To create a table, you should use the file "table.sql"

- The tests was implemented in CRUD and DB functions