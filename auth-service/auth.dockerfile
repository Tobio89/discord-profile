FROM alpine:latest

RUN mkdir /app

COPY ./dist/auth-service /app

CMD [ "/app/auth-service"]
