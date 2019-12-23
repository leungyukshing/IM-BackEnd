docker stop $(docker ps -aq)
docker rm $(docker ps -aq)
docker rmi $(docker images -q)
docker build -t test -f docker/Dockerfile .
docker run --rm -it -d -p 3004:8080 test