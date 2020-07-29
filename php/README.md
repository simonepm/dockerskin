### BUILD

docker build . -t hello-php:latest

### RUN

docker run -it -p 8080:8080 --rm --init hello-php:latest
