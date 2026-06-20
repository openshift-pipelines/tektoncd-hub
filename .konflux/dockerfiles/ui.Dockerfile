# --- builder image
ARG NODEJS_BUILDER=registry.access.redhat.com/ubi9/nodejs-22@sha256:2e15adce9fe27151461c3e131a05a93deceee458c75c0b6e83ab3050bf669087
ARG RUNTIME=registry.access.redhat.com/ubi9/nginx-124@sha256:fc282f9fc249d9b9606d1d93769e1b6e1c042221fdae285b53032b59ba9690d0

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
ARG VERSION=1.20

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

# Copy static nginx fragment to the correct location
COPY --from=builder --chown=1001 $REMOTE_SOURCE/ui/image/config-js.conf "${NGINX_DEFAULT_CONF_PATH}"/config-js.conf

CMD /usr/bin/start.sh

LABEL \
    com.redhat.component="openshift-pipelines-hub-ui-rhel9-container" \
    cpe="cpe:/a:redhat:openshift_pipelines:1.20::el9" \
    description="Red Hat OpenShift Pipelines tektoncd-hub ui" \
    io.k8s.description="Red Hat OpenShift Pipelines tektoncd-hub ui" \
    io.k8s.display-name="Red Hat OpenShift Pipelines tektoncd-hub ui" \
    io.openshift.tags="tekton,openshift,tektoncd-hub,ui" \
    maintainer="pipelines-extcomm@redhat.com" \
    name="openshift-pipelines/pipelines-hub-ui-rhel9" \
    summary="Red Hat OpenShift Pipelines tektoncd-hub ui" \
    version="v1.20.5"
