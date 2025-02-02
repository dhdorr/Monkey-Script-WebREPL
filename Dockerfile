FROM golang

WORKDIR /app

COPY / ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-monkey-repl

EXPOSE 80

CMD [ "/docker-monkey-repl" ]