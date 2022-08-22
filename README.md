# streetmarketAPI

- To start the API, execute "go build" and "./feira-api"
- To use the script to read a CSV file and insert in database, remove the comment block at line 12 in main.go
- To make a request, use "Postman" or another software which supports REST calls

- Its possible to check all logs on "smlogs.csv" at logs path

- To create a table, you should use the file "table.sql"

- The tests was implemented in CRUD and DB functions

routes:
POST, PUT (UPDATE) AND DELETE:
- /streetmarket/ - This endpoint is gonna make a crud of a streetmarket (depends what you put on request)

The response statuses can be:
- 200 - OK
- 400 - Bad request (if some required field is empty)
- 500 - Internal server error (generic error)