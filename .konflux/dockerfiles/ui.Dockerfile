# --- builder image
ARG NODEJS_BUILDER=registry.redhat.io/ubi9/nodejs-22@sha256:ec60ac766c3959176acb5a0d4237be73457e39f6b36faba1cbf16eb9625c04d0
ARG RUNTIME=registry.access.redhat.com/ubi9/nginx-124@sha256:f3e0afbc6e1459a63b87ded4666b6ea3fa59a5d146d6bfa39995af936b1fb9f3 

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
ARG VERSION=hub-ui-next

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
