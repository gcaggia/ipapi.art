# ipapi.art
Real-time IP address API written in Golang running with Docker.

Link: [ipapi.art](http://ipapi.art "ipapi url")

## How to use it

### With Golang

You will need to install go first.
https://golang.org/doc/install

```sh
$ git clone https://github.com/gcaggia/ipapi.art.git
$ cd ipapi.art
$ go run main.go
```

It will start the application on: [localhost:18000](http://localhost:18000)

Then, to build and deploy the app: 
```sh
$ go build main.go
$ ./main
```

### With Docker

A dockerfile has been written to make this app running anywhere.

Basicaly, with docker, it creates a small container (around 6 MB) and the application is ready for production.

To create the container: 
```sh
$ docker build -t ipapi:1.0 .
$ docker run -d -p 18000:18000 ipapi:1.0
```
