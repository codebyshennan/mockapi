FROM golang:1.18-alpine
RUN apk add bash curl make git gcc libc-dev openssh-client-default ca-certificates

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

ARG APP_ENV=production
ARG MONGO_URI
ARG DB_NAME
ARG JWT_SIGNING_KEY
ARG API_MODE
ARG GOOGLE_CLIENT_ID

ENV APP_ENV ${APP_ENV}
ENV MONGO_URI ${MONGO_URI}
ENV DB_NAME ${DB_NAME}
ENV JWT_SIGNING_KEY ${JWT_SIGNING_KEY}
ENV API_MODE ${API_MODE}
ENV GOOGLE_CLIENT_ID ${GOOGLE_CLIENT_ID}

COPY . .
RUN go build -o /server.out ./cmd/server
CMD [ "/server.out" ]
