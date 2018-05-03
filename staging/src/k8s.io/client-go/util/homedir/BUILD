package(default_visibility = ["//visibility:public"])

load(
    "@io_bazel_rules_go//go:def.bzl",
    "go_library",
)

go_library(
    name = "go_default_library",
    srcs = select({
        "@io_bazel_rules_go//go/platform:android": [
            "homedir.go",
        ],
        "@io_bazel_rules_go//go/platform:darwin": [
            "homedir.go",
        ],
        "@io_bazel_rules_go//go/platform:dragonfly": [
            "homedir.go",
        ],
        "@io_bazel_rules_go//go/platform:freebsd": [
            "homedir.go",
        ],
        "@io_bazel_rules_go//go/platform:linux": [
            "homedir.go",
        ],
        "@io_bazel_rules_go//go/platform:nacl": [
            "homedir.go",
        ],
        "@io_bazel_rules_go//go/platform:netbsd": [
            "homedir.go",
        ],
        "@io_bazel_rules_go//go/platform:openbsd": [
            "homedir.go",
        ],
        "@io_bazel_rules_go//go/platform:plan9": [
            "homedir.go",
        ],
        "@io_bazel_rules_go//go/platform:solaris": [
            "homedir.go",
        ],
        "@io_bazel_rules_go//go/platform:windows": [
            "homedir_windows.go",
        ],
        "//conditions:default": [],
    }),
    importpath = "k8s.io/client-go/util/homedir",
    deps = select({
        "@io_bazel_rules_go//go/platform:windows": [
            "//vendor/golang.org/x/sys/windows:go_default_library",
        ],
        "//conditions:default": [],
    }),
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
)
