workspace(name = "kubevirt")

load("//third_party:deps.bzl", "deps")

deps()

# register crosscompiler toolchains
load("//bazel/toolchain:toolchain.bzl", "register_all_toolchains")

register_all_toolchains()

load(
    "@bazel_tools//tools/build_defs/repo:http.bzl",
    "http_archive",
    "http_file",
)
load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

http_archive(
    name = "rules_python",
    sha256 = "934c9ceb552e84577b0faf1e5a2f0450314985b4d8712b2b70717dc679fdc01b",
    urls = [
        "https://github.com/bazelbuild/rules_python/releases/download/0.3.0/rules_python-0.3.0.tar.gz",
        "https://storage.googleapis.com/builddeps/934c9ceb552e84577b0faf1e5a2f0450314985b4d8712b2b70717dc679fdc01b",
    ],
)

# Additional bazel rules
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

http_archive(
    name = "rules_proto",
    sha256 = "bc12122a5ae4b517fa423ea03a8d82ea6352d5127ea48cb54bc324e8ab78493c",
    strip_prefix = "rules_proto-af6481970a34554c6942d993e194a9aed7987780",
    urls = [
        "https://github.com/bazelbuild/rules_proto/archive/af6481970a34554c6942d993e194a9aed7987780.tar.gz",
        "https://storage.googleapis.com/builddeps/bc12122a5ae4b517fa423ea03a8d82ea6352d5127ea48cb54bc324e8ab78493c",
    ],
)

load("@rules_proto//proto:repositories.bzl", "rules_proto_dependencies", "rules_proto_toolchains")

rules_proto_dependencies()

rules_proto_toolchains()

http_archive(
    name = "io_bazel_rules_go",
    sha256 = "2b1641428dff9018f9e85c0384f03ec6c10660d935b750e3fa1492a281a53b0f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.29.0/rules_go-v0.29.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.29.0/rules_go-v0.29.0.zip",
        "https://storage.googleapis.com/builddeps/2b1641428dff9018f9e85c0384f03ec6c10660d935b750e3fa1492a281a53b0f",
    ],
)

http_archive(
    name = "bazel_gazelle",
    sha256 = "de69a09dc70417580aabf20a28619bb3ef60d038470c7cf8442fafcf627c21cb",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.24.0/bazel-gazelle-v0.24.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.24.0/bazel-gazelle-v0.24.0.tar.gz",
        "https://storage.googleapis.com/builddeps/de69a09dc70417580aabf20a28619bb3ef60d038470c7cf8442fafcf627c21cb",
    ],
)

http_archive(
    name = "io_bazel_rules_docker",
    sha256 = "95d39fd84ff4474babaf190450ee034d958202043e366b9fc38f438c9e6c3334",
    strip_prefix = "rules_docker-0.16.0",
    urls = [
        "https://github.com/bazelbuild/rules_docker/releases/download/v0.16.0/rules_docker-v0.16.0.tar.gz",
        "https://storage.googleapis.com/builddeps/95d39fd84ff4474babaf190450ee034d958202043e366b9fc38f438c9e6c3334",
    ],
)

http_archive(
    name = "com_github_ash2k_bazel_tools",
    sha256 = "80ba082177c93e43a7c085a8566c7f11654dbae41da7da0da52e0ed2e917cd12",
    strip_prefix = "bazel-tools-6e2a416f565062955735edcfae881cdba2b7abf7",
    urls = [
        "https://github.com/ash2k/bazel-tools/archive/6e2a416f565062955735edcfae881cdba2b7abf7.zip",
        "https://storage.googleapis.com/builddeps/80ba082177c93e43a7c085a8566c7f11654dbae41da7da0da52e0ed2e917cd12",
    ],
)

# Disk images
http_file(
    name = "alpine_image",
    sha256 = "e97eaedb3bff39a081d1d7e67629d5c0e8fb39677d6a9dd1eaf2752e39061e02",
    urls = [
        "https://dl-cdn.alpinelinux.org/alpine/v3.15/releases/x86_64/alpine-virt-3.15.0-x86_64.iso",
        "https://storage.googleapis.com/builddeps/e97eaedb3bff39a081d1d7e67629d5c0e8fb39677d6a9dd1eaf2752e39061e02",
    ],
)

http_file(
    name = "alpine_image_aarch64",
    sha256 = "f302cf1b2dbbd0661b8f53b167f24131c781b86ab3ae059654db05cd62d3c39c",
    urls = [
        "https://dl-cdn.alpinelinux.org/alpine/v3.15/releases/aarch64/alpine-virt-3.15.0-aarch64.iso",
        "https://storage.googleapis.com/builddeps/f302cf1b2dbbd0661b8f53b167f24131c781b86ab3ae059654db05cd62d3c39c",
    ],
)

http_file(
    name = "cirros_image",
    sha256 = "932fcae93574e242dc3d772d5235061747dfe537668443a1f0567d893614b464",
    urls = [
        "https://download.cirros-cloud.net/0.5.2/cirros-0.5.2-x86_64-disk.img",
        "https://storage.googleapis.com/builddeps/932fcae93574e242dc3d772d5235061747dfe537668443a1f0567d893614b464",
    ],
)

http_file(
    name = "cirros_image_aarch64",
    sha256 = "889c1117647b3b16cfc47957931c6573bf8e755fc9098fdcad13727b6c9f2629",
    urls = [
        "https://download.cirros-cloud.net/0.5.2/cirros-0.5.2-aarch64-disk.img",
        "https://storage.googleapis.com/builddeps/889c1117647b3b16cfc47957931c6573bf8e755fc9098fdcad13727b6c9f2629",
    ],
)

http_file(
    name = "virtio_win_image",
    sha256 = "b8a4bc66835c43091a85d35a10b59bd8b1b62b55ea9f02ec754f68bd32e82c0e",
    urls = [
        "https://fedorapeople.org/groups/virt/virtio-win/direct-downloads/archive-virtio/virtio-win-0.1.217-1/virtio-win-0.1.217.iso",
        "https://storage.googleapis.com/builddeps/b8a4bc66835c43091a85d35a10b59bd8b1b62b55ea9f02ec754f68bd32e82c0e",
    ],
)

http_archive(
    name = "bazeldnf",
    sha256 = "c37709d05ad7eae4d32d7a525f098fd026483ada5e11cdf84d47028222796605",
    strip_prefix = "bazeldnf-0.5.2",
    urls = [
        "https://github.com/rmohr/bazeldnf/archive/v0.5.2.tar.gz",
        "https://storage.googleapis.com/builddeps/c37709d05ad7eae4d32d7a525f098fd026483ada5e11cdf84d47028222796605",
    ],
)

load(
    "@io_bazel_rules_go//go:deps.bzl",
    "go_register_toolchains",
    "go_rules_dependencies",
)
load("@bazeldnf//:deps.bzl", "bazeldnf_dependencies", "rpm")

go_rules_dependencies()

go_register_toolchains(
    go_version = "1.17.8",
    nogo = "@//:nogo_vet",
)

load("@com_github_ash2k_bazel_tools//goimports:deps.bzl", "goimports_dependencies")

goimports_dependencies()

load(
    "@bazel_gazelle//:deps.bzl",
    "gazelle_dependencies",
    "go_repository",
)

gazelle_dependencies()

load("@com_github_bazelbuild_buildtools//buildifier:deps.bzl", "buildifier_dependencies")

buildifier_dependencies()

load(
    "@bazel_tools//tools/build_defs/repo:git.bzl",
    "git_repository",
)

# Winrmcli dependencies
go_repository(
    name = "com_github_masterzen_winrmcli",
    commit = "c85a68ee8b6e3ac95af2a5fd62d2f41c9e9c5f32",
    importpath = "github.com/masterzen/winrm-cli",
)

# Winrmcp deps
go_repository(
    name = "com_github_packer_community_winrmcp",
    commit = "c76d91c1e7db27b0868c5d09e292bb540616c9a2",
    importpath = "github.com/packer-community/winrmcp",
)

go_repository(
    name = "com_github_masterzen_winrm_cli",
    commit = "6f0c57dee4569c04f64c44c335752b415e5d73a7",
    importpath = "github.com/masterzen/winrm-cli",
)

go_repository(
    name = "com_github_masterzen_winrm",
    commit = "1d17eaf15943ca3554cdebb3b1b10aaa543a0b7e",
    importpath = "github.com/masterzen/winrm",
)

go_repository(
    name = "com_github_nu7hatch_gouuid",
    commit = "179d4d0c4d8d407a32af483c2354df1d2c91e6c3",
    importpath = "github.com/nu7hatch/gouuid",
)

go_repository(
    name = "com_github_dylanmei_iso8601",
    commit = "2075bf119b58e5576c6ed9f867b8f3d17f2e54d4",
    importpath = "github.com/dylanmei/iso8601",
)

go_repository(
    name = "com_github_gofrs_uuid",
    commit = "abfe1881e60ef34074c1b8d8c63b42565c356ed6",
    importpath = "github.com/gofrs/uuid",
)

go_repository(
    name = "com_github_christrenkamp_goxpath",
    commit = "c5096ec8773dd9f554971472081ddfbb0782334e",
    importpath = "github.com/ChrisTrenkamp/goxpath",
)

go_repository(
    name = "com_github_azure_go_ntlmssp",
    commit = "4a21cbd618b459155f8b8ee7f4491cd54f5efa77",
    importpath = "github.com/Azure/go-ntlmssp",
)

go_repository(
    name = "com_github_masterzen_simplexml",
    commit = "31eea30827864c9ab643aa5a0d5b2d4988ec8409",
    importpath = "github.com/masterzen/simplexml",
)

go_repository(
    name = "org_golang_x_crypto",
    commit = "4def268fd1a49955bfb3dda92fe3db4f924f2285",
    importpath = "golang.org/x/crypto",
)

# override rules_docker issue with this dependency
# rules_docker 0.16 uses 0.1.4, bit since there the checksum changed, which is very weird, going with 0.1.4.1 to
go_repository(
    name = "com_github_google_go_containerregistry",
    importpath = "github.com/google/go-containerregistry",
    sha256 = "bc0136a33f9c1e4578a700f7afcdaa1241cfff997d6bba695c710d24c5ae26bd",
    strip_prefix = "google-go-containerregistry-efb2d62",
    type = "tar.gz",
    urls = ["https://api.github.com/repos/google/go-containerregistry/tarball/efb2d62d93a7705315b841d0544cb5b13565ff2a"],  # v0.1.4.1
)

# bazel docker rules
load(
    "@io_bazel_rules_docker//container:container.bzl",
    "container_image",
    "container_pull",
)
load(
    "@io_bazel_rules_docker//repositories:repositories.bzl",
    container_repositories = "repositories",
)

container_repositories()

load("@io_bazel_rules_docker//repositories:deps.bzl", container_deps = "deps")

container_deps()

# Pull base image fedora31
# WARNING: please update any automated process to push this image to quay.io
# instead of index.docker.io
container_pull(
    name = "fedora",
    digest = "sha256:5e2b864cfe165fa7da6606b29a9e60549eb7cc9ae7fb574614110d1494b0f0c2",
    registry = "quay.io",
    repository = "kubevirtci/fedora",
    tag = "31",
)

# As rpm package in https://dl.fedoraproject.org/pub/fedora/linux/releases/31 is empty, we use fedora 32 here.
# TODO add fedora image to quay.io
container_pull(
    name = "fedora_aarch64",
    digest = "sha256:425676dd30f2c85ba3593b82040ce03341cd6dc4e38838e57c8bc5eef95b5f81",
    registry = "index.docker.io",
    repository = "library/fedora",
    tag = "32",
)

# Pull go_image_base
container_pull(
    name = "go_image_base",
    digest = "sha256:f65536ce108fcc41cdcd5cb101006fcb82b9a1527409263feb9e34032f00bda0",
    registry = "gcr.io",
    repository = "distroless/base",
)

container_pull(
    name = "go_image_base_aarch64",
    digest = "sha256:789c477fbd30a7d85435450306e54f20c53938e40af644284a229d852db30dde",
    registry = "gcr.io",
    repository = "distroless/base",
)

# Pull nfs-server image
# WARNING: please update any automated process to push this image to quay.io
# instead of index.docker.io
# TODO build nfs-server for multi-arch
container_pull(
    name = "nfs-server",
    digest = "sha256:8c1fa882dddb2885c4152e9ce632c466f4b8dce29339455e9b6bfe71f0a3d3ef",
    registry = "quay.io",
    repository = "kubevirtci/nfs-ganesha",  # see https://github.com/slintes/docker-nfs-ganesha
)

container_pull(
    name = "nfs-server_aarch64",
    digest = "sha256:8c1fa882dddb2885c4152e9ce632c466f4b8dce29339455e9b6bfe71f0a3d3ef",
    registry = "quay.io",
    repository = "kubevirtci/nfs-ganesha",  # see https://github.com/slintes/docker-nfs-ganesha
)

# Pull fedora container-disk preconfigured with ci tooling
# like stress and qemu guest agent pre-configured
# TODO build fedora_with_test_tooling for multi-arch
container_pull(
    name = "fedora_with_test_tooling",
    digest = "sha256:da6c118dbd9ac643713c1737cbaa43dcc7386b269b4beb0984413168f3a5f2d3",
    registry = "quay.io",
    repository = "kubevirtci/fedora-with-test-tooling",
)

container_pull(
    name = "alpine_with_test_tooling",
    digest = "sha256:d1dab23ed46af711acb33e54b1dd2a7c6dfaab24227346a487748057e2c81d11",
    registry = "quay.io",
    repository = "kubevirtci/alpine-with-test-tooling-container-disk",
    tag = "2206291207-35b9c64",
)

container_pull(
    name = "fedora_with_test_tooling_aarch64",
    digest = "sha256:9b1371260c05086a24ac9effdbedca9759c885ea8db93de7f0339df3bcd5a5c3",
    registry = "quay.io",
    repository = "kubevirtci/fedora-with-test-tooling",
)

container_pull(
    name = "alpine-ext-kernel-boot-demo-container-base",
    digest = "sha256:a2ddb2f568bf3814e594a14bc793d5a655a61d5983f3561d60d02afa7bbc56b4",
    registry = "quay.io",
    repository = "kubevirt/alpine-ext-kernel-boot-demo",
)

# TODO build fedora_realtime for multi-arch
container_pull(
    name = "fedora_realtime",
    digest = "sha256:437f4e02986daf0058239f4a282d32304dcac629d5d1b4c75a74025f1ce22811",
    registry = "quay.io",
    repository = "kubevirt/fedora-realtime-container-disk",
)

load(
    "@io_bazel_rules_docker//go:image.bzl",
    _go_image_repos = "repositories",
)

_go_image_repos()

http_archive(
    name = "io_bazel_rules_container_rpm",
    sha256 = "151261f1b81649de6e36f027c945722bff31176f1340682679cade2839e4b1e1",
    strip_prefix = "rules_container_rpm-0.0.5",
    urls = [
        "https://github.com/rmohr/rules_container_rpm/archive/v0.0.5.tar.gz",
        "https://storage.googleapis.com/builddeps/151261f1b81649de6e36f027c945722bff31176f1340682679cade2839e4b1e1",
    ],
)

http_file(
    name = "libguestfs-appliance",
    sha256 = "51d38a062d1b91bd7cb3dd8e68354aae86f6a889b4bb68a358b3ab55030dc0c9",
    urls = [
        "https://storage.googleapis.com/kubevirt-prow/devel/release/kubevirt/libguestfs-appliance/appliance-1.44.0-linux-4.18.0-338-centos8.tar.xz",
    ],
)

# Get container-disk-v1alpha RPM's
http_file(
    name = "qemu-img",
    sha256 = "669250ad47aad5939cf4d1b88036fd95a94845d8e0bbdb05e933f3d2fe262fea",
    urls = ["https://storage.googleapis.com/builddeps/669250ad47aad5939cf4d1b88036fd95a94845d8e0bbdb05e933f3d2fe262fea"],
)

# some repos which are not part of go_rules anymore
go_repository(
    name = "com_github_golang_glog",
    importpath = "github.com/golang/glog",
    sum = "h1:VKtxabqXZkF25pY9ekfRL6a582T4P37/31XEstQ5p58=",
    version = "v0.0.0-20160126235308-23def4e6c14b",
)

go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    sum = "h1:M5a8xTlYTxwMn5ZFkwhRabsygDY5G8TYLyQDBxJNAxE=",
    version = "v1.30.0",
)

go_repository(
    name = "org_golang_x_net",
    importpath = "golang.org/x/net",
    sum = "h1:oWX7TPOiFAMXLq8o0ikBYfCJVlRHBcsciT5bXOrH628=",
    version = "v0.0.0-20190311183353-d8887717615a",
)

go_repository(
    name = "org_golang_x_text",
    importpath = "golang.org/x/text",
    sum = "h1:g61tztE5qeGQ89tm6NTjjM9VPIm088od1l6aSorWRWg=",
    version = "v0.3.0",
)

register_toolchains("//:py_toolchain")

go_repository(
    name = "org_golang_x_mod",
    build_file_generation = "on",
    build_file_proto_mode = "disable",
    importpath = "golang.org/x/mod",
    sum = "h1:RM4zey1++hCTbCVQfnWeKs9/IEsaBLA8vTkd0WVtmH4=",
    version = "v0.3.0",
)

go_repository(
    name = "org_golang_x_xerrors",
    build_file_generation = "on",
    build_file_proto_mode = "disable",
    importpath = "golang.org/x/xerrors",
    sum = "h1:go1bK/D/BFZV2I8cIQd1NKEZ+0owSTG1fDTci4IqFcE=",
    version = "v0.0.0-20200804184101-5ec99f83aff1",
)

bazeldnf_dependencies()

rpm(
    name = "SDL2-0__2.0.14-5.el8.x86_64",
    sha256 = "bd69354c64f58177a303264a39ad51cb13332daf624e73cc5df428ed8e416798",
    urls = ["https://pkgs.dyn.su/el8/extras/x86_64/SDL2-2.0.14-5.el8.x86_64.rpm"],
)

rpm(
    name = "acl-0__2.2.53-1.el8.aarch64",
    sha256 = "47c2cc5872174c548de1096dc5673ee91349209d89e0193a4793955d6865b3b1",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/acl-2.2.53-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/47c2cc5872174c548de1096dc5673ee91349209d89e0193a4793955d6865b3b1",
    ],
)

rpm(
    name = "acl-0__2.2.53-1.el8.x86_64",
    sha256 = "227de6071cd3aeca7e10ad386beaf38737d081e06350d02208a3f6a2c9710385",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/acl-2.2.53-1.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/227de6071cd3aeca7e10ad386beaf38737d081e06350d02208a3f6a2c9710385",
    ],
)

rpm(
    name = "adwaita-cursor-theme-0__3.28.0-3.el8.x86_64",
    sha256 = "ef9022d1bcd60a523c64bb3c754ba08dfcf6fed1c441b682115d6bd636e888d5",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/adwaita-cursor-theme-3.28.0-3.el8.noarch.rpm"],
)

rpm(
    name = "adwaita-icon-theme-0__3.28.0-3.el8.x86_64",
    sha256 = "362f2807e1c4c081e9ff2bfa754d7e7afbee795d718512c29f768834c987d014",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/adwaita-icon-theme-3.28.0-3.el8.noarch.rpm"],
)

rpm(
    name = "at-spi2-atk-0__2.26.2-1.el8.x86_64",
    sha256 = "1d63dc4f754e976a73376541ecda8dfb6ed52d79b70f4307b6e465222f41b0cc",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/at-spi2-atk-2.26.2-1.el8.x86_64.rpm"],
)

rpm(
    name = "at-spi2-core-0__2.28.0-1.el8.x86_64",
    sha256 = "7db76e337e1b033da9452240607ac4270030584f6cb922e4f435f56e2fecc82a",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/at-spi2-core-2.28.0-1.el8.x86_64.rpm"],
)

rpm(
    name = "atk-0__2.28.1-1.el8.x86_64",
    sha256 = "e2e78b8bfed233002b701921ed6fb395cd51b625ee1460ee4bd7ab7b16c0e70a",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/atk-2.28.1-1.el8.x86_64.rpm"],
)

rpm(
    name = "audit-libs-0__3.0.7-4.el8.aarch64",
    sha256 = "2b05f70005d024a2b540a56afd9e05729c07c9dee120ff01100a21e21781f017",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/audit-libs-3.0.7-4.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/2b05f70005d024a2b540a56afd9e05729c07c9dee120ff01100a21e21781f017",
    ],
)

rpm(
    name = "audit-libs-0__3.0.7-4.el8.x86_64",
    sha256 = "b37099679b46f9a15d20b7c54fdd993388a8b84105f76869494c1be17140b512",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/audit-libs-3.0.7-4.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/b37099679b46f9a15d20b7c54fdd993388a8b84105f76869494c1be17140b512",
    ],
)

rpm(
    name = "augeas-libs-0__1.12.0-8.el8.x86_64",
    sha256 = "8d871c7339ed515b012497d8fe97bda5252649c14edfce27ade65ccd1edb16df",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/augeas-libs-1.12.0-8.el8.x86_64.rpm"],
)

rpm(
    name = "autogen-libopts-0__5.18.12-8.el8.aarch64",
    sha256 = "a69b87111415322e6586ba6b35494d77af7d9d58b2d9dfaf0360e4f827622dd2",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/autogen-libopts-5.18.12-8.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/a69b87111415322e6586ba6b35494d77af7d9d58b2d9dfaf0360e4f827622dd2",
    ],
)

rpm(
    name = "autogen-libopts-0__5.18.12-8.el8.x86_64",
    sha256 = "c73af033015bfbdbe8a43e162b098364d148517d394910f8db5d33b76b93aa48",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/autogen-libopts-5.18.12-8.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/c73af033015bfbdbe8a43e162b098364d148517d394910f8db5d33b76b93aa48",
    ],
)

rpm(
    name = "avahi-libs-0__0.7-20.el8.x86_64",
    sha256 = "53f3477bb6481ecf4e26ae156ec3edf90fd6d27546f54c10f9ba7a4d123c5275",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/avahi-libs-0.7-20.el8.x86_64.rpm"],
)

rpm(
    name = "basesystem-0__11-5.el8.aarch64",
    sha256 = "48226934763e4c412c1eb65df314e6879720b4b1ebcb3d07c126c9526639cb68",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/basesystem-11-5.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/48226934763e4c412c1eb65df314e6879720b4b1ebcb3d07c126c9526639cb68",
    ],
)

rpm(
    name = "basesystem-0__11-5.el8.x86_64",
    sha256 = "48226934763e4c412c1eb65df314e6879720b4b1ebcb3d07c126c9526639cb68",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/basesystem-11-5.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/48226934763e4c412c1eb65df314e6879720b4b1ebcb3d07c126c9526639cb68",
    ],
)

rpm(
    name = "bash-0__4.4.20-4.el8.aarch64",
    sha256 = "cb47111790ede91e0f1fb34817a27123a97e0304e7f7b6df06731fd391859f45",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/bash-4.4.20-4.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/cb47111790ede91e0f1fb34817a27123a97e0304e7f7b6df06731fd391859f45",
    ],
)

rpm(
    name = "bash-0__4.4.20-4.el8.x86_64",
    sha256 = "a104837b8aea5214122cf09c2de436db8f528812c1361c39f2d7471343dc509b",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/bash-4.4.20-4.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/a104837b8aea5214122cf09c2de436db8f528812c1361c39f2d7471343dc509b",
    ],
)

rpm(
    name = "binutils-0__2.30-117.el8.aarch64",
    sha256 = "10cc7e5ae3939eb78ef345127f05428eb003482c91dff1506121bde6228ed55f",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/binutils-2.30-117.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/10cc7e5ae3939eb78ef345127f05428eb003482c91dff1506121bde6228ed55f",
    ],
)

rpm(
    name = "binutils-0__2.30-119.el8.x86_64",
    sha256 = "e420b35cd541cf2fde38208b9a3d304b34b4b7adc051838d4700a6e4ce23f6ac",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/binutils-2.30-119.el8.x86_64.rpm"],
)

rpm(
    name = "bzip2-0__1.0.6-26.el8.aarch64",
    sha256 = "b18d9f23161d7d5de93fa72a56c645762deefbc0f3e5a095bb8d9e3cf09521e6",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/bzip2-1.0.6-26.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/b18d9f23161d7d5de93fa72a56c645762deefbc0f3e5a095bb8d9e3cf09521e6",
    ],
)

rpm(
    name = "bzip2-0__1.0.6-26.el8.x86_64",
    sha256 = "78596f457c3d737a97a4edfe9a03a01f593606379c281701ab7f7eba13ecaf18",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/bzip2-1.0.6-26.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/78596f457c3d737a97a4edfe9a03a01f593606379c281701ab7f7eba13ecaf18",
    ],
)

rpm(
    name = "bzip2-libs-0__1.0.6-26.el8.aarch64",
    sha256 = "a4451cae0e8a3307228ed8ac7dc9bab7de77fcbf2004141daa7f986f5dc9b381",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/bzip2-libs-1.0.6-26.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/a4451cae0e8a3307228ed8ac7dc9bab7de77fcbf2004141daa7f986f5dc9b381",
    ],
)

rpm(
    name = "bzip2-libs-0__1.0.6-26.el8.x86_64",
    sha256 = "19d66d152b745dbd49cea9d21c52aec0ec4d4321edef97a342acd3542404fa31",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/bzip2-libs-1.0.6-26.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/19d66d152b745dbd49cea9d21c52aec0ec4d4321edef97a342acd3542404fa31",
    ],
)

rpm(
    name = "ca-certificates-0__2021.2.50-82.el8.aarch64",
    sha256 = "1fad1d1f8b56e6967863aeb60f5fa3615e6a35b0f6532d8a23066e6823b50860",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/ca-certificates-2021.2.50-82.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/1fad1d1f8b56e6967863aeb60f5fa3615e6a35b0f6532d8a23066e6823b50860",
    ],
)

rpm(
    name = "ca-certificates-0__2022.2.54-80.2.el8.x86_64",
    sha256 = "3200d42d5585afa93a94600614a82b6e804139b06fff151576a53effd221e12b",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/ca-certificates-2022.2.54-80.2.el8.noarch.rpm"],
)

rpm(
    name = "cairo-0__1.15.12-6.el8.x86_64",
    sha256 = "8d94b1b954d06a5443c4e8036c1d51121a6724c1508f37539bbff96dbf806a92",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/cairo-1.15.12-6.el8.x86_64.rpm"],
)

rpm(
    name = "cairo-gobject-0__1.15.12-6.el8.x86_64",
    sha256 = "698694ee5d5fd001b7086373a16a97da4093e27fee169beb760d25b1dacb5d04",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/cairo-gobject-1.15.12-6.el8.x86_64.rpm"],
)

rpm(
    name = "centos-gpg-keys-1__8-6.el8.aarch64",
    sha256 = "567dd699e703dc6f5fa6ddb5548bf0dbd3bda08a0a6b1d10b32fa19012409cd0",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/centos-gpg-keys-8-6.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/567dd699e703dc6f5fa6ddb5548bf0dbd3bda08a0a6b1d10b32fa19012409cd0",
    ],
)

rpm(
    name = "centos-gpg-keys-1__8-6.el8.x86_64",
    sha256 = "567dd699e703dc6f5fa6ddb5548bf0dbd3bda08a0a6b1d10b32fa19012409cd0",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/centos-gpg-keys-8-6.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/567dd699e703dc6f5fa6ddb5548bf0dbd3bda08a0a6b1d10b32fa19012409cd0",
    ],
)

rpm(
    name = "centos-stream-release-0__8.6-1.el8.aarch64",
    sha256 = "3b3b86cb51f62632995ace850fbed9efc65381d639f1e1c5ceeff7ccf2dd6151",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/centos-stream-release-8.6-1.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/3b3b86cb51f62632995ace850fbed9efc65381d639f1e1c5ceeff7ccf2dd6151",
    ],
)

rpm(
    name = "centos-stream-release-0__8.6-1.el8.x86_64",
    sha256 = "3b3b86cb51f62632995ace850fbed9efc65381d639f1e1c5ceeff7ccf2dd6151",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/centos-stream-release-8.6-1.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/3b3b86cb51f62632995ace850fbed9efc65381d639f1e1c5ceeff7ccf2dd6151",
    ],
)

rpm(
    name = "centos-stream-repos-0__8-6.el8.aarch64",
    sha256 = "ff0a2d1fb5b00e9a26b05a82675d0dcdf0378ee5476f9ae765b32399c2ee561f",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/centos-stream-repos-8-6.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/ff0a2d1fb5b00e9a26b05a82675d0dcdf0378ee5476f9ae765b32399c2ee561f",
    ],
)

rpm(
    name = "centos-stream-repos-0__8-6.el8.x86_64",
    sha256 = "ff0a2d1fb5b00e9a26b05a82675d0dcdf0378ee5476f9ae765b32399c2ee561f",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/centos-stream-repos-8-6.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/ff0a2d1fb5b00e9a26b05a82675d0dcdf0378ee5476f9ae765b32399c2ee561f",
    ],
)

