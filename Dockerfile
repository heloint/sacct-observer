FROM golang:latest
LABEL base.image="golang:latest"
LABEL version="0.1"
LABEL software="sacct-observer"
LABEL software.version="0.1"
LABEL license="MIT"
LABEL maintainer="Daniel Majer"

# Temporary environment variables so that apt installation does not complain.
# https://github.com/phusion/baseimage-docker/issues/319
ENV DEBIAN_FRONTEND noninteractive
ENV TZ="Europe/Madrid"
ENV DEBCONF_NOWARNINGS="yes"

RUN apt-get update && apt-get install -y apt-utils && apt-get upgrade -y

RUN apt-get install -y build-essential \
                       cmake

WORKDIR /app
COPY . .
RUN make build
RUN make install
RUN make get-autocomplete
