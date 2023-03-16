# Test Data Import

## 1. Confluence

JSON from TestData Generator (custom format mixed with AAS format)
https://confluence.catena-x.net/display/PL/BoMAsPlanned+test+dataset

1. Take from https://catenax.io/schema/TestDataContainer/1.0.0

2. Filter for BPNL, take these objects
"bpnl": "BPNL00000003AYRE", (e.g. line 38)
 -> Inside this objects there are the sub models, that must be registered in the AAS registry

E.g.
"urn:bamm:io.catenax.part_as_planned:1.0.0#PartAsPlanned": [
        {
          "validityPeriod": {
            "validFrom": "2013-04-11T05:30:04.000Z",
            "validTo": "2025-04-23T19:59:03.000Z"
          },
          "catenaXId": "urn:uuid:4f7b1cf2-a598-4027-bc78-63f6d8e55699",
          "partTypeInformation": {
            "manufacturerPartId": "7A047C7-01",
            "classification": "product",
            "nameAtManufacturer": "N Tier A CathodeMaterial"
          }
        }
      ],

> register "urn:bamm:io.catenax.part_as_planned:1.0.0#PartAsPlanned" in registry and edc
> use line 16 to 27 is the actual data

3. For each submodel
-> import into EDC as Asset, PolicyDefition, ContractDefinition
-> import into aas registry as shell (containing these sub models)


Asset:
- CatenaX ID is first part of asset ID. Second part is random UUID

Payload for registry:

global asset id is the catena-x (shell id)
identifaction is the most important ID
localIdentifier in specificAssetIds


endpoints
base path edc + asset id + default suffix

https://app.swaggerhub.com/apis/BaSyx/
basyx_asset_administration_shell_repository_http_rest_api/v1


for registry import

registry_url="https://semantics.int.demo.catena-x.net/registry/registry"

token_url="https://centralidp.int.demo.catena-x.net/auth/realms/CX-Central/protocol/openid-connect/token"

send aas payload with POST to 
curl --location --request POST "${registry_url}/shell-descriptors" \ --silent --output /dev/null --write-out %{http_code} \ --header "Content-Type: application/json" \ --header "Authorization: Bearer ${access_token}" \ -d @"$/registry.json"

