docker stop $(docker ps -aq)
docker build -t test -f Dockerfile .
docker run --rm -it -d -p 3004:8080 --network host test