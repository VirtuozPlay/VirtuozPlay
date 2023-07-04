# Documentation API GraphQL

[🇬🇧 English Version](graphql.fr.md)

## Introduction

Ce projet utilise une API GraphQL pour communiquer avec le front.
La différence entre une API REST et une API GraphQL est que cette dernière permet au frontend de demander uniquement les
données dont il a besoin, et rien de plus.
Nous n'avons besoin de définir le schéma de l'API qu'une seule fois, et le frontend peut demander les données dont il a
besoin, sans avoir à créer un nouvel endpoint pour chaque requête.

L'API est disponible à l'endpoint `/graphql`.

## Schema

Le schéma est défini dans le fichier `graph/schema.graphqls` et est la spécification de toutes les données disponibles
Voci un exemple de schéma qui expose un objet `virtuozPlay` à la racine de l'API, qui contient une `version`:

```graphql
type Query {
  virtuozPlay: VirtuozPlay!
}

type VirtuozPlay {
  version: String!
}
```

## Backend

Le backend est responsable de l'implémentation du schéma défini dans le fichier `graph/schema.graphqls`.
Pour ce faire, nous implémentons des resolvers de requête dans le fichier `graph/schema.resolvers.go`, notez que ce fichier
est partiellement auto-généré et que seul le corps des fonctions de résolution doit être modifié.

Pour créer des résolveurs pour le schéma, vous devez exécuter la commande suivante:

```shell
buffalo task gqlgen:generate
```

Cela exécute [gqlgen](https://gqlgen.com/) sur le schéma qui générera automatiquement les types Go correspondants ainsi
que les corps vides des résolveurs.
Vous pouvez ensuite implémenter les résolveurs dans le fichier `graph/schema.resolvers.go`.

### Exemple

Après avoir exécuté `buffalo task gqlgen:generate`, vous devriez voir le code suivant dans le fichier `graph/schema.resolvers.go` :

```go
// [...]

// VirtuozPlay is the resolver for the virtuozPlay field.
func (r *queryResolver) VirtuozPlay(ctx context.Context) (*model.VirtuozPlay, error) {
    panic(fmt.Errorf("not implemented: VirtuozPlay - virtuozPlay"))
}

// [...]
```

Remplacez l'appel à `panic` par le code suivant :

```go
	return &model.VirtuozPlay{Version: "0.1.0"}, nil
```

Celà fera en sorte qu'une requête sur `virtuozPlay.version` retourne le JSON suivant :

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

Pour consommer l'API depuis le frontend, nous utilisons la bibliothèque [Apollo Client](https://www.apollographql.com/docs/react/).

Quand vous avez besoin de faire une requête ou une mutation :

1. Créez un document GraphQL dans les dossiers `front/gql/queries` ou `front/gql/mutations`.
2. Exécutez `yarn codegen` pour créer les types TypeScript correspondants et l'objet `DocumentNode`.
3. Utilisez l'objet `DocumentNode` généré dans votre composant Vue avec le composant [&lt;GraphQL&gt;](/front/components/GraphQL.vue) (`<GraphQL>` est un wrapper fin autour du composant [&lt;ApolloQuery&gt;](https://v4.apollo.vuejs.org/guide-components/) pour le rendre plus TypeScript-friendly).
4. Utilisez l'objet `QueryResult` dans votre template.

### Example

Créez la requête suivante dans `front/gql/queries/GetVirtuozPlayVersion.graphql` :

```graphql
query GetVirtuozPlayVersion {
  virtuozPlay {
    version
  }
}
```

Lancez `yarn codegen` pour générer les types TypeScript correspondants et l'objet `DocumentNode`.
Il devrait y avoir un fichier `front/gql/queries/GetVirtuozPlayVersion.ts` avec un contenu similaire à celui-ci :

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

Maintenant, vous pouvez utiliser l'objet `DocumentNode` dans votre composant Vue :

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

Si vous lancez l'application avec `buffalo dev`, vous devriez voir `VirtuozPlay version 0.1.0` dans votre navigateur.
