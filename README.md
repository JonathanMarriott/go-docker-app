# Simple Maths Operations Web Server

A containerized web server written in Go, which serves pages that compute the answer to maths operations.

Build the development image with 
```
docker build -t mathapp-dev .
```

and run the image in a container with the host directory mounted by running

```
docker run -it --rm -p 8080:8080 -v $PWD/src:/go/src/mathapp mathapp-dev
```
