<script setup lang="ts">
import GraphQL from '@/components/GraphQL.vue';
import { GetSongDocument, GetSongQuery } from '@/gql/queries/GetSong';
import type { QueryResult } from '@apollo/client';
</script>

<template>
    <GraphQL :query="GetSongDocument">
        <template #default="{ data }: QueryResult<GetSongQuery>">
            <div v-if="data" class="result apollo">VirtuozPlay version {{ data.songs[0] }}</div>
        </template>
    </GraphQL>
</template>
