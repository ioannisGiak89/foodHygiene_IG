# FSA Ratings

## Environments

Demo: N service at the moment

## Prerequisites

General:

* [GO >= 1.8](https://golang.org/doc/devel/release.html#go1.8)
* [Bootstrap 4](https://v4-alpha.getbootstrap.com/)
* [jQuery v3.2.1](https://jquery.com/)

To run with docker:

* Docker (1.12+)d

To run without Docker:

* GO >= 1.8

## Build and Run

Ensure you have all of the prerequisites listed above before starting this.

### Docker

Navigate to project's directory, you'll first need to build the image:

```
$ docker build -t ioannisgiak/foodhygiene .
```

After the image is build you'll just need to run it. The application listens on port `8080`:

```
$ docker run --rm -p 8080:8080 ioannisgiak/foodhygiene
```

You can now access the application on http://localhost:8080/home

(Note: if you're using Docker Machine, you'll need to also specify the IP for the machine to connect.
 You can find out the IP address with `$ docker-machine ip <machine-name>`)

### Go build

Make sure that you copy the project to:

```
$ $GOPATH/src/github.com/ioannisGiak89/foodHygiene_IG/
```

Navigate to project's directory and:

```
$ go build main.go
$ ./main
```

You can now access the application on http://localhost:8080/home

## Test

Navigate to project's directory and run:

```
$ go test ./...
```

## TODO

* Log errors to a file in order to retrieve debug information
* Improve cache. Add the time that the data cached, in order to avoid calling the api 2 times within few seconds
* Create an error page 
* Create helpers to make testing easier 
* Do some more testing

