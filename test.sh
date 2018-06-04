#!/bin/sh
docker-compose -f ./docker-compose.test.yml up

docker cp app_server_1:/go/src/github.com/IvanoffDan/event-sourcing/coverage.out ./coverage.out

/usr/lib/node_modules/codeclimate-test-reporter/bin/codeclimate.js < ./coverage.out