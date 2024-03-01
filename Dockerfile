FROM debian:bullseye
RUN printf "deb http://deb.debian.org/debian bullseye-backports main non-free\ndeb-src http://deb.debian.org/debian bullseye-backports main non-free" > /etc/apt/sources.list.d/backports.list
RUN apt-get update
RUN apt-get install -y -t bullseye-backports golang
RUN apt-get install -y ca-certificates postgresql-client
RUN mkdir -p /opt/chaostasks/bin
COPY src /opt/chaostasks/src
RUN ls /opt/chaostasks
WORKDIR /opt/chaostasks/src/chaostasks
RUN go build
RUN mv ./chaostasks /opt/chaostasks/bin/chaostasks
# Copy scripts
COPY scripts/2-import-tasks.sh /opt/chaostasks/bin/2-import-tasks.sh
ARG PORT
EXPOSE $PORT
CMD /opt/chaostasks/bin/chaostasks
