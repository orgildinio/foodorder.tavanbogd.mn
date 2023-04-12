#FROM golang:latest AS builder
FROM golang:latest

ENV LD_LIBRARY_PATH=/usr/local/lib:$LD_LIBRARY_PATH

COPY libsqlite3/usr/lib/x86_64-linux-gnu/libsqlite3.so.0.8.6 /usr/lib/x86_64-linux-gnu/
COPY libsqlite3/usr/lib/x86_64-linux-gnu/libsqlite3.so.0 /usr/lib/x86_64-linux-gnu/
RUN ln -s /usr/lib/x86_64-linux-gnu/libsqlite3.so.0 /usr/lib/x86_64-linux-gnu/libsqlite3.so

COPY x64/libcrypto.so.1.0.0 /usr/lib
RUN ln -s /usr/lib/libcrypto.so.1.0.0 /usr/lib/libcrypto.so

COPY x64/libssl.so.1.0.0 /usr/lib
RUN ln -s /usr/lib/libssl.so.1.0.0 /usr/lib/libssl.so

COPY x64/libicudata.so.53.1 /usr/lib
RUN ln -s /usr/lib/libicudata.so.53.1 /usr/lib/libicudata.so.53 \
    && ln -s /usr/lib/libicudata.so.53.1 /usr/lib/libicudata.so.5 \
    && ln -s /usr/lib/libicudata.so.53.1 /usr/lib/libicudata.so

COPY x64/libicui18n.so.53.1 /usr/lib
RUN ln -s /usr/lib/libicui18n.so.53.1 /usr/lib/libicui18n.so.53 \
    && ln -s /usr/lib/libicui18n.so.53.1 /usr/lib/libicui18n.so.5 \
    && ln -s /usr/lib/libicui18n.so.53.1 /usr/lib/libicui18n.so

COPY x64/libicuuc.so.53.1 /usr/lib
RUN ln -s /usr/lib/libicuuc.so.53.1 /usr/lib/libicuuc.so.53 \
    && ln -s /usr/lib/libicuuc.so.53.1 /usr/lib/libicuuc.so.5 \
    && ln -s /usr/lib/libicuuc.so.53.1 /usr/lib/libicuuc.so

COPY x64/libQt5Core.so.5.4.1 /usr/lib
RUN ln -s /usr/lib/libQt5Core.so.5.4.1 /usr/lib/libQt5Core.so.5.4 \
    && ln -s /usr/lib/libQt5Core.so.5.4.1 /usr/lib/libQt5Core.so.5 \
    && ln -s /usr/lib/libQt5Core.so.5.4.1 /usr/lib/libQt5Core.so

COPY x64/libQt5Network.so.5.4.1 /usr/lib
RUN ln -s /usr/lib/libQt5Network.so.5.4.1 /usr/lib/libQt5Network.so.5.4 \
    && ln -s /usr/lib/libQt5Network.so.5.4.1 /usr/lib/libQt5Network.so.5 \
    && ln -s /usr/lib/libQt5Network.so.5.4.1 /usr/lib/libQt5Network.so

COPY x64/libQt5Script.so.5.4.1 /usr/lib
RUN ln -s /usr/lib/libQt5Script.so.5.4.1 /usr/lib/libQt5Script.so.5.4 \
    && ln -s /usr/lib/libQt5Script.so.5.4.1 /usr/lib/libQt5Script.so.5 \
    && ln -s /usr/lib/libQt5Script.so


# Add user
RUN useradd -ms /bin/bash ebarimtuser
USER ebarimtuser

# Set working directory
WORKDIR /home/ebarimtuser/app

# Copy the rest of the project files into the container at /app
COPY --chown=ebarimtuser:ebarimtuser . .

# Download dependencies
RUN go mod download

# Build the project
RUN go build -o main

RUN chmod +x /home/ebarimtuser/app/start.sh
RUN chmod +x /home/ebarimtuser/app/main
# Expose port 3000 for the application

ENV DB_CONNECTION=postgres \
    DB_HOST=127.0.0.1 \
    DB_PORT=5432 \
    DB_DATABASE=postgres \
    DB_USERNAME=postgres \
    DB_PASSWORD=postgres \
    DB_DEBUG=false \
    APP_NAME=LAMBDA \
    APP_PORT=8080 \
    APP_MIGRATE=true \
    APP_SEED=true \
    SYSADMIN_LOGIN=superadmin \
    SYSADMIN_PASSWORD=superadmin \
    SYSADMIN_EMAIL=null \
    JWT_SECRET=BIzaSyBUEDlwdKQ0AZOnHOkZv2MIw0GlRjql6V4 \
    JWT_TTL=72 \
    MAIL_MAILER=smtp \
    MAIL_HOST=smtp.gmail.com \
    MAIL_PORT=587 \
    MAIL_USERNAME=null \
    MAIL_PASSWORD=null \
    MAIL_ENCRYPTION=tls \
    MAIL_FROM_ADDRESS=null \
    MAIL_FROM_NAME=LAMBDA \
    GRAPHQL_DEBUG=true \
    LAMBDA_ROOT="./node_modules/@lambda-platform/lambda-builder"
# Run the application when the container starts
CMD ["sh", "./start.sh"]
