FROM techknowlogick/xgo:go-1.15.6@sha256:e58c2e9ee05ab5257a3e2948717a38093e3ffe8fabed542227aeda25a74ca15e

RUN mkdir -p /build

# Inject the customized build script
ADD build.sh /build.sh
ENV BUILD /build.sh
RUN chmod +x $BUILD

ENTRYPOINT ["/build.sh"]
