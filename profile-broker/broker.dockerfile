FROM alpine:latest

RUN mkdir /app

COPY ./dist/profile-broker /app

CMD [ "/app/profile-broker"]
