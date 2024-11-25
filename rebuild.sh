docker-compose down
docker build -t webapp webapp
docker build -t server server
docker-compose up -d