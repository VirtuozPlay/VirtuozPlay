<script setup lang="ts">
import {
    GetPostPerformanceStatsDocument,
    GetPostPerformanceStatsQuery,
    GetPostPerformanceStatsQueryVariables,
} from '@/gql/queries/GetPostPerformanceStats';
import type { ApolloQueryResult } from '@apollo/client/core/types';
import GraphQL from '@/components/GraphQL.vue';
import BigStatistic from '@/components/stats/BigStatistic.vue';

defineProps<{
    performanceId: string;
}>();

const timeFormat = new Intl.DateTimeFormat('fr-FR', {
    minute: 'numeric',
    second: 'numeric',
});
const percentFormat = new Intl.NumberFormat('fr-FR', {
    style: 'percent',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0,
});
</script>

<template>
    <GraphQL
        :query="GetPostPerformanceStatsDocument"
        :variables="{ id: performanceId } satisfies GetPostPerformanceStatsQueryVariables"
        tag="main"
        aria-label="Post Performance Stats"
        class="mx-auto my-8"
    >
        <template #default="{ loading, error, data }: ApolloQueryResult<GetPostPerformanceStatsQuery>">
            <h2 class="text-3xl font-bold">Results</h2>
            <div v-if="loading">Loading...</div>
            <div v-else-if="error">
                <p class="text-red-900">Le résulat demandé n'existe pas ou à été supprimé</p>
                <RouterLink :to="{ name: 'collection' }">Retour a la collection</RouterLink>
            </div>

            <RouterLink
                v-if="data && data.performance"
                :to="{ name: 'collection', params: { songIdOrName: performanceId } }"
                class="text-2xl font-bold hover:underline"
                >{{ data.performance.song.title }}
            </RouterLink>

            <div class='flex flex-row gap-4 my-2' v-if='data && data.performance'>
                <BigStatistic name="temps" :value="timeFormat.format(new Date(data.performance.duration * 100))" />
                <BigStatistic name="auteur" v-if='data.performance.author' :value="data.performance.author.name" />
                <BigStatistic name="precision" :value="percentFormat.format(0.97)" />
            </div>
        </template>
    </GraphQL>
</template>
