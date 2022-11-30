#!/usr/bin/env bash

# Copyright 2014 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

set -o errexit
set -o nounset
set -o pipefail

REPO_ROOT=$(dirname "${BASH_SOURCE[0]}")/..
# shellcheck source=hack/common-vars.sh
source "${REPO_ROOT}/hack/common-vars.sh"

make --directory="${REPO_ROOT}" "${KUSTOMIZE##*/}"

flavors_dir="${REPO_ROOT}/templates/flavors/"
ci_dir="${REPO_ROOT}/templates/test/ci/"
dev_dir="${REPO_ROOT}/templates/test/dev/"

# NOTE: On macOS, the kustomize commands run by xargs that use the -I
# replacement string may only be 255 characters. Commands longer than that will
# not expand the replacement string when that would go over that limit,
# resulting in generated manifests like "templates/test/ci/cluster-template-{}".
#
# Currently, the longest command is *exactly* 255 characters, so making these
# commands any longer or adding a new template with a longer name will run into
# this problem.

find "${flavors_dir}"* -maxdepth 0 -type d -print0 | xargs -0 -I {} basename {} | grep -v base | xargs -I {} sh -c "${KUSTOMIZE} build --load-restrictor LoadRestrictionsNone --reorder none ${flavors_dir}{} > ${REPO_ROOT}/templates/cluster-template-{}.yaml"
# move the default template to the default file expected by clusterctl
mv "${REPO_ROOT}/templates/cluster-template-default.yaml" "${REPO_ROOT}/templates/cluster-template.yaml"

find "${ci_dir}"* -maxdepth 0 -type d -print0 | xargs -0 -I {} basename {} | grep -v patches | xargs -I {} sh -c "${KUSTOMIZE} build --load-restrictor LoadRestrictionsNone --reorder none ${ci_dir}{} > ${ci_dir}cluster-template-{}.yaml"
find "${dev_dir}"* -maxdepth 0 -type d -print0 | xargs -0 -I {} basename {} | grep -v patches | xargs -I {} sh -c "${KUSTOMIZE} build --load-restrictor LoadRestrictionsNone --reorder none ${dev_dir}{} > ${dev_dir}cluster-template-{}.yaml"
