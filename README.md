# Item Application

This app is design to practice Golang. Application is created using Mux module and routed to perform CRUD operations.

Endpoints
```
/item       GET, POST, PUT, DELETE
/items      GET
```


## To run the app

*With Docker*
1. In the root directory, execute `docker build -t <image name> .`
2. Once compiled, execute `docker run -d -p 4040:4040 <image name>`
3. App should be runnning and exposed to `http://localhost:4040`

*Without Docker*
1. In the root directory, execute `go intall` to download necessary dependencies
2. Then, execute `go run main.go`
3. App should be runnning and exposed to `http://localhost:4040`