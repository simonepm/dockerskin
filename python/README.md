### BUILD

docker build . -t hello-py:latest

### RUN

docker run -it -p 8080:8080 --rm --init hello-py:latest
