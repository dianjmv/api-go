FROM ubuntu:latest
LABEL authors="dannyrivadeneira"

ENTRYPOINT ["top", "-b"]