FROM golang AS builder
COPY . /perlance
RUN cd cmd && \
    go build && \
    mv cmd / && \
    rm /perlance

FROM builder
CMD ./cmd
