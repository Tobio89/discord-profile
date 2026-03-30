FROM alpine:latest

RUN mkdir /app

COPY ./dist/discord-bot /app
COPY dev.env /app

CMD [ "/app/discord-bot"]
