FROM alpine:latest

RUN mkdir /app

COPY ./dist/http-broker /app

CMD [ "/app/http-broker"]
