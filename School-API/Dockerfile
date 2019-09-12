FROM golang as builder

# Labels
LABEL maintainer="Vaibhav Kulshrestha"
LABEL version="1.0.0"
LABEL description="API Service"

# Dependencies
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/gorilla/handlers
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/go-sql-driver/mysql

# Path
WORKDIR /src/github.com/vaibhav-k/Go-WebService/School-API
ADD . /src/github.com/vaibhav-k/Go-WebService/School-API

CMD ["go", "run", "main.go"]