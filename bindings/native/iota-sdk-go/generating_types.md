The IOTA SDK offers a message bus which is working soley with JSON messages.

The Go library is a PoC and mainly concentrates on the wallet/ledger functionality.

To bootstrap the library quicker, all typed JSON models are taken from the TypeScript binding.

The models were exported as a JSON schema, which can be used to generate Go structs.

It is most likely not 100% compatible or without faults, and should therefore not be used in a production environment. 

Expect to find mismatches in different number types (int vs uint vs bigint) as the Json Schema only supports `number`.

It is possible to define a `minimum` of `0` which schema generators might use, but this here doesn't.


# Generating the Go types

## Create a JSON Schema from the TypeScript types

```bash
cd <root>/bindings/nodejs/
npm i
cd <root>/bindings/native/iota-sdk-go
```

### Create the schema 

`npx ts-json-schema-generator --path '../../nodejs/lib/types/**/*.ts' -f "../../nodejs/tsconfig.json" -o schema.json`

Patch the `schema.json` file. Remove:

```
"nativeTokens": {
    "items": {
      "items": [
        {
          "type": "string"
        },
        {
          "$ref": "#/definitions/HexEncodedAmount"
        }
      ],
      "maxItems": 2,
      "minItems": 2,
      "type": "array"
    },
    "type": "array"
  },

```
**in line ~4019** at `SendNativeTokensParams` 

otherwise the next step will fail. 

The nativeTokens object needs to be added manually to the `types.go`

## Create Go structs

```bash
go install github.com/atombender/go-jsonschema/cmd/gojsonschema@latest
gojsonschema schema.json -p iota_sdk_go -o types.go
```

It will warn about duplicates, but the types are complete.