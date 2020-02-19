FROM postgres:11.5-alpine

COPY 001_data.sql /docker-entrypoint-initdb.d/001_data.sql

ENV PGDATA=/data

EXPOSE 5432