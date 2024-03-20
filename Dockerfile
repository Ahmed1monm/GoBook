FROM ubuntu:latest
LABEL authors="monem"

ENTRYPOINT ["top", "-b"]