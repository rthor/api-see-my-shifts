FROM postgres:10-alpine

RUN mkdir -p /docker-entrypoint-initdb.d

COPY ./sql/*.sql /docker-entrypoint-initdb.d/

