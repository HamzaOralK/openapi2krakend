FROM alpine:latest
RUN mkdir /openapi2krakend && \
    mkdir /openapi2krakend/swagger && \
    mkdir /openapi2krakend/output
COPY ../docker/init-script.sh /openapi2krakend/init-script.sh
COPY ../build/openapi2krakend /openapi2krakend/openapi2krakend
RUN chmod +x /openapi2krakend/init-script.sh && \
    chmod +x /openapi2krakend/openapi2krakend
