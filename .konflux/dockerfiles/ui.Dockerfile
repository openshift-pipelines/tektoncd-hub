# --- builder image
ARG NODEJS_BUILDER=registry.redhat.io/ubi9/nodejs-22@sha256:f34e139a493588f03104e21fcd6b182d8884b7ae2cee7d91e56a8fff8a106b2f
ARG RUNTIME=registry.access.redhat.com/ubi9/nginx-124@sha256:07547973a7e10860b474a331727ab7ecd9f30a2ab15f4a7ff024c3fec7dc80ba 

FROM $NODEJS_BUILDER AS builder

ARG REMOTE_SOURCE=/go/src/github.com/tektoncd/hub

WORKDIR $REMOTE_SOURCE

COPY upstream .
USER root
RUN chmod -R g+w $REMOTE_SOURCE
COPY .konflux/patches patches/
RUN set -e; for f in patches/*.patch; do echo foo ${f}; [[ -f ${f} ]] || continue; git apply ${f}; done

WORKDIR $REMOTE_SOURCE/ui

RUN npm clean-install --legacy-peer-deps --ignore-scripts --max-old-space-size=4096 --force --no-optional && \
    npm run build

# --- runtime image
FROM $RUNTIME
ARG REMOTE_SOURCE=/go/src/github.com/tektoncd/hub

COPY --from=builder $REMOTE_SOURCE/ui/build /opt/app-root/src
COPY --from=builder --chown=1001 $REMOTE_SOURCE/ui/image/start.sh /usr/bin/
ARG VERSION=hub-ui-next

USER root
RUN fips-mode-setup --enable && \
    update-crypto-policies --set FIPS && \
    echo "Verifying FIPS kernel parameter:" && \
    cat /proc/sys/crypto/fips_enabled && \
    echo "Verifying OpenSSL FIPS status:" && \
    openssl version -a | grep -i fips && \
    (openssl md5 /dev/null || echo "MD5 test passed (expected failure in FIPS mode)")


# Use /tmp/config for writable config.js
ENV BASE_PATH="/tmp/config"
ENV CONFIG_DIR="/tmp/config"


USER root
# Create writable directory for config.js and ensure proper permissions
RUN mkdir -p /tmp/config && \
    chmod +x /usr/bin/start.sh && \
    chgrp -R 0 /tmp/config && \
    chmod -R g=u /tmp/config && \
    echo 'location = /config.js { alias /tmp/config/config.js; }' > "${NGINX_DEFAULT_CONF_PATH}"/config-location.conf
USER nginx

EXPOSE 8080

COPY --from=builder --chown=1001 $REMOTE_SOURCE/ui/image/location.locations "${NGINX_DEFAULT_CONF_PATH}"/location.conf

CMD /usr/bin/start.sh

LABEL \
    com.redhat.component="openshift-pipelines-hub-ui-container" \
    name="openshift-pipelines/pipelines-hub-ui-rhel9" \
    version=$VERSION \
    summary="Red Hat OpenShift Pipelines Hub UI" \
    maintainer="pipelines-extcomm@redhat.com" \
    description="Red Hat OpenShift Pipelines Hub UI" \
    io.k8s.display-name="Red Hat OpenShift Pipelines Hub UI" \
    io.k8s.description="Red Hat OpenShift Pipelines Hub UI" \
    io.openshift.tags="pipelines,tekton,openshift"
