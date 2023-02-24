
## code generator

```bash
export GOPATH=$pwd
execDir=~/code/cloud-native/code-generator
"$execDir"/generate-groups.sh all \
  kluster/pkg/client \
  kluster/pkg/apis \
  viveksingh.dev:v1alpha1 \
  --go-header-file "$execDir"/hack/boilerplate.go.txt
```
