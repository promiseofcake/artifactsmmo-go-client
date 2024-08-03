# artifactsmmo-go-client

Generated Go Client based upon <https://docs.artifactsmmo.com/api_guide/openapi_spec>

## Usage

You can use `/client` as a generated API Client for your own projects.

I downgraded the OpenAPI spec to 3.0 in order to use <https://github.com/oapi-codegen/oapi-codegen/v2> for a cleaner client.

## Updating Spec

For now updating the spec is relying on specific versions of tools because I am
not automating this.

* [openapi-down-convert](https://github.com/apiture/openapi-down-convert)
* sed (FreeBSD)
* jq -- cross platform
