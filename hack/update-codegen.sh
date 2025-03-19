set -o errexit
set -o nounset
set -o pipefail
set -x

SCRIPT_ROOT=$(dirname ${BASH_SOURCE})/..
CODEGEN_PKG=~/go/pkg/mod/k8s.io/code-generator@v0.30.10

source "${CODEGEN_PKG}/kube_codegen.sh"

THIS_PKG="github.com/kubeovn/kube-ovn"

kube::codegen::gen_helpers \
    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
    "${SCRIPT_ROOT}/pkg/apis"

kube::codegen::gen_client \
    --with-watch \
    --output-dir "${SCRIPT_ROOT}/pkg/client" \
    --output-pkg "${THIS_PKG}/pkg/client" \
    --boilerplate "${SCRIPT_ROOT}/hack/boilerplate.go.txt" \
    "${SCRIPT_ROOT}/pkg/apis"
