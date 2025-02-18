ARG GO_BUILDER=brew.registry.redhat.io/rh-osbs/openshift-golang-builder:v1.22
ARG RUNTIME=registry.access.redhat.com/ubi9/ubi-minimal:latest@sha256:66b99214cb9733e77c4a12cc3e3cbbe76769a213f4e2767f170a4f0fdf9db490

FROM $GO_BUILDER AS builder

WORKDIR /go/src/github.com/tektoncd/hub

COPY upstream .
COPY .konflux/patches patches/
RUN set -e; for f in patches/*.patch; do echo ${f}; [[ -f ${f} ]] || continue; git apply ${f}; done
COPY head HEAD

ENV GODEBUG="http2server=0"
RUN go build -ldflags="-X 'knative.dev/pkg/changeset.rev=$(cat HEAD)'" -mod=vendor -tags disable_gcp -v -o /tmp/hub-db-migration \
    ./api/cmd/db

FROM $RUNTIME
ARG VERSION=hub-main

COPY --from=builder /tmp/hub-db-migration /ko-app/hub-db-migration
COPY head ${KO_DATA_PATH}/HEAD

EXPOSE 8000

LABEL \
    com.redhat.component="openshift-pipelines-hub-db-migration-container" \
    name="openshift-pipelines/pipelines-hub-db-migration-rhel8" \
    version=$VERSION \
    summary="Red Hat OpenShift Pipelines Hub DB Migration" \
    maintainer="pipelines-extcomm@redhat.com" \
    description="Red Hat OpenShift Pipelines Hub DB Migration" \
    io.k8s.display-name="Red Hat OpenShift Pipelines Hub DB Migration" \
    io.k8s.description="Red Hat OpenShift Pipelines Hub DB Migration" \
    io.openshift.tags="pipelines,tekton,openshift"


RUN groupadd -r -g 65532 nonroot && useradd --no-log-init -r -u 65532 -g nonroot nonroot
USER 65532

ENTRYPOINT ["/ko-app/hub-db-migration"]
