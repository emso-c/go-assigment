# Use an official Postgres image as a parent image
FROM postgres:latest

# Set environment variables for the database
ENV POSTGRES_USER=${POSTGRES_USER} \
    POSTGRES_PASSWORD=${POSTGRES_PASSWORD} \
    POSTGRES_DB=${POSTGRES_DB}

# Copy the SQL script to initialize the database
COPY *.sql /docker-entrypoint-initdb.d/

# Expose the default Postgres port
EXPOSE 5432
