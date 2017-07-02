[![forthebadge](http://forthebadge.com/images/badges/built-with-love.svg)](https://github.com/daidokoro/gzs3)
# gzs3 - zip your Git Repo in S3

[![codebeat badge](https://codebeat.co/badges/13cd8045-d01f-4425-8125-49e45f899a2b)](https://codebeat.co/projects/github-com-daidokoro-gzs3-master)

[![asciicast](https://asciinema.org/a/rrDghylQhrx8FN34p33tSbJTL.png)](https://asciinema.org/a/rrDghylQhrx8FN34p33tSbJTL?speed=1.5)

__gzs3__(Git Zip to S3) was born of a very specific need to automate the zipping and uploading of lambda functions to S3 from a Git Repository for Cloudformation Deployments.

__gzs3__ is capable of zipping any repo and storing it in s3, not just repos containing lambda functions.


### Required

- AWS Credentials need to be configured to allow access to the S3 Bucket being written to.


## How it Works!

__gzs3__ does not create a single file on disc, instead, the repo and file zip operation are all handled in memory.

1. Your repo is cloned into memory
2. A Zip file is then created (still memory)
3. The Zip file is written to s3

All you need to do is create a __gzs3file__ in root of your repo, containing the following:

```yaml
# bucket name
bucket: somebucket

# the name of the zip and key/prefix to stored it under in s3
key: some/key.zip
```


Then simply call the repo using the CLI tool.

```
Clone git Repo & ZIP to AWS S3...

Usage:
  gzs3 [flags]

Examples:
gzs3 git@github.com/some/repo.git

Flags:
      --debug             debug mode
  -h, --help              help for gzs3
      --no-colors         disable color output
      --password string   git password
  -p, --profile string    configured AWS profile (default "default")
  -i, --ssh-rsa string    path to git SSH id_rsa (default "/Users/shaunremekie/.ssh/id_rsa")
  -u, --user string       git username
```

__Try__:

```
$ gzs3 git@github.com:some/repo.git
```

__Or__ via http(s)

```
$ gzs3 https://github.com/some/repo.git
```


__Note__: It's important to note that larger repos may consume a larger amount of memory to create a Zip file.

## Installation

If you have Golang installed:

`go get github.com/daidokoro/gzs3`


## TODO:
- Write some tests!!!
- Support for _tar_ archive

# Contributing

Fork -> Patch -> Push -> Pull Request

_Pull requests welcomed...._
