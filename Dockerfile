FROM ubuntu:latest

WORKDIR /app

COPY webserver /app/webserver

RUN chmod +x /app/webserver

EXPOSE 8080

CMD ["/app/webserver"]
