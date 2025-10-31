# --- builder image
ARG NODEJS_BUILDER=registry.access.redhat.com/ubi9/nodejs-20@sha256:95a493933ed5f5f3d4b4dcf933002714d3f2384104504304137c3e6c13d29fd8
ARG RUNTIME=registry.access.redhat.com/ubi9/nginx-124@sha256:aa73fdb10af2bf24611ba714a412c2e65cec88a00eee628a0f2a75e564ec18f2

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
ENV BASE_PATH="/opt/app-root/src"
ARG VERSION=hub-ui-1.19.0

USER root
RUN fips-mode-setup --enable && \
    update-crypto-policies --set FIPS && \
    echo "Verifying FIPS kernel parameter:" && \
    cat /proc/sys/crypto/fips_enabled && \
    echo "Verifying OpenSSL FIPS status:" && \
    openssl version -a | grep -i fips && \
    (openssl md5 /dev/null || echo "MD5 test passed (expected failure in FIPS mode)")

RUN chmod ugo+rw /opt/app-root/src/config.js && \
    chown nginx:nginx /opt/app-root/src/config.js && \
    chmod +x /usr/bin/start.sh
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
    io.openshift.tags="pipelines,tekton,openshift" \
    cpe="cpe:/a:redhat:openshift_pipelines:1.19::el9"
