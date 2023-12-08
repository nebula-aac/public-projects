/**
 * @generated SignedSource<<29a9fe4ec8db0fa302eb81e07d97c8c4>>
 * @lightSyntaxTransform
 * @nogrep
 */

/* tslint:disable */
/* eslint-disable */
// @ts-nocheck

import { ConcreteRequest, Query } from 'relay-runtime';
export type UuidQuery$variables = {
  name: string;
};
export type UuidQuery$data = {
  readonly findNinja: {
    readonly name: string | null | undefined;
    readonly rank: string | null | undefined;
  } | null | undefined;
};
export type UuidQuery = {
  response: UuidQuery$data;
  variables: UuidQuery$variables;
};

const node: ConcreteRequest = (function(){
var v0 = [
  {
    "defaultValue": null,
    "kind": "LocalArgument",
    "name": "name"
  }
],
v1 = [
  {
    "alias": null,
    "args": [
      {
        "kind": "Variable",
        "name": "name",
        "variableName": "name"
      }
    ],
    "concreteType": "Ninja",
    "kind": "LinkedField",
    "name": "findNinja",
    "plural": false,
    "selections": [
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "name",
        "storageKey": null
      },
      {
        "alias": null,
        "args": null,
        "kind": "ScalarField",
        "name": "rank",
        "storageKey": null
      }
    ],
    "storageKey": null
  }
];
return {
  "fragment": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Fragment",
    "metadata": null,
    "name": "UuidQuery",
    "selections": (v1/*: any*/),
    "type": "Query",
    "abstractKey": null
  },
  "kind": "Request",
  "operation": {
    "argumentDefinitions": (v0/*: any*/),
    "kind": "Operation",
    "name": "UuidQuery",
    "selections": (v1/*: any*/)
  },
  "params": {
    "cacheID": "e01f530eb9d7c97e5ea620de889d3152",
    "id": null,
    "metadata": {},
    "name": "UuidQuery",
    "operationKind": "query",
    "text": "query UuidQuery(\n  $name: String!\n) {\n  findNinja(name: $name) {\n    name\n    rank\n  }\n}\n"
  }
};
})();

(node as any).hash = "486c79f495a47b34ac4eebee8cbb5c61";

export default node;
