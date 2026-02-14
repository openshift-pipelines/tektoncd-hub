# Rebuild trigger: 1.15.4 release 2026-01-19
# --- builder image
ARG NODEJS_BUILDER=registry.redhat.io/ubi8/nodejs-20@sha256:2b073286348c505b277b587c63aa64077ca57abb5544aed520ae00b4186bf23b
ARG RUNTIME=registry.redhat.io/ubi8/nginx-124@sha256:9f6c14ee78e4d9229f0b91464890ebc6b381a0fbd942d92fa2fc339330b07c33

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
ARG VERSION=hub-ui-1.15.4

USER root
RUN chmod ugo+rw /opt/app-root/src/config.js && \
    chown nginx:nginx /opt/app-root/src/config.js && \
    chmod +x /usr/bin/start.sh
USER nginx

EXPOSE 8080

COPY --from=builder --chown=1001 $REMOTE_SOURCE/ui/image/location.locations "${NGINX_DEFAULT_CONF_PATH}"/location.conf

CMD /usr/bin/start.sh

LABEL \
    com.redhat.component="openshift-pipelines-hub-ui-container" \
    name="openshift-pipelines/pipelines-hub-ui-rhel8" \
    version=$VERSION \
    summary="Red Hat OpenShift Pipelines Hub UI" \
    maintainer="pipelines-extcomm@redhat.com" \
    description="Red Hat OpenShift Pipelines Hub UI" \
    io.k8s.display-name="Red Hat OpenShift Pipelines Hub UI" \
    io.k8s.description="Red Hat OpenShift Pipelines Hub UI" \
    io.openshift.tags="pipelines,tekton,openshift" \
    cpe="cpe:/a:redhat:openshift_pipelines:1.15::el8"
# trigger rebuild 2026-02-14
