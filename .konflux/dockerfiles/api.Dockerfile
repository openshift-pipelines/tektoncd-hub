# Rebuild trigger: 1.15.4 release 2026-01-19
ARG GO_BUILDER=registry.redhat.io/ubi8/go-toolset:latest@sha256:f10be871f461b5b2920559e55cf3e594dc1758941b835b7a62b6bd1f08464181
ARG RUNTIME=registry.redhat.io/ubi8/ubi@sha256:b9d42de318b28ca7025c0c912769a344db82a89f5e3cf84b90d8f75a7ea11310

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

# Red Hat's UBI8 CDN occasionally serves momentarily inconsistent repodata
# across the baseos/appstream repos (e.g. mismatched openssl/openssl-libs
# builds), causing a transient dependency-resolution failure here. Refresh
# the metadata cache and retry a few times to ride out that skew instead of
# failing the whole image build on a transient upstream glitch.
RUN for i in 1 2 3 4 5; do \
        dnf clean all && dnf makecache --refresh && \
        dnf install -y openssh-clients git shadow-utils && exit 0; \
        echo "dnf install attempt ${i} failed, retrying in 15s..."; \
        sleep 15; \
    done; \
    echo "dnf install failed after 5 attempts"; exit 1

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
    version="v1.15.6"

RUN groupadd -r -g 65532 nonroot && useradd --no-log-init -r -u 65532 -g nonroot nonroot
USER 65532

ENTRYPOINT ["/ko-app/hub-api-server"]# trigger rebuild 2026-02-14
