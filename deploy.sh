docker stop $(docker ps -aq)
docker build -t test -f Dockerfile .
docker run --rm -it -d --network="host" test