rpm(
    name = "chkconfig-0__1.19.1-1.el8.aarch64",
    sha256 = "be370bfc2f375cdbfc1079b19423142236770cf67caf74cdb12a7aef8a29c8c5",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/chkconfig-1.19.1-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/be370bfc2f375cdbfc1079b19423142236770cf67caf74cdb12a7aef8a29c8c5",
    ],
)

rpm(
    name = "chkconfig-0__1.19.1-1.el8.x86_64",
    sha256 = "561b5fdadd60370b5d0a91b7ed35df95d7f60650cbade8c7e744323982ac82db",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/chkconfig-1.19.1-1.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/561b5fdadd60370b5d0a91b7ed35df95d7f60650cbade8c7e744323982ac82db",
    ],
)

rpm(
    name = "colord-libs-0__1.4.2-1.el8.x86_64",
    sha256 = "96af3978c25911547aafb3031d1e144484a080ba5fbe960a65235e22beb4d2fc",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/colord-libs-1.4.2-1.el8.x86_64.rpm"],
)

rpm(
    name = "coreutils-single-0__8.30-13.el8.aarch64",
    sha256 = "0f560179f5b79ee62e0d71efb8d67f0d8eca9b31b752064a507c1052985e1251",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/coreutils-single-8.30-13.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/0f560179f5b79ee62e0d71efb8d67f0d8eca9b31b752064a507c1052985e1251",
    ],
)

rpm(
    name = "coreutils-single-0__8.30-14.el8.x86_64",
    sha256 = "b3179217ec738dead7ddd6b62c60cc83247a2464c834d5a60817ca3bb03d6aaa",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/coreutils-single-8.30-14.el8.x86_64.rpm"],
)

rpm(
    name = "cpp-0__8.5.0-15.el8.aarch64",
    sha256 = "36bb703e9305764b2075c56d79f98d4ff86a8a9dbcb59c2ce2a8eef37b4b98a2",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/cpp-8.5.0-15.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/36bb703e9305764b2075c56d79f98d4ff86a8a9dbcb59c2ce2a8eef37b4b98a2",
    ],
)

rpm(
    name = "cpp-0__8.5.0-17.el8.x86_64",
    sha256 = "200c5683d5629e42124fbaad472dd034ade7f84ab833b0f0d3305dea89404799",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/cpp-8.5.0-17.el8.x86_64.rpm"],
)

rpm(
    name = "cracklib-0__2.9.6-15.el8.aarch64",
    sha256 = "54efb853142572e1c2872e351838fc3657b662722ff6b2913d1872d4752a0eb8",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/cracklib-2.9.6-15.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/54efb853142572e1c2872e351838fc3657b662722ff6b2913d1872d4752a0eb8",
    ],
)

rpm(
    name = "cracklib-0__2.9.6-15.el8.x86_64",
    sha256 = "dbbc9e20caabc30070354d91f61f383081f6d658e09d3c09e6df8764559e5aca",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/cracklib-2.9.6-15.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/dbbc9e20caabc30070354d91f61f383081f6d658e09d3c09e6df8764559e5aca",
    ],
)

rpm(
    name = "cracklib-dicts-0__2.9.6-15.el8.aarch64",
    sha256 = "d61741af0ffe96c55f588dd164b9c3c93e7c7175c7e616db25990ab3e16e0f22",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/cracklib-dicts-2.9.6-15.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/d61741af0ffe96c55f588dd164b9c3c93e7c7175c7e616db25990ab3e16e0f22",
    ],
)

rpm(
    name = "cracklib-dicts-0__2.9.6-15.el8.x86_64",
    sha256 = "f1ce23ee43c747a35367dada19ca200a7758c50955ccc44aa946b86b647077ca",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/cracklib-dicts-2.9.6-15.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/f1ce23ee43c747a35367dada19ca200a7758c50955ccc44aa946b86b647077ca",
    ],
)

rpm(
    name = "crypto-policies-0__20211116-1.gitae470d6.el8.aarch64",
    sha256 = "8fb69892af346bacf18e8f8e7e8098e09c6ef9547abab9c39f7e729db06c3d1e",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/crypto-policies-20211116-1.gitae470d6.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/8fb69892af346bacf18e8f8e7e8098e09c6ef9547abab9c39f7e729db06c3d1e",
    ],
)

rpm(
    name = "crypto-policies-0__20221215-1.gitece0092.el8.x86_64",
    sha256 = "29d99b0985833aea0b2590036dcbb03e225877c30a18c707f2e149eaf5ba3697",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/crypto-policies-20221215-1.gitece0092.el8.noarch.rpm"],
)

rpm(
    name = "cryptsetup-libs-0__2.3.7-2.el8.aarch64",
    sha256 = "15a9d91ba7f5c192bee3e0d511e9b501c109a53c68120987e3f79ed88b1f69b5",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/cryptsetup-libs-2.3.7-2.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/15a9d91ba7f5c192bee3e0d511e9b501c109a53c68120987e3f79ed88b1f69b5",
    ],
)

rpm(
    name = "cryptsetup-libs-0__2.3.7-3.el8.x86_64",
    sha256 = "a010a71279e625f8c86d1a94f610e5ff5d8031a6b45a55cdfea61e110abb0f4b",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/cryptsetup-libs-2.3.7-3.el8.x86_64.rpm"],
)

rpm(
    name = "cups-libs-1__2.2.6-50.el8.x86_64",
    sha256 = "ea3e76d52c039a789dc825090c69984261d81d094a8bcb5f4d9fe488387171cd",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/cups-libs-2.2.6-50.el8.x86_64.rpm"],
)

rpm(
    name = "curl-0__7.61.1-25.el8.aarch64",
    sha256 = "56d7d77a32456f4c6b84ae4c6251d7ddfe2fb7097f9ecf8ba5e5834f7b7611c7",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/curl-7.61.1-25.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/56d7d77a32456f4c6b84ae4c6251d7ddfe2fb7097f9ecf8ba5e5834f7b7611c7",
    ],
)

rpm(
    name = "curl-0__7.61.1-27.el8.x86_64",
    sha256 = "393ccd1d727fd7513e820435069800b01f2281df3a33a4c75ede6ba49d774ebf",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/curl-7.61.1-27.el8.x86_64.rpm"],
)

rpm(
    name = "cyrus-sasl-0__2.1.27-6.el8_5.aarch64",
    sha256 = "e7acd635ac3d42260807c3fd6eab8713e3177b88bceadd79fe10d0719bfbff00",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/cyrus-sasl-2.1.27-6.el8_5.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/e7acd635ac3d42260807c3fd6eab8713e3177b88bceadd79fe10d0719bfbff00",
    ],
)

rpm(
    name = "cyrus-sasl-0__2.1.27-6.el8_5.x86_64",
    sha256 = "65a62affe9c99e597aabf117b8439a363761686c496723bc492dbfdcb6f60692",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/cyrus-sasl-2.1.27-6.el8_5.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/65a62affe9c99e597aabf117b8439a363761686c496723bc492dbfdcb6f60692",
    ],
)

rpm(
    name = "cyrus-sasl-gssapi-0__2.1.27-6.el8_5.aarch64",
    sha256 = "9fac42ea86802ebaf480d7373155a019d0a85dfd8093189d17194334af466a15",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/cyrus-sasl-gssapi-2.1.27-6.el8_5.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/9fac42ea86802ebaf480d7373155a019d0a85dfd8093189d17194334af466a15",
    ],
)

rpm(
    name = "cyrus-sasl-gssapi-0__2.1.27-6.el8_5.x86_64",
    sha256 = "6c9a8d9adc93d1be7db41fe7327c4dcce144cefad3008e580f5e9cadb6155eb4",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/cyrus-sasl-gssapi-2.1.27-6.el8_5.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/6c9a8d9adc93d1be7db41fe7327c4dcce144cefad3008e580f5e9cadb6155eb4",
    ],
)

rpm(
    name = "cyrus-sasl-lib-0__2.1.27-6.el8_5.aarch64",
    sha256 = "984998500ff0d60cb8756fee9eaeb82a001b7323b1130955770f2fa824f8a937",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/cyrus-sasl-lib-2.1.27-6.el8_5.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/984998500ff0d60cb8756fee9eaeb82a001b7323b1130955770f2fa824f8a937",
    ],
)

rpm(
    name = "cyrus-sasl-lib-0__2.1.27-6.el8_5.x86_64",
    sha256 = "5bd6e1201d8b10c6f01f500c43f63204f1d2ec8a4d8ce53c741e611c81ffb404",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/cyrus-sasl-lib-2.1.27-6.el8_5.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/5bd6e1201d8b10c6f01f500c43f63204f1d2ec8a4d8ce53c741e611c81ffb404",
    ],
)

rpm(
    name = "daxctl-libs-0__71.1-4.el8.x86_64",
    sha256 = "332af3c063fdb03d95632dc5010712c4e9ca7416f3049c901558c5aa0c6e445b",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/daxctl-libs-71.1-4.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/332af3c063fdb03d95632dc5010712c4e9ca7416f3049c901558c5aa0c6e445b",
    ],
)

rpm(
    name = "dbus-1__1.12.8-19.el8.aarch64",
    sha256 = "17a00a6a87fb07b3ff1541e42319bfa38a281a0f74b16010ae6e41037a0bcb53",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/dbus-1.12.8-19.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/17a00a6a87fb07b3ff1541e42319bfa38a281a0f74b16010ae6e41037a0bcb53",
    ],
)

rpm(
    name = "dbus-1__1.12.8-24.el8.x86_64",
    sha256 = "feba20c1a54cd905cba7ad79665814b084b71fd391f88458d36cc99a0e4786b9",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/dbus-1.12.8-24.el8.x86_64.rpm"],
)

rpm(
    name = "dbus-common-1__1.12.8-19.el8.aarch64",
    sha256 = "aac8490975c287223e920d58276f6a08c89f92743245a7f2bf31b702b17a82a9",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/dbus-common-1.12.8-19.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/aac8490975c287223e920d58276f6a08c89f92743245a7f2bf31b702b17a82a9",
    ],
)

rpm(
    name = "dbus-common-1__1.12.8-24.el8.x86_64",
    sha256 = "5fb132e3a6b3fcedbb13de4ef5004d8c1ee4722cd42f17712e69fbdc1ae70572",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/dbus-common-1.12.8-24.el8.noarch.rpm"],
)

rpm(
    name = "dbus-daemon-1__1.12.8-19.el8.aarch64",
    sha256 = "b4d0b8eb5f65f9739109050a84feda6eea99099556dbf946d3d5be8ae801fb17",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/dbus-daemon-1.12.8-19.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/b4d0b8eb5f65f9739109050a84feda6eea99099556dbf946d3d5be8ae801fb17",
    ],
)

rpm(
    name = "dbus-daemon-1__1.12.8-24.el8.x86_64",
    sha256 = "6b5611899424c5382d9917d74148473535e0e7b9dc7ef8dd74e410b28b5d9342",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/dbus-daemon-1.12.8-24.el8.x86_64.rpm"],
)

rpm(
    name = "dbus-libs-1__1.12.8-19.el8.aarch64",
    sha256 = "6e59ed0aae536a2d8af9279fb9b263bb1e60c3501acd395584b67dc2d0bf668d",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/dbus-libs-1.12.8-19.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/6e59ed0aae536a2d8af9279fb9b263bb1e60c3501acd395584b67dc2d0bf668d",
    ],
)

rpm(
    name = "dbus-libs-1__1.12.8-24.el8.x86_64",
    sha256 = "4687b9ae45e0bb542c76694db9473c21e88961abc47237156cd9147eaf524be7",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/dbus-libs-1.12.8-24.el8.x86_64.rpm"],
)

rpm(
    name = "dbus-tools-1__1.12.8-19.el8.aarch64",
    sha256 = "90634aa2ac243dbed52d88f20a04afcf8b2298a175d8eb115fedf04fe77af21f",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/dbus-tools-1.12.8-19.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/90634aa2ac243dbed52d88f20a04afcf8b2298a175d8eb115fedf04fe77af21f",
    ],
)

rpm(
    name = "dbus-tools-1__1.12.8-24.el8.x86_64",
    sha256 = "a35c85304f8c360779b7488dcc687a95f24a71327de6f33db758f418e0b491b6",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/dbus-tools-1.12.8-24.el8.x86_64.rpm"],
)

rpm(
    name = "device-mapper-8__1.02.181-5.el8.aarch64",
    sha256 = "b9b3a6977b2a26674f6b83608e695b0378e3c78d9244e3a911024beacdcc0e21",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/device-mapper-1.02.181-5.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/b9b3a6977b2a26674f6b83608e695b0378e3c78d9244e3a911024beacdcc0e21",
    ],
)

rpm(
    name = "device-mapper-8__1.02.181-6.el8.x86_64",
    sha256 = "8e89a7c9e0b011917c8c360e625c3a2bfa3a81b82e9ac961977aa09a98b9da27",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/device-mapper-1.02.181-6.el8.x86_64.rpm"],
)

rpm(
    name = "device-mapper-event-8__1.02.181-6.el8.x86_64",
    sha256 = "a89244d4420b679d3d567e09578e68ff4159b7eec3968a3ecc39a9c0521c5ec3",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/device-mapper-event-1.02.181-6.el8.x86_64.rpm"],
)

rpm(
    name = "device-mapper-event-libs-8__1.02.181-6.el8.x86_64",
    sha256 = "e759885e81245b164bf418caf38f2358eebe56f68094d848f915ba957cf04a47",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/device-mapper-event-libs-1.02.181-6.el8.x86_64.rpm"],
)

rpm(
    name = "device-mapper-libs-8__1.02.181-5.el8.aarch64",
    sha256 = "2d2d686fae3a652633ca1c23cbd643553f29a93d4b3b3b4cbce23dd96c03faf9",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/device-mapper-libs-1.02.181-5.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/2d2d686fae3a652633ca1c23cbd643553f29a93d4b3b3b4cbce23dd96c03faf9",
    ],
)

rpm(
    name = "device-mapper-libs-8__1.02.181-6.el8.x86_64",
    sha256 = "bdd83ee5a458034908007edae7c55aa7ebc1138f0306877356f9d2f1215dd065",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/device-mapper-libs-1.02.181-6.el8.x86_64.rpm"],
)

rpm(
    name = "device-mapper-multipath-libs-0__0.8.4-26.el8.aarch64",
    sha256 = "cca401deac042f41d5a76cdfd0ef26080848300cb98a767e5a307a3ac7cd303b",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/device-mapper-multipath-libs-0.8.4-26.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/cca401deac042f41d5a76cdfd0ef26080848300cb98a767e5a307a3ac7cd303b",
    ],
)

rpm(
    name = "device-mapper-multipath-libs-0__0.8.4-34.el8.x86_64",
    sha256 = "f4b4bb1ed8a724f4468b7640aae81dd474fb276d0d4fb69877216f2d3ede4761",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/device-mapper-multipath-libs-0.8.4-34.el8.x86_64.rpm"],
)

rpm(
    name = "device-mapper-persistent-data-0__0.9.0-7.el8.x86_64",
    sha256 = "609c2bf12ce2994a0753177e334cde294a96750903c24d8583e7a0674c80485e",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/device-mapper-persistent-data-0.9.0-7.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/609c2bf12ce2994a0753177e334cde294a96750903c24d8583e7a0674c80485e",
    ],
)

rpm(
    name = "diffutils-0__3.6-6.el8.aarch64",
    sha256 = "8cbebc0fa970ceca4f479ee292eaad155084987be2cf7f97bbafe4a529319c98",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/diffutils-3.6-6.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/8cbebc0fa970ceca4f479ee292eaad155084987be2cf7f97bbafe4a529319c98",
    ],
)

rpm(
    name = "diffutils-0__3.6-6.el8.x86_64",
    sha256 = "c515d78c64a93d8b469593bff5800eccd50f24b16697ab13bdce81238c38eb77",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/diffutils-3.6-6.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/c515d78c64a93d8b469593bff5800eccd50f24b16697ab13bdce81238c38eb77",
    ],
)

rpm(
    name = "dmidecode-1__3.3-4.el8.x86_64",
    sha256 = "c1347fe2d5621a249ea230e9e8ff2774e538031070a225245154a75428ec67a5",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/dmidecode-3.3-4.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/c1347fe2d5621a249ea230e9e8ff2774e538031070a225245154a75428ec67a5",
    ],
)

rpm(
    name = "e2fsprogs-0__1.45.6-5.el8.aarch64",
    sha256 = "b916de2e7ea8fc3b0b381e0afe4353ab401b82885cea5afec0551232beb30fe2",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/e2fsprogs-1.45.6-5.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/b916de2e7ea8fc3b0b381e0afe4353ab401b82885cea5afec0551232beb30fe2",
    ],
)

rpm(
    name = "e2fsprogs-0__1.45.6-5.el8.x86_64",
    sha256 = "baa1ec089da85bf196f6e1e135727bb540f27ee7fe39d08bb17b712e59f4db8a",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/e2fsprogs-1.45.6-5.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/baa1ec089da85bf196f6e1e135727bb540f27ee7fe39d08bb17b712e59f4db8a",
    ],
)

rpm(
    name = "e2fsprogs-libs-0__1.45.6-5.el8.aarch64",
    sha256 = "0ec196d820abc43432cfa52c887c880b27b63619c6785dc30daed0e091c5bb76",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/e2fsprogs-libs-1.45.6-5.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/0ec196d820abc43432cfa52c887c880b27b63619c6785dc30daed0e091c5bb76",
    ],
)

rpm(
    name = "e2fsprogs-libs-0__1.45.6-5.el8.x86_64",
    sha256 = "035c5ed68339e632907c3f952098cdc9181ab9138239473903000e6a50446d98",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/e2fsprogs-libs-1.45.6-5.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/035c5ed68339e632907c3f952098cdc9181ab9138239473903000e6a50446d98",
    ],
)

rpm(
    name = "edk2-aarch64-0__20220126gitbb1bba3d77-2.el8.aarch64",
    sha256 = "0985ef697fbe90b66dbb0f70bfb4d0022f97255a36479e8d9ae4dd0489afd01a",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/edk2-aarch64-20220126gitbb1bba3d77-2.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/0985ef697fbe90b66dbb0f70bfb4d0022f97255a36479e8d9ae4dd0489afd01a",
    ],
)

rpm(
    name = "edk2-ovmf-0__20220126gitbb1bba3d77-2.el8.x86_64",
    sha256 = "a360d8e0ac13460ebab244e3063d6a9e2fb4d3a6bc2eb501534e5bfe9d0cff1e",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/edk2-ovmf-20220126gitbb1bba3d77-2.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/a360d8e0ac13460ebab244e3063d6a9e2fb4d3a6bc2eb501534e5bfe9d0cff1e",
    ],
)

rpm(
    name = "elfutils-default-yama-scope-0__0.187-4.el8.aarch64",
    sha256 = "3c89377bb7409293f0dc8ada62071fe2e3cf042ae2b5ca7cf09faf77394b5187",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/elfutils-default-yama-scope-0.187-4.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/3c89377bb7409293f0dc8ada62071fe2e3cf042ae2b5ca7cf09faf77394b5187",
    ],
)

rpm(
    name = "elfutils-default-yama-scope-0__0.188-3.el8.x86_64",
    sha256 = "fa1c01e489744a0bc3127d7996b9a2527347ace9c97c04d146c2331fd0acb926",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/elfutils-default-yama-scope-0.188-3.el8.noarch.rpm"],
)

rpm(
    name = "elfutils-libelf-0__0.187-4.el8.aarch64",
    sha256 = "bfdfc37f2dd1052d4067937724a6ef6a9858a9c1b3c1aacf1e9085a83e99e1b4",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/elfutils-libelf-0.187-4.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/bfdfc37f2dd1052d4067937724a6ef6a9858a9c1b3c1aacf1e9085a83e99e1b4",
    ],
)

rpm(
    name = "elfutils-libelf-0__0.188-3.el8.x86_64",
    sha256 = "746cb30b5c69ddfe1c525b165a036866fbbb38091d9e1565c9b1a3c4fa48f74c",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/elfutils-libelf-0.188-3.el8.x86_64.rpm"],
)

rpm(
    name = "elfutils-libs-0__0.187-4.el8.aarch64",
    sha256 = "682c1b9f11d68cdec87ea746ea0d5861f3afcf2159aa732854625bfa180bbaee",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/elfutils-libs-0.187-4.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/682c1b9f11d68cdec87ea746ea0d5861f3afcf2159aa732854625bfa180bbaee",
    ],
)

rpm(
    name = "elfutils-libs-0__0.188-3.el8.x86_64",
    sha256 = "1ece046828213af0a1fadaf44cfc456441e3ce2440a1ee32ed22a640b7f87510",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/elfutils-libs-0.188-3.el8.x86_64.rpm"],
)

rpm(
    name = "ethtool-2__5.13-2.el8.aarch64",
    sha256 = "5bdb69b9c4161ba3d4846082686ee8edce640b7c6ff0bbf1c1eae12084661c24",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/ethtool-5.13-2.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/5bdb69b9c4161ba3d4846082686ee8edce640b7c6ff0bbf1c1eae12084661c24",
    ],
)

rpm(
    name = "ethtool-2__5.13-2.el8.x86_64",
    sha256 = "f1af67b33961ddf98360e5ce855910d2dee534bffe953068f27ad96b846a2fb7",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/ethtool-5.13-2.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/f1af67b33961ddf98360e5ce855910d2dee534bffe953068f27ad96b846a2fb7",
    ],
)

rpm(
    name = "expat-0__2.2.5-10.el8.1.x86_64",
    sha256 = "83b28b1ab0999914cbe8913a62b6c16deca4d989ee5fd0579546d6b6cf070c4a",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/expat-2.2.5-10.el8.1.x86_64.rpm"],
)

rpm(
    name = "expat-0__2.2.5-9.el8.aarch64",
    sha256 = "4ca97fb015687a8f2ac442f581d1c42154662b4336e0f34c71be2659cb716fc8",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/expat-2.2.5-9.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/4ca97fb015687a8f2ac442f581d1c42154662b4336e0f34c71be2659cb716fc8",
    ],
)

rpm(
    name = "file-0__5.33-21.el8.x86_64",
    sha256 = "202e8164df8a6110d58692fa25eaf1d1078a988372943ae73536333237dc3818",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/file-5.33-21.el8.x86_64.rpm"],
)

rpm(
    name = "file-libs-0__5.33-21.el8.x86_64",
    sha256 = "9a51006d0e557e456eb9fc03ff7ed236633d32823dbd46984aca96f379e09f21",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/file-libs-5.33-21.el8.x86_64.rpm"],
)

rpm(
    name = "filesystem-0__3.8-6.el8.aarch64",
    sha256 = "e6c3fa94860eda0bc2ae6b1b78acd1159cbed355a03e7bec8b3defa1d90782b6",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/filesystem-3.8-6.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/e6c3fa94860eda0bc2ae6b1b78acd1159cbed355a03e7bec8b3defa1d90782b6",
    ],
)

rpm(
    name = "filesystem-0__3.8-6.el8.x86_64",
    sha256 = "50bdb81d578914e0e88fe6b13550b4c30aac4d72f064fdcd78523df7dd2f64da",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/filesystem-3.8-6.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/50bdb81d578914e0e88fe6b13550b4c30aac4d72f064fdcd78523df7dd2f64da",
    ],
)

rpm(
    name = "findutils-1__4.6.0-20.el8.aarch64",
    sha256 = "985479064966d05aa82010ed5b8905942e47e2bebb919c9c1bd004a28addad1d",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/findutils-4.6.0-20.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/985479064966d05aa82010ed5b8905942e47e2bebb919c9c1bd004a28addad1d",
    ],
)

rpm(
    name = "findutils-1__4.6.0-20.el8.x86_64",
    sha256 = "811eb112646b7d87773c65af47efdca975468f3e5df44aa9944e30de24d83890",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/findutils-4.6.0-20.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/811eb112646b7d87773c65af47efdca975468f3e5df44aa9944e30de24d83890",
    ],
)

rpm(
    name = "fontconfig-0__2.13.1-4.el8.x86_64",
    sha256 = "1d2c61493d72419e85272e8cbc1a0bf3232c81b9bed4707d68f2bbeef2391a55",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/fontconfig-2.13.1-4.el8.x86_64.rpm"],
)

rpm(
    name = "fontpackages-filesystem-0__1.44-22.el8.x86_64",
    sha256 = "700b9050aa490b5eca6d1f8630cbebceb122fce11c370689b5ccb37f5a43ee2f",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/fontpackages-filesystem-1.44-22.el8.noarch.rpm"],
)

rpm(
    name = "freetype-0__2.9.1-9.el8.x86_64",
    sha256 = "0097dc947c987310bb5cbcb9976594eac1e1d111e065ffee150abc2d69b8d709",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/freetype-2.9.1-9.el8.x86_64.rpm"],
)

rpm(
    name = "fribidi-0__1.0.4-9.el8.x86_64",
    sha256 = "6540f56f1d5f191d91e8445d7156bf7fae954c18f07c7191bd0cb0ef38455e05",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/fribidi-1.0.4-9.el8.x86_64.rpm"],
)

rpm(
    name = "fuse-0__2.9.7-16.el8.x86_64",
    sha256 = "c208aa2f2f216a2172b1d9fa82bcad1b201e62f9a3101f4d52fb3de54ed28596",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/fuse-2.9.7-16.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/c208aa2f2f216a2172b1d9fa82bcad1b201e62f9a3101f4d52fb3de54ed28596",
    ],
)

rpm(
    name = "fuse-common-0__3.3.0-16.el8.x86_64",
    sha256 = "d637dfd117080f52f1a60444b6c09aaf65a535844cacce05945d1d691b8d7043",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/fuse-common-3.3.0-16.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/d637dfd117080f52f1a60444b6c09aaf65a535844cacce05945d1d691b8d7043",
    ],
)

rpm(
    name = "fuse-libs-0__2.9.7-16.el8.aarch64",
    sha256 = "6970abceb1e040a2a37a13faeaf2a4204c79a57d5bc8273ed276b385be813afb",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/fuse-libs-2.9.7-16.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/6970abceb1e040a2a37a13faeaf2a4204c79a57d5bc8273ed276b385be813afb",
    ],
)

rpm(
    name = "fuse-libs-0__2.9.7-16.el8.x86_64",
    sha256 = "77fff0f92a55307b7df2334bc9cc2998c024586abd96286a251919b0509f0473",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/fuse-libs-2.9.7-16.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/77fff0f92a55307b7df2334bc9cc2998c024586abd96286a251919b0509f0473",
    ],
)

rpm(
    name = "gawk-0__4.2.1-4.el8.aarch64",
    sha256 = "75594a09076ad901d5afb1027c74aae945f77e0e357e7d4f46148cbcbd1d0ae4",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/gawk-4.2.1-4.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/75594a09076ad901d5afb1027c74aae945f77e0e357e7d4f46148cbcbd1d0ae4",
    ],
)

rpm(
    name = "gawk-0__4.2.1-4.el8.x86_64",
    sha256 = "ff4438c2dff5bf933d7874fd55f131ca6ee067f8fb4324c89719d63e60b40aba",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/gawk-4.2.1-4.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/ff4438c2dff5bf933d7874fd55f131ca6ee067f8fb4324c89719d63e60b40aba",
    ],
)

rpm(
    name = "gcc-0__8.5.0-15.el8.aarch64",
    sha256 = "347dbe82b51689eda62164b0ffdabb2dadf26f170c7430c32936d3ee87a67693",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/gcc-8.5.0-15.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/347dbe82b51689eda62164b0ffdabb2dadf26f170c7430c32936d3ee87a67693",
    ],
)

rpm(
    name = "gcc-0__8.5.0-17.el8.x86_64",
    sha256 = "9a48117583d7c9a4bf29dc8fba41757b49efe152cf262e3d51088c43454387e7",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/gcc-8.5.0-17.el8.x86_64.rpm"],
)

rpm(
    name = "gdbm-1__1.18-2.el8.aarch64",
    sha256 = "c032e3863180bb2247ddc0e02cd54be72099137af21452e2dc25ddd03f9a5395",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/gdbm-1.18-2.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/c032e3863180bb2247ddc0e02cd54be72099137af21452e2dc25ddd03f9a5395",
    ],
)

rpm(
    name = "gdbm-1__1.18-2.el8.x86_64",
    sha256 = "fa1751b26519b9637cf3f0a25ea1874eb2df005dde1e1371a3f13d0c9a38b9ca",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/gdbm-1.18-2.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/fa1751b26519b9637cf3f0a25ea1874eb2df005dde1e1371a3f13d0c9a38b9ca",
    ],
)

rpm(
    name = "gdbm-libs-1__1.18-2.el8.aarch64",
    sha256 = "bdb64aec2a4ea8a2c70652cd57e5f88353079042402e7662e0e89934d3737562",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/gdbm-libs-1.18-2.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/bdb64aec2a4ea8a2c70652cd57e5f88353079042402e7662e0e89934d3737562",
    ],
)

