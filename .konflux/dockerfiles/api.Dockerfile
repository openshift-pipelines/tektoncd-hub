ARG GO_BUILDER=registry.access.redhat.com/ubi9/go-toolset:9.7-1771417345
ARG RUNTIME=registry.redhat.io/ubi9/ubi-minimal@sha256:c7d44146f826037f6873d99da479299b889473492d3c1ab8af86f08af04ec8a0

FROM $GO_BUILDER AS builder

WORKDIR /go/src/github.com/tektoncd/hub

COPY upstream .
COPY .konflux/patches patches/
RUN set -e; for f in patches/*.patch; do echo ${f}; [[ -f ${f} ]] || continue; git apply ${f}; done
COPY head HEAD
 
ENV GODEBUG="http2server=0" 
ENV GOEXPERIMENT=strictfipsruntime
RUN go build -ldflags="-X 'knative.dev/pkg/changeset.rev=$(cat HEAD)'" -mod=vendor -tags disable_gcp,strictfipsruntime -v -o /tmp/hub-api-server \
    ./api/cmd/api

FROM $RUNTIME
ARG VERSION=hub-api-next

RUN microdnf install -y openssh-clients git shadow-utils

COPY --from=builder /tmp/hub-api-server /ko-app/hub-api-server
COPY head ${KO_DATA_PATH}/HEAD

EXPOSE 8000

LABEL \
      com.redhat.component="openshift-pipelines-hub-api-rhel9-container" \
      cpe="cpe:/a:redhat:openshift_pipelines:1.22::el9" \
      description="Red Hat OpenShift Pipelines tektoncd-hub api" \
      io.k8s.description="Red Hat OpenShift Pipelines tektoncd-hub api" \
      io.k8s.display-name="Red Hat OpenShift Pipelines tektoncd-hub api" \
      io.openshift.tags="tekton,openshift,tektoncd-hub,api" \
      maintainer="pipelines-extcomm@redhat.com" \
      name="openshift-pipelines/pipelines-hub-api-rhel9" \
      summary="Red Hat OpenShift Pipelines tektoncd-hub api" \
      version="v1.22.0"

RUN groupadd -r -g 65532 nonroot && useradd --no-log-init -r -u 65532 -g nonroot nonroot
USER 65532

ENTRYPOINT ["/ko-app/hub-api-server"]
