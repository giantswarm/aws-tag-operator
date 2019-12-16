FROM alpine:3.10

RUN apk add --no-cache ca-certificates

ADD ./aws-tag-operator /aws-tag-operator

ENTRYPOINT ["/aws-tag-operator"]
