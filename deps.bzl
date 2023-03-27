load("@bazel_gazelle//:deps.bzl", "go_repository")

def go_dependencies():
    go_repository(
        name = "com_github_masterminds_cookoo",
        importpath = "github.com/Masterminds/cookoo",
        sum = "h1:zwplWkfGEd4NxiL0iZHh5Jh1o25SUJTKWLfv2FkXh6o=",
        version = "v1.3.0",
    )
    go_repository(
        name = "com_github_pkg_errors",
        importpath = "github.com/pkg/errors",
        sum = "h1:FEBLx1zS214owpjy7qsBeixbURkuhQAwrK5UwLGTwt4=",
        version = "v0.9.1",
    )
