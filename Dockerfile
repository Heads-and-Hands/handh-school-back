FROM golang

RUN go get -u github.com/go-sql-driver/mysql
RUN go get -u github.com/gorilla/mux
RUN go get -u github.com/jinzhu/gorm
RUN go get -u github.com/qor/qor
RUN go get -u github.com/qor/admin
RUN go get -u github.com/mitchellh/panicwrap

ADD . /go/src/handh-school-back
RUN go install handh-school-back

ENTRYPOINT /go/bin/handh-school-back

