# Build Docker image

load("@io_bazel_rules_docker//container:container.bzl", "container_image")

container_image(
  name = "image",
  base = "@go_image_base//image",
  files = ["//server"],
  cmd = ["/server"],
  visibility = [
    "//visibility:public",
  ],
)

# Deploy to k8s

load("@k8s_deploy//:defaults.bzl", "k8s_deploy")

k8s_deploy(
  name = "deployment",
  template = ":deployment.yaml",
  images = {
    "server:dev": ":image"
  },
)