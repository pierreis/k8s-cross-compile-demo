load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

# ================================================================
# Service
# ================================================================

go_library(
  name = "go_default_library",
  srcs = [
      "main.go",
  ],
  importpath = "github.com/pierreis/k8s_cross_compile_demo",
  deps = [],
  visibility = [
    "//visibility:private"
  ],
)

go_binary(
  name = "server",
  embed = [
    ":go_default_library"
  ],
  visibility = [
    "//visibility:public"
  ],
  #goarch = "amd64",
  #goos = "linux",
  #pure = "on",
)