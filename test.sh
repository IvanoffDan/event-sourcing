echo Building Docker containers. . .
docker-compose build
echo Starting Docker. . .
docker-compose -f ./docker-compose.test.yml up
echo Done