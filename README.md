# dt-version-gen

A simple script to generate a semver compatible version string from an ISO 8601 date string.

## Example output
```bash
$ docker run --rm -it daskaesebrot/dt-version-gen version-gen "2023-05-08T16:34:12Z"
DATETIME_VERSION_MAJOR=2023
DATETIME_VERSION_MINOR=508
DATETIME_VERSION_PATCH=1634
DATETIME_VERSION_FULL=2023.508.1634
```

## Usage

Recommended usage is via a GitLab CI/CD pipelines.

### GitLab CI/CD

Also see: https://docs.gitlab.com/ee/ci/variables/#pass-an-environment-variable-to-another-job

```yaml
stages:
  - setup
  - publish

# execute the version generator first
prep-vars:
  stage: setup
  image: daskaesebrot/dt-version-gen
  # append generated variables to existing pipeline environment
  script:
    - version-gen "$CI_PIPELINE_CREATED_AT" >> build.env
  
  # persist the generated variables
  artifacts:
    reports:
      dotenv: build.env

# use the generated variables in your job
publish-docker:
  stage: publish
  image: docker:stable-dind
  services:
    - docker:stable-dind
  
  # important - retrieve persisted variables
  needs:
    - job: prep-vars
      artifacts: true
    
  script:
    - docker login -u $CI_REGISTRY_USER -p $CI_REGISTRY_PASSWORD $CI_REGISTRY
    - docker pull $CI_REGISTRY_IMAGE:latest || true
    - |
      docker build \
        --cache-from $CI_REGISTRY_IMAGE:latest \
        --tag $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA \
        --tag $CI_REGISTRY_IMAGE:$DATETIME_VERSION_FULL \
        --tag $CI_REGISTRY_IMAGE:latest \
        .
    - docker push $CI_REGISTRY_IMAGE:$CI_COMMIT_SHA
    - docker push $CI_REGISTRY_IMAGE:$DATETIME_VERSION_FULL
    - docker push $CI_REGISTRY_IMAGE:latest
```