rpm(
    name = "gdbm-libs-1__1.18-2.el8.x86_64",
    sha256 = "eddcea96342c8cfaa60b79fc2c66cb8c5b0038c3b11855abe55e659b2cad6199",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/gdbm-libs-1.18-2.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/eddcea96342c8cfaa60b79fc2c66cb8c5b0038c3b11855abe55e659b2cad6199",
    ],
)

rpm(
    name = "gdk-pixbuf2-0__2.36.12-5.el8.x86_64",
    sha256 = "94cb8dceb47a5b01e3c0542ea3b48601d720325da28e6e6d89ae529e4fddcd97",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/gdk-pixbuf2-2.36.12-5.el8.x86_64.rpm"],
)

rpm(
    name = "gdk-pixbuf2-modules-0__2.36.12-5.el8.x86_64",
    sha256 = "f689bcf0759fcf467ef007f8a640e2829c05879a0a1ff7d3a3fd36130230f2c9",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/gdk-pixbuf2-modules-2.36.12-5.el8.x86_64.rpm"],
)

rpm(
    name = "gettext-0__0.19.8.1-17.el8.aarch64",
    sha256 = "5f0c37488d3017b052039ddb8d9189a38c252af97884264959334237109c7e7c",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/gettext-0.19.8.1-17.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/5f0c37488d3017b052039ddb8d9189a38c252af97884264959334237109c7e7c",
    ],
)

rpm(
    name = "gettext-0__0.19.8.1-17.el8.x86_64",
    sha256 = "829c842bbd79dca18d37198414626894c44e5b8faf0cce0054ca0ba6623ae136",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/gettext-0.19.8.1-17.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/829c842bbd79dca18d37198414626894c44e5b8faf0cce0054ca0ba6623ae136",
    ],
)

rpm(
    name = "gettext-libs-0__0.19.8.1-17.el8.aarch64",
    sha256 = "882f23e0250a2d4aea49abb4ec8e11a9a3869ccdd812c796b6f85341ff9d30a2",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/gettext-libs-0.19.8.1-17.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/882f23e0250a2d4aea49abb4ec8e11a9a3869ccdd812c796b6f85341ff9d30a2",
    ],
)

rpm(
    name = "gettext-libs-0__0.19.8.1-17.el8.x86_64",
    sha256 = "ade52756aaf236e77dadd6cf97716821141c2759129ca7808524ab79607bb4c4",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/gettext-libs-0.19.8.1-17.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/ade52756aaf236e77dadd6cf97716821141c2759129ca7808524ab79607bb4c4",
    ],
)

rpm(
    name = "glib-networking-0__2.56.1-1.1.el8.x86_64",
    sha256 = "a7f9ae54f45ca4fcecf78d9885d12a789f7325119794178bfa2814c6185a953d",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/glib-networking-2.56.1-1.1.el8.x86_64.rpm"],
)

rpm(
    name = "glib2-0__2.56.4-159.el8.aarch64",
    sha256 = "daac37a432b09faa6dd1e330c3595f6a70c53bff23a71fbce8df33c72e9fde24",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/glib2-2.56.4-159.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/daac37a432b09faa6dd1e330c3595f6a70c53bff23a71fbce8df33c72e9fde24",
    ],
)

rpm(
    name = "glib2-0__2.56.4-160.el8.x86_64",
    sha256 = "143ad1fc5d281408bab3e2a29650b923d82f3008c8cf9101741b179c2c8f7565",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/glib2-2.56.4-160.el8.x86_64.rpm"],
)

rpm(
    name = "glibc-0__2.28-208.el8.aarch64",
    sha256 = "4e03038e95b2c9b380b2767b1f0144eeb596aff00a431e325fc3534b80a7a0a1",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/glibc-2.28-208.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/4e03038e95b2c9b380b2767b1f0144eeb596aff00a431e325fc3534b80a7a0a1",
    ],
)

rpm(
    name = "glibc-0__2.28-224.el8.x86_64",
    sha256 = "d435b2974794c7acd6b263676ae5f80fdf9ffaee30ba92f86eb0f0dbc07740db",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/glibc-2.28-224.el8.x86_64.rpm"],
)

rpm(
    name = "glibc-common-0__2.28-208.el8.aarch64",
    sha256 = "f4ce83dc2efac25d1e30c1953a1876a3c5f50fc9a4a7f58a8da13ec99d40243b",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/glibc-common-2.28-208.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/f4ce83dc2efac25d1e30c1953a1876a3c5f50fc9a4a7f58a8da13ec99d40243b",
    ],
)

rpm(
    name = "glibc-common-0__2.28-224.el8.x86_64",
    sha256 = "21c3069e5de0ffa8800b2e03112079582025c121a6da2fe66c63595db1f4e63b",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/glibc-common-2.28-224.el8.x86_64.rpm"],
)

rpm(
    name = "glibc-devel-0__2.28-208.el8.aarch64",
    sha256 = "2faa2af7573fd93eec6128f13ce1e304113a4de34966f4a2dc329f9d1997594a",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/glibc-devel-2.28-208.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/2faa2af7573fd93eec6128f13ce1e304113a4de34966f4a2dc329f9d1997594a",
    ],
)

rpm(
    name = "glibc-devel-0__2.28-224.el8.x86_64",
    sha256 = "56d7065e10a42be59be01f87434424cde755f9d2ce5dc60e2d141a37bc5a4b95",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/glibc-devel-2.28-224.el8.x86_64.rpm"],
)

rpm(
    name = "glibc-headers-0__2.28-208.el8.aarch64",
    sha256 = "2e212bdea807cdb84d2fee6114630eb579b3f9eb4f59e655fae1c78f0ed7f593",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/glibc-headers-2.28-208.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/2e212bdea807cdb84d2fee6114630eb579b3f9eb4f59e655fae1c78f0ed7f593",
    ],
)

rpm(
    name = "glibc-headers-0__2.28-224.el8.x86_64",
    sha256 = "43689ece9edd84b0ce6a7aee837fda99fed8412a508774c8cc22ec5e768ec137",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/glibc-headers-2.28-224.el8.x86_64.rpm"],
)

rpm(
    name = "glibc-minimal-langpack-0__2.28-208.el8.aarch64",
    sha256 = "0f9ce08544295951b1bff828ddf1296d2608de1a0d784f83910e2205ebe8faea",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/glibc-minimal-langpack-2.28-208.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/0f9ce08544295951b1bff828ddf1296d2608de1a0d784f83910e2205ebe8faea",
    ],
)

rpm(
    name = "glibc-minimal-langpack-0__2.28-224.el8.x86_64",
    sha256 = "fb700012e70b7d8597d2a3941c52878693041caef3849f7351c722967bf8f019",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/glibc-minimal-langpack-2.28-224.el8.x86_64.rpm"],
)

rpm(
    name = "glibc-static-0__2.28-208.el8.aarch64",
    sha256 = "6ef41b3231619ef931b19fd85d39791388e144955e376443b649ec402ee7c5e6",
    urls = [
        "http://mirror.centos.org/centos/8-stream/PowerTools/aarch64/os/Packages/glibc-static-2.28-208.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/6ef41b3231619ef931b19fd85d39791388e144955e376443b649ec402ee7c5e6",
    ],
)

rpm(
    name = "glibc-static-0__2.28-224.el8.x86_64",
    sha256 = "c423443fdd4b6b5c7e6a0a548906bbe26904d83c5131e48e7a185803c684df41",
    urls = ["http://mirror.centos.org/centos/8-stream/PowerTools/x86_64/os/Packages/glibc-static-2.28-224.el8.x86_64.rpm"],
)

rpm(
    name = "gmp-1__6.1.2-10.el8.aarch64",
    sha256 = "8d407f8ad961169fca2ee5e22e824cbc2d2b5fedca9701896cc492d4cb788603",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/gmp-6.1.2-10.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/8d407f8ad961169fca2ee5e22e824cbc2d2b5fedca9701896cc492d4cb788603",
    ],
)

rpm(
    name = "gmp-1__6.1.2-10.el8.x86_64",
    sha256 = "3b96e2c7d5cd4b49bfde8e52c8af6ff595c91438e50856e468f14a049d8511e2",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/gmp-6.1.2-10.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/3b96e2c7d5cd4b49bfde8e52c8af6ff595c91438e50856e468f14a049d8511e2",
    ],
)

rpm(
    name = "gnupg2-0__2.2.20-3.el8.x86_64",
    sha256 = "8c44c980dd9a6a42ccb93578d7e6e1940d36d2da0a5a99d783189c43b2ad6d5f",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/gnupg2-2.2.20-3.el8.x86_64.rpm"],
)

rpm(
    name = "gnutls-0__3.6.16-4.el8.aarch64",
    sha256 = "f97d55f7bdf6fe126e7a1446563af7ee4c1bb7ee3a2a9b12b6df1cdd344da47e",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/gnutls-3.6.16-4.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/f97d55f7bdf6fe126e7a1446563af7ee4c1bb7ee3a2a9b12b6df1cdd344da47e",
    ],
)

rpm(
    name = "gnutls-0__3.6.16-5.el8.x86_64",
    sha256 = "4bd8fc9616f01f02cf1b17cccf4ae4d072f5adbd0c159b04203c87e8fb74b013",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/gnutls-3.6.16-5.el8.x86_64.rpm"],
)

rpm(
    name = "gnutls-dane-0__3.6.16-4.el8.aarch64",
    sha256 = "df78e84002d6ba09e37901b2f85f462a160beda734e98876c8baba0c71caf638",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/gnutls-dane-3.6.16-4.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/df78e84002d6ba09e37901b2f85f462a160beda734e98876c8baba0c71caf638",
    ],
)

rpm(
    name = "gnutls-dane-0__3.6.16-5.el8.x86_64",
    sha256 = "bb65a2fb02d9d77c983ae1ecd2a64b211c96804d25fdc8e5b6575a8a19d8c59e",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/gnutls-dane-3.6.16-5.el8.x86_64.rpm"],
)

rpm(
    name = "gnutls-utils-0__3.6.16-4.el8.aarch64",
    sha256 = "1421e7f87f559b398b9bd289ee10c79b38a0505613761b4499ad9747aafb7da6",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/gnutls-utils-3.6.16-4.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/1421e7f87f559b398b9bd289ee10c79b38a0505613761b4499ad9747aafb7da6",
    ],
)

rpm(
    name = "gnutls-utils-0__3.6.16-5.el8.x86_64",
    sha256 = "fc7abd04a01d77c7f0207b4ffd1edbd9a5ebdfb2e5154351abb481a11fdaf534",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/gnutls-utils-3.6.16-5.el8.x86_64.rpm"],
)

rpm(
    name = "gobject-introspection-0__1.56.1-1.el8.x86_64",
    sha256 = "7e2804a4494d4179ed50c0c99da1e30c3a6abf8db889c1412b458943cff0e3e5",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/gobject-introspection-1.56.1-1.el8.x86_64.rpm"],
)

rpm(
    name = "google-droid-sans-mono-fonts-0__20120715-13.el8.x86_64",
    sha256 = "6c8ecdbfbd3115ee563584b1a15b3ce4f9fbf1dc5efb145699256778bfb32a49",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/google-droid-sans-mono-fonts-20120715-13.el8.noarch.rpm"],
)

rpm(
    name = "google-noto-fonts-common-0__20161022-7.el8.1.x86_64",
    sha256 = "d90f65b0b7c294e6114387dfcc06e60fe2c8a473f6df691bc468fc909ae2b2de",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/google-noto-fonts-common-20161022-7.el8.1.noarch.rpm"],
)

rpm(
    name = "google-noto-serif-fonts-0__20161022-7.el8.1.x86_64",
    sha256 = "60cce46a0ed5397386cd6fd2b8442b57d4b3f172517181858271dd7d2e9d9672",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/google-noto-serif-fonts-20161022-7.el8.1.noarch.rpm"],
)

rpm(
    name = "graphite2-0__1.3.10-10.el8.x86_64",
    sha256 = "0f9c3ee5f54ed296f99219bd70fa4f869c4c9986e3766d813a76a0cc5ecee24e",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/graphite2-1.3.10-10.el8.x86_64.rpm"],
)

rpm(
    name = "grep-0__3.1-6.el8.aarch64",
    sha256 = "7ffd6e95b0554466e97346b2f41fb5279aedcb29ae07828f63d06a8dedd7cd51",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/grep-3.1-6.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/7ffd6e95b0554466e97346b2f41fb5279aedcb29ae07828f63d06a8dedd7cd51",
    ],
)

rpm(
    name = "grep-0__3.1-6.el8.x86_64",
    sha256 = "3f8ffe48bb481a5db7cbe42bf73b839d872351811e5df41b2f6697c61a030487",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/grep-3.1-6.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/3f8ffe48bb481a5db7cbe42bf73b839d872351811e5df41b2f6697c61a030487",
    ],
)

rpm(
    name = "groff-base-0__1.22.3-18.el8.x86_64",
    sha256 = "b00855013100d3796e9ed6d82b1ab2d4dc7f4a3a3fa2e186f6de8523577974a0",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/groff-base-1.22.3-18.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/b00855013100d3796e9ed6d82b1ab2d4dc7f4a3a3fa2e186f6de8523577974a0",
    ],
)

rpm(
    name = "gsettings-desktop-schemas-0__3.32.0-6.el8.x86_64",
    sha256 = "4f05013bb8d2d2173d83dc667cafe942bdd0299fb21cb6bebe0f306d92df1842",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/gsettings-desktop-schemas-3.32.0-6.el8.x86_64.rpm"],
)

rpm(
    name = "gtk-update-icon-cache-0__3.22.30-11.el8.x86_64",
    sha256 = "a19e319c6cc171bf9cebc136eb3044f323e74d85533ef9c9097def020896c96c",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/gtk-update-icon-cache-3.22.30-11.el8.x86_64.rpm"],
)

rpm(
    name = "gtk3-0__3.22.30-11.el8.x86_64",
    sha256 = "eb526290fa7222d257ead3b4edc247aac3c6aefcecc6e2a30754b7dfa5b7ed37",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/gtk3-3.22.30-11.el8.x86_64.rpm"],
)

rpm(
    name = "gzip-0__1.9-13.el8.aarch64",
    sha256 = "80ee79fb497c43c06d3c54bf432e6391c5ae19ae43241111f3be4113ea49fa96",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/gzip-1.9-13.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/80ee79fb497c43c06d3c54bf432e6391c5ae19ae43241111f3be4113ea49fa96",
    ],
)

rpm(
    name = "gzip-0__1.9-13.el8.x86_64",
    sha256 = "1cc189e4991fc6b3526f7eebc9f798b8922e70d60a12ba499b6e0329eb473cea",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/gzip-1.9-13.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/1cc189e4991fc6b3526f7eebc9f798b8922e70d60a12ba499b6e0329eb473cea",
    ],
)

rpm(
    name = "harfbuzz-0__1.7.5-3.el8.x86_64",
    sha256 = "49c652f3d967e944b9d0ad9dea63e8942626d3b9f40fde12cfb0d3e924a82053",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/harfbuzz-1.7.5-3.el8.x86_64.rpm"],
)

rpm(
    name = "hexedit-0__1.2.13-12.el8.x86_64",
    sha256 = "4538e44d3ebff3f9323b59171767bca2b7f5244dd90141de101856ad4f4643f5",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/hexedit-1.2.13-12.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/4538e44d3ebff3f9323b59171767bca2b7f5244dd90141de101856ad4f4643f5",
    ],
)

rpm(
    name = "hicolor-icon-theme-0__0.17-2.el8.x86_64",
    sha256 = "22cfaa4da6a1d6c8026dc96078c6cb681c9d1b17dc09116a87d5bd4641ab799f",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/hicolor-icon-theme-0.17-2.el8.noarch.rpm"],
)

rpm(
    name = "hivex-0__1.3.18-23.module_el8.6.0__plus__983__plus__a7505f3f.x86_64",
    sha256 = "d24f86d286bd2294de8b3c2931c3f851495cd12f76a24705425635f55eaf1147",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/hivex-1.3.18-23.module_el8.6.0+983+a7505f3f.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/d24f86d286bd2294de8b3c2931c3f851495cd12f76a24705425635f55eaf1147",
    ],
)

rpm(
    name = "hwdata-0__0.314-8.14.el8.x86_64",
    sha256 = "56b38843fe7ae6317a64f9b3c013e381ab31ad8d5f2c4c5dcb63eca268e99e3a",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/hwdata-0.314-8.14.el8.noarch.rpm"],
)

rpm(
    name = "ibus-libs-0__1.5.19-14.el8.x86_64",
    sha256 = "e00f4e9d415d48446b83754dbdf1c78bd1604dfd10dc9035877131cfd78d1109",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/ibus-libs-1.5.19-14.el8.x86_64.rpm"],
)

rpm(
    name = "info-0__6.5-7.el8_5.aarch64",
    sha256 = "24a7e6f02ac095d965832203d0c8a9ee13aea301ef8572bb1ecdace435c796be",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/info-6.5-7.el8_5.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/24a7e6f02ac095d965832203d0c8a9ee13aea301ef8572bb1ecdace435c796be",
    ],
)

rpm(
    name = "info-0__6.5-7.el8_5.x86_64",
    sha256 = "63f03261cc8109b2fb61002ca50c93e52acb9cfd8382d139e8de6623394051e8",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/info-6.5-7.el8_5.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/63f03261cc8109b2fb61002ca50c93e52acb9cfd8382d139e8de6623394051e8",
    ],
)

rpm(
    name = "iproute-0__5.18.0-1.el8.aarch64",
    sha256 = "7ec84f47ebaed2388e48e27d9566a43609c7c384bbfbc3f0497c6bc314f618a5",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/iproute-5.18.0-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/7ec84f47ebaed2388e48e27d9566a43609c7c384bbfbc3f0497c6bc314f618a5",
    ],
)

rpm(
    name = "iproute-0__5.18.0-1.el8.x86_64",
    sha256 = "7ae4b834f060d111db19fa3cf6f6266d4c6fb56992b0347145799d7ff9f03d3c",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/iproute-5.18.0-1.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/7ae4b834f060d111db19fa3cf6f6266d4c6fb56992b0347145799d7ff9f03d3c",
    ],
)

rpm(
    name = "iproute-tc-0__5.18.0-1.el8.aarch64",
    sha256 = "8696d818b8ead9df0a2d66cf8e1fe03affd19899dd86e451267603faade5a161",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/iproute-tc-5.18.0-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/8696d818b8ead9df0a2d66cf8e1fe03affd19899dd86e451267603faade5a161",
    ],
)

rpm(
    name = "iproute-tc-0__5.18.0-1.el8.x86_64",
    sha256 = "bca80255b377f2a715c1fa2023485cd8fd03f2bab2a873faa0e5879082bca1c9",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/iproute-tc-5.18.0-1.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/bca80255b377f2a715c1fa2023485cd8fd03f2bab2a873faa0e5879082bca1c9",
    ],
)

rpm(
    name = "iptables-0__1.8.4-22.el8.aarch64",
    sha256 = "ddc326783f717ba44a9872faf2c82e2ea0fde7de021b2f144417889f84ce614e",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/iptables-1.8.4-22.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/ddc326783f717ba44a9872faf2c82e2ea0fde7de021b2f144417889f84ce614e",
    ],
)

rpm(
    name = "iptables-0__1.8.4-24.el8.x86_64",
    sha256 = "e4d26dec2832a8177e76d0d287a70dfaa57499ebf954610c215e449b9190492e",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/iptables-1.8.4-24.el8.x86_64.rpm"],
)

rpm(
    name = "iptables-libs-0__1.8.4-22.el8.aarch64",
    sha256 = "2c99183b888b75ae3b1d7665836757fb7a1ba130b03454177331f9efe33ca630",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/iptables-libs-1.8.4-22.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/2c99183b888b75ae3b1d7665836757fb7a1ba130b03454177331f9efe33ca630",
    ],
)

rpm(
    name = "iptables-libs-0__1.8.4-24.el8.x86_64",
    sha256 = "cf70e436e2fe912f419579500fd30512be5420009d63a82aacc47767b32901d5",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/iptables-libs-1.8.4-24.el8.x86_64.rpm"],
)

rpm(
    name = "iputils-0__20180629-10.el8.aarch64",
    sha256 = "7a40254a162ab0117a106ed2a08b824a2f2186b14e56257a5e848ae070cee0f1",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/iputils-20180629-10.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/7a40254a162ab0117a106ed2a08b824a2f2186b14e56257a5e848ae070cee0f1",
    ],
)

rpm(
    name = "iputils-0__20180629-10.el8.x86_64",
    sha256 = "66358ff76f9f26f6dbc403e479ab9389326d56233c5114daef316f589990c941",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/iputils-20180629-10.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/66358ff76f9f26f6dbc403e479ab9389326d56233c5114daef316f589990c941",
    ],
)

rpm(
    name = "ipxe-roms-qemu-0__20200823-7.git4bd064de.el8.x86_64",
    sha256 = "16bd1421408aa674e4d96d45787e652d97d11375b5818e35e37c1c1d8bf05cfa",
    urls = ["https://pkgs.dyn.su/el8/extras/x86_64/qemu/ipxe-roms-qemu-20200823-7.git4bd064de.el8.noarch.rpm"],
)

rpm(
    name = "isl-0__0.16.1-6.el8.aarch64",
    sha256 = "b9bd73b0edcd9573548853bd44f5a58919d9de77d8b1304a4176c7fad726b472",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/isl-0.16.1-6.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/b9bd73b0edcd9573548853bd44f5a58919d9de77d8b1304a4176c7fad726b472",
    ],
)

rpm(
    name = "isl-0__0.16.1-6.el8.x86_64",
    sha256 = "0cbdbdf53c8c12f48493bdae47d2bda45425011e67801a5827d164d6e10759ae",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/isl-0.16.1-6.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/0cbdbdf53c8c12f48493bdae47d2bda45425011e67801a5827d164d6e10759ae",
    ],
)

rpm(
    name = "jansson-0__2.14-1.el8.aarch64",
    sha256 = "69b4dd56ca16ed4ac5840e0d39a29d2e0b050905a349e1aceae4ec511a11b792",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/jansson-2.14-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/69b4dd56ca16ed4ac5840e0d39a29d2e0b050905a349e1aceae4ec511a11b792",
    ],
)

rpm(
    name = "jansson-0__2.14-1.el8.x86_64",
    sha256 = "f825b85b4506a740fb2f85b9a577c51264f3cfe792dd8b2bf8963059cc77c3c4",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/jansson-2.14-1.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/f825b85b4506a740fb2f85b9a577c51264f3cfe792dd8b2bf8963059cc77c3c4",
    ],
)

rpm(
    name = "jasper-libs-0__2.0.14-5.el8.x86_64",
    sha256 = "974e1dd8cbcd7bbe7f2d7cb305eef0258db5b38de552778f377fc6e4758320ce",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/jasper-libs-2.0.14-5.el8.x86_64.rpm"],
)

rpm(
    name = "jbigkit-libs-0__2.1-14.el8.x86_64",
    sha256 = "d5714bdf7c4e3b66540a43c5d58328dcef4f2360e58c578f2c5bbfce66651ad7",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/jbigkit-libs-2.1-14.el8.x86_64.rpm"],
)

rpm(
    name = "json-c-0__0.13.1-3.el8.aarch64",
    sha256 = "3bb6aa6c7aa0c3186c3dbce23661ec709c43c0e87a22a7e952148f515e2bfc82",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/json-c-0.13.1-3.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/3bb6aa6c7aa0c3186c3dbce23661ec709c43c0e87a22a7e952148f515e2bfc82",
    ],
)

rpm(
    name = "json-c-0__0.13.1-3.el8.x86_64",
    sha256 = "5035057553b61cb389c67aa2c29d99c8e0c1677369dad179d683942ccee90b3f",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/json-c-0.13.1-3.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/5035057553b61cb389c67aa2c29d99c8e0c1677369dad179d683942ccee90b3f",
    ],
)

rpm(
    name = "json-glib-0__1.4.4-1.el8.aarch64",
    sha256 = "01e70480bb032d5e0b60c5e732d4302d3a0ce73d1502a1729280d2b36e7e1c1a",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/json-glib-1.4.4-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/01e70480bb032d5e0b60c5e732d4302d3a0ce73d1502a1729280d2b36e7e1c1a",
    ],
)

rpm(
    name = "json-glib-0__1.4.4-1.el8.x86_64",
    sha256 = "98a6386df94fc9595365c3ecbc630708420fa68d1774614a723dec4a55e84b9c",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/json-glib-1.4.4-1.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/98a6386df94fc9595365c3ecbc630708420fa68d1774614a723dec4a55e84b9c",
    ],
)

rpm(
    name = "kernel-headers-0__4.18.0-408.el8.aarch64",
    sha256 = "208e7b141b8ad93ee6bd748f5c4117ed5a947b4ff48071d4fcdb826670aad76a",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/kernel-headers-4.18.0-408.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/208e7b141b8ad93ee6bd748f5c4117ed5a947b4ff48071d4fcdb826670aad76a",
    ],
)

rpm(
    name = "kernel-rxl-headers-0__5.4.209-1.el8.x86_64",
    sha256 = "53469bf837e4749a92df674ea07ed727e3ff0476d9133013d4003d9eeaf7517d",
    urls = ["https://pkgs.dyn.su/el8/extras/x86_64/kernel-rxl-headers-5.4.209-1.el8.x86_64.rpm"],
)

rpm(
    name = "keyutils-libs-0__1.5.10-9.el8.aarch64",
    sha256 = "c5af4350099a98929777412fb23e74c3bd2d7d8bbd09c2969a59d45937738aad",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/keyutils-libs-1.5.10-9.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/c5af4350099a98929777412fb23e74c3bd2d7d8bbd09c2969a59d45937738aad",
    ],
)

rpm(
    name = "keyutils-libs-0__1.5.10-9.el8.x86_64",
    sha256 = "423329269c719b96ada88a27325e1923e764a70672e0dc6817e22eff07a9af7b",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/keyutils-libs-1.5.10-9.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/423329269c719b96ada88a27325e1923e764a70672e0dc6817e22eff07a9af7b",
    ],
)

rpm(
    name = "kmod-0__25-19.el8.aarch64",
    sha256 = "056e83e9da3c6a582e83634b66c3ead78f1729f4b9dbd9970dbf3bfdc45edb54",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/kmod-25-19.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/056e83e9da3c6a582e83634b66c3ead78f1729f4b9dbd9970dbf3bfdc45edb54",
    ],
)

rpm(
    name = "kmod-0__25-19.el8.x86_64",
    sha256 = "37c299fdaa42efb0d653ba5e22c83bd20833af1244b66ed6ea880e75c1672dd2",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/kmod-25-19.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/37c299fdaa42efb0d653ba5e22c83bd20833af1244b66ed6ea880e75c1672dd2",
    ],
)

rpm(
    name = "kmod-libs-0__25-19.el8.aarch64",
    sha256 = "053b443be1bb0cbbc6da3314775391950350106462cc1dae01c7aed4358bf852",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/kmod-libs-25-19.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/053b443be1bb0cbbc6da3314775391950350106462cc1dae01c7aed4358bf852",
    ],
)

rpm(
    name = "kmod-libs-0__25-19.el8.x86_64",
    sha256 = "46a2ddc6067ed12089f04f2255c57117992807d707e280fc002f3ce786fc2abf",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/kmod-libs-25-19.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/46a2ddc6067ed12089f04f2255c57117992807d707e280fc002f3ce786fc2abf",
    ],
)

rpm(
    name = "krb5-libs-0__1.18.2-21.el8.aarch64",
    sha256 = "30f23e30b9e0de1c62a6b1d9f7031f7d5b263b458ad43c43915ea41a34711a92",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/krb5-libs-1.18.2-21.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/30f23e30b9e0de1c62a6b1d9f7031f7d5b263b458ad43c43915ea41a34711a92",
    ],
)

rpm(
    name = "krb5-libs-0__1.18.2-22.el8.x86_64",
    sha256 = "1dc1106dda34b328115dff7b2eca007dd93befb0bfa6a66c619f4b5637f6e004",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/krb5-libs-1.18.2-22.el8.x86_64.rpm"],
)

rpm(
    name = "lcms2-0__2.9-2.el8.x86_64",
    sha256 = "7ff6c7d3b55a8ea9c33e4e50e0420184a27f91badbd20ca7a24432b5ac60371c",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/lcms2-2.9-2.el8.x86_64.rpm"],
)

rpm(
    name = "less-0__530-1.el8.x86_64",
    sha256 = "f94172554b8ceeab97b560d0b05c2e2df4b2e737471adce6eca82fd3209be254",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/less-530-1.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/f94172554b8ceeab97b560d0b05c2e2df4b2e737471adce6eca82fd3209be254",
    ],
)

rpm(
    name = "libX11-0__1.6.8-5.el8.x86_64",
    sha256 = "2ab1fef0235ca1cafbe23ad6c9dbe3cd5d48ab99561f7e880456606a1ea78ae4",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libX11-1.6.8-5.el8.x86_64.rpm"],
)

