<script setup lang="ts">
import GraphQL from '@/components/GraphQL.vue';
import { GetSongDocument, GetSongQuery } from '@/gql/queries/GetSong';
import type { ApolloQueryResult } from '@apollo/client/core/types';
</script>

<template>
    <GraphQL :query="GetSongDocument">
        <template #default="{ data }: ApolloQueryResult<GetSongQuery>">
            <div v-if="data" class="result apollo">VirtuozPlay version {{ data.songs[0] }}</div>
        </template>
    </GraphQL>
</template>
