#/bin/bash
docker rmi mew_server:0.0.1
docker build -t mew_server:0.0.1 .
docker-compose up -d
docker-compose ps
docker logs moneybackward-mew_server-1