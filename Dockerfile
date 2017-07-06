FROM ubuntu:latest
MAINTAINER Tryout Team

RUN apt-get update \
    && apt-get -y --force-yes install git

RUN git clone https://github.com/globocom/tryout-agent.git
