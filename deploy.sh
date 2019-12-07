docker build -t test -f docker/Dockerfile .
docker run --rm -it -d -p 3004:8080 test