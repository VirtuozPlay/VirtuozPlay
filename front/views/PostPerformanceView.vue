<script setup lang="ts">
import {
    GetPostPerformanceStatsDocument,
    GetPostPerformanceStatsQuery,
    GetPostPerformanceStatsQueryVariables,
} from '@/gql/queries/GetPostPerformanceStats';
import type { ApolloQueryResult } from '@apollo/client/core/types';
import GraphQL from '@/components/GraphQL.vue';
import {
    ArcElement,
    CategoryScale,
    Chart as ChartJS,
    Legend,
    LinearScale,
    LineElement,
    PointElement,
    Title,
    Tooltip,
} from 'chart.js';
import TextualButton from '@/components/inputs/TextualButton.vue';
import Shadow from '@/utilities/chart/shadow';
import PostPerformanceStats from '@/components/PostPerformanceStats.vue';

ChartJS.register(ArcElement, CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend, Shadow);

defineProps<{
    performanceId: string;
}>();
</script>

<template>
    <GraphQL
        :query="GetPostPerformanceStatsDocument"
        :variables="{ id: performanceId } satisfies GetPostPerformanceStatsQueryVariables"
        tag="main"
        aria-label="Post Performance Stats"
        class="mx-4 md:mx-auto my-4 h-screen"
    >
        <template #default="{ loading, error, data }: ApolloQueryResult<GetPostPerformanceStatsQuery>">
            <h2 class="text-3xl font-bold">Résultats</h2>

            <div v-if="loading">Chargement...</div>
            <div v-else-if="error">
                <p class="text-red-900">Le résulat demandé n'existe pas ou à été supprimé</p>
                <RouterLink :to="{ name: 'collection' }">Retour a la collection</RouterLink>
            </div>

            <RouterLink
                v-if="data && data.performance"
                :to="{ name: 'collection', params: { songIdOrName: performanceId } }"
                class="text-2xl font-bold hover:underline"
            >
                {{ data.performance.song.title }}
            </RouterLink>

            <PostPerformanceStats v-if="data && data.performance" :performance="data.performance" />

            <div class="w-full flex flex-wrap justify-center">
                <TextualButton aria-label="Réessayer" hover-color="#FAFF00">Réessayer</TextualButton>
            </div>
        </template>
    </GraphQL>
</template>
