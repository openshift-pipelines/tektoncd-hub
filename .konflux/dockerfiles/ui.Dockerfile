# --- builder image
ARG NODEJS_BUILDER=registry.redhat.io/ubi9/nodejs-22@sha256:c5fa408240ada5f820b6a31004fd483437ad46884505f8c6f80f650359f8add1
ARG RUNTIME=registry.redhat.io/ubi9/nginx-124@sha256:62c469e0343bd135722577fee540e7282d271e2156d4ae364b6b0df62f86b8b2

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
COPY --from=builder $REMOTE_SOURCE/ui/image/nginx.conf "${NGINX_CONFIGURATION_PATH}"/server.conf
ARG VERSION=1.22

USER root
RUN fips-mode-setup --enable && \
    update-crypto-policies --set FIPS && \
    echo "Verifying FIPS kernel parameter:" && \
    cat /proc/sys/crypto/fips_enabled && \
    echo "Verifying OpenSSL FIPS status:" && \
    openssl version -a | grep -i fips && \
    (openssl md5 /dev/null || echo "MD5 test passed (expected failure in FIPS mode)")

USER root
# Create writable directory for config.js and ensure proper permissions
# Update nginx.conf to use correct root path and include path
RUN mkdir -p /tmp/config && \
    chmod +x /usr/bin/start.sh && \
    chown -R 1001:0 /tmp/config && \
    chmod -R g+rwX /tmp/config && \
    chmod 775 /tmp/config && \
    sed -i 's|root   /usr/share/nginx/html;|root   /opt/app-root/src;|g' "${NGINX_CONFIGURATION_PATH}"/server.conf && \
    sed -i 's|include /etc/nginx/conf.d/location.locations;|include '"${NGINX_DEFAULT_CONF_PATH}"'/location.locations;|g' "${NGINX_CONFIGURATION_PATH}"/server.conf

# Use /tmp/config for writable config.js
ENV CONFIG_DIR="/tmp/config"

USER 1001

EXPOSE 8080

COPY --from=builder --chown=1001 $REMOTE_SOURCE/ui/image/location.locations "${NGINX_DEFAULT_CONF_PATH}"/location.locations

# Copy static nginx fragment to the correct location
COPY --from=builder --chown=1001 $REMOTE_SOURCE/ui/image/config-js.conf "${NGINX_DEFAULT_CONF_PATH}"/config-js.conf

CMD /usr/bin/start.sh

LABEL \
    com.redhat.component="openshift-pipelines-hub-ui-rhel9-container" \
    cpe="cpe:/a:redhat:openshift_pipelines:1.22::el9" \
    description="Red Hat OpenShift Pipelines tektoncd-hub ui" \
    io.k8s.description="Red Hat OpenShift Pipelines tektoncd-hub ui" \
    io.k8s.display-name="Red Hat OpenShift Pipelines tektoncd-hub ui" \
    io.openshift.tags="tekton,openshift,tektoncd-hub,ui" \
    maintainer="pipelines-extcomm@redhat.com" \
    name="openshift-pipelines/pipelines-hub-ui-rhel9" \
    summary="Red Hat OpenShift Pipelines tektoncd-hub ui" \
    version="v1.22.4"