rpm(
    name = "libX11-common-0__1.6.8-5.el8.x86_64",
    sha256 = "53760c2d7e17f31bd1f999cb448e902d4ba68eff0f99f6203d85998cd4c44918",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libX11-common-1.6.8-5.el8.noarch.rpm"],
)

rpm(
    name = "libX11-xcb-0__1.6.8-5.el8.x86_64",
    sha256 = "d8d58813823c960f344bdebf4ed888b53781c81175e97badd814a86dc811b362",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libX11-xcb-1.6.8-5.el8.x86_64.rpm"],
)

rpm(
    name = "libXau-0__1.0.9-3.el8.x86_64",
    sha256 = "49d972c660b9238dd1d58ab5952285b77e440820bf4563cce2b5ecd2f6ceba78",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libXau-1.0.9-3.el8.x86_64.rpm"],
)

rpm(
    name = "libXcomposite-0__0.4.4-14.el8.x86_64",
    sha256 = "4a93949edfc613fd6f42844e7c29951f357f085ac36541a676c421165914b4d4",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libXcomposite-0.4.4-14.el8.x86_64.rpm"],
)

rpm(
    name = "libXcursor-0__1.1.15-3.el8.x86_64",
    sha256 = "6d916871352bd0448efdfed35f3597f3f3a04591ca9f97e47af30ee1fddd5418",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libXcursor-1.1.15-3.el8.x86_64.rpm"],
)

rpm(
    name = "libXdamage-0__1.1.4-14.el8.x86_64",
    sha256 = "b4579c48c59ef7874191c6815b5321e7427ccd8f9cc7ac8624b7033485be82f2",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libXdamage-1.1.4-14.el8.x86_64.rpm"],
)

rpm(
    name = "libXext-0__1.3.4-1.el8.x86_64",
    sha256 = "9869db60ee2b6d8f2115937857fb0586802720a75e043baf21514011a9fa79aa",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libXext-1.3.4-1.el8.x86_64.rpm"],
)

rpm(
    name = "libXfixes-0__5.0.3-7.el8.x86_64",
    sha256 = "81f7df4c736963636c9ebab7441ca4f4e41a7483ef6e7b2ac0d1bf37afe52a14",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libXfixes-5.0.3-7.el8.x86_64.rpm"],
)

rpm(
    name = "libXft-0__2.3.3-1.el8.x86_64",
    sha256 = "ab754d37388e0ecb52152e41c9560392dd0d504939f850ff25d9794090f8b101",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libXft-2.3.3-1.el8.x86_64.rpm"],
)

rpm(
    name = "libXi-0__1.7.10-1.el8.x86_64",
    sha256 = "4a795d275c32ba03551178a088185af08391e5e137c6669eb8601e56c905631b",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libXi-1.7.10-1.el8.x86_64.rpm"],
)

rpm(
    name = "libXinerama-0__1.1.4-1.el8.x86_64",
    sha256 = "f07f41680f3d3a29981415c5c6d83e40da2958abbadc2417be6fbc467259be9b",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libXinerama-1.1.4-1.el8.x86_64.rpm"],
)

rpm(
    name = "libXrandr-0__1.5.2-1.el8.x86_64",
    sha256 = "f81bfbb4ad309a6f934aaf8ee3d11847360557c4f5e7198eff3e74e73d0d76fc",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libXrandr-1.5.2-1.el8.x86_64.rpm"],
)

rpm(
    name = "libXrender-0__0.9.10-7.el8.x86_64",
    sha256 = "11ac209220f3a53a762adebb4eeb43190e02ef7cdad2c54bcb474b206f7eb6f2",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libXrender-0.9.10-7.el8.x86_64.rpm"],
)

rpm(
    name = "libXtst-0__1.2.3-7.el8.x86_64",
    sha256 = "a19ecb3f3814649210b4667cf82ebca98b0d00e1b8222bc2f5aca2cc062999e6",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libXtst-1.2.3-7.el8.x86_64.rpm"],
)

rpm(
    name = "libXxf86vm-0__1.1.4-9.el8.x86_64",
    sha256 = "5813a48905fafc027e4b71b8113e654f23c963a9526015ec4fd0738b68de264a",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libXxf86vm-1.1.4-9.el8.x86_64.rpm"],
)

rpm(
    name = "libacl-0__2.2.53-1.el8.aarch64",
    sha256 = "c4cfed85e5a0db903ad134b4327b1714e5453fcf5c4348ec93ab344860a970ef",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libacl-2.2.53-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/c4cfed85e5a0db903ad134b4327b1714e5453fcf5c4348ec93ab344860a970ef",
    ],
)

rpm(
    name = "libacl-0__2.2.53-1.el8.x86_64",
    sha256 = "4973664648b7ed9278bf29074ec6a60a9f660aa97c23a283750483f64429d5bb",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libacl-2.2.53-1.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/4973664648b7ed9278bf29074ec6a60a9f660aa97c23a283750483f64429d5bb",
    ],
)

rpm(
    name = "libaio-0__0.3.112-1.el8.aarch64",
    sha256 = "3bcb1ade26c217ead2da81c92b7ef78026c4a78383d28b6e825a7b840cae97fa",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libaio-0.3.112-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/3bcb1ade26c217ead2da81c92b7ef78026c4a78383d28b6e825a7b840cae97fa",
    ],
)

rpm(
    name = "libaio-0__0.3.112-1.el8.x86_64",
    sha256 = "2c63399bee449fb6e921671a9bbf3356fda73f890b578820f7d926202e98a479",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libaio-0.3.112-1.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/2c63399bee449fb6e921671a9bbf3356fda73f890b578820f7d926202e98a479",
    ],
)

rpm(
    name = "libarchive-0__3.3.3-4.el8.aarch64",
    sha256 = "0dd36d8de0c8f40cbb01d9d1fc072eebf28967302b1eed287d7ad958aa383673",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libarchive-3.3.3-4.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/0dd36d8de0c8f40cbb01d9d1fc072eebf28967302b1eed287d7ad958aa383673",
    ],
)

rpm(
    name = "libarchive-0__3.3.3-5.el8.x86_64",
    sha256 = "d2e208573fde1934bd11c52a45edd6c360d365e0c675b43043fe863a248f5f5b",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libarchive-3.3.3-5.el8.x86_64.rpm"],
)

rpm(
    name = "libasan-0__8.5.0-15.el8.aarch64",
    sha256 = "34e627e042580439b22395344a15dbfb7fe0ce7a93530217ce38134278084c60",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libasan-8.5.0-15.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/34e627e042580439b22395344a15dbfb7fe0ce7a93530217ce38134278084c60",
    ],
)

rpm(
    name = "libassuan-0__2.5.1-3.el8.x86_64",
    sha256 = "b49e8c674e462e3f494e825c5fca64002008cbf7a47bf131aa98b7f41678a6eb",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libassuan-2.5.1-3.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/b49e8c674e462e3f494e825c5fca64002008cbf7a47bf131aa98b7f41678a6eb",
    ],
)

rpm(
    name = "libatomic-0__8.5.0-15.el8.aarch64",
    sha256 = "58ea796ac4166da751068de1e250378e83b016586e08e2b2fb85d5903387f3b4",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libatomic-8.5.0-15.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/58ea796ac4166da751068de1e250378e83b016586e08e2b2fb85d5903387f3b4",
    ],
)

rpm(
    name = "libattr-0__2.4.48-3.el8.aarch64",
    sha256 = "6a6db7eab6e53dccc54116d2ddf86b02db4cff332a58b868f7ba778a99666c58",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libattr-2.4.48-3.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/6a6db7eab6e53dccc54116d2ddf86b02db4cff332a58b868f7ba778a99666c58",
    ],
)

rpm(
    name = "libattr-0__2.4.48-3.el8.x86_64",
    sha256 = "a02e1344ccde1747501ceeeff37df4f18149fb79b435aa22add08cff6bab3a5a",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libattr-2.4.48-3.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/a02e1344ccde1747501ceeeff37df4f18149fb79b435aa22add08cff6bab3a5a",
    ],
)

rpm(
    name = "libblkid-0__2.32.1-35.el8.aarch64",
    sha256 = "4080ea4327979c8b461009107f2ab6090532e0cd1b944c3287c03bdbb9c9f022",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libblkid-2.32.1-35.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/4080ea4327979c8b461009107f2ab6090532e0cd1b944c3287c03bdbb9c9f022",
    ],
)

rpm(
    name = "libblkid-0__2.32.1-39.el8.x86_64",
    sha256 = "8368b8462e9763cdc7a9586ca1a266d3858aafa5ad82c473cb4e1c0bf2d6c755",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libblkid-2.32.1-39.el8.x86_64.rpm"],
)

rpm(
    name = "libbpf-0__0.5.0-1.el8.aarch64",
    sha256 = "1ecce335e1821b021b9fcfc8ffe1093a75f474249503510cf2bc499c61848cbb",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libbpf-0.5.0-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/1ecce335e1821b021b9fcfc8ffe1093a75f474249503510cf2bc499c61848cbb",
    ],
)

rpm(
    name = "libbpf-0__0.5.0-1.el8.x86_64",
    sha256 = "4d25308c27041d8a88a3340be12591e9bd46c9aebbe4195ee5d2f712d63ce033",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libbpf-0.5.0-1.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/4d25308c27041d8a88a3340be12591e9bd46c9aebbe4195ee5d2f712d63ce033",
    ],
)

rpm(
    name = "libburn-0__1.4.8-3.el8.aarch64",
    sha256 = "5ae88291a28b2a86efb6cdc8ff67baaf73dad1428c858c8b0fa9e8df0f0f041c",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/libburn-1.4.8-3.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/5ae88291a28b2a86efb6cdc8ff67baaf73dad1428c858c8b0fa9e8df0f0f041c",
    ],
)

rpm(
    name = "libburn-0__1.4.8-3.el8.x86_64",
    sha256 = "d4b0815ced6c1ec209b78fee4e2c1ee74efcd401d5462268b47d94a28ebfaf31",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libburn-1.4.8-3.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/d4b0815ced6c1ec209b78fee4e2c1ee74efcd401d5462268b47d94a28ebfaf31",
    ],
)

rpm(
    name = "libcap-0__2.48-4.el8.aarch64",
    sha256 = "f1fb5fe3b85ce5016a7882ccd9640b80f8fd6fbad1c44dc02076a8cdf33fc33d",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libcap-2.48-4.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/f1fb5fe3b85ce5016a7882ccd9640b80f8fd6fbad1c44dc02076a8cdf33fc33d",
    ],
)

rpm(
    name = "libcap-0__2.48-4.el8.x86_64",
    sha256 = "34f69bed9ae0f5ba314a62172e8cfd9cf6795cb0c3bd29f15d174fc2a0acbb5b",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libcap-2.48-4.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/34f69bed9ae0f5ba314a62172e8cfd9cf6795cb0c3bd29f15d174fc2a0acbb5b",
    ],
)

rpm(
    name = "libcap-ng-0__0.7.11-1.el8.aarch64",
    sha256 = "cbbbb1771fe9cfaa3284837e5e02cd2101190504ea0baa0278c9cfb2b169073c",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libcap-ng-0.7.11-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/cbbbb1771fe9cfaa3284837e5e02cd2101190504ea0baa0278c9cfb2b169073c",
    ],
)

rpm(
    name = "libcap-ng-0__0.7.11-1.el8.x86_64",
    sha256 = "15c3c696ec2e21f48e951f426d3c77b53b579605b8dd89843b35c9ab9b1d7e69",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libcap-ng-0.7.11-1.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/15c3c696ec2e21f48e951f426d3c77b53b579605b8dd89843b35c9ab9b1d7e69",
    ],
)

rpm(
    name = "libcom_err-0__1.45.6-5.el8.aarch64",
    sha256 = "bdd5ab69772a43725e1f8397e8142094bdd28b21b65ff02da74a8fc986424f3c",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libcom_err-1.45.6-5.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/bdd5ab69772a43725e1f8397e8142094bdd28b21b65ff02da74a8fc986424f3c",
    ],
)

rpm(
    name = "libcom_err-0__1.45.6-5.el8.x86_64",
    sha256 = "4e4f13acac0477f0a121812107a9939ea2164eebab052813f1618d5b7df5d87a",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libcom_err-1.45.6-5.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/4e4f13acac0477f0a121812107a9939ea2164eebab052813f1618d5b7df5d87a",
    ],
)

rpm(
    name = "libconfig-0__1.5-9.el8.x86_64",
    sha256 = "a4a2c7c0e2f454abae61dddbf4286a0b3617a8159fd20659bddbcedd8eaaa80c",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libconfig-1.5-9.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/a4a2c7c0e2f454abae61dddbf4286a0b3617a8159fd20659bddbcedd8eaaa80c",
    ],
)

rpm(
    name = "libcroco-0__0.6.12-4.el8_2.1.aarch64",
    sha256 = "0022ec2580783f68e603e9d4751478c28f2b383c596b4e896469077748771bfe",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libcroco-0.6.12-4.el8_2.1.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/0022ec2580783f68e603e9d4751478c28f2b383c596b4e896469077748771bfe",
    ],
)

rpm(
    name = "libcroco-0__0.6.12-4.el8_2.1.x86_64",
    sha256 = "87f2a4d80cf4f6a958f3662c6a382edefc32a5ad2c364a7f3c40337cf2b1e8ba",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libcroco-0.6.12-4.el8_2.1.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/87f2a4d80cf4f6a958f3662c6a382edefc32a5ad2c364a7f3c40337cf2b1e8ba",
    ],
)

rpm(
    name = "libcurl-minimal-0__7.61.1-25.el8.aarch64",
    sha256 = "2852cffc539a2178e52304b24c83ded856a7da3dbc76c0f21c7db522c72b03b1",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libcurl-minimal-7.61.1-25.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/2852cffc539a2178e52304b24c83ded856a7da3dbc76c0f21c7db522c72b03b1",
    ],
)

rpm(
    name = "libcurl-minimal-0__7.61.1-27.el8.x86_64",
    sha256 = "9466c9d951f5f41ae966c6c63de4fdb02813d207dcf8fd9a0fd9aa5e69abddf6",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libcurl-minimal-7.61.1-27.el8.x86_64.rpm"],
)

rpm(
    name = "libdatrie-0__0.2.9-7.el8.x86_64",
    sha256 = "7d43fda5ced8faf64d09cb3c47dcb6c9aa1fd936fc49f8609af29780c7a75f90",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libdatrie-0.2.9-7.el8.x86_64.rpm"],
)

rpm(
    name = "libdb-0__5.3.28-42.el8_4.aarch64",
    sha256 = "7ab75211c6fca91324039d3c2eb73903f2da73c17d6edaf8e997462ce4fbb46c",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libdb-5.3.28-42.el8_4.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/7ab75211c6fca91324039d3c2eb73903f2da73c17d6edaf8e997462ce4fbb46c",
    ],
)

rpm(
    name = "libdb-0__5.3.28-42.el8_4.x86_64",
    sha256 = "058f77432592f4337039cbb7a4e5f680020d8b85a477080c01d96a7728de6934",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libdb-5.3.28-42.el8_4.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/058f77432592f4337039cbb7a4e5f680020d8b85a477080c01d96a7728de6934",
    ],
)

rpm(
    name = "libdb-utils-0__5.3.28-42.el8_4.aarch64",
    sha256 = "84d0f5ae6a2bb4855d800c8e26be44bd06ac5f3c286a7877310bddabec12477a",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libdb-utils-5.3.28-42.el8_4.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/84d0f5ae6a2bb4855d800c8e26be44bd06ac5f3c286a7877310bddabec12477a",
    ],
)

rpm(
    name = "libdb-utils-0__5.3.28-42.el8_4.x86_64",
    sha256 = "ceb3dbd9e0d39d3e6b566eaf05359de4dd9a18d09da9238f2319f66f7cfebf7b",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libdb-utils-5.3.28-42.el8_4.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/ceb3dbd9e0d39d3e6b566eaf05359de4dd9a18d09da9238f2319f66f7cfebf7b",
    ],
)

rpm(
    name = "libdrm-0__2.4.114-1.el8.x86_64",
    sha256 = "af65274314c0e0423fd6430d19f79a0f11ec3f3f23fba1c10ea7ebdf47443cc9",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libdrm-2.4.114-1.el8.x86_64.rpm"],
)

rpm(
    name = "libepoxy-0__1.5.8-1.el8.x86_64",
    sha256 = "a825b6169fbd3377aed37ee63114aff24ac1a0ae123710c4559a56564fb0c15a",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libepoxy-1.5.8-1.el8.x86_64.rpm"],
)

rpm(
    name = "liberation-fonts-common-1__2.00.3-7.el8.x86_64",
    sha256 = "7e48c1ce8c9a48c0b7ca7795aafe83fe9871bd1ae88dfd31435bf318baac8f64",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/liberation-fonts-common-2.00.3-7.el8.noarch.rpm"],
)

rpm(
    name = "liberation-mono-fonts-1__2.00.3-7.el8.x86_64",
    sha256 = "fdbff1e8b002da671142b914c8f3750e51933f3de2ef2eeba2ef457e60d5e289",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/liberation-mono-fonts-2.00.3-7.el8.noarch.rpm"],
)

rpm(
    name = "libevent-0__2.1.8-5.el8.aarch64",
    sha256 = "a7fed3b521d23e60539dcbd548bda2a62f0d745a99dd5feeb43b6539f7f88232",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libevent-2.1.8-5.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/a7fed3b521d23e60539dcbd548bda2a62f0d745a99dd5feeb43b6539f7f88232",
    ],
)

rpm(
    name = "libevent-0__2.1.8-5.el8.x86_64",
    sha256 = "746bac6bb011a586d42bd82b2f8b25bac72c9e4bbd4c19a34cf88eadb1d83873",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libevent-2.1.8-5.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/746bac6bb011a586d42bd82b2f8b25bac72c9e4bbd4c19a34cf88eadb1d83873",
    ],
)

rpm(
    name = "libfdisk-0__2.32.1-35.el8.aarch64",
    sha256 = "408ebe12192393e73949b086db7a3d157868a1d9a5e08af5814a4f7f1a31f10d",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libfdisk-2.32.1-35.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/408ebe12192393e73949b086db7a3d157868a1d9a5e08af5814a4f7f1a31f10d",
    ],
)

rpm(
    name = "libfdisk-0__2.32.1-39.el8.x86_64",
    sha256 = "7e8c4ea7f5ea7339bba2db65fa8737e7acba431aa9a50b11cb741fa61aaf374d",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libfdisk-2.32.1-39.el8.x86_64.rpm"],
)

rpm(
    name = "libfdt-0__1.6.0-1.el8.aarch64",
    sha256 = "a2f3c86d18ee25ce4764a1df0854c63b615db37291ef9780e649f0123a92acf5",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/libfdt-1.6.0-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/a2f3c86d18ee25ce4764a1df0854c63b615db37291ef9780e649f0123a92acf5",
    ],
)

rpm(
    name = "libfdt-0__1.6.0-1.el8.x86_64",
    sha256 = "1788b4786715c45a1ac90ca9f413ef51f2cdd03170a981e0ef13eab204f44429",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libfdt-1.6.0-1.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/1788b4786715c45a1ac90ca9f413ef51f2cdd03170a981e0ef13eab204f44429",
    ],
)

rpm(
    name = "libffi-0__3.1-23.el8.aarch64",
    sha256 = "ba34d0bb067722c37dd4367534d82aa18c659facbfd17952f8d826e8662cb7c1",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libffi-3.1-23.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/ba34d0bb067722c37dd4367534d82aa18c659facbfd17952f8d826e8662cb7c1",
    ],
)

rpm(
    name = "libffi-0__3.1-24.el8.x86_64",
    sha256 = "3a0b75d820053e5a75f3a9a04fa2b78a7ac559140d7ce98f0d684cd8433ece81",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libffi-3.1-24.el8.x86_64.rpm"],
)

rpm(
    name = "libgcc-0__8.5.0-15.el8.aarch64",
    sha256 = "f62a7bd6b2ce584a9ee3561513053372db492efd867333b27f7ba9a3844ff553",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libgcc-8.5.0-15.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/f62a7bd6b2ce584a9ee3561513053372db492efd867333b27f7ba9a3844ff553",
    ],
)

rpm(
    name = "libgcc-0__8.5.0-17.el8.x86_64",
    sha256 = "862a1396e875c95be53adb2796048e455d740f218b67f224bf21f8fd214cc73a",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libgcc-8.5.0-17.el8.x86_64.rpm"],
)

rpm(
    name = "libgcrypt-0__1.8.5-7.el8.aarch64",
    sha256 = "88a32029615cc5986884cbab1b5c137e455b9ef08b23c6219b9ec9b42079be88",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libgcrypt-1.8.5-7.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/88a32029615cc5986884cbab1b5c137e455b9ef08b23c6219b9ec9b42079be88",
    ],
)

rpm(
    name = "libgcrypt-0__1.8.5-7.el8.x86_64",
    sha256 = "01541f1263532f80114111a44f797d6a8eed75744db997e85fddd021e636c5bb",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libgcrypt-1.8.5-7.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/01541f1263532f80114111a44f797d6a8eed75744db997e85fddd021e636c5bb",
    ],
)

rpm(
    name = "libglvnd-1__1.3.4-1.el8.x86_64",
    sha256 = "a94d8debdf9e1f20dc561baaa1c5903ef85c511f2b647092b5d8908ccfbf6a6a",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libglvnd-1.3.4-1.el8.x86_64.rpm"],
)

rpm(
    name = "libglvnd-egl-1__1.3.4-1.el8.x86_64",
    sha256 = "0c7e300aae2f33e48ae5bedbbcf9c6b50af18477d9493075c73355c7fe080b43",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libglvnd-egl-1.3.4-1.el8.x86_64.rpm"],
)

rpm(
    name = "libglvnd-glx-1__1.3.4-1.el8.x86_64",
    sha256 = "bf40ab7dbe4ae55fb5403204df6b9b27013898cdb450da39e8e19a2c4229aea5",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libglvnd-glx-1.3.4-1.el8.x86_64.rpm"],
)

rpm(
    name = "libgomp-0__8.5.0-15.el8.aarch64",
    sha256 = "edb71029b4d451240f53399652c872035ebab3237bfa4d416e010be58bc8a056",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libgomp-8.5.0-15.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/edb71029b4d451240f53399652c872035ebab3237bfa4d416e010be58bc8a056",
    ],
)

rpm(
    name = "libgomp-0__8.5.0-17.el8.x86_64",
    sha256 = "23fafd607c2bb0237f6ca0125992c564677f188d48c8d1c55d3dfd56aee9c37b",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libgomp-8.5.0-17.el8.x86_64.rpm"],
)

rpm(
    name = "libgpg-error-0__1.31-1.el8.aarch64",
    sha256 = "b953729a0a2be24749aeee9f00853fdc3227737971cf052a999a37ac36387cd9",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libgpg-error-1.31-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/b953729a0a2be24749aeee9f00853fdc3227737971cf052a999a37ac36387cd9",
    ],
)

rpm(
    name = "libgpg-error-0__1.31-1.el8.x86_64",
    sha256 = "845a0732d9d7a01b909124cd8293204764235c2d856227c7a74dfa0e38113e34",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libgpg-error-1.31-1.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/845a0732d9d7a01b909124cd8293204764235c2d856227c7a74dfa0e38113e34",
    ],
)

rpm(
    name = "libguestfs-1__1.44.0-5.module_el8.6.0__plus__1087__plus__b42c8331.x86_64",
    sha256 = "a0cbdc5c27f1d45480b2c4b28caac267a9a879de19091efa057119705611cbef",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libguestfs-1.44.0-5.module_el8.6.0+1087+b42c8331.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/a0cbdc5c27f1d45480b2c4b28caac267a9a879de19091efa057119705611cbef",
    ],
)

rpm(
    name = "libguestfs-tools-1__1.44.0-5.module_el8.6.0__plus__1087__plus__b42c8331.x86_64",
    sha256 = "fb8f81a46a30e7254f614f5b0376af1fef45c9082b2e6f88061e61cc046de99f",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libguestfs-tools-1.44.0-5.module_el8.6.0+1087+b42c8331.noarch.rpm",
        "https://storage.googleapis.com/builddeps/fb8f81a46a30e7254f614f5b0376af1fef45c9082b2e6f88061e61cc046de99f",
    ],
)

rpm(
    name = "libguestfs-tools-c-1__1.44.0-5.module_el8.6.0__plus__1087__plus__b42c8331.x86_64",
    sha256 = "61bb7c563c80a44fcce4bf9c1004539cf33165700f94a3ee384483345f60edc2",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libguestfs-tools-c-1.44.0-5.module_el8.6.0+1087+b42c8331.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/61bb7c563c80a44fcce4bf9c1004539cf33165700f94a3ee384483345f60edc2",
    ],
)

rpm(
    name = "libgusb-0__0.3.0-1.el8.x86_64",
    sha256 = "ab7bc1a828168006b88934bea949ab2b29b837b0a431f7da1a12147f64f6ddb5",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libgusb-0.3.0-1.el8.x86_64.rpm"],
)

rpm(
    name = "libibverbs-0__37.2-1.el8.aarch64",
    sha256 = "588f7bddcde691b08b8cf652de96f5e3334231ff4d3697353fbd53d25eadad03",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libibverbs-37.2-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/588f7bddcde691b08b8cf652de96f5e3334231ff4d3697353fbd53d25eadad03",
    ],
)

rpm(
    name = "libibverbs-0__41.0-1.el8.x86_64",
    sha256 = "888b1ce059dfaf1b8277cac3529970114ba1cadc75fbcf9410f3031451ab7e30",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libibverbs-41.0-1.el8.x86_64.rpm"],
)

rpm(
    name = "libidn2-0__2.2.0-1.el8.aarch64",
    sha256 = "b62589101a60a365ef34447cae78f62e6dba560d403dc56c87036709ea00ad88",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libidn2-2.2.0-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/b62589101a60a365ef34447cae78f62e6dba560d403dc56c87036709ea00ad88",
    ],
)

rpm(
    name = "libidn2-0__2.2.0-1.el8.x86_64",
    sha256 = "7e08785bd3cc0e09f9ab4bf600b98b705203d552cbb655269a939087987f1694",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libidn2-2.2.0-1.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/7e08785bd3cc0e09f9ab4bf600b98b705203d552cbb655269a939087987f1694",
    ],
)

rpm(
    name = "libisoburn-0__1.4.8-4.el8.aarch64",
    sha256 = "3ff828ef16f6033227d71207bc1b00983b826172fe7c575cd7590a72d846d831",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/libisoburn-1.4.8-4.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/3ff828ef16f6033227d71207bc1b00983b826172fe7c575cd7590a72d846d831",
    ],
)

rpm(
    name = "libisoburn-0__1.4.8-4.el8.x86_64",
    sha256 = "7aa030310250b462d90895d8c04ce47695722d86f5470930fdf8bfba0570c4dc",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libisoburn-1.4.8-4.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/7aa030310250b462d90895d8c04ce47695722d86f5470930fdf8bfba0570c4dc",
    ],
)

rpm(
    name = "libisofs-0__1.4.8-3.el8.aarch64",
    sha256 = "2e5435efba38348be8d33a43e5abbffc85f7c5a9504ebe6451b87c44006b3b4c",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/libisofs-1.4.8-3.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/2e5435efba38348be8d33a43e5abbffc85f7c5a9504ebe6451b87c44006b3b4c",
    ],
)

rpm(
    name = "libisofs-0__1.4.8-3.el8.x86_64",
    sha256 = "66b7bcc256b62736f7b3d33fa65c6a91a17e08c61484a7c3748f4f86b4589bc7",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libisofs-1.4.8-3.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/66b7bcc256b62736f7b3d33fa65c6a91a17e08c61484a7c3748f4f86b4589bc7",
    ],
)

rpm(
    name = "libjpeg-turbo-0__1.5.3-12.el8.x86_64",
    sha256 = "94b6e9d7ebd6d3bee36ac8b5c381a305bc16158a7de5bf7b71ddf2f41f10f03c",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libjpeg-turbo-1.5.3-12.el8.x86_64.rpm"],
)

rpm(
    name = "libksba-0__1.3.5-8.el8.x86_64",
    sha256 = "8054ca806450e99f1a65d52315229d036cb495ffddfd3f9fccb44e05d0108b46",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libksba-1.3.5-8.el8.x86_64.rpm"],
)

rpm(
    name = "libmnl-0__1.0.4-6.el8.aarch64",
    sha256 = "fbe4f2cb2660ebe3cb90a73c7dfbd978059af138356e46c9a93049761c0467ef",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libmnl-1.0.4-6.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/fbe4f2cb2660ebe3cb90a73c7dfbd978059af138356e46c9a93049761c0467ef",
    ],
)

