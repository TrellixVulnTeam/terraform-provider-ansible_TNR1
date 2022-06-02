FROM ubuntu:20.04

VOLUME [ "/opt" ]
WORKDIR /opt

RUN apt update && apt upgrade -y &&\
    apt install -y python3-dev python3-pip zlib1g-dev &&\
    pip3 install pyinstaller

ENTRYPOINT "/bin/bash" "-c"
