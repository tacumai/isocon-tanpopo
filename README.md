```
$ docker build -t mygo .
$ docker run --name app -p 8080:8080 -v $(pwd):/go/src/app mygo
```