rpm(
    name = "libmnl-0__1.0.4-6.el8.x86_64",
    sha256 = "30fab73ee155f03dbbd99c1e30fe59dfba4ae8fdb2e7213451ccc36d6918bfcc",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libmnl-1.0.4-6.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/30fab73ee155f03dbbd99c1e30fe59dfba4ae8fdb2e7213451ccc36d6918bfcc",
    ],
)

rpm(
    name = "libmodman-0__2.0.1-17.el8.x86_64",
    sha256 = "c3b8c553b166491d3114793e198cd1aad95e494d177af8d0dc7180b8b841124d",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libmodman-2.0.1-17.el8.x86_64.rpm"],
)

rpm(
    name = "libmount-0__2.32.1-35.el8.aarch64",
    sha256 = "0c4101fb30bfa09d02e950aba244857cdcdc291a9f195029ed383be7c815f43d",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libmount-2.32.1-35.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/0c4101fb30bfa09d02e950aba244857cdcdc291a9f195029ed383be7c815f43d",
    ],
)

rpm(
    name = "libmount-0__2.32.1-39.el8.x86_64",
    sha256 = "199c6968b0caa6fbe7f413c704e8a22f9915c54883809fd4f61c6327e2eb45c0",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libmount-2.32.1-39.el8.x86_64.rpm"],
)

rpm(
    name = "libmpc-0__1.1.0-9.1.el8.aarch64",
    sha256 = "9701bd94db9b467e11590b2de375a122ab61aa8d624be7df22631a6da91c79e4",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/libmpc-1.1.0-9.1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/9701bd94db9b467e11590b2de375a122ab61aa8d624be7df22631a6da91c79e4",
    ],
)

rpm(
    name = "libmpc-0__1.1.0-9.1.el8.x86_64",
    sha256 = "93c2232d1885ec6265159f4669aeb13335a80e74d3ae0832f624678d87ea3638",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libmpc-1.1.0-9.1.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/93c2232d1885ec6265159f4669aeb13335a80e74d3ae0832f624678d87ea3638",
    ],
)

rpm(
    name = "libnetfilter_conntrack-0__1.0.6-5.el8.aarch64",
    sha256 = "4e43b0f85746f74064b082fdf6914ba4e9fe386651b1c39aeaecc702b2a59fc0",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libnetfilter_conntrack-1.0.6-5.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/4e43b0f85746f74064b082fdf6914ba4e9fe386651b1c39aeaecc702b2a59fc0",
    ],
)

rpm(
    name = "libnetfilter_conntrack-0__1.0.6-5.el8.x86_64",
    sha256 = "224100af3ecfc80c416796ec02c7c4dd113a38d42349d763485f3b42f260493f",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libnetfilter_conntrack-1.0.6-5.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/224100af3ecfc80c416796ec02c7c4dd113a38d42349d763485f3b42f260493f",
    ],
)

rpm(
    name = "libnfnetlink-0__1.0.1-13.el8.aarch64",
    sha256 = "8422fbc84108abc9a89fe98cef9cd18ad1788b4dc6a9ec0bba1836b772fcaeda",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libnfnetlink-1.0.1-13.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/8422fbc84108abc9a89fe98cef9cd18ad1788b4dc6a9ec0bba1836b772fcaeda",
    ],
)

rpm(
    name = "libnfnetlink-0__1.0.1-13.el8.x86_64",
    sha256 = "cec98aa5fbefcb99715921b493b4f92d34c4eeb823e9c8741aa75e280def89f1",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libnfnetlink-1.0.1-13.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/cec98aa5fbefcb99715921b493b4f92d34c4eeb823e9c8741aa75e280def89f1",
    ],
)

rpm(
    name = "libnftnl-0__1.1.5-5.el8.aarch64",
    sha256 = "00522e43ce63cf63468052e627a429ededac0815212c644f4eadda88b990c3ee",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libnftnl-1.1.5-5.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/00522e43ce63cf63468052e627a429ededac0815212c644f4eadda88b990c3ee",
    ],
)

rpm(
    name = "libnftnl-0__1.1.5-5.el8.x86_64",
    sha256 = "293e1f0f44a9c1d5dedbe831dff3049fad9e88c5f0e281d889f427603ac51fa6",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libnftnl-1.1.5-5.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/293e1f0f44a9c1d5dedbe831dff3049fad9e88c5f0e281d889f427603ac51fa6",
    ],
)

rpm(
    name = "libnghttp2-0__1.33.0-3.el8_2.1.aarch64",
    sha256 = "23e9ff009c2316652c3bcd96a8b69b5bc26f2acd46214f652a7ce26a572cbabb",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libnghttp2-1.33.0-3.el8_2.1.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/23e9ff009c2316652c3bcd96a8b69b5bc26f2acd46214f652a7ce26a572cbabb",
    ],
)

rpm(
    name = "libnghttp2-0__1.33.0-3.el8_2.1.x86_64",
    sha256 = "0126a384853d46484dec98601a4cb4ce58b2e0411f8f7ef09937174dd5975bac",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libnghttp2-1.33.0-3.el8_2.1.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/0126a384853d46484dec98601a4cb4ce58b2e0411f8f7ef09937174dd5975bac",
    ],
)

rpm(
    name = "libnl3-0__3.7.0-1.el8.aarch64",
    sha256 = "8c8dd63daf7ad4c6322a4316fceb256f1cfd2d8244bce515bbae539b4314a643",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libnl3-3.7.0-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/8c8dd63daf7ad4c6322a4316fceb256f1cfd2d8244bce515bbae539b4314a643",
    ],
)

rpm(
    name = "libnl3-0__3.7.0-1.el8.x86_64",
    sha256 = "9ce7aa4d7bd810448d9fb3aa85a66cca00950f7c2c59bc9721ced3e4f3ad2885",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libnl3-3.7.0-1.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/9ce7aa4d7bd810448d9fb3aa85a66cca00950f7c2c59bc9721ced3e4f3ad2885",
    ],
)

rpm(
    name = "libnsl2-0__1.2.0-2.20180605git4a062cf.el8.aarch64",
    sha256 = "b33276781f442757afd5e066ead95ec79927f2aed608a368420f230d5ee28686",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libnsl2-1.2.0-2.20180605git4a062cf.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/b33276781f442757afd5e066ead95ec79927f2aed608a368420f230d5ee28686",
    ],
)

rpm(
    name = "libnsl2-0__1.2.0-2.20180605git4a062cf.el8.x86_64",
    sha256 = "5846c73edfa2ff673989728e9621cce6a1369eb2f8a269ac5205c381a10d327a",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libnsl2-1.2.0-2.20180605git4a062cf.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/5846c73edfa2ff673989728e9621cce6a1369eb2f8a269ac5205c381a10d327a",
    ],
)

rpm(
    name = "libpcap-14__1.9.1-5.el8.aarch64",
    sha256 = "239019a8aadb26e4b015d99f7fe49e80c2d1dfa227f7c71322dca2a2a85c2de1",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libpcap-1.9.1-5.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/239019a8aadb26e4b015d99f7fe49e80c2d1dfa227f7c71322dca2a2a85c2de1",
    ],
)

rpm(
    name = "libpcap-14__1.9.1-5.el8.x86_64",
    sha256 = "7f429477c26b4650a3eca4a27b3972ff0857c843bdb4d8fcb02086da111ce5fd",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libpcap-1.9.1-5.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/7f429477c26b4650a3eca4a27b3972ff0857c843bdb4d8fcb02086da111ce5fd",
    ],
)

rpm(
    name = "libpciaccess-0__0.14-1.el8.x86_64",
    sha256 = "759386be8f49257266ac614432b762b8e486a89aac5d5f7a581a0330efb59c77",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libpciaccess-0.14-1.el8.x86_64.rpm"],
)

rpm(
    name = "libpkgconf-0__1.4.2-1.el8.aarch64",
    sha256 = "8f3e34df67e6c4a20bd7617f17d1199f0441a626fbab8059ddc6bf06c7ff4e78",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libpkgconf-1.4.2-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/8f3e34df67e6c4a20bd7617f17d1199f0441a626fbab8059ddc6bf06c7ff4e78",
    ],
)

rpm(
    name = "libpkgconf-0__1.4.2-1.el8.x86_64",
    sha256 = "a76ff4cf270d2e38106a4bba1880c3a0899d186cd4e1986d7e97c01b934e13b7",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libpkgconf-1.4.2-1.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/a76ff4cf270d2e38106a4bba1880c3a0899d186cd4e1986d7e97c01b934e13b7",
    ],
)

rpm(
    name = "libpmem-0__1.12.1-1.module_el8.8.0__plus__1231__plus__994ef5f7.x86_64",
    sha256 = "631f555b4816b73e9f0c5cbf76136d587a93ca03ba735747ac03fc6c6a73bad2",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libpmem-1.12.1-1.module_el8.8.0+1231+994ef5f7.x86_64.rpm"],
)

rpm(
    name = "libpng-2__1.6.34-5.el8.aarch64",
    sha256 = "d7bd4e7a7ff4424266c0f6030bf444de0bea88d0540ff4caf4f7f6c2bac175f6",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libpng-1.6.34-5.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/d7bd4e7a7ff4424266c0f6030bf444de0bea88d0540ff4caf4f7f6c2bac175f6",
    ],
)

rpm(
    name = "libpng-2__1.6.34-5.el8.x86_64",
    sha256 = "cc2f054cf7ef006faf0b179701838ff8632c3ac5f45a0199a13f9c237f632b82",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libpng-1.6.34-5.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/cc2f054cf7ef006faf0b179701838ff8632c3ac5f45a0199a13f9c237f632b82",
    ],
)

rpm(
    name = "libproxy-0__0.4.15-5.2.el8.x86_64",
    sha256 = "c9597eecf39a25497b2ac3c69bc9777eda05b9eaa6d5d29d004a81d71a45d0d7",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libproxy-0.4.15-5.2.el8.x86_64.rpm"],
)

rpm(
    name = "libpwquality-0__1.4.4-3.el8.aarch64",
    sha256 = "64e55ddddc1dd27e05097c9222e73052f6f20f9d2f7605f46922b7756adeb0b5",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libpwquality-1.4.4-3.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/64e55ddddc1dd27e05097c9222e73052f6f20f9d2f7605f46922b7756adeb0b5",
    ],
)

rpm(
    name = "libpwquality-0__1.4.4-5.el8.x86_64",
    sha256 = "4a7159ebfb7914f23f009981a38fcbec8368b243b20dfed6326a6dade95cf3a2",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libpwquality-1.4.4-5.el8.x86_64.rpm"],
)

rpm(
    name = "librdmacm-0__37.2-1.el8.aarch64",
    sha256 = "b2a6c52cb312817b7c4010d7af46df4cf378a680bb611770e85dae2f259fa96c",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/librdmacm-37.2-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/b2a6c52cb312817b7c4010d7af46df4cf378a680bb611770e85dae2f259fa96c",
    ],
)

rpm(
    name = "librdmacm-0__41.0-1.el8.x86_64",
    sha256 = "caf52cd9c97677b5684730ad61f8abe464cfc41d332b3f4d4887fb2e8ea87916",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/librdmacm-41.0-1.el8.x86_64.rpm"],
)

rpm(
    name = "libseccomp-0__2.5.2-1.el8.aarch64",
    sha256 = "2460f610a00c11b7070ff75d27fb22fab4b8d67c856da2ffb097cf3eff28f365",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libseccomp-2.5.2-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/2460f610a00c11b7070ff75d27fb22fab4b8d67c856da2ffb097cf3eff28f365",
    ],
)

rpm(
    name = "libseccomp-0__2.5.2-1.el8.x86_64",
    sha256 = "4a6322832274a9507108719de9af48406ee0fcfc54c9906b9450e1ae231ede4b",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libseccomp-2.5.2-1.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/4a6322832274a9507108719de9af48406ee0fcfc54c9906b9450e1ae231ede4b",
    ],
)

rpm(
    name = "libselinux-0__2.9-5.el8.aarch64",
    sha256 = "9474fe348bd9e3a7a6ffe7813538e979e80ddb970b074e4e79bd122b4ece8b64",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libselinux-2.9-5.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/9474fe348bd9e3a7a6ffe7813538e979e80ddb970b074e4e79bd122b4ece8b64",
    ],
)

rpm(
    name = "libselinux-0__2.9-8.el8.x86_64",
    sha256 = "67f7412ebbbc65ec953aa4e99489c04f821c9645fe048c3ee170040663535dc2",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libselinux-2.9-8.el8.x86_64.rpm"],
)

rpm(
    name = "libselinux-utils-0__2.9-5.el8.aarch64",
    sha256 = "e4613455147d283b222fcff5ef0f85b3a1a323893ed884db8950e51936e97c52",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libselinux-utils-2.9-5.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/e4613455147d283b222fcff5ef0f85b3a1a323893ed884db8950e51936e97c52",
    ],
)

rpm(
    name = "libselinux-utils-0__2.9-8.el8.x86_64",
    sha256 = "d54bc5c131a6b41d8d69235dcb33ddb8a96df549f3da7b3020bf4dbdee65d71e",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libselinux-utils-2.9-8.el8.x86_64.rpm"],
)

rpm(
    name = "libsemanage-0__2.9-8.el8.aarch64",
    sha256 = "8d45efec5a7de58fa3c951c06823a401f9004a97e29d6a7a990114d3d29da7ba",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libsemanage-2.9-8.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/8d45efec5a7de58fa3c951c06823a401f9004a97e29d6a7a990114d3d29da7ba",
    ],
)

rpm(
    name = "libsemanage-0__2.9-9.el8.x86_64",
    sha256 = "7b8293193b1dda6c408c04074c4b501faf37ff9e4a4b6cd1ca2cce81d5bb67bf",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libsemanage-2.9-9.el8.x86_64.rpm"],
)

rpm(
    name = "libsepol-0__2.9-3.el8.aarch64",
    sha256 = "e9d2e6252228076c270850b51b7205baed31c1c3c2ccdb9d3280c9b0de5d652a",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libsepol-2.9-3.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/e9d2e6252228076c270850b51b7205baed31c1c3c2ccdb9d3280c9b0de5d652a",
    ],
)

rpm(
    name = "libsepol-0__2.9-3.el8.x86_64",
    sha256 = "f91e372ffa25c4c82ae7e001565cf5ff73048c407083493555025fdb5fc4c14a",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libsepol-2.9-3.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/f91e372ffa25c4c82ae7e001565cf5ff73048c407083493555025fdb5fc4c14a",
    ],
)

rpm(
    name = "libsigsegv-0__2.11-5.el8.aarch64",
    sha256 = "b377f4e8bcdc750ed0be94f97bdbfbb12843c458fbc1d5d507f92ad04aaf592b",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libsigsegv-2.11-5.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/b377f4e8bcdc750ed0be94f97bdbfbb12843c458fbc1d5d507f92ad04aaf592b",
    ],
)

rpm(
    name = "libsigsegv-0__2.11-5.el8.x86_64",
    sha256 = "02d728cf74eb47005babeeab5ac68ca04472c643203a1faef0037b5f33710fe2",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libsigsegv-2.11-5.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/02d728cf74eb47005babeeab5ac68ca04472c643203a1faef0037b5f33710fe2",
    ],
)

rpm(
    name = "libsmartcols-0__2.32.1-35.el8.aarch64",
    sha256 = "f4e88f36b11fe6529e0720758857276ea5ff93174afa8be006c70ec0bcca7fbd",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libsmartcols-2.32.1-35.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/f4e88f36b11fe6529e0720758857276ea5ff93174afa8be006c70ec0bcca7fbd",
    ],
)

rpm(
    name = "libsmartcols-0__2.32.1-39.el8.x86_64",
    sha256 = "9277921d36c7164667fb6be5fe191adec82ef6f6e50551ccc7a26c9f3a5cc67b",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libsmartcols-2.32.1-39.el8.x86_64.rpm"],
)

rpm(
    name = "libsoup-0__2.62.3-3.el8.x86_64",
    sha256 = "b97273e313e5234cef54eb7fa9bd12249194b83664e28d6bfd724e69717e9c1f",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libsoup-2.62.3-3.el8.x86_64.rpm"],
)

rpm(
    name = "libss-0__1.45.6-5.el8.aarch64",
    sha256 = "68b0f490ced8811f8b25423c7bd2d81b26301317e4445705c4b280283a50b8e9",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libss-1.45.6-5.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/68b0f490ced8811f8b25423c7bd2d81b26301317e4445705c4b280283a50b8e9",
    ],
)

rpm(
    name = "libss-0__1.45.6-5.el8.x86_64",
    sha256 = "f489f5eaaddbdedae046e4ddfe93947cdd636533ca8d35820bf5c92ae5dd3037",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libss-1.45.6-5.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/f489f5eaaddbdedae046e4ddfe93947cdd636533ca8d35820bf5c92ae5dd3037",
    ],
)

rpm(
    name = "libssh-0__0.9.6-3.el8.aarch64",
    sha256 = "4e7b5c73bf2ff1dc42904d96b86891ab3d2ccc27ba0e6d71de4984f9b1e71d65",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libssh-0.9.6-3.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/4e7b5c73bf2ff1dc42904d96b86891ab3d2ccc27ba0e6d71de4984f9b1e71d65",
    ],
)

rpm(
    name = "libssh-0__0.9.6-5.el8.x86_64",
    sha256 = "2e3184566602b90579267f98eaa460042b99791b5e242a0c7acaf613eb1d6b76",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libssh-0.9.6-5.el8.x86_64.rpm"],
)

rpm(
    name = "libssh-config-0__0.9.6-3.el8.aarch64",
    sha256 = "e9e954ba21bac58e3aebaf52bf824758fe4c2ad09d75171b3009a214bd52bbec",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libssh-config-0.9.6-3.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/e9e954ba21bac58e3aebaf52bf824758fe4c2ad09d75171b3009a214bd52bbec",
    ],
)

rpm(
    name = "libssh-config-0__0.9.6-5.el8.x86_64",
    sha256 = "55c673e4ae29ab8df6e023451ca409bc9808aec9ca00773c7d7bb1b1837ee2a8",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libssh-config-0.9.6-5.el8.noarch.rpm"],
)

rpm(
    name = "libsss_idmap-0__2.7.3-1.el8.aarch64",
    sha256 = "93bb0f876d19b1cfa61acceaf725cc72d35fe341ab4972a9dfd588ed3a563d2e",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libsss_idmap-2.7.3-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/93bb0f876d19b1cfa61acceaf725cc72d35fe341ab4972a9dfd588ed3a563d2e",
    ],
)

rpm(
    name = "libsss_idmap-0__2.8.1-1.el8.x86_64",
    sha256 = "b7fb9d90fce8edec5edb27014eac9aa4bc53b6b3451439049377dfcce411a8b5",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libsss_idmap-2.8.1-1.el8.x86_64.rpm"],
)

rpm(
    name = "libsss_nss_idmap-0__2.7.3-1.el8.aarch64",
    sha256 = "256d5b1d23bab41df0150964b9c45892beeebb07fa8345e27ffcf92ae5a4b6d9",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libsss_nss_idmap-2.7.3-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/256d5b1d23bab41df0150964b9c45892beeebb07fa8345e27ffcf92ae5a4b6d9",
    ],
)

rpm(
    name = "libsss_nss_idmap-0__2.8.1-1.el8.x86_64",
    sha256 = "86cb5a2faf024d4036c11feee140577d3959f311ade1e0c5717289f4e3327c06",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libsss_nss_idmap-2.8.1-1.el8.x86_64.rpm"],
)

rpm(
    name = "libstdc__plus____plus__-0__8.5.0-15.el8.aarch64",
    sha256 = "91d6f78ddeab3c6df90479eeca76e77450605983619a54c01faaa8ede3767214",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libstdc++-8.5.0-15.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/91d6f78ddeab3c6df90479eeca76e77450605983619a54c01faaa8ede3767214",
    ],
)

rpm(
    name = "libstdc__plus____plus__-0__8.5.0-17.el8.x86_64",
    sha256 = "42be13127a6ad3138d494e8a848d839369e41f95d190ca1acb2786f779328dec",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libstdc++-8.5.0-17.el8.x86_64.rpm"],
)

rpm(
    name = "libtasn1-0__4.13-3.el8.aarch64",
    sha256 = "3401ccfb7fd08c12578b6257b4dac7e94ba5f4cd70fc6a234fd90bb99d1bb108",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libtasn1-4.13-3.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/3401ccfb7fd08c12578b6257b4dac7e94ba5f4cd70fc6a234fd90bb99d1bb108",
    ],
)

rpm(
    name = "libtasn1-0__4.13-3.el8.x86_64",
    sha256 = "e8d9697a8914226a2d3ed5a4523b85e8e70ac09cf90aae05395e6faee9858534",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libtasn1-4.13-3.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/e8d9697a8914226a2d3ed5a4523b85e8e70ac09cf90aae05395e6faee9858534",
    ],
)

rpm(
    name = "libthai-0__0.1.27-2.el8.x86_64",
    sha256 = "91bbf9cd7d7ae62682a3d24a889512daef57c3c4b41866336c42af6361702fef",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libthai-0.1.27-2.el8.x86_64.rpm"],
)

rpm(
    name = "libtiff-0__4.0.9-26.el8.x86_64",
    sha256 = "c7a992d4ad5da9df1cd12953bc0b84a4f3cb1e09ccb80b541eb7f15301b2c369",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libtiff-4.0.9-26.el8.x86_64.rpm"],
)

rpm(
    name = "libtirpc-0__1.1.4-7.el8.aarch64",
    sha256 = "b8e1ecf3484660710fed69be5b185ad955b8458d5012a71172cd15fbb9001083",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libtirpc-1.1.4-7.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/b8e1ecf3484660710fed69be5b185ad955b8458d5012a71172cd15fbb9001083",
    ],
)

rpm(
    name = "libtirpc-0__1.1.4-8.el8.x86_64",
    sha256 = "bcade31f01063824b3a3e77218caaedd16532413282978c437c82b81c2991e4e",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libtirpc-1.1.4-8.el8.x86_64.rpm"],
)

rpm(
    name = "libtpms-0__0.9.1-0.20211126git1ff6fe1f43.module_el8.6.0__plus__1087__plus__b42c8331.aarch64",
    sha256 = "6e7acc637f2b2697a18419655b9a0d4b9d06b1a7a380c4546de8064c1be035d5",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/libtpms-0.9.1-0.20211126git1ff6fe1f43.module_el8.6.0+1087+b42c8331.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/6e7acc637f2b2697a18419655b9a0d4b9d06b1a7a380c4546de8064c1be035d5",
    ],
)

rpm(
    name = "libtpms-0__0.9.1-1.20211126git1ff6fe1f43.module_el8.7.0__plus__1218__plus__f626c2ff.x86_64",
    sha256 = "22948530ccb9782fb07a6fadbe1904e7c8d9863d6f097d3fb210a7b63d4843fd",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libtpms-0.9.1-1.20211126git1ff6fe1f43.module_el8.7.0+1218+f626c2ff.x86_64.rpm"],
)

rpm(
    name = "libubsan-0__8.5.0-15.el8.aarch64",
    sha256 = "f17b6540d94e217baf503abe38e9ff08132872c7d35c15048e8891fe0cefedb1",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libubsan-8.5.0-15.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/f17b6540d94e217baf503abe38e9ff08132872c7d35c15048e8891fe0cefedb1",
    ],
)

rpm(
    name = "libunistring-0__0.9.9-3.el8.aarch64",
    sha256 = "707429ccb3223628d55097a162cd0d3de1bd00b48800677c1099931b0f019e80",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libunistring-0.9.9-3.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/707429ccb3223628d55097a162cd0d3de1bd00b48800677c1099931b0f019e80",
    ],
)

rpm(
    name = "libunistring-0__0.9.9-3.el8.x86_64",
    sha256 = "20bb189228afa589141d9c9d4ed457729d13c11608305387602d0b00ed0a3093",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libunistring-0.9.9-3.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/20bb189228afa589141d9c9d4ed457729d13c11608305387602d0b00ed0a3093",
    ],
)

rpm(
    name = "libusbx-0__1.0.23-4.el8.aarch64",
    sha256 = "ae797d004f3cafb89773fcc8a3f0d6d046546b7cb3f9741be200d095c637706f",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libusbx-1.0.23-4.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/ae797d004f3cafb89773fcc8a3f0d6d046546b7cb3f9741be200d095c637706f",
    ],
)

rpm(
    name = "libusbx-0__1.0.23-4.el8.x86_64",
    sha256 = "7e704756a93f07feec345a9748204e78994ce06a4667a2ef35b44964ff754306",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libusbx-1.0.23-4.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/7e704756a93f07feec345a9748204e78994ce06a4667a2ef35b44964ff754306",
    ],
)

rpm(
    name = "libutempter-0__1.1.6-14.el8.aarch64",
    sha256 = "8f6d9839a758fdacfdb4b4b0731e8023b8bbb0b633bd32dbf21c2ce85a933a8a",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libutempter-1.1.6-14.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/8f6d9839a758fdacfdb4b4b0731e8023b8bbb0b633bd32dbf21c2ce85a933a8a",
    ],
)

rpm(
    name = "libutempter-0__1.1.6-14.el8.x86_64",
    sha256 = "c8c54c56bff9ca416c3ba6bccac483fb66c81a53d93a19420088715018ed5169",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libutempter-1.1.6-14.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/c8c54c56bff9ca416c3ba6bccac483fb66c81a53d93a19420088715018ed5169",
    ],
)

rpm(
    name = "libuuid-0__2.32.1-35.el8.aarch64",
    sha256 = "fffa3128caf9f988a57a1a6c8ece0e0907571c17f8219721b7e3bfadb0dfa33e",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libuuid-2.32.1-35.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/fffa3128caf9f988a57a1a6c8ece0e0907571c17f8219721b7e3bfadb0dfa33e",
    ],
)

rpm(
    name = "libuuid-0__2.32.1-39.el8.x86_64",
    sha256 = "558002d6a6d0369bd68dd2df750149f99db98b0a981769f0f1b21072bc49d189",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libuuid-2.32.1-39.el8.x86_64.rpm"],
)

rpm(
    name = "libverto-0__0.3.2-2.el8.aarch64",
    sha256 = "1a8478fe342782d95f29253a2845bdb3e88ced25b5e6b029cecc52a43df1932b",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libverto-0.3.2-2.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/1a8478fe342782d95f29253a2845bdb3e88ced25b5e6b029cecc52a43df1932b",
    ],
)

rpm(
    name = "libverto-0__0.3.2-2.el8.x86_64",
    sha256 = "96b8ea32c5e9b3275788525ecbf35fd6ac1ae137754a2857503776512d4db58a",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libverto-0.3.2-2.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/96b8ea32c5e9b3275788525ecbf35fd6ac1ae137754a2857503776512d4db58a",
    ],
)

rpm(
    name = "libvirt-client-0__8.0.0-2.module_el8.6.0__plus__1087__plus__b42c8331.aarch64",
    sha256 = "fd736b99c4910c52e7bffd34532ece859819ea1e4ad2dc616a554fe630eb8d3a",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/libvirt-client-8.0.0-2.module_el8.6.0+1087+b42c8331.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/fd736b99c4910c52e7bffd34532ece859819ea1e4ad2dc616a554fe630eb8d3a",
    ],
)

rpm(
    name = "libvirt-client-0__8.0.0-2.module_el8.6.0__plus__1087__plus__b42c8331.x86_64",
    sha256 = "722f30f8e4a8240662ec03c4bfc1320de88908738ca77fa4fa05e87627821bb1",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libvirt-client-8.0.0-2.module_el8.6.0+1087+b42c8331.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/722f30f8e4a8240662ec03c4bfc1320de88908738ca77fa4fa05e87627821bb1",
    ],
)

rpm(
    name = "libvirt-daemon-0__8.0.0-2.module_el8.6.0__plus__1087__plus__b42c8331.aarch64",
    sha256 = "734437ae41c5c705ab1da476ee9521a57124261727c16f398c8a1bdd8be44922",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/libvirt-daemon-8.0.0-2.module_el8.6.0+1087+b42c8331.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/734437ae41c5c705ab1da476ee9521a57124261727c16f398c8a1bdd8be44922",
    ],
)

rpm(
    name = "libvirt-daemon-0__8.0.0-2.module_el8.6.0__plus__1087__plus__b42c8331.x86_64",
    sha256 = "0429a9e9d8eb98c5ebd689993a3ca8f14949ae45be5a290fce8bbe9c4ad68850",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libvirt-daemon-8.0.0-2.module_el8.6.0+1087+b42c8331.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/0429a9e9d8eb98c5ebd689993a3ca8f14949ae45be5a290fce8bbe9c4ad68850",
    ],
)

rpm(
    name = "libvirt-daemon-driver-qemu-0__8.0.0-2.module_el8.6.0__plus__1087__plus__b42c8331.aarch64",
    sha256 = "77a8a98da56eeaf7cdfe11bdc6b01e42f9eea16b0c04f1abfe7fbafe216a4a66",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/libvirt-daemon-driver-qemu-8.0.0-2.module_el8.6.0+1087+b42c8331.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/77a8a98da56eeaf7cdfe11bdc6b01e42f9eea16b0c04f1abfe7fbafe216a4a66",
    ],
)

