FROM golang:alpine
WORKDIR /usr/local/api
RUN apk update \
  && apk add git \
  && apk add bash \
  && go get github.com/cosmtrek/air

RUN air init

COPY . .

CMD [ "air", "-d" ]