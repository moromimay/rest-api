FROM golang

# RUN mkdir /rest-api
# WORKDIR /rest-api

ADD . /go/src/rest-api
ENV db_name redcoin
ENV db_pass test
ENV db_user root
ENV db_host localhost
ENV db_port 3306
ENV token_password thisIsTheJwtPassword

RUN go get github.com/gorilla/mux
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/jinzhu/gorm
RUN go get golang.org/x/crypto/bcrypt
RUN go get github.com/jinzhu/gorm/dialects/mysql
RUN go get github.com/joho/godotenv

RUN go install rest-api

ENTRYPOINT /go/bin/rest-api

EXPOSE 8000