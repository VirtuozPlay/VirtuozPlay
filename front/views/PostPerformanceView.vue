<script setup lang="ts">
import {
    GetPostPerformanceStatsDocument,
    GetPostPerformanceStatsQuery,
    GetPostPerformanceStatsQueryVariables,
} from '@/gql/queries/GetPostPerformanceStats';
import type { ApolloQueryResult } from '@apollo/client/core/types';
import GraphQL from '@/components/GraphQL.vue';

defineProps<{
    performanceId: string;
}>();
</script>

<template>
    <p>Performance id is {{ performanceId }}</p>

    <GraphQL
        :query="GetPostPerformanceStatsDocument"
        :variables="{ id: performanceId } satisfies GetPostPerformanceStatsQueryVariables"
    >
        <template #default="{ loading, error, data }: ApolloQueryResult<GetPostPerformanceStatsQuery>">
            <div v-if="loading">Loading...</div>

            <!-- Error -->
            <div v-else-if="error">An error occurred: {{ error }}</div>

            <!-- Result -->
            <ul v-else-if="data && data.performance">
                <li>Song Title: {{ data.performance.song.title }}</li>
                <li>Duration: {{ data.performance.duration }}</li>
                <li>Author: {{ data.performance.author?.name ?? 'No author' }}</li>
            </ul>
        </template>
    </GraphQL>
</template>
