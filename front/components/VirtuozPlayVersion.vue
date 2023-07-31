<script setup lang="ts">
import GraphQL from '@/components/GraphQL.vue';
import type { ApolloQueryResult } from '@apollo/client/core/types';
import { GetVirtuozPlayVersionDocument, GetVirtuozPlayVersionQuery } from '@/gql/queries/GetVirtuozPlayVersion';
</script>

<template>
    <GraphQL :query="GetVirtuozPlayVersionDocument">
        <template #default="{ loading, error, data }: ApolloQueryResult<GetVirtuozPlayVersionQuery>">
            <!-- Loading -->
            <div v-if="loading" class="loading apollo">Loading...</div>

            <!-- Error -->
            <div v-else-if="error" class="error apollo">An error occurred: {{ error }}</div>

            <!-- Result -->
            <div v-else-if="data" class="result apollo">VirtuozPlay version {{ data.virtuozPlay.version }}</div>

            <!-- No result -->
            <div v-else class="no-result apollo">No result :(</div>
        </template>
    </GraphQL>
</template>
