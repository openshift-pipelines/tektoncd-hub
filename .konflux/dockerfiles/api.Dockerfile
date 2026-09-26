# Rebuild trigger: 1.15.4 release 2026-01-19
ARG GO_BUILDER=registry.access.redhat.com/ubi8/go-toolset:latest@sha256:93b85580bf0737559318b12baa82cf034db7fce7a6c095c593ff156c09b50339
ARG RUNTIME=registry.redhat.io/ubi8/ubi:latest@sha256:f62cf5375e9e17dbb719ab70d6ab3abfe83445a2ed31d46580b70ee2f0440617

FROM $GO_BUILDER AS builder

WORKDIR /go/src/github.com/tektoncd/hub

COPY upstream .
COPY .konflux/patches patches/
RUN set -e; for f in patches/*.patch; do echo ${f}; [[ -f ${f} ]] || continue; git apply ${f}; done
COPY head HEAD

ENV GODEBUG="http2server=0" GOTOOLCHAIN=auto
RUN go build -ldflags="-X 'knative.dev/pkg/changeset.rev=$(cat HEAD)'" -mod=vendor -tags disable_gcp -v -o /tmp/hub-api-server \
    ./api/cmd/api

FROM $RUNTIME
ARG VERSION=1.15

RUN dnf install -y openssh-clients git shadow-utils

COPY --from=builder /tmp/hub-api-server /ko-app/hub-api-server
COPY head ${KO_DATA_PATH}/HEAD

EXPOSE 8000

LABEL \
    com.redhat.component="openshift-pipelines-hub-api-rhel8-container" \
    cpe="cpe:/a:redhat:openshift_pipelines:1.15::el8" \
    description="Red Hat OpenShift Pipelines tektoncd-hub api" \
    io.k8s.description="Red Hat OpenShift Pipelines tektoncd-hub api" \
    io.k8s.display-name="Red Hat OpenShift Pipelines tektoncd-hub api" \
    io.openshift.tags="tekton,openshift,tektoncd-hub,api" \
    maintainer="pipelines-extcomm@redhat.com" \
    name="openshift-pipelines/pipelines-hub-api-rhel8" \
    summary="Red Hat OpenShift Pipelines tektoncd-hub api" \
    version="v1.15.5"

RUN groupadd -r -g 65532 nonroot && useradd --no-log-init -r -u 65532 -g nonroot nonroot
USER 65532

ENTRYPOINT ["/ko-app/hub-api-server"]# trigger rebuild 2026-02-14
