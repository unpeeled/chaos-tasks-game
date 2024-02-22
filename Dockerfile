FROM debian:bullseye
RUN apt-get update
RUN apt-get install -y ca-certificates golang postgresql-client
RUN mkdir -p /opt/chaostasks/bin
COPY src /opt/chaostasks/src
RUN ls /opt/chaostasks
WORKDIR /opt/chaostasks/src/chaostasks
RUN go build
RUN mv ./chaostasks /opt/chaostasks/bin/chaostasks
# Copy scripts
COPY scripts/2-import-tasks.sh /opt/chaostasks/bin/2-import-tasks.sh
EXPOSE 3000
CMD /opt/chaostasks/bin/chaostasks