## Bitly-go-fiber
```
RUN: go mod init .
RUN: go mod tidy
```

``` golang
# Running with nodemon
nodemon --exec go run main.go --signal SIGTERM
```

### Endpoints
``` golang
router.Get("/goly",GetGollies)             // all golinks
      .Get("/goly/:id", GetGoly)           // single golink
      .Post("/goly", CreateGoly)           // create golink
      .Patch("/goly", UpdateGoly)          // update golink
      .Delete("/goly/:id", DeletaGoly)     // delete single golink
      .Get("/r/:redirect", Redirect)       // get redirect link and update click
