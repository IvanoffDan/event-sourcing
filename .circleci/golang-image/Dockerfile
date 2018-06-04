FROM docker:17.05.0-ce-git

MAINTAINER ivanoff.daniil@gmail.com

RUN apk add --update --no-cache py-pip=9.0.0-r1 \
                                nodejs

RUN pip install docker-compose==1.21.1

RUN npm install -g codeclimate-test-reporter
