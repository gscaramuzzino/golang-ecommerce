# Dockerfile
FROM postgres:13-alpine

# set environment variables
ENV POSTGRES_USER=user
ENV POSTGRES_PASSWORD=secret
ENV POSTGRES_DB=ecommerce

# add initialization scripts
COPY ./init.sql /docker-entrypoint-initdb.d/
