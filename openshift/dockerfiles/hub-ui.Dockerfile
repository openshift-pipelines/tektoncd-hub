# --- builder image
ARG GO_BUILDER=registry.access.redhat.com/ubi8/nodejs-16:latest
ARG RUNTIME=registry.access.redhat.com/ubi8/nginx-122:latest

FROM $GO_BUILDER AS BUILD

ARG REMOTE_SOURCE=/go/src/github.com/tektoncd/hub
WORKDIR /go/src/github.com/tektoncd/hub
COPY upstream .
COPY --chown=1001 $REMOTE_SOURCE
COPY --chown=1001 patches patches/
RUN set -e; for f in patches/*.patch; do echo foo ${f}; [[ -f ${f} ]] || continue; git apply ${f}; done

WORKDIR $REMOTE_SOURCE/ui

RUN source $CACHITO_ENV_FILE && \
    npm clean-install --legacy-peer-deps && \
    npm run build

# --- runtime image
FROM $RUNTIME
ARG VERSION=hub-main

COPY --from=BUILD $REMOTE_SOURCE/ui/build /opt/app-root/src
COPY --chown=1001 $REMOTE_SOURCE/ui/image/start.sh /usr/bin/
ENV BASE_PATH="/opt/app-root/src"

USER root
RUN chmod ugo+rw /opt/app-root/src/config.js && \
    chown nginx:nginx /opt/app-root/src/config.js && \
    chmod +x /usr/bin/start.sh
USER nginx

EXPOSE 8080

COPY --chown=1001 $REMOTE_SOURCE/ui/image/location.locations "${NGINX_DEFAULT_CONF_PATH}"/location.conf

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
    io.openshift.tags="pipelines,tekton,openshift"
