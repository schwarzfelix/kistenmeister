docker-compose down
docker build -t webapp webapp
docker build -t server server
docker network create -d bridge service-network
docker-compose up -d