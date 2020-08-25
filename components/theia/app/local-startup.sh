#!/bin/bash

OLD_WD="$PWD"
# Create precoditions for startup.sh
if [ ! -d /var/git/ ]; then
    sudo mkdir -p /var/git/
    sudo chown -R $USER /var/git
fi
rm -Rf /var/git/*


set -a
# These values correspond to those in the server: local-theia-mock.ts
export GITPOD_GIT_CLONE_LOCATION="/var/git"
# Use somefox credentials and fixed repository
export GITPOD_GIT_REPO_CLONE_URI="https://x-access-token:30155e44111b78a6fcb137fcc2d9b532979c1843@github.com/TypeFox/gitpod-test-repo.git"
export GITPOD_GIT_CLONE_BRANCH=12test
#   - GITPOD_GIT_CLONE_COMMIT   # Not needed as BRANCH is set
export GITPOD_GIT_REPO_NAME=gitpod-test-repo
export GITPOD_GIT_USER_NAME=somefox
export GITPOD_GIT_USER_EMAIL=somefox@example.org
export GITPOD_THEIA_PORT=23000

# export GITPOD_WORKSPACE_ID="a12-321" # ISSUE workspace
export GITPOD_WORKSPACE_ID="a12-345" # PR workspace
export GITPOD_HOST="http://gitpod.io-local"
export GITPOD_INTERVAL=10000
export GITPOD_TASKS='[{"command": "echo \"hello world\""}]';

# Set to prevent bash config manipulation
export THEIA_LOCAL=true
export THEIA_ARGS_ADD="--log-level=debug"
set +a


###############################################################################
# Clone workspace
###############################################################################
cd "$GITPOD_GIT_CLONE_LOCATION"                                         && \
#echo "Cleaning stuff left behind by former crashes..."                  && \
rm -Rf ./*                                                              && \

# Clone repo
# TODO Use 'clone --depth 1' and 'fetch --unshallow', to reduce initial delay, later?
if [ -n "$GITPOD_GIT_CLONE_BRANCH" ]; then
    echo "Git clone branch: $GITPOD_GIT_CLONE_BRANCH"                   && \
    git clone --branch "$GITPOD_GIT_CLONE_BRANCH"   \
                "$GITPOD_GIT_REPO_CLONE_URI"        \
                "$GITPOD_GIT_REPO_NAME" --progress                      && \
    cd "$GITPOD_GIT_REPO_NAME"
else
    echo "Git clone commit: $GITPOD_GIT_CLONE_COMMIT"                   && \
    git clone "$GITPOD_GIT_REPO_CLONE_URI"          \
                "$GITPOD_GIT_REPO_NAME" --progress                      && \
    cd "$GITPOD_GIT_REPO_NAME"                                          && \
    git checkout "$GITPOD_GIT_CLONE_COMMIT"
fi      


cd "$OLD_WD"
exec ./startup.sh
