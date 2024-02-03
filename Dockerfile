FROM debian:bullseye
RUN apt-get update
RUN apt-get install -y ca-certificates golang
RUN mkdir -p /opt/chaostasks/bin
COPY src /opt/chaostasks/src
RUN ls /opt/chaostasks
WORKDIR /opt/chaostasks/src/chaostasks
RUN go build
RUN mv ./chaostasks /opt/chaostasks/bin/chaostasks
EXPOSE 3000
CMD /opt/chaostasks/bin/chaostasks
