# semver-generator

A simple script to generate a semver compatible version string from an ISO 8601 date string.

## Usage

Recommended usage is via a docker container or in GitLab CI/CD pipelines.

### Docker

```bash
$ docker run --rm das-kaesebrot:dt-version-gen "2023-04-24T09:12:23"
DATETIME_VERSION_MAJOR=2023
DATETIME_VERSION_MINOR=424
DATETIME_VERSION_PATCH=91223
DATETIME_VERSION_FULL=2023.424.91223
```
### GitLab CI/CD

```yaml
stages:
  - setup
  - build

prep-vars:
  stage: setup
  image: das-kaesebrot:dt-version-gen
  script:
    - version-gen "$CI_PIPELINE_CREATED_AT" >> build.env
```