# --- builder image
FROM registry.access.redhat.com/ubi8/nodejs-16:latest as BUILD

COPY --chown=1001 $REMOTE_SOURCE $REMOTE_SOURCE_DIR
WORKDIR $REMOTE_SOURCE_DIR/app
COPY --chown=1001 patches patches/
RUN set -e; for f in patches/*.patch; do echo foo ${f}; [[ -f ${f} ]] || continue; git apply ${f}; done
WORKDIR $REMOTE_SOURCE_DIR/app/ui

RUN source $CACHITO_ENV_FILE && \
    npm clean-install --legacy-peer-deps && \
    npm run build

# --- runtime image
FROM registry.access.redhat.com/ubi8/nginx-122:latest

COPY --from=BUILD $REMOTE_SOURCE_DIR/app/ui/build /opt/app-root/src
COPY --chown=1001 $REMOTE_SOURCE/app/ui/image/start.sh /usr/bin/
ENV BASE_PATH="/opt/app-root/src"

USER root
RUN chmod ugo+rw /opt/app-root/src/config.js && \
    chown nginx:nginx /opt/app-root/src/config.js && \
    chmod +x /usr/bin/start.sh
USER nginx

EXPOSE 8080

COPY --chown=1001 $REMOTE_SOURCE/app/ui/image/location.locations "${NGINX_DEFAULT_CONF_PATH}"/location.conf

CMD /usr/bin/start.sh

LABEL \
    com.redhat.component="openshift-pipelines-hub-ui-container" \
    name="openshift-pipelines/pipelines-hub-ui-rhel8" \
    version="${CI_CONTAINER_VERSION}" \
    summary="Red Hat OpenShift Pipelines Hub UI" \
    maintainer="pipelines-extcomm@redhat.com" \
    description="Red Hat OpenShift Pipelines Hub UI" \
    io.k8s.display-name="Red Hat OpenShift Pipelines Hub UI"
    io.k8s.description="Red Hat OpenShift Pipelines Hub UI" \
    io.openshift.tags="pipelines,tekton,openshift"
