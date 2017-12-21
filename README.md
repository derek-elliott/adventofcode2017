# Advent of Code 2017
## Install
There is a Dockerfile included that will build the binary and put it in a Scratch image.  It can be built with the following command:
```
docker build -t advent:2017 .
```

## Usage
The Docker image will need to have the port 3000 exposed, otherwise it doesn't need any other arguments.  You can run it with the following command:
```
docker run -p3000:3000 advent:2017
```

The root path will list all the available endpoints and the HTTP methods they accept.  There are scripts in the `scripts` directory to transform the data into a json request.

## Contribute

Submit PRs if you'd like, but I wouldn't bother if I were you.

## License

MIT © Derek Elliott
