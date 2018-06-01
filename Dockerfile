FROM golang:1.10

COPY . ${GOPATH}/src/github.com/IvanoffDan/event-sourcing

WORKDIR ${GOPATH}/src/github.com/IvanoffDan/event-sourcing

RUN go get github.com/kardianos/govendor

RUN govendor fetch +missing

RUN go build -o ${GOPATH}/bin/app ${GOPATH}/src/github.com/IvanoffDan/event-sourcing/main.go

CMD if [ ${ENV} = production ]; \
	then \
	${GOPATH}/bin/app; \
	else \
	go get github.com/IvanoffDan/fresh && \
	fresh -c ./runner.conf; \
	fi