docker compose -f go/docker-compose.yaml down --volumes
sleep 2
docker compose -f go/docker-compose.yaml up -d
echo "Go up"
docker compose -f kafka/docker-compose.yaml down --volumes
sleep 2
docker compose -f kafka/docker-compose.yaml up -d
echo "Kafka up"
docker compose -f nextjs/docker-compose.yaml down --volumes
sleep 2
docker compose -f nextjs/docker-compose.yaml up -d
echo "Nextjs up"
docker compose -f nestjs/docker-compose.yaml down --volumes
sleep 2
docker compose -f nestjs/docker-compose.yaml up -d
echo "Nestjs up"
