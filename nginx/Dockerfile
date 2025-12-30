#
# Builder
#
ARG NGINX_VERSION
FROM --platform=$BUILDPLATFORM nginx:${NGINX_VERSION} AS builder
RUN apt-get update && \
    apt-get install -y \
        cmake \
        build-essential \
        libssl-dev \
        zlib1g-dev \
        libpcre2-dev \
        pkg-config \
        libc-ares-dev \
        libre2-dev \
        git

WORKDIR /usr/src
RUN git clone https://github.com/nginx/nginx.git --branch release-${NGINX_VERSION}

WORKDIR /usr/src/nginx
RUN auto/configure --with-compat

WORKDIR /usr/src
RUN git clone https://github.com/waduhek/nginx-otel-module.git nginx-otel

WORKDIR /usr/src/nginx-otel
RUN mkdir build && \
    cd build && \
    cmake -DNGX_OTEL_NGINX_BUILD_DIR=/usr/src/nginx/objs .. && \
    make -j$(nproc)


#
# Runner
#
FROM --platform=$BUILDPLATFORM nginx:${NGINX_VERSION} AS runner
RUN apt-get update && apt-get install -y libc-ares2
COPY --from=builder /usr/src/nginx-otel/build/ngx_otel_module.so /etc/nginx/modules/
COPY nginx.conf /etc/nginx/nginx.conf
