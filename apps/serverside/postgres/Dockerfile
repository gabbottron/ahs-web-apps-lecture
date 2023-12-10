FROM postgres:latest
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=password
ENV POSTGRES_DB=webapp
COPY ./init-scripts /docker-entrypoint-initdb.d/

