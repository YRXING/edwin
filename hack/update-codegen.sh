#!/usr/bin/env bash
#表示有报错即退出 跟set -e含义一样
set -o errexit
#执行脚本的时候，如果遇到不存在的变量，Bash 默认忽略它 ,跟 set -u含义一样
set -o nounset
# 只要一个子命令失败，整个管道命令就失败，脚本就会终止执行
set -o pipefail

MODULE=github.com/YRXING/edwin

#api包
APIS_PKG=pkg/apis

#代码生出输出，生成Resource时指定的group一样
OUTPUT_PKG=pkg/client

# group-version such as cronjob:v1
GROUP=edwin
VERSION=v1
GROUP_VERSION=$GROUP:$VERSION

SCRIPT_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
CODEGEN_PKG=${CODEGEN_PKG:-$(cd "${SCRIPT_ROOT}"; ls -d -1 ./vendor/k8s.io/code-generator 2>/dev/null || echo ../code-generator)}

# generate the code with:
# --output-base    because this script should also be able to run inside the vendor dir of
#                  k8s.io/kubernetes. The output-base is needed for the generators to output into the vendor dir
#                  instead of the $GOPATH directly. For normal projects this can be dropped.
#client,informer,lister(注意: code-generator 生成的deepcopy不适配 kubebuilder 所生成的api)
bash "${CODEGEN_PKG}"/generate-groups.sh "deepcopy,client,informer,lister" \
  ${MODULE}/${OUTPUT_PKG} ${MODULE}/${APIS_PKG} \
  ${GROUP_VERSION} \
  --go-header-file "${SCRIPT_ROOT}"/hack/boilerplate.go.txt
#  --output-base "${SCRIPT_ROOT}"
#  --output-base "${SCRIPT_ROOT}/../../.."
