FROM golang:1.22.3-alpine3.19

# go dependencies
RUN apk add make
RUN apk add git
RUN apk add gcc
RUN apk add libc-dev
RUN apk add pkgconfig

# dev env dependencies
RUN apk add build-base \
  freetype-dev \
  libpng-dev \
  openblas-dev \
  libffi-dev\
  redis \
  bash 

ENV DOCKER_CLI_HINTS=off
WORKDIR /app

CMD [ "tail", "-f", "/dev/null" ]
