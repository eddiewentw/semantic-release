# Semantic Release

[![github-action](https://github.com/eddiewentw/semantic-release/workflows/Release/badge.svg)](https://github.com/eddiewentw/semantic-release/actions) [![latest-tag](https://img.shields.io/github/v/tag/eddiewentw/semantic-release.svg)](https://github.com/eddiewentw/semantic-release/releases) [![conventional-commits](https://img.shields.io/badge/Conventional%20Commits-1.0.0-yellow.svg)](https://conventionalcommits.org)

Inspired by [standard-version](https://github.com/conventional-changelog/standard-version). Standard-version is awesome, when I develop a JS project, it is always my favorite. But sometimes when I develop a non-JS project, I don't want to integrate npm system to use that. That's why I start this repository. It's developed by Golang and it offers a binary file to execute.

## Usage

If you're using golang, you can ....

```bash
$ go run github.com/eddiewentw/semantic-release/cmd
```

If you're not using golang ....

There are binary files in [release](https://github.com/eddiewentw/semantic-release/releases) page. You can download it and use in your project.

```bash
# make it executable
$ chmod +x ./scripts/semantic-release

$ ./scripts/semantic-release
```

Please open an issue to tell me if you have a better idea to improve this usage.

### Debug

```bash
$ /semantic-release --debug
```

Print more information in every action.

### Dry run

```bash
$ /semantic-release --dry-run
```

Use this flag when you want to know what the next version is.

### First release

```bash
$ /semantic-release --first-release
```

Use this flag for your first release. First release will be "v1.0.0".

## How does it work

Versioning rule follows [Conventional Commits](https://www.conventionalcommits.org/en/v1.0.0/). When you execute it, it does these steps.

1. Write the new version into [.semantic-version](./.semantic-version) file.
2. Commit this change and tag it.

```bash
commit 26a24221dd549ed21b9887a4d75c84a7856d27a5 (tag: v0.3.0, origin/release)
Author: GitHub Action <action@github.com>
Date:   Sat Jan 11 06:05:00 2020 +0000

    chore(release): v0.3.0

diff --git a/.semantic-version b/.semantic-version
index 22c08f7..268b033 100644
--- a/.semantic-version
+++ b/.semantic-version
@@ -1 +1 @@
-v0.2.1
+v0.3.0
```

[.semantic-version](./.semantic-version) file only contains version number, your application can read this file to get the current version information.
