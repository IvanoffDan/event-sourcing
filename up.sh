#!/bin/bash

echo Building Docker containers. . .
docker-compose build
echo Starting Docker. . .
docker-compose up
echo Done