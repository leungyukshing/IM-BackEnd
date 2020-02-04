docker stop $(docker ps -aq)
docker build -t test -f Dockerfile .
docker run --rm -it -d -p 8080:8080 --net=host test /bin/bash # -c 'go run /go/src/github.com/backend/main.go'