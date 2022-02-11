 FROM golang:1.17-alpine as build

 WORKDIR /src
 COPY . .
 RUN go build -o app


 FROM alpine as runtime 
 VOLUME [ "/logs" ]
 COPY --from=build /src/app /usr/local/bin/app
 COPY start.sh /
 RUN chmod +x /start.sh
 ENTRYPOINT [ "./start.sh" ]
