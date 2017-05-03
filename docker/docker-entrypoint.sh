#!/bin/bash

set -eo pipefail

# Environment variable and params (alphabetical order)
# BACKEND_SCHEME =
# BACKEND_TARGET =

exec /simple-proxy --backend-scheme ${BACKEND_SCHEME} --backend-target ${BACKEND_TARGET}
return
