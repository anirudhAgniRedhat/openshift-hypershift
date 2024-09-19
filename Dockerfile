FROM registry.ci.openshift.org/openshift/release:rhel-9-release-golang-1.22-openshift-4.18 as builder

WORKDIR /hypershift

COPY . .

RUN make hypershift hypershift-operator product-cli

FROM registry.access.redhat.com/ubi9:latest
COPY --from=builder /hypershift/bin/hypershift \
                    /hypershift/bin/hcp \
                    /hypershift/bin/hypershift-operator \
     /usr/bin/

ENTRYPOINT ["/usr/bin/hypershift"]

LABEL io.openshift.hypershift.control-plane-operator-subcommands=true
LABEL io.openshift.hypershift.control-plane-operator-skips-haproxy=true
LABEL io.openshift.hypershift.ignition-server-healthz-handler=true
LABEL io.openshift.hypershift.control-plane-operator-manages-ignition-server=true
LABEL io.openshift.hypershift.control-plane-operator-manages.cluster-machine-approver=true
LABEL io.openshift.hypershift.control-plane-operator-manages.cluster-autoscaler=true
LABEL io.openshift.hypershift.control-plane-operator-manages.decompress-decode-config=true
LABEL io.openshift.hypershift.control-plane-operator-creates-aws-sg=true
LABEL io.openshift.hypershift.control-plane-operator-applies-management-kas-network-policy-label=true
LABEL io.openshift.hypershift.restricted-psa=true
LABEL io.openshift.hypershift.control-plane-pki-operator-signs-csrs=true
LABEL io.openshift.hypershift.hosted-cluster-config-operator-reports-node-count=true
