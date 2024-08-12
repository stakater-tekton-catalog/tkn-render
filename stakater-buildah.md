# stakater-buildah

### url
https://raw.githubusercontent.com//0.0.5/stakater-buildah.yaml

### Version
Current version is: 0.0.5

## Description
Buildah task builds source into a container image and then pushes it to a container registry. Buildah Task builds source into a container image using Project Atomic's Buildah build tool.It uses Buildah's support for building from Dockerfiles, using its buildah bud command.This command executes the directives in the Dockerfile to assemble a container image, then pushes that image to a container registry.

### Params
| Name | Description | Type | Default |
|------|-------------|------|---------|
| `IMAGE` | Reference of the image buildah will produce. | string | <no value> |
| `BUILDER_IMAGE` | The location of the buildah builder image. | string | registry.redhat.io/rhel8/buildah@sha256:180c4d9849b6ab0e5465d30d4f3a77765cf0d852ca1cb1efb59d6e8c9f90d467 |
| `STORAGE_DRIVER` | Set buildah storage driver | string | overlay |
| `DOCKERFILE` | Path to the Dockerfile to build. | string | ./Dockerfile |
| `CONTEXT` | Path to the directory to use as context. | string | . |
| `TLSVERIFY` | Verify the TLS on the registry endpoint (for push/pull to a non-TLS registry) | string | true |
| `FORMAT` | The format of the built container, oci or docker | string | oci |
| `BUILD_EXTRA_ARGS` | Extra parameters passed for the build command when building images. | string |  |
| `PUSH_EXTRA_ARGS` | Extra parameters passed for the push command when pushing images. | string |  |
| `BUILD_IMAGE` | Flag specifying whether image should be built again. | string | true |
| `IMAGE_REGISTRY` | Image registry url. | string |  |
| `CURRENT_GIT_TAG` | Current version of the application/image in dev. | string |  |


### Results
| Name | Description | Type |
|------|-------------|------|
| IMAGE_DIGEST | Digest of the image just built. | <no value>


### Changelog
<nil>
[]