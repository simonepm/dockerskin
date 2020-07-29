### BUILD

docker build . -t hello-node:latest

### RUN

docker run -it -p 8080:8080 --rm --init hello-node:latest