rpm(
    name = "libvirt-daemon-driver-qemu-0__8.0.0-2.module_el8.6.0__plus__1087__plus__b42c8331.x86_64",
    sha256 = "d34af964ae21ad21c7e8f97f7f05e2b362744e77270930d8e41a98bced9d91e7",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libvirt-daemon-driver-qemu-8.0.0-2.module_el8.6.0+1087+b42c8331.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/d34af964ae21ad21c7e8f97f7f05e2b362744e77270930d8e41a98bced9d91e7",
    ],
)

rpm(
    name = "libvirt-devel-0__8.0.0-2.module_el8.6.0__plus__1087__plus__b42c8331.aarch64",
    sha256 = "aa47408e4c1499bc03442a6873444ea7d4cd3b62bf59118ff30da2e9db29369f",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/libvirt-devel-8.0.0-2.module_el8.6.0+1087+b42c8331.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/aa47408e4c1499bc03442a6873444ea7d4cd3b62bf59118ff30da2e9db29369f",
    ],
)

rpm(
    name = "libvirt-devel-0__8.0.0-2.module_el8.6.0__plus__1087__plus__b42c8331.x86_64",
    sha256 = "b14d075708f66875be58adb67be1f2ba3b7c1c1c89c87b3656c07b3b6ee03ded",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libvirt-devel-8.0.0-2.module_el8.6.0+1087+b42c8331.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/b14d075708f66875be58adb67be1f2ba3b7c1c1c89c87b3656c07b3b6ee03ded",
    ],
)

rpm(
    name = "libvirt-libs-0__8.0.0-2.module_el8.6.0__plus__1087__plus__b42c8331.aarch64",
    sha256 = "7feb59b591f71783999b5ec9256ef61da19e5e3cdaae46bec162781cdab4b074",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/libvirt-libs-8.0.0-2.module_el8.6.0+1087+b42c8331.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/7feb59b591f71783999b5ec9256ef61da19e5e3cdaae46bec162781cdab4b074",
    ],
)

rpm(
    name = "libvirt-libs-0__8.0.0-2.module_el8.6.0__plus__1087__plus__b42c8331.x86_64",
    sha256 = "ba3daa6361d8b7a0f673840088f81f8aa994f811de1cc95c8c6e1c4baf31ebed",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libvirt-libs-8.0.0-2.module_el8.6.0+1087+b42c8331.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/ba3daa6361d8b7a0f673840088f81f8aa994f811de1cc95c8c6e1c4baf31ebed",
    ],
)

rpm(
    name = "libwayland-client-0__1.21.0-1.el8.x86_64",
    sha256 = "bf1b7055999f0961fcd23fb29d07678c9d6bf1f9c57f42b06b6237b84a3f5aa9",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libwayland-client-1.21.0-1.el8.x86_64.rpm"],
)

rpm(
    name = "libwayland-cursor-0__1.21.0-1.el8.x86_64",
    sha256 = "ed32158e75e2f3decf8089f5de5dbdf21915c881293a795f5e77cfba3d3af403",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libwayland-cursor-1.21.0-1.el8.x86_64.rpm"],
)

rpm(
    name = "libwayland-egl-0__1.21.0-1.el8.x86_64",
    sha256 = "aa7b2f9d27c75f0844bdbcd02c325aafb79756f1b422fd8d6c229afd4c9c79ad",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libwayland-egl-1.21.0-1.el8.x86_64.rpm"],
)

rpm(
    name = "libwayland-server-0__1.21.0-1.el8.x86_64",
    sha256 = "86b1b725f8b725706cbad9d44d0c896a52b249b3e7b556814128dabc03cef023",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libwayland-server-1.21.0-1.el8.x86_64.rpm"],
)

rpm(
    name = "libxcb-0__1.13.1-1.el8.x86_64",
    sha256 = "0221e6e3671c2bd130e9519a7b352404b7e510584b4707d38e1a733e19c7f74f",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libxcb-1.13.1-1.el8.x86_64.rpm"],
)

rpm(
    name = "libxcrypt-0__4.1.1-6.el8.aarch64",
    sha256 = "4948420ee35381c71c619fab4b8deabfa93c04e7c5729620b02e4382a50550ad",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libxcrypt-4.1.1-6.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/4948420ee35381c71c619fab4b8deabfa93c04e7c5729620b02e4382a50550ad",
    ],
)

rpm(
    name = "libxcrypt-0__4.1.1-6.el8.x86_64",
    sha256 = "645853feb85c921d979cb9cf9109663528429eda63cf5a1e31fe578d3d7e713a",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libxcrypt-4.1.1-6.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/645853feb85c921d979cb9cf9109663528429eda63cf5a1e31fe578d3d7e713a",
    ],
)

rpm(
    name = "libxcrypt-devel-0__4.1.1-6.el8.aarch64",
    sha256 = "c561c433a3c295f5d7a49e79a43e4cc96094ed15bcc2fa271bf31f5a6deeacd1",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libxcrypt-devel-4.1.1-6.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/c561c433a3c295f5d7a49e79a43e4cc96094ed15bcc2fa271bf31f5a6deeacd1",
    ],
)

rpm(
    name = "libxcrypt-devel-0__4.1.1-6.el8.x86_64",
    sha256 = "6d84082741a4b7f1a98872a7ee8f12efca835b3dbcb15401aa1b5eccfc674bd4",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libxcrypt-devel-4.1.1-6.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/6d84082741a4b7f1a98872a7ee8f12efca835b3dbcb15401aa1b5eccfc674bd4",
    ],
)

rpm(
    name = "libxcrypt-static-0__4.1.1-6.el8.aarch64",
    sha256 = "a8268856b30e6700f0f67651a6a43449b1e5fccaff512a95280d305468e44dfc",
    urls = [
        "http://mirror.centos.org/centos/8-stream/PowerTools/aarch64/os/Packages/libxcrypt-static-4.1.1-6.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/a8268856b30e6700f0f67651a6a43449b1e5fccaff512a95280d305468e44dfc",
    ],
)

rpm(
    name = "libxcrypt-static-0__4.1.1-6.el8.x86_64",
    sha256 = "599cded5497aa6155c409321f3bb88b7a820341e1d502eac80bf17447283a29b",
    urls = [
        "http://mirror.centos.org/centos/8-stream/PowerTools/x86_64/os/Packages/libxcrypt-static-4.1.1-6.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/599cded5497aa6155c409321f3bb88b7a820341e1d502eac80bf17447283a29b",
    ],
)

rpm(
    name = "libxkbcommon-0__0.9.1-1.el8.aarch64",
    sha256 = "3aca03c788af2ecf8ef39421f246769d7ef7f37260ee9421fc68c1d1cc913600",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/libxkbcommon-0.9.1-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/3aca03c788af2ecf8ef39421f246769d7ef7f37260ee9421fc68c1d1cc913600",
    ],
)

rpm(
    name = "libxkbcommon-0__0.9.1-1.el8.x86_64",
    sha256 = "e03d462995326a4477dcebc8c12eae3c1776ce2f095617ace253c0c492c89082",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libxkbcommon-0.9.1-1.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/e03d462995326a4477dcebc8c12eae3c1776ce2f095617ace253c0c492c89082",
    ],
)

rpm(
    name = "libxml2-0__2.9.7-14.el8.aarch64",
    sha256 = "e9cfed7ab4e4fbce2d0e170b80c1cb3ebf199386350e12a851b2ced13b3b0cc1",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libxml2-2.9.7-14.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/e9cfed7ab4e4fbce2d0e170b80c1cb3ebf199386350e12a851b2ced13b3b0cc1",
    ],
)

rpm(
    name = "libxml2-0__2.9.7-15.el8.x86_64",
    sha256 = "fd99e5a3ef51c11b1380bb3ea1d906a9677032dd80fe3a5fc274e1e9407a8efb",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libxml2-2.9.7-15.el8.x86_64.rpm"],
)

rpm(
    name = "libxshmfence-0__1.3-2.el8.x86_64",
    sha256 = "bfb818e14cfa05d800f1131366ee8fd0c30ab0c735470c870e62dabb7d3f1073",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/libxshmfence-1.3-2.el8.x86_64.rpm"],
)

rpm(
    name = "libzstd-0__1.4.4-1.el8.aarch64",
    sha256 = "b560a8a185100a7c80e6c32f69ba65ce17004156f7218cf183249b15c13295cc",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/libzstd-1.4.4-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/b560a8a185100a7c80e6c32f69ba65ce17004156f7218cf183249b15c13295cc",
    ],
)

rpm(
    name = "libzstd-0__1.4.4-1.el8.x86_64",
    sha256 = "7c2dc6044f13fe4ae04a4c1620da822a6be591b5129bf68ba98a3d8e9092f83b",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/libzstd-1.4.4-1.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/7c2dc6044f13fe4ae04a4c1620da822a6be591b5129bf68ba98a3d8e9092f83b",
    ],
)

rpm(
    name = "llvm-compat-libs-0__14.0.6-1.module_el8.8.0__plus__1224__plus__64629835.x86_64",
    sha256 = "d5b56f06a379eff11206eaec28a263df36cd8afbb833d57f56320badca590b59",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/llvm-compat-libs-14.0.6-1.module_el8.8.0+1224+64629835.x86_64.rpm"],
)

rpm(
    name = "llvm-libs-0__14.0.6-1.module_el8.7.0__plus__1198__plus__0c3eb6e2.x86_64",
    sha256 = "0cadca04a1c5ce79dcddf36329f3022ba02cbda04d0282f1ff550eba8d817527",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/llvm-libs-14.0.6-1.module_el8.7.0+1198+0c3eb6e2.x86_64.rpm"],
)

rpm(
    name = "lua-libs-0__5.3.4-12.el8.aarch64",
    sha256 = "2ef9801e4453de316429be284d4f6cb12f4d7662e7c6224dbf2341e3cfc5fab6",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/lua-libs-5.3.4-12.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/2ef9801e4453de316429be284d4f6cb12f4d7662e7c6224dbf2341e3cfc5fab6",
    ],
)

rpm(
    name = "lua-libs-0__5.3.4-12.el8.x86_64",
    sha256 = "0268af0ee5754fb90fcf71b00fb737f1bf5b3c54c9ff312f13df8c2201311cfe",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/lua-libs-5.3.4-12.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/0268af0ee5754fb90fcf71b00fb737f1bf5b3c54c9ff312f13df8c2201311cfe",
    ],
)

rpm(
    name = "lvm2-8__2.03.14-6.el8.x86_64",
    sha256 = "d66449a34c08cf0d22fae47507c032fa4f51401d4ea6aafc70fa606f3a548019",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/lvm2-2.03.14-6.el8.x86_64.rpm"],
)

rpm(
    name = "lvm2-libs-8__2.03.14-6.el8.x86_64",
    sha256 = "dce1d014dd3107351c1c6918ffd4de4a88fbaebed210c00a4a3f0c1966c3aabf",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/lvm2-libs-2.03.14-6.el8.x86_64.rpm"],
)

rpm(
    name = "lz4-libs-0__1.8.3-3.el8_4.aarch64",
    sha256 = "db9075646bed11355faf8b425c655a40a55436715a9f401f60e205ddd66edfeb",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/lz4-libs-1.8.3-3.el8_4.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/db9075646bed11355faf8b425c655a40a55436715a9f401f60e205ddd66edfeb",
    ],
)

rpm(
    name = "lz4-libs-0__1.8.3-3.el8_4.x86_64",
    sha256 = "8ecac05bb0ec99f91026f2361f7443b9be3272582193a7836884ec473bf8f423",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/lz4-libs-1.8.3-3.el8_4.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/8ecac05bb0ec99f91026f2361f7443b9be3272582193a7836884ec473bf8f423",
    ],
)

rpm(
    name = "lzo-0__2.08-14.el8.aarch64",
    sha256 = "6809839757bd05082ca1b8d23eac617898eda3ce34844a0d31b0a030c8cc6653",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/lzo-2.08-14.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/6809839757bd05082ca1b8d23eac617898eda3ce34844a0d31b0a030c8cc6653",
    ],
)

rpm(
    name = "lzo-0__2.08-14.el8.x86_64",
    sha256 = "5c68635cb03533a38d4a42f6547c21a1d5f9952351bb01f3cf865d2621a6e634",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/lzo-2.08-14.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/5c68635cb03533a38d4a42f6547c21a1d5f9952351bb01f3cf865d2621a6e634",
    ],
)

rpm(
    name = "lzop-0__1.03-20.el8.aarch64",
    sha256 = "003b309833a1ed94ad97ed62f04c2fcda4a20fb8b7b5933c36459974f4e4986c",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/lzop-1.03-20.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/003b309833a1ed94ad97ed62f04c2fcda4a20fb8b7b5933c36459974f4e4986c",
    ],
)

rpm(
    name = "lzop-0__1.03-20.el8.x86_64",
    sha256 = "04eae61018a5be7656be832797016f97cd7b6e19d56f58cb658cd3969dedf2b0",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/lzop-1.03-20.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/04eae61018a5be7656be832797016f97cd7b6e19d56f58cb658cd3969dedf2b0",
    ],
)

rpm(
    name = "mesa-dri-drivers-0__22.3.0-1.el8.x86_64",
    sha256 = "533da3d6b9440a997c8bfb36072ca8cb5a5d880d59d0f249c48bbaf3da77b03c",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/mesa-dri-drivers-22.3.0-1.el8.x86_64.rpm"],
)

rpm(
    name = "mesa-filesystem-0__22.3.0-1.el8.x86_64",
    sha256 = "8621f6e5aa14675722a2c2402fec04f95fa2ab137889593fa017a3d42524e40e",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/mesa-filesystem-22.3.0-1.el8.x86_64.rpm"],
)

rpm(
    name = "mesa-libEGL-0__22.3.0-1.el8.x86_64",
    sha256 = "50ab415cfba02d1d24859ef79a6118d01ffe87dd1cda6b3afb43447847882ce2",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/mesa-libEGL-22.3.0-1.el8.x86_64.rpm"],
)

rpm(
    name = "mesa-libGL-0__22.3.0-1.el8.x86_64",
    sha256 = "34c9c1b4fadc4fdc91fbb05b9c52e2dbd5fc44f014ba2cdd48cd18380717fad1",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/mesa-libGL-22.3.0-1.el8.x86_64.rpm"],
)

rpm(
    name = "mesa-libgbm-0__22.3.0-1.el8.x86_64",
    sha256 = "d23996e6c193cdd4c37971a79dd2b5fbec64ac476ae920b9c49aef91406ba609",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/mesa-libgbm-22.3.0-1.el8.x86_64.rpm"],
)

rpm(
    name = "mesa-libglapi-0__22.3.0-1.el8.x86_64",
    sha256 = "dbdaa8e1d30a19c0e60a2dd4848d170cf3e5714a5f7166ee89035ccbfeb4b5dc",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/mesa-libglapi-22.3.0-1.el8.x86_64.rpm"],
)

rpm(
    name = "mpfr-0__3.1.6-1.el8.aarch64",
    sha256 = "97a998a1b93c21bf070f9a9a1dbb525234b00fccedfe67de8967cd9ec7132eb1",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/mpfr-3.1.6-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/97a998a1b93c21bf070f9a9a1dbb525234b00fccedfe67de8967cd9ec7132eb1",
    ],
)

rpm(
    name = "mpfr-0__3.1.6-1.el8.x86_64",
    sha256 = "e7f0c34f83c1ec2abb22951779e84d51e234c4ba0a05252e4ffd8917461891a5",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/mpfr-3.1.6-1.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/e7f0c34f83c1ec2abb22951779e84d51e234c4ba0a05252e4ffd8917461891a5",
    ],
)

rpm(
    name = "ncurses-0__6.1-9.20180224.el8.x86_64",
    sha256 = "fc22ce73243e2f926e72967c28de57beabfa3720e51248b9a39e40207fbc6c8a",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/ncurses-6.1-9.20180224.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/fc22ce73243e2f926e72967c28de57beabfa3720e51248b9a39e40207fbc6c8a",
    ],
)

rpm(
    name = "ncurses-base-0__6.1-9.20180224.el8.aarch64",
    sha256 = "41716536ea16798238ac89fbc3041b3f9dc80f9a64ea4b19d6e67ad2c909269a",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/ncurses-base-6.1-9.20180224.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/41716536ea16798238ac89fbc3041b3f9dc80f9a64ea4b19d6e67ad2c909269a",
    ],
)

rpm(
    name = "ncurses-base-0__6.1-9.20180224.el8.x86_64",
    sha256 = "41716536ea16798238ac89fbc3041b3f9dc80f9a64ea4b19d6e67ad2c909269a",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/ncurses-base-6.1-9.20180224.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/41716536ea16798238ac89fbc3041b3f9dc80f9a64ea4b19d6e67ad2c909269a",
    ],
)

rpm(
    name = "ncurses-libs-0__6.1-9.20180224.el8.aarch64",
    sha256 = "b938a6facc8d8a3de12b369871738bb531c822b1ec5212501b06bcaaf6cd25fa",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/ncurses-libs-6.1-9.20180224.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/b938a6facc8d8a3de12b369871738bb531c822b1ec5212501b06bcaaf6cd25fa",
    ],
)

rpm(
    name = "ncurses-libs-0__6.1-9.20180224.el8.x86_64",
    sha256 = "54609dd070a57a14a6103f0c06bea99bb0a4e568d1fbc6a22b8ba67c954d90bf",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/ncurses-libs-6.1-9.20180224.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/54609dd070a57a14a6103f0c06bea99bb0a4e568d1fbc6a22b8ba67c954d90bf",
    ],
)

rpm(
    name = "ndctl-libs-0__71.1-4.el8.x86_64",
    sha256 = "d1518d8f29a72c8c9501f67929258405cf25fd4be365fd905acc57b846d49c8a",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/ndctl-libs-71.1-4.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/d1518d8f29a72c8c9501f67929258405cf25fd4be365fd905acc57b846d49c8a",
    ],
)

rpm(
    name = "nettle-0__3.4.1-7.el8.aarch64",
    sha256 = "5441222132ae52cd31063e9b9e3bb40f2e5711dfb0c84315b4aec2907278a075",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/nettle-3.4.1-7.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/5441222132ae52cd31063e9b9e3bb40f2e5711dfb0c84315b4aec2907278a075",
    ],
)

rpm(
    name = "nettle-0__3.4.1-7.el8.x86_64",
    sha256 = "fe9a848502c595e0b7acc699d69c24b9c5ad0ac58a0b3933cd228f3633de31cb",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/nettle-3.4.1-7.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/fe9a848502c595e0b7acc699d69c24b9c5ad0ac58a0b3933cd228f3633de31cb",
    ],
)

rpm(
    name = "nftables-1__0.9.3-26.el8.aarch64",
    sha256 = "22cacdb52fb6a31659789b5190f8e6db27ca1dddd9b67f3c6b2c1db917ef882f",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/nftables-0.9.3-26.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/22cacdb52fb6a31659789b5190f8e6db27ca1dddd9b67f3c6b2c1db917ef882f",
    ],
)

rpm(
    name = "nftables-1__0.9.3-26.el8.x86_64",
    sha256 = "813d7c361e77b394f6f05fb29983c3ee6c2dd2e8fe8b857e2bdb6b9914e0c129",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/nftables-0.9.3-26.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/813d7c361e77b394f6f05fb29983c3ee6c2dd2e8fe8b857e2bdb6b9914e0c129",
    ],
)

rpm(
    name = "nmap-ncat-2__7.70-7.el8.aarch64",
    sha256 = "a1da5a886825ef4cb2342db821cf1da3133860a99fd436b113affbd0e8dd9f81",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/nmap-ncat-7.70-7.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/a1da5a886825ef4cb2342db821cf1da3133860a99fd436b113affbd0e8dd9f81",
    ],
)

rpm(
    name = "nmap-ncat-2__7.70-8.el8.x86_64",
    sha256 = "01f8398a2bcb3b258bc51f219ec7d3fb9c408c91170659919f136edea2b1cc32",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/nmap-ncat-7.70-8.el8.x86_64.rpm"],
)

rpm(
    name = "npth-0__1.5-4.el8.x86_64",
    sha256 = "168ab5dbc86b836b8742b2e63eee51d074f1d790728e3d30b0c59fff93cf1d8d",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/npth-1.5-4.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/168ab5dbc86b836b8742b2e63eee51d074f1d790728e3d30b0c59fff93cf1d8d",
    ],
)

rpm(
    name = "numactl-libs-0__2.0.12-13.el8.aarch64",
    sha256 = "5f2d7a8db99ad318df35e60d43e5e7f462294c00ffa3d7c24207c16bfd3a6619",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/numactl-libs-2.0.12-13.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/5f2d7a8db99ad318df35e60d43e5e7f462294c00ffa3d7c24207c16bfd3a6619",
    ],
)

rpm(
    name = "numactl-libs-0__2.0.12-13.el8.x86_64",
    sha256 = "b7b71ba34b3af893dc0acbb9d2228a2307da849d38e1c0007bd3d64f456640af",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/numactl-libs-2.0.12-13.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/b7b71ba34b3af893dc0acbb9d2228a2307da849d38e1c0007bd3d64f456640af",
    ],
)

rpm(
    name = "numad-0__0.5-26.20150602git.el8.aarch64",
    sha256 = "5b580f1a1c2193384a7c4c5171200d1e6f4ca6a19e6a01a327a75d03db916484",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/numad-0.5-26.20150602git.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/5b580f1a1c2193384a7c4c5171200d1e6f4ca6a19e6a01a327a75d03db916484",
    ],
)

rpm(
    name = "numad-0__0.5-26.20150602git.el8.x86_64",
    sha256 = "5d975c08273b1629683275c32f16e52ca8e37e6836598e211092c915d38878bf",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/numad-0.5-26.20150602git.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/5d975c08273b1629683275c32f16e52ca8e37e6836598e211092c915d38878bf",
    ],
)

rpm(
    name = "open-sans-fonts-0__1.10-6.el8.x86_64",
    sha256 = "ce7ef2eb3a548c0f137d10bed6799f855add5d76af08abb1dd7607d0a1f9c1ac",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/open-sans-fonts-1.10-6.el8.noarch.rpm"],
)

rpm(
    name = "openldap-0__2.4.46-18.el8.aarch64",
    sha256 = "254200cc7c35fefbeab3de24c36f94dec10f913ea2199b6d6c769f0fc8a10546",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/openldap-2.4.46-18.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/254200cc7c35fefbeab3de24c36f94dec10f913ea2199b6d6c769f0fc8a10546",
    ],
)

rpm(
    name = "openldap-0__2.4.46-18.el8.x86_64",
    sha256 = "95327d6c83a370a12c125767403496435d20a94b70ee395eabfc356270d2ada9",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/openldap-2.4.46-18.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/95327d6c83a370a12c125767403496435d20a94b70ee395eabfc356270d2ada9",
    ],
)

rpm(
    name = "openssl-libs-1__1.1.1k-6.el8.aarch64",
    sha256 = "2e0099f5ad83b66f4a8287b6f1c40bd2d7265a035d10075d4538688e75713c99",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/openssl-libs-1.1.1k-6.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/2e0099f5ad83b66f4a8287b6f1c40bd2d7265a035d10075d4538688e75713c99",
    ],
)

rpm(
    name = "openssl-libs-1__1.1.1k-7.el8.x86_64",
    sha256 = "7b42ba3855f29955fe204ad7c189a832a5b1423a32abcda079d8ef2f787c8e73",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/openssl-libs-1.1.1k-7.el8.x86_64.rpm"],
)

rpm(
    name = "p11-kit-0__0.23.22-1.el8.aarch64",
    sha256 = "cfee10a5ca5613896a4e84716aa393094fd97c09f2c585c9aa921e6063783867",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/p11-kit-0.23.22-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/cfee10a5ca5613896a4e84716aa393094fd97c09f2c585c9aa921e6063783867",
    ],
)

rpm(
    name = "p11-kit-0__0.23.22-1.el8.x86_64",
    sha256 = "6a67c8721fe24af25ec56c6aae956a190d8463e46efed45adfbbd800086550c7",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/p11-kit-0.23.22-1.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/6a67c8721fe24af25ec56c6aae956a190d8463e46efed45adfbbd800086550c7",
    ],
)

rpm(
    name = "p11-kit-trust-0__0.23.22-1.el8.aarch64",
    sha256 = "3fc181bf0f076fef283fdb63d36e7b84930c8822fa67dff6e1ccea9987d6dbf3",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/p11-kit-trust-0.23.22-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/3fc181bf0f076fef283fdb63d36e7b84930c8822fa67dff6e1ccea9987d6dbf3",
    ],
)

rpm(
    name = "p11-kit-trust-0__0.23.22-1.el8.x86_64",
    sha256 = "d218619a4859e002fe677703bc1767986314cd196ae2ac397ed057f3bec36516",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/p11-kit-trust-0.23.22-1.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/d218619a4859e002fe677703bc1767986314cd196ae2ac397ed057f3bec36516",
    ],
)

rpm(
    name = "pam-0__1.3.1-22.el8.aarch64",
    sha256 = "b900edf1f702460be4a6b2e402e02887068fe9172b88256660b8c20b89a772d5",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/pam-1.3.1-22.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/b900edf1f702460be4a6b2e402e02887068fe9172b88256660b8c20b89a772d5",
    ],
)

rpm(
    name = "pam-0__1.3.1-25.el8.x86_64",
    sha256 = "1dd647b181f70dfa8a3e742a9942f3b134c17a721f890057b756691f2389333c",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/pam-1.3.1-25.el8.x86_64.rpm"],
)

rpm(
    name = "pango-0__1.42.4-8.el8.x86_64",
    sha256 = "1e74c391edf2f383b5c236e65ddd15bcf83883975b8d08b70808d2e14916d496",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/pango-1.42.4-8.el8.x86_64.rpm"],
)

rpm(
    name = "passt-0__0.git.2022_08_04.b516d15-0.el8.aarch64",
    sha256 = "b5e57c05a358097536196e7eb6156f5f9205ea0900d0096678d141a6c7029250",
    urls = [
        "https://download.copr.fedorainfracloud.org/results/sbrivio/passt/centos-stream-8-aarch64/04703016-passt/passt-0.git.2022_08_04.b516d15-0.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/b5e57c05a358097536196e7eb6156f5f9205ea0900d0096678d141a6c7029250",
    ],
)

rpm(
    name = "pcre-0__8.42-6.el8.aarch64",
    sha256 = "5591faa4f51dc97067292938883b771d75ec2b3a749ec956eddc0408e689c369",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/pcre-8.42-6.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/5591faa4f51dc97067292938883b771d75ec2b3a749ec956eddc0408e689c369",
    ],
)

rpm(
    name = "pcre-0__8.42-6.el8.x86_64",
    sha256 = "876e9e99b0e50cb2752499045bafa903dd29e5c491d112daacef1ae16f614dad",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/pcre-8.42-6.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/876e9e99b0e50cb2752499045bafa903dd29e5c491d112daacef1ae16f614dad",
    ],
)

rpm(
    name = "pcre2-0__10.32-3.el8.aarch64",
    sha256 = "b8e4367f28a53ec70a6b8a329a5bda886374eddde5f55c9467e1783d4158b5d1",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/pcre2-10.32-3.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/b8e4367f28a53ec70a6b8a329a5bda886374eddde5f55c9467e1783d4158b5d1",
    ],
)

rpm(
    name = "pcre2-0__10.32-3.el8.x86_64",
    sha256 = "2f865747024d26b91d5a9f2f35dd1b04e1039d64e772d0371b437145cd7beceb",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/pcre2-10.32-3.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/2f865747024d26b91d5a9f2f35dd1b04e1039d64e772d0371b437145cd7beceb",
    ],
)

rpm(
    name = "perl-Carp-0__1.42-396.el8.x86_64",
    sha256 = "d03b9f4b9848e3a88d62bcf6e536d659c325b2dc03b2136be7342b5fe5e2b6a9",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-Carp-1.42-396.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/d03b9f4b9848e3a88d62bcf6e536d659c325b2dc03b2136be7342b5fe5e2b6a9",
    ],
)

rpm(
    name = "perl-Encode-4__2.97-3.el8.x86_64",
    sha256 = "d2b0e4b28a5aac754f6caa119d5479a64816f93c059e0ac564e46391264e2234",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-Encode-2.97-3.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/d2b0e4b28a5aac754f6caa119d5479a64816f93c059e0ac564e46391264e2234",
    ],
)

rpm(
    name = "perl-Errno-0__1.28-421.el8.x86_64",
    sha256 = "8d9b26f17e427dc497032b1897b9296c4ca37fa1b96d9c459b42516d72ef06a1",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-Errno-1.28-421.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/8d9b26f17e427dc497032b1897b9296c4ca37fa1b96d9c459b42516d72ef06a1",
    ],
)

