# curld
Got a failing GET request to a black box database?

curld systematically removes parameters, attempts the GET request and tells you what's making the request fail. Useful if you've got a request with many parameters.

## Build & Run
```
go build && ./curld
```
