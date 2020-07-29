### BUILD

docker build . -t hello-go:latest

### RUN

docker run -it -p 8080:8080 --rm --init hello-go:latest