rpm(
    name = "perl-Exporter-0__5.72-396.el8.x86_64",
    sha256 = "7edc503f5a919c489b651757095d8031982d530cc88088fdaeb743188364e9b0",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-Exporter-5.72-396.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/7edc503f5a919c489b651757095d8031982d530cc88088fdaeb743188364e9b0",
    ],
)

rpm(
    name = "perl-File-Path-0__2.15-2.el8.x86_64",
    sha256 = "e83928bd4552ecdf8e71d283e2358c7eccd006d284ba31fbc9c89e407989fd60",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-File-Path-2.15-2.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/e83928bd4552ecdf8e71d283e2358c7eccd006d284ba31fbc9c89e407989fd60",
    ],
)

rpm(
    name = "perl-File-Temp-0__0.230.600-1.el8.x86_64",
    sha256 = "e269f7d33abbb790311ffa95fa7df9766cac8bf31ace24fce6ed732ba0db19ae",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-File-Temp-0.230.600-1.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/e269f7d33abbb790311ffa95fa7df9766cac8bf31ace24fce6ed732ba0db19ae",
    ],
)

rpm(
    name = "perl-Getopt-Long-1__2.50-4.el8.x86_64",
    sha256 = "da4c6daa0d5406bc967cc89b02a69689491f42c543aceea1a31136f0f1a8d991",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-Getopt-Long-2.50-4.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/da4c6daa0d5406bc967cc89b02a69689491f42c543aceea1a31136f0f1a8d991",
    ],
)

rpm(
    name = "perl-HTTP-Tiny-0__0.074-1.el8.x86_64",
    sha256 = "a1af93a1b62e8ca05b7597d5749a2b3d28735a86928f0432064fec61db1ff844",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-HTTP-Tiny-0.074-1.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/a1af93a1b62e8ca05b7597d5749a2b3d28735a86928f0432064fec61db1ff844",
    ],
)

rpm(
    name = "perl-IO-0__1.38-421.el8.x86_64",
    sha256 = "7ff911df218c38953660d4a09f9864364e2433b9aaf8283db8b7d5214411e28a",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-IO-1.38-421.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/7ff911df218c38953660d4a09f9864364e2433b9aaf8283db8b7d5214411e28a",
    ],
)

rpm(
    name = "perl-MIME-Base64-0__3.15-396.el8.x86_64",
    sha256 = "5642297bf32bb174173917dd10fd2a3a2ef7277c599f76c0669c5c448f10bdaf",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-MIME-Base64-3.15-396.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/5642297bf32bb174173917dd10fd2a3a2ef7277c599f76c0669c5c448f10bdaf",
    ],
)

rpm(
    name = "perl-PathTools-0__3.74-1.el8.x86_64",
    sha256 = "512245f7741790b36b03562469b9262f4dedfb8862dfa2d42e64598bb205d4c9",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-PathTools-3.74-1.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/512245f7741790b36b03562469b9262f4dedfb8862dfa2d42e64598bb205d4c9",
    ],
)

rpm(
    name = "perl-Pod-Escapes-1__1.07-395.el8.x86_64",
    sha256 = "545cd23ad8e4f71a5109551093668fd4b5e1a50d6a60364ce0f04f64eecd99d1",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-Pod-Escapes-1.07-395.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/545cd23ad8e4f71a5109551093668fd4b5e1a50d6a60364ce0f04f64eecd99d1",
    ],
)

rpm(
    name = "perl-Pod-Perldoc-0__3.28-396.el8.x86_64",
    sha256 = "0225dc3999e3d7b1bb57186a2fc93c98bd1e4e08e062fb51c966e1f2a2c91bb4",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-Pod-Perldoc-3.28-396.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/0225dc3999e3d7b1bb57186a2fc93c98bd1e4e08e062fb51c966e1f2a2c91bb4",
    ],
)

rpm(
    name = "perl-Pod-Simple-1__3.35-395.el8.x86_64",
    sha256 = "51c3ee5d824bdde0a8faa10c99841c2590c0c26edfb17125aa97945a688c83ed",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-Pod-Simple-3.35-395.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/51c3ee5d824bdde0a8faa10c99841c2590c0c26edfb17125aa97945a688c83ed",
    ],
)

rpm(
    name = "perl-Pod-Usage-4__1.69-395.el8.x86_64",
    sha256 = "794f970f498af07b37f914c19ad5dedc6b6c2f89d343af9dd1768d17232555de",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-Pod-Usage-1.69-395.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/794f970f498af07b37f914c19ad5dedc6b6c2f89d343af9dd1768d17232555de",
    ],
)

rpm(
    name = "perl-Scalar-List-Utils-3__1.49-2.el8.x86_64",
    sha256 = "3db0d05ca5ba00981312f3a3ddcbabf466c2f1fc639cbf29482bb2cd952df456",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-Scalar-List-Utils-1.49-2.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/3db0d05ca5ba00981312f3a3ddcbabf466c2f1fc639cbf29482bb2cd952df456",
    ],
)

rpm(
    name = "perl-Socket-4__2.027-3.el8.x86_64",
    sha256 = "de138a9614191af63b9603cf0912d4ffd9bd9e5b122c2d0a78ae0eac009a602f",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-Socket-2.027-3.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/de138a9614191af63b9603cf0912d4ffd9bd9e5b122c2d0a78ae0eac009a602f",
    ],
)

rpm(
    name = "perl-Storable-1__3.11-3.el8.x86_64",
    sha256 = "0c3007b68a37325866aaade4ae076232bca15e268f66c3d3b3a6d236bb85e1e9",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-Storable-3.11-3.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/0c3007b68a37325866aaade4ae076232bca15e268f66c3d3b3a6d236bb85e1e9",
    ],
)

rpm(
    name = "perl-Sys-Guestfs-1__1.44.0-5.module_el8.6.0__plus__1087__plus__b42c8331.x86_64",
    sha256 = "8e01d8cca7a1297980a36db1b56835cce506c08450d12b7b21e11bfa58ad22bb",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/perl-Sys-Guestfs-1.44.0-5.module_el8.6.0+1087+b42c8331.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/8e01d8cca7a1297980a36db1b56835cce506c08450d12b7b21e11bfa58ad22bb",
    ],
)

rpm(
    name = "perl-Term-ANSIColor-0__4.06-396.el8.x86_64",
    sha256 = "f4e3607f242bbca7ec2379822ca961860e6d9c276da51c6e2dfd17a29469ec78",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-Term-ANSIColor-4.06-396.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/f4e3607f242bbca7ec2379822ca961860e6d9c276da51c6e2dfd17a29469ec78",
    ],
)

rpm(
    name = "perl-Term-Cap-0__1.17-395.el8.x86_64",
    sha256 = "6bbb721dd2c411c85c75f7477b14c54c776d78ee9b93557615e919ef47577440",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-Term-Cap-1.17-395.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/6bbb721dd2c411c85c75f7477b14c54c776d78ee9b93557615e919ef47577440",
    ],
)

rpm(
    name = "perl-Text-ParseWords-0__3.30-395.el8.x86_64",
    sha256 = "2975de6545b4ca7907ae368a1716c531764e4afccbf27fb0a694d90e983c38e2",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-Text-ParseWords-3.30-395.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/2975de6545b4ca7907ae368a1716c531764e4afccbf27fb0a694d90e983c38e2",
    ],
)

rpm(
    name = "perl-Text-Tabs__plus__Wrap-0__2013.0523-395.el8.x86_64",
    sha256 = "7e50a5d0f2fbd8c95375f72f5772c7731186e999a447121b8247f448b065a4ef",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-Text-Tabs+Wrap-2013.0523-395.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/7e50a5d0f2fbd8c95375f72f5772c7731186e999a447121b8247f448b065a4ef",
    ],
)

rpm(
    name = "perl-Time-Local-1__1.280-1.el8.x86_64",
    sha256 = "1edcf2b441ddf21417ef2b33e1ab2a30900758819335d7fabafe3b16bb3eab62",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-Time-Local-1.280-1.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/1edcf2b441ddf21417ef2b33e1ab2a30900758819335d7fabafe3b16bb3eab62",
    ],
)

rpm(
    name = "perl-Unicode-Normalize-0__1.25-396.el8.x86_64",
    sha256 = "99678a57c35343d8b2e2a502efcccc17bde3e40d97d7d2c5f988af8d3aa166d0",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-Unicode-Normalize-1.25-396.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/99678a57c35343d8b2e2a502efcccc17bde3e40d97d7d2c5f988af8d3aa166d0",
    ],
)

rpm(
    name = "perl-constant-0__1.33-396.el8.x86_64",
    sha256 = "7559c097998db5e5d14dab1a7a1637a5749e9dab234ca68d17c9c21f8cfbf8d6",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-constant-1.33-396.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/7559c097998db5e5d14dab1a7a1637a5749e9dab234ca68d17c9c21f8cfbf8d6",
    ],
)

rpm(
    name = "perl-hivex-0__1.3.18-23.module_el8.6.0__plus__983__plus__a7505f3f.x86_64",
    sha256 = "42db01e9df5ba75147ad2a0cfb37f5f6c37ae980260d218dc93a0ead8cab7983",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/perl-hivex-1.3.18-23.module_el8.6.0+983+a7505f3f.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/42db01e9df5ba75147ad2a0cfb37f5f6c37ae980260d218dc93a0ead8cab7983",
    ],
)

rpm(
    name = "perl-interpreter-4__5.26.3-421.el8.x86_64",
    sha256 = "4618427acf4bcfa66ec91cccf995d938e1ed0f87b1088d7d948a9993a6d15b29",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-interpreter-5.26.3-421.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/4618427acf4bcfa66ec91cccf995d938e1ed0f87b1088d7d948a9993a6d15b29",
    ],
)

rpm(
    name = "perl-libintl-perl-0__1.29-2.el8.x86_64",
    sha256 = "8b8c1ce375e1d8dd73f905e99bd452243ec194dd707a36fa5bdea7a252165c60",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/perl-libintl-perl-1.29-2.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/8b8c1ce375e1d8dd73f905e99bd452243ec194dd707a36fa5bdea7a252165c60",
    ],
)

rpm(
    name = "perl-libs-4__5.26.3-421.el8.x86_64",
    sha256 = "d3a5510385cd4b2d53d70942e4fb4c149917aac2ce2df881c28ae2afdcd26619",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-libs-5.26.3-421.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/d3a5510385cd4b2d53d70942e4fb4c149917aac2ce2df881c28ae2afdcd26619",
    ],
)

rpm(
    name = "perl-macros-4__5.26.3-421.el8.x86_64",
    sha256 = "5969bb5bd8b28a6cead135cfbdae89ac60f649b29f88a1daac3016eea47dc45b",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-macros-5.26.3-421.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/5969bb5bd8b28a6cead135cfbdae89ac60f649b29f88a1daac3016eea47dc45b",
    ],
)

rpm(
    name = "perl-parent-1__0.237-1.el8.x86_64",
    sha256 = "f5e73bbd776a2426a796971d8d38664f2e94898479fb76947dccdd28cf9fe1d0",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-parent-0.237-1.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/f5e73bbd776a2426a796971d8d38664f2e94898479fb76947dccdd28cf9fe1d0",
    ],
)

rpm(
    name = "perl-podlators-0__4.11-1.el8.x86_64",
    sha256 = "78d17ed089151e7fa3d1a3cdbbac8ca3b1b5c484fae5ba025642cc9107991037",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-podlators-4.11-1.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/78d17ed089151e7fa3d1a3cdbbac8ca3b1b5c484fae5ba025642cc9107991037",
    ],
)

rpm(
    name = "perl-threads-1__2.21-2.el8.x86_64",
    sha256 = "2e3da17b1c1685edea9c52bdaa0d77c019d6144c765fc6b3b1c783d98f634f96",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-threads-2.21-2.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/2e3da17b1c1685edea9c52bdaa0d77c019d6144c765fc6b3b1c783d98f634f96",
    ],
)

rpm(
    name = "perl-threads-shared-0__1.58-2.el8.x86_64",
    sha256 = "b4a14dc0e3550da946d7ca65e54d19fc805e30c6c3dbf5ef3fc077d1d94e6d71",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/perl-threads-shared-1.58-2.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/b4a14dc0e3550da946d7ca65e54d19fc805e30c6c3dbf5ef3fc077d1d94e6d71",
    ],
)

rpm(
    name = "pixman-0__0.38.4-2.el8.aarch64",
    sha256 = "038eba8224034c5090cd08184c68a25ff8037dee804ad3eae0109a1cf4096078",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/pixman-0.38.4-2.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/038eba8224034c5090cd08184c68a25ff8037dee804ad3eae0109a1cf4096078",
    ],
)

rpm(
    name = "pixman-0__0.38.4-2.el8.x86_64",
    sha256 = "e496740940bd0b4d6f6537feaaffff57580624f6629c736c7f5e415259dc6cbe",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/pixman-0.38.4-2.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/e496740940bd0b4d6f6537feaaffff57580624f6629c736c7f5e415259dc6cbe",
    ],
)

rpm(
    name = "pkgconf-0__1.4.2-1.el8.aarch64",
    sha256 = "9a2c046a45d46e681f417f3b438d4bb5a21e1b93deacb59d906b8aa08a7535ad",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/pkgconf-1.4.2-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/9a2c046a45d46e681f417f3b438d4bb5a21e1b93deacb59d906b8aa08a7535ad",
    ],
)

rpm(
    name = "pkgconf-0__1.4.2-1.el8.x86_64",
    sha256 = "dd08de48d25573f0a8492cf858ce8c37abb10eb560975d9df0e45a7f91b3b41d",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/pkgconf-1.4.2-1.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/dd08de48d25573f0a8492cf858ce8c37abb10eb560975d9df0e45a7f91b3b41d",
    ],
)

rpm(
    name = "pkgconf-m4-0__1.4.2-1.el8.aarch64",
    sha256 = "56187f25e8ae7c2a5ce228d13c6e93b9c6a701960d61dff8ad720a8879b6059e",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/pkgconf-m4-1.4.2-1.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/56187f25e8ae7c2a5ce228d13c6e93b9c6a701960d61dff8ad720a8879b6059e",
    ],
)

rpm(
    name = "pkgconf-m4-0__1.4.2-1.el8.x86_64",
    sha256 = "56187f25e8ae7c2a5ce228d13c6e93b9c6a701960d61dff8ad720a8879b6059e",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/pkgconf-m4-1.4.2-1.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/56187f25e8ae7c2a5ce228d13c6e93b9c6a701960d61dff8ad720a8879b6059e",
    ],
)

rpm(
    name = "pkgconf-pkg-config-0__1.4.2-1.el8.aarch64",
    sha256 = "aadca7b635ac2b30c3463a4edfe38eaee2c6064181cb090694619186747f3950",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/pkgconf-pkg-config-1.4.2-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/aadca7b635ac2b30c3463a4edfe38eaee2c6064181cb090694619186747f3950",
    ],
)

rpm(
    name = "pkgconf-pkg-config-0__1.4.2-1.el8.x86_64",
    sha256 = "bf5319e42dbe96c24cd64c974b17f422847cc658c4461d9d61cfe76ad76e9c67",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/pkgconf-pkg-config-1.4.2-1.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/bf5319e42dbe96c24cd64c974b17f422847cc658c4461d9d61cfe76ad76e9c67",
    ],
)

rpm(
    name = "platform-python-0__3.6.8-47.el8.aarch64",
    sha256 = "43ffa547514ccad75bc69b6fdc402cc133234b33da4a62ddacc3c51ebf738fd0",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/platform-python-3.6.8-47.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/43ffa547514ccad75bc69b6fdc402cc133234b33da4a62ddacc3c51ebf738fd0",
    ],
)

rpm(
    name = "platform-python-0__3.6.8-49.el8.x86_64",
    sha256 = "039a83cb2d6a5cdf8d002809a701b684459f31b6ecd8011571116c1a641ee7e2",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/platform-python-3.6.8-49.el8.x86_64.rpm"],
)

rpm(
    name = "platform-python-pip-0__9.0.3-22.el8.aarch64",
    sha256 = "f66c6d22a96febc3907247a6350097cceeaf77abcb628574052dfdb1a4411607",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/platform-python-pip-9.0.3-22.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/f66c6d22a96febc3907247a6350097cceeaf77abcb628574052dfdb1a4411607",
    ],
)

rpm(
    name = "platform-python-pip-0__9.0.3-22.el8.x86_64",
    sha256 = "f66c6d22a96febc3907247a6350097cceeaf77abcb628574052dfdb1a4411607",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/platform-python-pip-9.0.3-22.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/f66c6d22a96febc3907247a6350097cceeaf77abcb628574052dfdb1a4411607",
    ],
)

rpm(
    name = "platform-python-setuptools-0__39.2.0-6.el8.aarch64",
    sha256 = "946ba273a3a3b6fdf140f3c03112918c0a556a5871c477f5dbbb98600e6ca557",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/platform-python-setuptools-39.2.0-6.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/946ba273a3a3b6fdf140f3c03112918c0a556a5871c477f5dbbb98600e6ca557",
    ],
)

rpm(
    name = "platform-python-setuptools-0__39.2.0-6.el8.x86_64",
    sha256 = "946ba273a3a3b6fdf140f3c03112918c0a556a5871c477f5dbbb98600e6ca557",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/platform-python-setuptools-39.2.0-6.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/946ba273a3a3b6fdf140f3c03112918c0a556a5871c477f5dbbb98600e6ca557",
    ],
)

rpm(
    name = "policycoreutils-0__2.9-19.el8.aarch64",
    sha256 = "84fdd1ce718c9893444c3dc8ae3f37ed243764623d984767d9136e11e01189d2",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/policycoreutils-2.9-19.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/84fdd1ce718c9893444c3dc8ae3f37ed243764623d984767d9136e11e01189d2",
    ],
)

rpm(
    name = "policycoreutils-0__2.9-20.el8.x86_64",
    sha256 = "341b432d82c58ba14392e38d0ab9aa1e9686a18d0be72491832e6dc697400b17",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/policycoreutils-2.9-20.el8.x86_64.rpm"],
)

rpm(
    name = "polkit-0__0.115-13.0.1.el8.2.aarch64",
    sha256 = "eef4d3b177ff36c7f1781fcb456bef44169484a29f5931f268486f15933e4b24",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/polkit-0.115-13.0.1.el8.2.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/eef4d3b177ff36c7f1781fcb456bef44169484a29f5931f268486f15933e4b24",
    ],
)

rpm(
    name = "polkit-0__0.115-13.0.1.el8.2.x86_64",
    sha256 = "8bfccf9235747eb132c1d10c2f26b5544a0db078019eb7911b88522131e16dc8",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/polkit-0.115-13.0.1.el8.2.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/8bfccf9235747eb132c1d10c2f26b5544a0db078019eb7911b88522131e16dc8",
    ],
)

rpm(
    name = "polkit-libs-0__0.115-13.0.1.el8.2.aarch64",
    sha256 = "dc74d77dfeb155b2708820c9a1d5cbb2c4c29de2c3a1cb76d0987e6bbbf40c9a",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/polkit-libs-0.115-13.0.1.el8.2.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/dc74d77dfeb155b2708820c9a1d5cbb2c4c29de2c3a1cb76d0987e6bbbf40c9a",
    ],
)

rpm(
    name = "polkit-libs-0__0.115-13.0.1.el8.2.x86_64",
    sha256 = "d957da6b452f7b15830ad9a73176d4f04d9c3e26e119b7f3f4f4060087bb9082",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/polkit-libs-0.115-13.0.1.el8.2.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/d957da6b452f7b15830ad9a73176d4f04d9c3e26e119b7f3f4f4060087bb9082",
    ],
)

rpm(
    name = "polkit-pkla-compat-0__0.1-12.el8.aarch64",
    sha256 = "d25d562fe77f391458903ebf0d9078b6d38af6d9ced39d902b9afc7e717d2234",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/polkit-pkla-compat-0.1-12.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/d25d562fe77f391458903ebf0d9078b6d38af6d9ced39d902b9afc7e717d2234",
    ],
)

rpm(
    name = "polkit-pkla-compat-0__0.1-12.el8.x86_64",
    sha256 = "e7ee4b6d6456cb7da0332f5a6fb8a7c47df977bcf616f12f0455413765367e89",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/polkit-pkla-compat-0.1-12.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/e7ee4b6d6456cb7da0332f5a6fb8a7c47df977bcf616f12f0455413765367e89",
    ],
)

rpm(
    name = "popt-0__1.18-1.el8.aarch64",
    sha256 = "2596d6cba62bf9594e4fbb07df31e2459eb6fca8e479fd0be2b32c7561e9ad95",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/popt-1.18-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/2596d6cba62bf9594e4fbb07df31e2459eb6fca8e479fd0be2b32c7561e9ad95",
    ],
)

rpm(
    name = "popt-0__1.18-1.el8.x86_64",
    sha256 = "3fc009f00388e66befab79be548ff3c7aa80ca70bd7f183d22f59137d8e2c2ae",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/popt-1.18-1.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/3fc009f00388e66befab79be548ff3c7aa80ca70bd7f183d22f59137d8e2c2ae",
    ],
)

rpm(
    name = "procps-ng-0__3.3.15-11.el8.x86_64",
    sha256 = "c051ea0fae01a366ff0625e59d27037a88bdb7d1640e91dc8dc5c531275bd831",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/procps-ng-3.3.15-11.el8.x86_64.rpm"],
)

rpm(
    name = "procps-ng-0__3.3.15-8.el8.aarch64",
    sha256 = "624a7d0fae4a4d7d3988726b8a435fe377e8c531b5e1024b1fa9cc3dc1cece2d",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/procps-ng-3.3.15-8.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/624a7d0fae4a4d7d3988726b8a435fe377e8c531b5e1024b1fa9cc3dc1cece2d",
    ],
)

rpm(
    name = "psmisc-0__23.1-5.el8.aarch64",
    sha256 = "e6852f9e715174c037c57ef9ee45a6318775968322c244185fc51f40a10dbdcc",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/psmisc-23.1-5.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/e6852f9e715174c037c57ef9ee45a6318775968322c244185fc51f40a10dbdcc",
    ],
)

rpm(
    name = "psmisc-0__23.1-5.el8.x86_64",
    sha256 = "9d433d8c058e59c891c0852b95b3b87795ea30a85889c77ba0b12f965517d626",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/psmisc-23.1-5.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/9d433d8c058e59c891c0852b95b3b87795ea30a85889c77ba0b12f965517d626",
    ],
)

rpm(
    name = "python3-libs-0__3.6.8-47.el8.aarch64",
    sha256 = "1ec95b8b8d4e226558d193bd46d3e928c143e41e5c0403a8868f872f7a7d2ad1",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/python3-libs-3.6.8-47.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/1ec95b8b8d4e226558d193bd46d3e928c143e41e5c0403a8868f872f7a7d2ad1",
    ],
)

rpm(
    name = "python3-libs-0__3.6.8-49.el8.x86_64",
    sha256 = "7ab22745f91c1832bf6a703fa2ff95dd5126949231a2608f3938e4e9359988fd",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/python3-libs-3.6.8-49.el8.x86_64.rpm"],
)

rpm(
    name = "python3-pip-0__9.0.3-22.el8.aarch64",
    sha256 = "ba83ca7667c98d265da7334a3ef7f786fbb48c85e32cdec11979778594750953",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/python3-pip-9.0.3-22.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/ba83ca7667c98d265da7334a3ef7f786fbb48c85e32cdec11979778594750953",
    ],
)

rpm(
    name = "python3-pip-0__9.0.3-22.el8.x86_64",
    sha256 = "ba83ca7667c98d265da7334a3ef7f786fbb48c85e32cdec11979778594750953",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/python3-pip-9.0.3-22.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/ba83ca7667c98d265da7334a3ef7f786fbb48c85e32cdec11979778594750953",
    ],
)

rpm(
    name = "python3-pip-wheel-0__9.0.3-22.el8.aarch64",
    sha256 = "772093492e290af496c3c8d4cf1d83d3288af49c4f0eb550f9c2489f96ecd89d",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/python3-pip-wheel-9.0.3-22.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/772093492e290af496c3c8d4cf1d83d3288af49c4f0eb550f9c2489f96ecd89d",
    ],
)

rpm(
    name = "python3-pip-wheel-0__9.0.3-22.el8.x86_64",
    sha256 = "772093492e290af496c3c8d4cf1d83d3288af49c4f0eb550f9c2489f96ecd89d",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/python3-pip-wheel-9.0.3-22.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/772093492e290af496c3c8d4cf1d83d3288af49c4f0eb550f9c2489f96ecd89d",
    ],
)

rpm(
    name = "python3-setuptools-0__39.2.0-6.el8.aarch64",
    sha256 = "c6f27b6e01d80e756408e3c1451e4af00e7d02da0aa24402644c0785118753fe",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/python3-setuptools-39.2.0-6.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/c6f27b6e01d80e756408e3c1451e4af00e7d02da0aa24402644c0785118753fe",
    ],
)

rpm(
    name = "python3-setuptools-0__39.2.0-6.el8.x86_64",
    sha256 = "c6f27b6e01d80e756408e3c1451e4af00e7d02da0aa24402644c0785118753fe",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/python3-setuptools-39.2.0-6.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/c6f27b6e01d80e756408e3c1451e4af00e7d02da0aa24402644c0785118753fe",
    ],
)

rpm(
    name = "python3-setuptools-wheel-0__39.2.0-6.el8.aarch64",
    sha256 = "b19bd4f106ce301ee21c860183cc1c2ef9c09bdf495059bdf16e8d8ccc71bbe8",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/python3-setuptools-wheel-39.2.0-6.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/b19bd4f106ce301ee21c860183cc1c2ef9c09bdf495059bdf16e8d8ccc71bbe8",
    ],
)

rpm(
    name = "python3-setuptools-wheel-0__39.2.0-6.el8.x86_64",
    sha256 = "b19bd4f106ce301ee21c860183cc1c2ef9c09bdf495059bdf16e8d8ccc71bbe8",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/python3-setuptools-wheel-39.2.0-6.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/b19bd4f106ce301ee21c860183cc1c2ef9c09bdf495059bdf16e8d8ccc71bbe8",
    ],
)

rpm(
    name = "python36-0__3.6.8-38.module_el8.5.0__plus__895__plus__a459eca8.aarch64",
    sha256 = "ab1d26bddf3f97decf17ac4a12c545add80be07bba1d7a1519481df24151e390",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/python36-3.6.8-38.module_el8.5.0+895+a459eca8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/ab1d26bddf3f97decf17ac4a12c545add80be07bba1d7a1519481df24151e390",
    ],
)

rpm(
    name = "python36-0__3.6.8-38.module_el8.5.0__plus__895__plus__a459eca8.x86_64",
    sha256 = "002b3672de2744c3f97ad8776d012952c058f9213a0cf8e01f7f9b8651b3e6af",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/python36-3.6.8-38.module_el8.5.0+895+a459eca8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/002b3672de2744c3f97ad8776d012952c058f9213a0cf8e01f7f9b8651b3e6af",
    ],
)

rpm(
    name = "qemu-common-15__6.1.0-15.el8.x86_64",
    sha256 = "cfba330c2cdd2f7ac2e98d2b8588ca4f6a9ac19a64374e6bdbef5c30839fa426",
    urls = ["https://pkgs.dyn.su/el8/extras/x86_64/qemu/qemu-common-6.1.0-15.el8.x86_64.rpm"],
)

rpm(
    name = "qemu-img-15__6.2.0-5.el8.x86_64",
    sha256 = "f808f215b2fedf08c20387e40174ed7bc7246e2fcf04c5e1129c3faa414d7ea7",
    urls = ["http://localhost:3031/qemu-img-6.2.0-5.el8.x86_64.rpm"],
)

rpm(
    name = "qemu-img-15__6.2.0-5.module_el8.6.0__plus__1087__plus__b42c8331.aarch64",
    sha256 = "af3133d3653a921ca543317bc1bc327fc3c853abfe71d7c8343af4bd8885cfaa",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/qemu-img-6.2.0-5.module_el8.6.0+1087+b42c8331.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/af3133d3653a921ca543317bc1bc327fc3c853abfe71d7c8343af4bd8885cfaa",
    ],
)

rpm(
    name = "qemu-kvm-common-15__6.2.0-5.el8.x86_64",
    sha256 = "08af674cfdf16f04b1003182ae4c988a468a37b670ae51f41ed703773965a717",
    urls = ["http://localhost:3031/qemu-kvm-common-6.2.0-5.el8.x86_64.rpm"],
)

rpm(
    name = "qemu-kvm-common-15__6.2.0-5.module_el8.6.0__plus__1087__plus__b42c8331.aarch64",
    sha256 = "e51be1ba77f9e5436483e748bea7dd141c26f5557764cbebbece8f175034a2ab",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/qemu-kvm-common-6.2.0-5.module_el8.6.0+1087+b42c8331.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/e51be1ba77f9e5436483e748bea7dd141c26f5557764cbebbece8f175034a2ab",
    ],
)

