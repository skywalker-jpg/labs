FROM ubuntu:latest

WORKDIR /app

RUN apt-get update && apt-get install -y ca-certificates



COPY ./bin/labs/labs /app/labs/labs
COPY ./config/config.yaml /app
COPY ./migrations /app/migrations

CMD ["/app/labs/labs"]