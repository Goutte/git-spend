#!/usr/bin/env sh

# Installation script for git-spend
#
# Usage:
#     curl https://raw.githubusercontent.com/Goutte/git-spend/main/install.sh | sh
#
# Will ask for password, if you want no interaction pipe the curl into "sudo sh" instead.
#
# I've set this up with sh, but I have no objection against using bash instead at some point.
# zsh is fine too, but perhaps has less support out of the box.
#
# Overview
# --------
# 1. Detect OS & ARCH
# 2. Download the appropriate binary
# 3. Install that binary
# 4. Install man pages
# 5. Install git hook?
#

# -----

set -e

DOWNLOAD_URL=https://github.com/Goutte/git-spend/releases/latest/download/git-spend
BINARY_INSTALL_PATH=/usr/local/bin
BINARY_FILENAME=git-spend

# -----

echo "You are about to install git-spend on your system."

# 2. Download the appropriate binary
echo "Let's download the latest release from github.com…"
curl --location ${DOWNLOAD_URL} > "${BINARY_FILENAME}"

# 3. Install that binary
echo "Installing ${BINARY_FILENAME} to ${BINARY_INSTALL_PATH}/${BINARY_FILENAME} requires superuser privileges…"
sudo install --preserve-timestamps "${BINARY_FILENAME}" "${BINARY_INSTALL_PATH}/${BINARY_FILENAME}"

# 4. Install man pages
echo "Installing man pages for ${BINARY_FILENAME}…"
sudo "${BINARY_INSTALL_PATH}/git-spend" man --install

# ---

echo "All done !"
echo ""
echo "You can now use:"
echo "    git spend sum"
echo "in git projects where you have '/spend <duration>' directives in commits."
echo ""
echo "PLEASE MAKE SURE THIS SOFTWARE IS NOT USED TO OPPRESS"
echo "AS IT WOULD BE AGAINST ITS LICENCE -- #CodersUnion"