rpm(
    name = "qemu-kvm-core-15__6.2.0-5.el8.x86_64",
    sha256 = "383704988a4bcfbfe5102e8b7b396eddbb148f29a0babc55714fd4d55fc2c260",
    urls = ["http://localhost:3031/qemu-kvm-core-6.2.0-5.el8.x86_64.rpm"],
)

rpm(
    name = "qemu-kvm-core-15__6.2.0-5.module_el8.6.0__plus__1087__plus__b42c8331.aarch64",
    sha256 = "28b10ff340e60d70ded17ce0b06dfe19962cf9f5c8e0c04d50bbd0becaeb99f2",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/qemu-kvm-core-6.2.0-5.module_el8.6.0+1087+b42c8331.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/28b10ff340e60d70ded17ce0b06dfe19962cf9f5c8e0c04d50bbd0becaeb99f2",
    ],
)

rpm(
    name = "qemu-kvm-hw-usbredir-15__6.2.0-5.el8.x86_64",
    sha256 = "77228fc0cc3a2a707f8f6940c346f2b0cd001277f2a4fe19451654e7ae19dc25",
    urls = ["http://localhost:3031/qemu-kvm-hw-usbredir-6.2.0-5.el8.x86_64.rpm"],
)

rpm(
    name = "qemu-ui-gtk-15__6.1.0-15.el8.x86_64",
    sha256 = "1ac3f09d6e3c2679b1f47c5797e2cd8afc65d1b6d1a351c0ffad62ff915e4f4f",
    urls = ["https://pkgs.dyn.su/el8/extras/x86_64/qemu/qemu-ui-gtk-6.1.0-15.el8.x86_64.rpm"],
)

rpm(
    name = "qemu-ui-opengl-15__6.1.0-15.el8.x86_64",
    sha256 = "c5eaa64224865c4bad83a21cdfccacdca98a4139ee534d9b6dc5b64ae4749b03",
    urls = ["https://pkgs.dyn.su/el8/extras/x86_64/qemu/qemu-ui-opengl-6.1.0-15.el8.x86_64.rpm"],
)

rpm(
    name = "qemu-ui-sdl-15__6.1.0-15.el8.x86_64",
    sha256 = "f7fc7ebd26d52c4cd4fb02420e47dca1bbe85a8e939e33f97660dcc7dbf70546",
    urls = ["https://pkgs.dyn.su/el8/extras/x86_64/qemu/qemu-ui-sdl-6.1.0-15.el8.x86_64.rpm"],
)

rpm(
    name = "readline-0__7.0-10.el8.aarch64",
    sha256 = "ef74f2c65ed0e38dd021177d6e59fcdf7fb8de8929b7544b7a6f0709eff6562c",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/readline-7.0-10.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/ef74f2c65ed0e38dd021177d6e59fcdf7fb8de8929b7544b7a6f0709eff6562c",
    ],
)

rpm(
    name = "readline-0__7.0-10.el8.x86_64",
    sha256 = "fea868a7d82a7b6f392260ed4afb472dc4428fd71eab1456319f423a845b5084",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/readline-7.0-10.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/fea868a7d82a7b6f392260ed4afb472dc4428fd71eab1456319f423a845b5084",
    ],
)

rpm(
    name = "rest-0__0.8.1-2.el8.x86_64",
    sha256 = "000221caf7172bad19efd8f8bca6d8b0cdf3174cc4601e07916e5e58caca7df9",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/rest-0.8.1-2.el8.x86_64.rpm"],
)

rpm(
    name = "rpm-0__4.14.3-23.el8.aarch64",
    sha256 = "d803f082920abc401f44b7220ce96f6f2b070b06dcfe6b5c34573b8c7bcc5267",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/rpm-4.14.3-23.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/d803f082920abc401f44b7220ce96f6f2b070b06dcfe6b5c34573b8c7bcc5267",
    ],
)

rpm(
    name = "rpm-0__4.14.3-25.el8.x86_64",
    sha256 = "490ff18c5aaebef35d8166534c2d98dcd2d78413f3a3c4bde582595975a8b89f",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/rpm-4.14.3-25.el8.x86_64.rpm"],
)

rpm(
    name = "rpm-libs-0__4.14.3-23.el8.aarch64",
    sha256 = "26fdda368fc8c50c774cebd9ddf4786ced58d8ee9b12e5ce57113205d147f0a1",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/rpm-libs-4.14.3-23.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/26fdda368fc8c50c774cebd9ddf4786ced58d8ee9b12e5ce57113205d147f0a1",
    ],
)

rpm(
    name = "rpm-libs-0__4.14.3-25.el8.x86_64",
    sha256 = "0caf7bbc302461fd2e198c3463d814580e96949dbac089e1a6d7926f2b0308b0",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/rpm-libs-4.14.3-25.el8.x86_64.rpm"],
)

rpm(
    name = "rpm-plugin-selinux-0__4.14.3-23.el8.aarch64",
    sha256 = "66c8e46bde5c784c083c7e674f72edb493394c9dedf59e7b40600968f083ca5c",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/rpm-plugin-selinux-4.14.3-23.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/66c8e46bde5c784c083c7e674f72edb493394c9dedf59e7b40600968f083ca5c",
    ],
)

rpm(
    name = "rpm-plugin-selinux-0__4.14.3-25.el8.x86_64",
    sha256 = "326ae46a22f01ec78749ea381a93e28fb308b0ec9ea7f2da360f6a27c6726f24",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/rpm-plugin-selinux-4.14.3-25.el8.x86_64.rpm"],
)

rpm(
    name = "seabios-0__1.15.0-1.module_el8.6.0__plus__1087__plus__b42c8331.x86_64",
    sha256 = "4d421d4139e7ad6e5a2ec8be8f347bc16a871571525d6b8d2ae251436d4bd89f",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/seabios-1.15.0-1.module_el8.6.0+1087+b42c8331.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/4d421d4139e7ad6e5a2ec8be8f347bc16a871571525d6b8d2ae251436d4bd89f",
    ],
)

rpm(
    name = "seabios-bin-0__1.15.0-1.module_el8.6.0__plus__1087__plus__b42c8331.x86_64",
    sha256 = "3c8d058cabbdad4e9780aab2f3770c8162bfc28f837dd6036690497b82101d3f",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/seabios-bin-1.15.0-1.module_el8.6.0+1087+b42c8331.noarch.rpm",
        "https://storage.googleapis.com/builddeps/3c8d058cabbdad4e9780aab2f3770c8162bfc28f837dd6036690497b82101d3f",
    ],
)

rpm(
    name = "seavgabios-bin-0__1.15.0-1.module_el8.6.0__plus__1087__plus__b42c8331.x86_64",
    sha256 = "34d9c5e00e88a00e8be874470dc2f1460f7957335fd0081936e8a17fcf66605c",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/seavgabios-bin-1.15.0-1.module_el8.6.0+1087+b42c8331.noarch.rpm",
        "https://storage.googleapis.com/builddeps/34d9c5e00e88a00e8be874470dc2f1460f7957335fd0081936e8a17fcf66605c",
    ],
)

rpm(
    name = "sed-0__4.5-5.el8.aarch64",
    sha256 = "806550c684c46a58a455953223fafbacc343e35e488d436bf963844944a33861",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/sed-4.5-5.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/806550c684c46a58a455953223fafbacc343e35e488d436bf963844944a33861",
    ],
)

rpm(
    name = "sed-0__4.5-5.el8.x86_64",
    sha256 = "5a09d6d967d12580c7e6ab92db35bcafd3426d6121ec60c78f54e3cd4961cd26",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/sed-4.5-5.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/5a09d6d967d12580c7e6ab92db35bcafd3426d6121ec60c78f54e3cd4961cd26",
    ],
)

rpm(
    name = "selinux-policy-0__3.14.3-105.el8.aarch64",
    sha256 = "6026b6a44def22c85959a5fe466cd62a7f000aa6baf3dfb925fb59c3e9dcc4e9",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/selinux-policy-3.14.3-105.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/6026b6a44def22c85959a5fe466cd62a7f000aa6baf3dfb925fb59c3e9dcc4e9",
    ],
)

rpm(
    name = "selinux-policy-0__3.14.3-113.el8.x86_64",
    sha256 = "480daab7442157b8652f9915d07ac1a53309cbebb1b8174150e5f905e0e52306",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/selinux-policy-3.14.3-113.el8.noarch.rpm"],
)

rpm(
    name = "selinux-policy-targeted-0__3.14.3-105.el8.aarch64",
    sha256 = "b544ae2bd1d437669944b90a63a90eb5cdc8f7896375877889074d62c8a9180d",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/selinux-policy-targeted-3.14.3-105.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/b544ae2bd1d437669944b90a63a90eb5cdc8f7896375877889074d62c8a9180d",
    ],
)

rpm(
    name = "selinux-policy-targeted-0__3.14.3-113.el8.x86_64",
    sha256 = "59cd1c9fe944627907b429e18d12085f82c86c5014e6343d236a51e0fe3a7606",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/selinux-policy-targeted-3.14.3-113.el8.noarch.rpm"],
)

rpm(
    name = "setup-0__2.12.2-7.el8.aarch64",
    sha256 = "0e5bdfebabb44848a9f37d2cc02a8a6a099b1c4c1644f4940718e55ce5b95464",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/setup-2.12.2-7.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/0e5bdfebabb44848a9f37d2cc02a8a6a099b1c4c1644f4940718e55ce5b95464",
    ],
)

rpm(
    name = "setup-0__2.12.2-9.el8.x86_64",
    sha256 = "0a0696aebfadbbeb229445c0828a83be763460d6af6a552b3bd533acde011644",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/setup-2.12.2-9.el8.noarch.rpm"],
)

rpm(
    name = "sgabios-bin-1__0.20180715git-7.el8.x86_64",
    sha256 = "883dbf88a32ceb46c82220b7e5c42a98f596771e088d69a73708b2bc2cc0a6ac",
    urls = ["https://pkgs.dyn.su/el8/extras/x86_64/qemu/sgabios-bin-0.20180715git-7.el8.noarch.rpm"],
)

rpm(
    name = "shadow-utils-2__4.6-17.el8.aarch64",
    sha256 = "c2ed285e2a2495b33e926c57e1917114c7898f2f4536866d643f206780a699af",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/shadow-utils-4.6-17.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/c2ed285e2a2495b33e926c57e1917114c7898f2f4536866d643f206780a699af",
    ],
)

rpm(
    name = "shadow-utils-2__4.6-17.el8.x86_64",
    sha256 = "fb3c71778fc23c4d3c91911c49e0a0d14c8a5192c431fc9ba07f2a14c938a172",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/shadow-utils-4.6-17.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/fb3c71778fc23c4d3c91911c49e0a0d14c8a5192c431fc9ba07f2a14c938a172",
    ],
)

rpm(
    name = "shared-mime-info-0__1.9-3.el8.x86_64",
    sha256 = "3f3cced089849779b46d7b1f27ae1b73e0b1144eca15716e4c19e4b54bb16f6c",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/shared-mime-info-1.9-3.el8.x86_64.rpm"],
)

rpm(
    name = "snappy-0__1.1.8-3.el8.aarch64",
    sha256 = "4731985b22fc7b733ff89be6c1423396f27c94a78bb09fc89be5c2200bee893c",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/snappy-1.1.8-3.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/4731985b22fc7b733ff89be6c1423396f27c94a78bb09fc89be5c2200bee893c",
    ],
)

rpm(
    name = "snappy-0__1.1.8-3.el8.x86_64",
    sha256 = "839c62cd7fc7e152decded6f28c80b5f7b8f34a5e319057867b38b26512cee67",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/snappy-1.1.8-3.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/839c62cd7fc7e152decded6f28c80b5f7b8f34a5e319057867b38b26512cee67",
    ],
)

rpm(
    name = "sqlite-libs-0__3.26.0-15.el8.aarch64",
    sha256 = "b3a0c27117c927795b1a3a1ef2c08c857a88199bcfad5603cd2303c9519671a4",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/sqlite-libs-3.26.0-15.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/b3a0c27117c927795b1a3a1ef2c08c857a88199bcfad5603cd2303c9519671a4",
    ],
)

rpm(
    name = "sqlite-libs-0__3.26.0-17.el8.x86_64",
    sha256 = "a44b1bd3d9f5a6b0654ba4ae2f8aa45aefec54c9377dfe4446ec1c0e2fd0ac89",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/sqlite-libs-3.26.0-17.el8.x86_64.rpm"],
)

rpm(
    name = "sssd-client-0__2.7.3-1.el8.aarch64",
    sha256 = "e7815948862ecfbcaddd3a0dfcaa4bd72c1a7124e917aa1f55ce235ae3feaa32",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/sssd-client-2.7.3-1.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/e7815948862ecfbcaddd3a0dfcaa4bd72c1a7124e917aa1f55ce235ae3feaa32",
    ],
)

rpm(
    name = "sssd-client-0__2.8.1-1.el8.x86_64",
    sha256 = "7e432b68b5e214b8364ad44d99e9be77b229677237027d2b4b1443032bc4e6cd",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/sssd-client-2.8.1-1.el8.x86_64.rpm"],
)

rpm(
    name = "swtpm-0__0.7.0-1.20211109gitb79fd91.module_el8.6.0__plus__1087__plus__b42c8331.aarch64",
    sha256 = "e004d8c0bf81bedece3539bea0e12961740ae7dcc69d39f3774d7a738902e063",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/swtpm-0.7.0-1.20211109gitb79fd91.module_el8.6.0+1087+b42c8331.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/e004d8c0bf81bedece3539bea0e12961740ae7dcc69d39f3774d7a738902e063",
    ],
)

rpm(
    name = "swtpm-0__0.7.0-4.20211109gitb79fd91.module_el8.7.0__plus__1218__plus__f626c2ff.x86_64",
    sha256 = "2125c4d6cb910e47daf45fbef10d75f93b5d30e64908b42dfc77aeee201feb60",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/swtpm-0.7.0-4.20211109gitb79fd91.module_el8.7.0+1218+f626c2ff.x86_64.rpm"],
)

rpm(
    name = "swtpm-libs-0__0.7.0-1.20211109gitb79fd91.module_el8.6.0__plus__1087__plus__b42c8331.aarch64",
    sha256 = "df014e3d5b857c279e5aac2c8147889710010e52abad3f0cef7945865533ec21",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/swtpm-libs-0.7.0-1.20211109gitb79fd91.module_el8.6.0+1087+b42c8331.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/df014e3d5b857c279e5aac2c8147889710010e52abad3f0cef7945865533ec21",
    ],
)

rpm(
    name = "swtpm-libs-0__0.7.0-4.20211109gitb79fd91.module_el8.7.0__plus__1218__plus__f626c2ff.x86_64",
    sha256 = "f29e2f9e3f3c4ba3cddbe4af4dc7db2e7ad0088db6e955da86dacb40d4e75466",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/swtpm-libs-0.7.0-4.20211109gitb79fd91.module_el8.7.0+1218+f626c2ff.x86_64.rpm"],
)

rpm(
    name = "swtpm-tools-0__0.7.0-1.20211109gitb79fd91.module_el8.6.0__plus__1087__plus__b42c8331.aarch64",
    sha256 = "dd923963bfa8c404d69dbc8d70014791355a08335930c3802d403a13b341838c",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/swtpm-tools-0.7.0-1.20211109gitb79fd91.module_el8.6.0+1087+b42c8331.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/dd923963bfa8c404d69dbc8d70014791355a08335930c3802d403a13b341838c",
    ],
)

rpm(
    name = "swtpm-tools-0__0.7.0-4.20211109gitb79fd91.module_el8.7.0__plus__1218__plus__f626c2ff.x86_64",
    sha256 = "bb88081e4d8978aaea3e902252be225211fc496f053ac721757a8b005c3ad86d",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/swtpm-tools-0.7.0-4.20211109gitb79fd91.module_el8.7.0+1218+f626c2ff.x86_64.rpm"],
)

rpm(
    name = "systemd-0__239-62.el8.aarch64",
    sha256 = "a637d2ce107c70a5315838441e93769ceb6c0216c759800f6387a3b14227dac8",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/systemd-239-62.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/a637d2ce107c70a5315838441e93769ceb6c0216c759800f6387a3b14227dac8",
    ],
)

rpm(
    name = "systemd-0__239-69.el8.x86_64",
    sha256 = "65ae1f1c2ee8c972446e667ca8e95f0366db1851bdc1ac8ca953e130c096f511",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/systemd-239-69.el8.x86_64.rpm"],
)

rpm(
    name = "systemd-container-0__239-62.el8.aarch64",
    sha256 = "50fee5eb1dccf0e1491dca78de066e088ed96159a84a83e43ad2342bb2636d7d",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/systemd-container-239-62.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/50fee5eb1dccf0e1491dca78de066e088ed96159a84a83e43ad2342bb2636d7d",
    ],
)

rpm(
    name = "systemd-container-0__239-69.el8.x86_64",
    sha256 = "a16764898802e77f8674e839898189b510b7c2cb4eb68f0bddd75e0555ed2a2f",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/systemd-container-239-69.el8.x86_64.rpm"],
)

rpm(
    name = "systemd-libs-0__239-62.el8.aarch64",
    sha256 = "a391a6450f791020f521c7f1483d0756dcc168eb6bbcf99e0b426234fb9ca39a",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/systemd-libs-239-62.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/a391a6450f791020f521c7f1483d0756dcc168eb6bbcf99e0b426234fb9ca39a",
    ],
)

rpm(
    name = "systemd-libs-0__239-69.el8.x86_64",
    sha256 = "fdbf203811dd13f8be1a9907b6d55b99079320f8ae440b0c7111fe43be21773a",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/systemd-libs-239-69.el8.x86_64.rpm"],
)

rpm(
    name = "systemd-pam-0__239-62.el8.aarch64",
    sha256 = "8e06e2569d86cbf4ed7a669decc0a4a5faf2a65cb1b84ee10cc1947946cff1d7",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/systemd-pam-239-62.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/8e06e2569d86cbf4ed7a669decc0a4a5faf2a65cb1b84ee10cc1947946cff1d7",
    ],
)

rpm(
    name = "systemd-pam-0__239-69.el8.x86_64",
    sha256 = "6f12c39f240ac9a13f1d15dadf0a1f38838638cb6df367f0cd3038284fd3a925",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/systemd-pam-239-69.el8.x86_64.rpm"],
)

rpm(
    name = "tar-2__1.30-6.el8.aarch64",
    sha256 = "ef568db2a1acf8da0aa45c2378fd517150d3c878b025c0c5e030471ddb548772",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/tar-1.30-6.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/ef568db2a1acf8da0aa45c2378fd517150d3c878b025c0c5e030471ddb548772",
    ],
)

rpm(
    name = "tar-2__1.30-8.el8.x86_64",
    sha256 = "55056b8b60c9382f41e97004ed64fb300c930b29f15b70394d6a5aeffe23010e",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/tar-1.30-8.el8.x86_64.rpm"],
)

rpm(
    name = "tzdata-0__2022a-2.el8.aarch64",
    sha256 = "0440f6795ede1959a5381056845a232db6991633aae371373e703d9c16e592e2",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/tzdata-2022a-2.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/0440f6795ede1959a5381056845a232db6991633aae371373e703d9c16e592e2",
    ],
)

rpm(
    name = "tzdata-0__2022g-1.el8.x86_64",
    sha256 = "e53697a9248574f061b517de5286ca8134edb1601e08bb1819bf5d301bb8b5a6",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/tzdata-2022g-1.el8.noarch.rpm"],
)

rpm(
    name = "unbound-libs-0__1.16.0-2.el8.aarch64",
    sha256 = "280dbefbd2d57d9fc8c39e6822f1bf65694a6c25ebcbec2fbafcb5ac5319649f",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/unbound-libs-1.16.0-2.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/280dbefbd2d57d9fc8c39e6822f1bf65694a6c25ebcbec2fbafcb5ac5319649f",
    ],
)

rpm(
    name = "unbound-libs-0__1.16.2-2.el8.x86_64",
    sha256 = "9d1fd4ba858e6788c32a6dd3adaa8db51fc4c1ae34366fe62ef136dbfa64a9b7",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/unbound-libs-1.16.2-2.el8.x86_64.rpm"],
)

rpm(
    name = "usbredir-0__0.12.0-2.el8.x86_64",
    sha256 = "0b6e50e9e9c68d0dbacc39e81c4a3a3a7ccf3afaddf40afb06ca86424a46ba23",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/usbredir-0.12.0-2.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/0b6e50e9e9c68d0dbacc39e81c4a3a3a7ccf3afaddf40afb06ca86424a46ba23",
    ],
)

rpm(
    name = "userspace-rcu-0__0.10.1-4.el8.aarch64",
    sha256 = "c4b53c8f1121938c2c5ae3fabd48b9d8f77c7d26f47a76f5c0eab3fd7f0a6cfc",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/userspace-rcu-0.10.1-4.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/c4b53c8f1121938c2c5ae3fabd48b9d8f77c7d26f47a76f5c0eab3fd7f0a6cfc",
    ],
)

rpm(
    name = "userspace-rcu-0__0.10.1-4.el8.x86_64",
    sha256 = "4025900345c5125fd6c10c1780275139f56b63be2bfac10be83628758c225dd0",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/userspace-rcu-0.10.1-4.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/4025900345c5125fd6c10c1780275139f56b63be2bfac10be83628758c225dd0",
    ],
)

rpm(
    name = "util-linux-0__2.32.1-35.el8.aarch64",
    sha256 = "68bae6f15242edce8c052351a681fb31e06a676f4320589c39227208bc6bc78a",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/util-linux-2.32.1-35.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/68bae6f15242edce8c052351a681fb31e06a676f4320589c39227208bc6bc78a",
    ],
)

rpm(
    name = "util-linux-0__2.32.1-39.el8.x86_64",
    sha256 = "071b1a3a157faed2cfb9a48ca0e43cda41ae9cebfd74926f68ccaa379497f278",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/util-linux-2.32.1-39.el8.x86_64.rpm"],
)

rpm(
    name = "vim-minimal-2__8.0.1763-19.el8.4.aarch64",
    sha256 = "4a921c33ca497386a80d4f6ace2ec54bc8e568c83f6197daa9a0f29b8a97fe1d",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/vim-minimal-8.0.1763-19.el8.4.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/4a921c33ca497386a80d4f6ace2ec54bc8e568c83f6197daa9a0f29b8a97fe1d",
    ],
)

rpm(
    name = "vim-minimal-2__8.2.2637-12.el8.x86_64",
    sha256 = "daed605e77a360d70f2bdfc8b20a8cfea3dbeffde7425009b52a5cbe5df321a8",
    urls = ["https://pkgs.dyn.su/el8/extras/x86_64/vim-minimal-8.2.2637-12.el8.x86_64.rpm"],
)

rpm(
    name = "vte-profile-0__0.52.4-2.el8.x86_64",
    sha256 = "487c0b4ead22f4728cc37d1a470e889bf7bf0f02172561aea406260c197de3bc",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/vte-profile-0.52.4-2.el8.x86_64.rpm"],
)

rpm(
    name = "vte291-0__0.52.4-2.el8.x86_64",
    sha256 = "902d5820b7cecb836117ab40a2416efdbb76e9e48d19868dffc4508ddff7479b",
    urls = ["http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/vte291-0.52.4-2.el8.x86_64.rpm"],
)

rpm(
    name = "which-0__2.21-18.el8.aarch64",
    sha256 = "c27e749065a42c812467155241ee9eedfcaae0f08f4cec952aa65194e98723d7",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/which-2.21-18.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/c27e749065a42c812467155241ee9eedfcaae0f08f4cec952aa65194e98723d7",
    ],
)

rpm(
    name = "which-0__2.21-18.el8.x86_64",
    sha256 = "0e4d5ee4cbea952903ee4febb1450caf92bf3c2d6ecac9d0dd8ac8611e9ff4db",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/which-2.21-18.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/0e4d5ee4cbea952903ee4febb1450caf92bf3c2d6ecac9d0dd8ac8611e9ff4db",
    ],
)

rpm(
    name = "xkeyboard-config-0__2.28-1.el8.aarch64",
    sha256 = "a2aeabb3962859069a78acc288bc3bffb35485428e162caafec8134f5ce6ca67",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/xkeyboard-config-2.28-1.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/a2aeabb3962859069a78acc288bc3bffb35485428e162caafec8134f5ce6ca67",
    ],
)

rpm(
    name = "xkeyboard-config-0__2.28-1.el8.x86_64",
    sha256 = "a2aeabb3962859069a78acc288bc3bffb35485428e162caafec8134f5ce6ca67",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/xkeyboard-config-2.28-1.el8.noarch.rpm",
        "https://storage.googleapis.com/builddeps/a2aeabb3962859069a78acc288bc3bffb35485428e162caafec8134f5ce6ca67",
    ],
)

rpm(
    name = "xorriso-0__1.4.8-4.el8.aarch64",
    sha256 = "4280064ab658525b486d7b8c2ca5f87aeef90002361a0925f2819fd7a7909500",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/xorriso-1.4.8-4.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/4280064ab658525b486d7b8c2ca5f87aeef90002361a0925f2819fd7a7909500",
    ],
)

rpm(
    name = "xorriso-0__1.4.8-4.el8.x86_64",
    sha256 = "3a232d848da1ace286efef6c8c9cf0fcfab2c47dd58968ddb6a24718629a6220",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/xorriso-1.4.8-4.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/3a232d848da1ace286efef6c8c9cf0fcfab2c47dd58968ddb6a24718629a6220",
    ],
)

rpm(
    name = "xz-0__5.2.4-4.el8.aarch64",
    sha256 = "c30b066af6b844602964858ef77b995e944ffbdd7a153a9c5c7fc30fd802b926",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/xz-5.2.4-4.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/c30b066af6b844602964858ef77b995e944ffbdd7a153a9c5c7fc30fd802b926",
    ],
)

rpm(
    name = "xz-0__5.2.4-4.el8.x86_64",
    sha256 = "99d7d4bfee1d5b55e08ee27c6869186531939f399d6c3ea33db191cae7e53f70",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/xz-5.2.4-4.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/99d7d4bfee1d5b55e08ee27c6869186531939f399d6c3ea33db191cae7e53f70",
    ],
)

rpm(
    name = "xz-libs-0__5.2.4-4.el8.aarch64",
    sha256 = "9498f961afe361c5f9e0eea0ce64f11071b1cb1afe30636cb888d109737ea16f",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/xz-libs-5.2.4-4.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/9498f961afe361c5f9e0eea0ce64f11071b1cb1afe30636cb888d109737ea16f",
    ],
)

rpm(
    name = "xz-libs-0__5.2.4-4.el8.x86_64",
    sha256 = "69d67ea8b4bd532f750ff0592f0098ace60470da0fd0e4056188fda37a268d42",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/xz-libs-5.2.4-4.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/69d67ea8b4bd532f750ff0592f0098ace60470da0fd0e4056188fda37a268d42",
    ],
)

rpm(
    name = "yajl-0__2.1.0-11.el8.aarch64",
    sha256 = "3ae671d2c8bfd1f53ea706e3969dd2dafd5a2960371e8b6f6083fb345985a491",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/aarch64/os/Packages/yajl-2.1.0-11.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/3ae671d2c8bfd1f53ea706e3969dd2dafd5a2960371e8b6f6083fb345985a491",
    ],
)

rpm(
    name = "yajl-0__2.1.0-11.el8.x86_64",
    sha256 = "55a094ffe9f378ef465619bf6f60e9f26b672f67236883565fb893de7675c163",
    urls = [
        "http://mirror.centos.org/centos/8-stream/AppStream/x86_64/os/Packages/yajl-2.1.0-11.el8.x86_64.rpm",
        "https://storage.googleapis.com/builddeps/55a094ffe9f378ef465619bf6f60e9f26b672f67236883565fb893de7675c163",
    ],
)

rpm(
    name = "zlib-0__1.2.11-19.el8.aarch64",
    sha256 = "4fe374ebed682fa00ac065080b6cef94b7fcfc5d422d77a4cfdf5eb878c60ced",
    urls = [
        "http://mirror.centos.org/centos/8-stream/BaseOS/aarch64/os/Packages/zlib-1.2.11-19.el8.aarch64.rpm",
        "https://storage.googleapis.com/builddeps/4fe374ebed682fa00ac065080b6cef94b7fcfc5d422d77a4cfdf5eb878c60ced",
    ],
)

rpm(
    name = "zlib-0__1.2.11-21.el8.x86_64",
    sha256 = "9aabeb4a75c05b98661200dc9f0f1c7c528af42b9535c7c133dd4c0c5f80d179",
    urls = ["http://mirror.centos.org/centos/8-stream/BaseOS/x86_64/os/Packages/zlib-1.2.11-21.el8.x86_64.rpm"],
)

rpm(
    name = "zlib-ng-0__2.0.1-3.el8.x86_64",
    sha256 = "0116bcac8f3b2684ba3186c93d2eb60a83bd33c23cc11da6475eec87e66e7007",
    urls = ["https://pkgs.dyn.su/el8/extras/x86_64/zlib-ng-2.0.1-3.el8.x86_64.rpm"],
)
