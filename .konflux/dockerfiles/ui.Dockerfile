# Rebuild trigger: 1.15.4 release 2026-01-19
# --- builder image
ARG NODEJS_BUILDER=registry.redhat.io/ubi8/nodejs-20@sha256:90ced1ab0ffdb355fb6becb2a7bbc15e163fc277fbcce326e0bd539fbe8fadb5
ARG RUNTIME=registry.access.redhat.com/ubi8/ubi-minimal:latest

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
ARG VERSION=1.15

USER root
RUN chmod ugo+rw /opt/app-root/src/config.js && \
    chown nginx:nginx /opt/app-root/src/config.js && \
    chmod +x /usr/bin/start.sh
USER nginx

EXPOSE 8080

COPY --from=builder --chown=1001 $REMOTE_SOURCE/ui/image/location.locations "${NGINX_DEFAULT_CONF_PATH}"/location.conf

CMD /usr/bin/start.sh

LABEL \
    com.redhat.component="openshift-pipelines-hub-ui-rhel8-container" \
    cpe="cpe:/a:redhat:openshift_pipelines:1.15::el9" \
    description="Red Hat OpenShift Pipelines tektoncd-hub ui" \
    io.k8s.description="Red Hat OpenShift Pipelines tektoncd-hub ui" \
    io.k8s.display-name="Red Hat OpenShift Pipelines tektoncd-hub ui" \
    io.openshift.tags="tekton,openshift,tektoncd-hub,ui" \
    maintainer="pipelines-extcomm@redhat.com" \
    name="openshift-pipelines/pipelines-hub-ui-rhel8" \
    summary="Red Hat OpenShift Pipelines tektoncd-hub ui" \
    version="v1.15.5"
# trigger rebuild 2026-02-14
