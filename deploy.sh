docker stop $(docker ps -aq)
docker rmi -f test:latest
docker build -t test -f Dockerfile .
docker run --rm -it -d -p 8080:8080 --net=host test
docker cp /root/bee $(docker ps -q):/go/src/github.com/backend
docker exec -i $(docker ps -q) /bin/bash -c 'cd /go/src/github.com/backend && ./run.sh && exit'