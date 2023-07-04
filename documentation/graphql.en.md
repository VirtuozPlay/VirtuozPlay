# GraphQL API Documentation

[🇫🇷 Version Française](graphql.en.md)

## Introduction

Ce project utilise une API GraphQL pour communiquer avec le front.
La différence entre une API REST et une API GraphQL est que cette dernière permet au frontend de demander uniquement les
données dont il a besoin, et rien de plus.
Nous n'avons besoin de définir le schéma de l'API qu'une seule fois, et le frontend peut demander les données dont il a
besoin, sans avoir à créer un nouvel endpoint pour chaque requête.

The API is available at the `/graphql` endpoint.

## Schema

The schema is defined in the `graph/schema.graphqls` file and is the specification of all the data available to the
frontend.
Here is an example of a schema that exposes a `virtuozPlay` object at the root of the API, which contains a `version`:

```graphql
type Query {
  virtuozPlay: VirtuozPlay!
}

type VirtuozPlay {
  version: String!
}
```

## Backend

The backend is responsible for implementing the schema defined in the `graph/schema.graphqls` file.
To do this, we implement query resolvers in the `graph/schema.resolvers.go` file, note that this file is partially
auto-generated any only the body of the resolver functions should be modified.

To create resolvers for the schema, you need to run the following command:

```shell
buffalo task gqlgen:generate
```

This runs [gqlgen](https://gqlgen.com/) against the schema which will automatically generate the corresponding Go types
as well resolvers stubs.
You may then implement the resolvers in the `graph/schema.resolvers.go` file.

### Example

After running `buffalo task gqlgen:generate`, you should see the following code in the `graph/schema.resolvers.go` file:

```go
// [...]

// VirtuozPlay is the resolver for the virtuozPlay field.
func (r *queryResolver) VirtuozPlay(ctx context.Context) (*model.VirtuozPlay, error) {
    panic(fmt.Errorf("not implemented: VirtuozPlay - virtuozPlay"))
}

// [...]
```

Replace the `panic` call with the following code:

```go
	return &model.VirtuozPlay{Version: "0.1.0"}, nil
```

This will make a query on `virtuozPlay.version` query return the following JSON:

```json
{
  "data": {
    "virtuozPlay": {
      "version": "0.1.0"
    }
  }
}
```

## Frontend

To consume the API from the frontend, we use the [Apollo Client](https://www.apollographql.com/docs/react/) library.

When you need to make a query or mutation:

1. Create a GraphQL document in the `front/gql/queries` or `front/gql/mutations` folders.
2. Run `yarn codegen` to create the corresponding TypeScript types and `DocumentNode` object.
3. Use the generated `DocumentNode` object in your Vue component
   with [&lt;GraphQL&gt; wrapper component](/front/components/GraphQL.vue) We make (indirect) (`<GraphQL>` is a thin
   wrapper around the [&lt;ApolloQuery&gt; component](https://v4.apollo.vuejs.org/guide-components/) to make it more TypeScript-friendly).
4. Use the `QueryResult` object in your template.

### Example

Create the following query in `front/gql/queries/GetVirtuozPlayVersion.graphql`:

```graphql
query GetVirtuozPlayVersion {
  virtuozPlay {
    version
  }
}
```

Run `yarn codegen` to generate the corresponding TypeScript types and `DocumentNode` object.
There should be a file at `front/gql/queries/GetVirtuozPlayVersion.ts` with a content similar to this:

```typescript
// ...

export declare const GetVirtuozPlayVersion: import('graphql').DocumentNode;

export const GetVirtuozPlayVersionDocument = gql`
  query GetVirtuozPlayVersion {
    virtuozPlay {
      version
    }
  }
`;

// ...
```

Now your query is ready to use in your Vue component:

```vue
<script setup lang="ts">
import GraphQL from '@/components/GraphQL.vue';
import type { QueryResult } from '@apollo/client';
import { GetVirtuozPlayVersionDocument, GetVirtuozPlayVersionQuery } from '@/gql/queries/GetVirtuozPlayVersion';
</script>

<template>
  <GraphQL :query="GetVirtuozPlayVersionDocument">
    <!-- The result of the query is available in the default slot, you should always put the correct type in QueryResult<...> -->
    <template #default="{ loading, error, data }: QueryResult<GetVirtuozPlayVersionQuery>">
      <!-- Loading indicator, optional -->
      <div v-if="loading" class="loading apollo">Loading...</div>

      <!-- Error indicator, optional -->
      <div v-else-if="error" class="error apollo">An error occurred: {{ error }}</div>

      <!-- Result -->
      <div v-else-if="data" class="result apollo">VirtuozPlay version {{ data.virtuozPlay.version }}</div>

      <!-- No result -->
      <div v-else class="no-result apollo">No result :(</div>
    </template>
  </GraphQL>
</template>
```

If the backend is up and running, you should now see `VirtuozPlay version 0.1.0` in your browser.
