FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY /bin/linux_amd64 .
EXPOSE 8080
CMD ./api -db-dsn=${DBDSN} -port=8080