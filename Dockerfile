FROM golang AS builder
COPY . /perlance
RUN cd /perlance/cmd && \
    go build && \
    mv cmd / && \
    rm -r /perlance

FROM builder
EXPOSE 8080
CMD /cmd
