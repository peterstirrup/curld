# curld
Got a failing GET request to a black box database?

curld systematically removes parameters and attempts the GET request. Useful if you've got a request with 10+ parameters.

## Build & Run
```
go build && ./curld
```