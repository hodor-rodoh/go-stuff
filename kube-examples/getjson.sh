#!/usr/bin/env bash

set -o errexit
set -o pipefail
# set -x
BIN_DIR=$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )

# kubectl config get-contexts
#
# kubectl --context *replace* get ingress --all-namespaces -o json | \
# jq -r '[.items[] | { app: .metadata.name, namespace: .metadata.namespace }]' > output.json
#
# kubectl --context *replace* get secrets --all-namespaces -o json | \
# jq -r '[.items[] | { app: .metadata.name, namespace: .metadata.namespace }]' > output.json

command_output() {

export KUBECONFIG=~/.kube/config2

  while :; do
    case $1 in
      dev_ingress|dev_ing)
        kubectl --context *replace* get ingress --all-namespaces -o json | \
        jq -r '[.items[] | { app: .metadata.name, namespace: .metadata.namespace }]' > output.json
        ;;
      dev_secrets|dev_sec)
        kubectl --context *replace* get secrets --all-namespaces -o json | \
        jq -r '[.items[] | { app: .metadata.name, namespace: .metadata.namespace }]' > output.json
        ;;
      qa_ingress|qa_ing)
        kubectl --context *replace* get ingress --all-namespaces -o json | \
        jq -r '[.items[] | { app: .metadata.name, namespace: .metadata.namespace }]' > output.json
        ;;
      qa_secrets|qa_sec)
        kubectl --context *replace* get secrets --all-namespaces -o json | \
        jq -r '[.items[] | { app: .metadata.name, namespace: .metadata.namespace }]' > output.json
        ;;
      stg_ingress|stg_ing)
        kubectl --context *replace* get ingress --all-namespaces -o json | \
        jq -r '[.items[] | { app: .metadata.name, namespace: .metadata.namespace }]' > output.json
        ;;
      stg_secrets|stg_sec)
        kubectl --context *replace* get secrets --all-namespaces -o json | \
        jq -r '[.items[] | { app: .metadata.name, namespace: .metadata.namespace }]' > output.json
        ;;
      prod_ingress|prd_ing)
        kubectl --context *replace* get ingress --all-namespaces -o json | \
        jq -r '[.items[] | { app: .metadata.name, namespace: .metadata.namespace }]' > output.json
        ;;
      prod_secrets|prd_sec)
        kubectl --context *replace* get secrets --all-namespaces -o json | \
        jq -r '[.items[] | { app: .metadata.name, namespace: .metadata.namespace }]' > output.json
        ;;
    esac
    shift
  done
}

show_help() {
cat << EOF
Usage: ${0##*/}
Description
    -h, --help
        Display this help and exit
    -prd, -p
        get data from prod
EOF

}

err() {
  printf " - ERROR : [$(date +'%Y-%m-%dT%H:%M:%S%z')]: ERROR: ===> %s \\n " "$*" >&2
  exit 1
}

main() {
  if [[ -z "$1" ]]; then
    show_help
    exit 1
  fi

  while :; do
    case $1 in
      -h|-\?|--help)
        show_help
        exit
        ;;
      -get|--g)
        command_output $2
        shift
        ;;
      *)
        err "Unknown option: $1"
    esac
    shift
  done
}

# go run parse.go

main "$@"
