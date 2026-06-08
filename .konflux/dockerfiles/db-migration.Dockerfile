ARG GO_BUILDER=registry.access.redhat.com/ubi9/go-toolset:latest
ARG RUNTIME=registry.redhat.io/ubi9/ubi-minimal@sha256:6fc28bcb6776e387d7a35a2056d9d2b985dc4e26031e98a2bd35a7137cd6fd71  

FROM $GO_BUILDER AS builder 

WORKDIR /go/src/github.com/tektoncd/hub

COPY upstream .
COPY .konflux/patches patches/
RUN set -e; for f in patches/*.patch; do echo ${f}; [[ -f ${f} ]] || continue; git apply ${f}; done
COPY head HEAD

ENV GODEBUG="http2server=0"
ENV GOEXPERIMENT=strictfipsruntime
RUN go build -ldflags="-X 'knative.dev/pkg/changeset.rev=$(cat HEAD)'" -mod=vendor -tags disable_gcp,strictfipsruntime -v -o /tmp/hub-db-migration \
    ./api/cmd/db

FROM $RUNTIME
ARG VERSION=1.23

COPY --from=builder /tmp/hub-db-migration /ko-app/hub-db-migration
COPY head ${KO_DATA_PATH}/HEAD

EXPOSE 8000

LABEL \
    com.redhat.component="openshift-pipelines-hub-db-migration-rhel9-container" \
    cpe="cpe:/a:redhat:openshift_pipelines:1.23::el9" \
    description="Red Hat OpenShift Pipelines tektoncd-hub db-migration" \
    io.k8s.description="Red Hat OpenShift Pipelines tektoncd-hub db-migration" \
    io.k8s.display-name="Red Hat OpenShift Pipelines tektoncd-hub db-migration" \
    io.openshift.tags="tekton,openshift,tektoncd-hub,db-migration" \
    maintainer="pipelines-extcomm@redhat.com" \
    name="openshift-pipelines/pipelines-hub-db-migration-rhel9" \
    summary="Red Hat OpenShift Pipelines tektoncd-hub db-migration" \
    version="v1.23.2"


RUN groupadd -r -g 65532 nonroot && useradd --no-log-init -r -u 65532 -g nonroot nonroot
USER 65532

ENTRYPOINT ["/ko-app/hub-db-migration"]
