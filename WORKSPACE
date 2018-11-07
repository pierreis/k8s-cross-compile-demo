workspace(name = "k8s_cross_compile_demo")

load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

# ================================================================
# Go support
# ================================================================

http_archive(
  name = "io_bazel_rules_go",
  urls = ["https://github.com/bazelbuild/rules_go/releases/download/0.16.2/rules_go-0.16.2.tar.gz"],
  sha256 = "f87fa87475ea107b3c69196f39c82b7bbf58fe27c62a338684c20ca17d1d8613",
)

http_archive(
  name = "bazel_gazelle",
  urls = ["https://github.com/bazelbuild/bazel-gazelle/releases/download/0.15.0/bazel-gazelle-0.15.0.tar.gz"],
  sha256 = "6e875ab4b6bf64a38c352887760f21203ab054676d9c1b274963907e0768740d",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")

go_rules_dependencies()

go_register_toolchains()

# ================================================================
# Docker
# ================================================================

http_archive(
  name = "io_bazel_rules_docker",
  sha256 = "29d109605e0d6f9c892584f07275b8c9260803bf0c6fcb7de2623b2bedc910bd",
  strip_prefix = "rules_docker-0.5.1",
  urls = ["https://github.com/bazelbuild/rules_docker/archive/v0.5.1.tar.gz"],
)

load("@io_bazel_rules_docker//container:container.bzl", "container_pull", container_repositories = "repositories")

container_repositories()

load("@io_bazel_rules_docker//go:image.bzl", go_image_repositories = "repositories")

go_image_repositories()

# ================================================================
# K8s
# ================================================================

git_repository(
  name = "io_bazel_rules_k8s",
  commit = "489310b9bc69a8a79ec4768be84a65037fe5ce9b",
  remote = "https://github.com/bazelbuild/rules_k8s.git",
)

load("@io_bazel_rules_k8s//k8s:k8s.bzl", "k8s_repositories")

k8s_repositories()

load("@io_bazel_rules_k8s//k8s:k8s.bzl", "k8s_defaults")

k8s_defaults(
  name = "k8s_deploy",
  kind = "deployment",
  cluster = "docker-for-desktop-cluster",
  #namespace = "{BUILD_USER}",
  namespace = "default",
  image_chroot = "localhost:5000/hurow-{BUILD_USER}/dev",
)

# ================================================================
# Docker containers
# ================================================================

container_pull(
  name = "golang",
  registry = "index.docker.io",
  repository = "library/golang",
  digest = "sha256:e7462ca504afc789d289f2bb5fd471815cc11833439d2fe4e61915b190045359",
)