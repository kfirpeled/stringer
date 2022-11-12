description = "stringer - stringer is a super fancy CLI (kidding)"
binaries = ["stringer"]
test = "stringer -v"

platform "darwin" "amd64" {
  source = "https://github.com/kfirpeled/stringer/releases/download/v${version}/stringer_${version}_${os}_${arch}.tar.gz"
}

platform "darwin" "arm64" {
  source = "https://github.com/kfirpeled/stringer/releases/download/v${version}/stringer_${version}_${os}_${arch}.tar.gz"
}

platform "linux" "amd64" {
  source = "https://github.com/kfirpeled/stringer/releases/download/v${version}/stringer_${version}_${os}_${arch}.tar.gz"
}

version "0.0.1" {
  auto-version {
    github-release = "kfirpeled/stringer"
  }
}
