# semver-generator

A simple script to generate a semver compatible version string from an ISO 8601 date string.

## Usage

Recommended usage is via a docker container or in GitLab CI/CD pipelines.

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