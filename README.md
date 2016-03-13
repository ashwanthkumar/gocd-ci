[![Build Status](https://snap-ci.com/ashwanthkumar/gocd-ci/branch/master/build_image)](https://snap-ci.com/ashwanthkumar/gocd-ci/branch/master)
# gocd-ci

`GoCD-CI` helps you build your projects using a `.gocd-ci.yml` file. You just have to download the binary and run it from the project's root.

## Features
- [x] Set environment variables
- [x] Run commands for your build
- [ ] Upload artifacts to GoCD after the build

## .gocd-ci.yml
```yml
---
  name: "GoCD CI"
  env:
    FOO: "BAR"
  cmd:
    - make setup
    - make test
    - make build
  artifacts:
   gocd-ci : .
```

- `name` - Human readable name for the project. Currently not being used.
- `env` - Map of Environment variables required for the application. You'll get `GO_*` variables for free. If none, it can be ignored from the yml file.
- `cmd` - List of commands to run as part of the build. Like [travis-ci](https://docs.travis-ci.com/user/customizing-the-build/#Customizing-the-Build-Step) we run all the commands even if one of them fail. Though the final exit code will be 1 if any command failed.
- `artifacts` - Map of artifacts to be exposed on the pipeline via GoCD. If none, it can be ignored from the yml file.
