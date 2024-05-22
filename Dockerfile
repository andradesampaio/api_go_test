FROM ubuntu:latest

EXPOSE 8080

WORKDIR /app

ENV HOST=localhost PORT=5432

ENV USER=root PASSWORD=root DBNAME=mydatabase

COPY ./cmd/main main

CMD [ "./main" ]