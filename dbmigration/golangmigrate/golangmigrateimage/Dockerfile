ARG BUILD_IMAGE_NAME=golang:1.24-alpine3.21
ARG PROD_IMAGE_NAME=alpine:3.21

FROM ${BUILD_IMAGE_NAME} AS build_image
ENV CGO_ENABLED=0
ENV PATH=/usr/local/go/bin:$PATH
ARG VERSION=v4
ARG SOURCE="file go_bindata github github_ee bitbucket aws_s3 google_cloud_storage godoc_vfs gitlab"
ARG DATABASE="postgres mysql redshift cassandra spanner cockroachdb yugabytedb clickhouse mongodb sqlserver firebird neo4j pgx pgx5 rqlite"
WORKDIR /go/src/migrate
RUN apk add --no-cache git gcc musl-dev make && \
    git clone https://github.com/golang-migrate/migrate.git ./
RUN go mod download && \
    go build -a -o ./build/migrate.linux-386 -ldflags="-s -w -X main.Version=${VERSION}" -tags "${DATABASE} ${SOURCE}" ./cmd/migrate

FROM ${PROD_IMAGE_NAME} AS release_image
ENV GOOS=linux
ARG USERNAME=dockeruser
ARG USERDIR=/opt/golang-migrate
RUN apk add --no-cache ca-certificates bash curl envsubst postgresql-client  && \
    addgroup --g 1500 groupcontainer && \
    adduser -u 1500 -G groupcontainer -D ${USERNAME} --home ${USERDIR}
WORKDIR ${USERDIR}
COPY --from=build_image /go/src/migrate/build/migrate.linux-386 /usr/local/bin/migrate
COPY . ./
RUN ln -s /usr/local/bin/migrate ./migrate && \
    mv env.default .env && \
    chmod +x ./migration.sh && \
    chown -R 1500 postgres clickhouse
USER ${USERNAME}

ENTRYPOINT ["./migration.sh"]
