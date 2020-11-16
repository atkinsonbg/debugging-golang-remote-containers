FROM golang:1.14.12-stretch
RUN apt-get update && \
    apt-get install -y exiftool
COPY main.go gobook.pdf ./
ENTRYPOINT ["go", "run", "main.go"]