<script setup lang="ts">
import GraphQL from '@/GraphQL.vue';
import type { QueryResult } from '@apollo/client';
</script>

<template>
    <GraphQL
        :query="
            (gql) => gql`
                {
                    virtuozPlay {
                        version
                    }
                }
            `
        "
    >
        <template v-slot="{ loading, error, data }: QueryResult<{ virtuozPlay: { version: string } }>">
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
