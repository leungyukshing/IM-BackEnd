#docker-compose -f ../docker/docker-compose.yml up
docker build -t test -f docker/Dockerfile .
docker run --rm -it -d -p 3002:8080 test