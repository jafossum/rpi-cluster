# MongoDB Dockerfile
# https://github.com/dockerfile/mongodb

FROM balenalib/raspberry-pi-debian

LABEL maintainer="John A. Fossum"

# Install MongoDB.
RUN apt-get update && apt-get install -y \
    mongodb \
    && rm -rf /var/lib/apt/lists/*

# Define mountable directories.
VOLUME ["/data/db"]

# Define working directory.
WORKDIR /data

# Expose ports.
#   - 27017: process
#   - 28017: http
EXPOSE 27017
EXPOSE 28017

# Define default command.
CMD ["mongod"]

