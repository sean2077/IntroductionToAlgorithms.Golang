# Minimal makefile for Sphinx documentation
#

# You can set these variables from the command line, and also
# from the environment for the first two.
SPHINXOPTS    ?=
SPHINXBUILD   ?= sphinx-build
SOURCEDIR     = .
BUILDDIR      = _build

# Put it first so that "make" without argument is like "make help".
help:
	@$(SPHINXBUILD) -M help "$(SOURCEDIR)" "$(BUILDDIR)" $(SPHINXOPTS) $(O)

.PHONY: help Makefile

# Catch-all target: route all unknown targets to Sphinx using the new
# "make mode" option.  $(O) is meant as a shortcut for $(SPHINXOPTS).
%: Makefile
	@$(SPHINXBUILD) -M $@ "$(SOURCEDIR)" "$(BUILDDIR)" $(SPHINXOPTS) $(O)

clean:
	rm -rf docs/* _build

github: clean html
	@cp -a _build/html/. docs

VERSION := 0.0.1
GIT_SHORT_HASH=$(shell git rev-parse --short HEAD)
BUILD_DATE=$(shell date -u '+%Y-%m-%d-%H:%M:%S')
IMAGE_TAG := v$(strip $(VERSION))-$(GIT_SHORT_HASH)$(IMAGE_SUFFIX)
RUNTIME_IMAGE := zhangxianbing/python-runtime:$(IMAGE_TAG)

runtime-image: ## 构建运行时镜像
	docker build -f dockerfiles/Dockerfile.runtime $(IMAGE_FLAGS) -t $(RUNTIME_IMAGE) .
	docker push $(RUNTIME_IMAGE)