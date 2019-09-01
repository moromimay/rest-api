FROM golang

ADD . /go/src/rest-api

RUN go get github.com/gorilla/mux
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/jinzhu/gorm
RUN go get golang.org/x/crypto/bcrypt
RUN go get github.com/jinzhu/gorm/dialects/mysql
RUN go get github.com/joho/godotenv

RUN go install rest-api

ENTRYPOINT /go/bin/rest-api

EXPOSE 8000