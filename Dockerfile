FROM alpine:latest

WORKDIR /app
COPY ./controller /app/controller
#RUN chmod 550 /app/controller

CMD ["/controller"]