# Rebuild trigger: 1.15.4 release 2026-01-19
ARG GO_BUILDER=registry.redhat.io/ubi8/go-toolset:latest@sha256:e0bd798876d458b2fc653cff7dbac26179fda5d7ab6ef4a1a7914c5f29730e62
ARG RUNTIME=registry.redhat.io/ubi8/ubi@sha256:bd097b288df365aec7bad694c87036534ce94d1a4974c943f417f757f8eb2ac4

FROM $GO_BUILDER AS builder

WORKDIR /go/src/github.com/tektoncd/hub

COPY upstream .
COPY .konflux/patches patches/
RUN set -e; for f in patches/*.patch; do echo ${f}; [[ -f ${f} ]] || continue; git apply ${f}; done
COPY head HEAD

ENV GODEBUG="http2server=0" GOTOOLCHAIN=auto
RUN go build -ldflags="-X 'knative.dev/pkg/changeset.rev=$(cat HEAD)'" -mod=vendor -tags disable_gcp -v -o /tmp/hub-db-migration \
    ./api/cmd/db

FROM $RUNTIME
ARG VERSION=1.15

COPY --from=builder /tmp/hub-db-migration /ko-app/hub-db-migration
COPY head ${KO_DATA_PATH}/HEAD

EXPOSE 8000

LABEL \
    com.redhat.component="openshift-pipelines-hub-db-migration-rhel8-container" \
    cpe="cpe:/a:redhat:openshift_pipelines:1.15::el8" \
    description="Red Hat OpenShift Pipelines tektoncd-hub db-migration" \
    io.k8s.description="Red Hat OpenShift Pipelines tektoncd-hub db-migration" \
    io.k8s.display-name="Red Hat OpenShift Pipelines tektoncd-hub db-migration" \
    io.openshift.tags="tekton,openshift,tektoncd-hub,db-migration" \
    maintainer="pipelines-extcomm@redhat.com" \
    name="openshift-pipelines/pipelines-hub-db-migration-rhel8" \
    summary="Red Hat OpenShift Pipelines tektoncd-hub db-migration" \
    version="v1.15.6"


RUN groupadd -r -g 65532 nonroot && useradd --no-log-init -r -u 65532 -g nonroot nonroot

USER 65532

ENTRYPOINT ["/ko-app/hub-db-migration"]# trigger rebuild 2026-02-14
