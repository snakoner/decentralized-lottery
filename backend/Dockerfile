FROM golang:1.23

RUN go version
# ENV GOPATH=/
RUN mkdir /build
WORKDIR /build
COPY . .

# install psql
# RUN apt-get update
# RUN apt-get -y install postgresql-client
# RUN chmod a+x wait-for-postgres.sh
# RUN chmod a+x init_db.sh

RUN go mod download
RUN go build -o app ./cmd/app/*

EXPOSE 8000

CMD ["./app"]