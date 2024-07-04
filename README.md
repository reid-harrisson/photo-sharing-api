# Simple Golang Backend (GIN)

Simple golang backend service using [Gin](https://gin-gonic.com/) framework.

<img url="https://raw.githubusercontent.com/gin-gonic/logo/master/color.png" width="100px" height="100px"><img/>

## Run Application

1.  Copy `.env.example` file and paste it within same directory
2.  Rename it as `.env`
3.  Run following command

```bash
swag init -g cmd/main.go
```

```bash
go run ./cmd
```

4.  Visit <http://localhost:8000/swagger/index.html> in your web browser to see swagger document
