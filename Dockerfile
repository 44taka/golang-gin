FROM golang:1.17-bullseye

RUN apt update
RUN apt -y install netcat

RUN mkdir -p /go/src/github.com/44taka/golang-gin

# TODO:go.modがあるからこの辺りいらない。整理すること。
RUN go get -u github.com/gin-gonic/gin
RUN go get -u github.com/cosmtrek/air
RUN go get -u gorm.io/driver/mysql
RUN go get -u gorm.io/driver/postgres
RUN go get -u gorm.io/gorm
RUN go get -u github.com/uudashr/gopkgs/v2/cmd/gopkgs \
  github.com/ramya-rao-a/go-outline \
  github.com/google/uuid \
  github.com/nsf/gocode \
  github.com/acroca/go-symbols \
  github.com/fatih/gomodifytags \
  github.com/josharian/impl \
  github.com/haya14busa/goplay/cmd/goplay \
  github.com/go-delve/delve/cmd/dlv \
  golang.org/x/lint/golint \
  golang.org/x/tools/gopls \
  github.com/dgrijalva/jwt-go \ 
  github.com/appleboy/gin-jwt/v2 \
  github.com/stretchr/testify \
  github.com/joho/godotenv
RUN go install github.com/golang/mock/mockgen@latest

COPY ./app /go/src/github.com/44taka/golang-gin
WORKDIR /go/src/github.com/44taka/golang-gin

RUN go build main.go

ENTRYPOINT ["./main"]