<script setup lang="ts">
import {
    GetPostPerformanceStatsDocument,
    GetPostPerformanceStatsQuery,
    GetPostPerformanceStatsQueryVariables,
} from '@/gql/queries/GetPostPerformanceStats';
import type { ApolloQueryResult } from '@apollo/client/core/types';
import { Line, Pie } from 'vue-chartjs';
import GraphQL from '@/components/GraphQL.vue';
import BigStatistic from '@/components/stats/BigStatistic.vue';
import {
    ArcElement,
    CategoryScale,
    Chart as ChartJS,
    ChartData,
    ChartOptions,
    Legend,
    LinearScale,
    LineElement,
    PointElement,
    Title,
    Tooltip,
} from 'chart.js';
import TextualButton from '@/components/inputs/TextualButton.vue';

ChartJS.register(ArcElement, CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend);

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

const chartData = {
    labels: ['Bien joué!', 'Pas terrible', 'Raté'],
    datasets: [
        {
            data: [50, 30, 20],
            backgroundColor: ['#74DBA3', '#FFCC15', '#E9121C'],
        },
    ],
} satisfies ChartData<'pie'>;
const chartOptions = {
    responsive: true,
    maintainAspectRatio: true,
    plugins: {
        legend: {
            position: 'bottom',
        },
    },
    cutout: '50%',
    borderColor: '#4C324D',
} satisfies ChartOptions<'pie'>;

const lineData = {
    datasets: [
        {
            label: 'Précision',
            data: [
                {
                    x: 0,
                    y: 0.5,
                },
                {
                    x: 1_000,
                    y: 0.8,
                },
                {
                    x: 31_000,
                    y: 0.9,
                },
            ],
            fill: false,
            borderColor: 'rgb(75, 192, 192)',
            tension: 0.1,
        },
    ],
} satisfies ChartData<'line'>;

const lineOptions = {
    responsive: true,
    maintainAspectRatio: false,
    scales: {
        y: {
            min: 0,
            max: 1,
            ticks: {
                stepSize: 0.2,
                callback: (value) => Number(value) * 100 + ' %',
            },
            border: {
                color: '#4C324D',
                width: 2,
            },
        },
        x: {
            min: 0,
            type: 'linear',
            ticks: {
                stepSize: 30_000,
                callback: (value) => Number(value) / 1000 + ' s',
            },
            border: {
                color: '#4C324D',
                width: 2,
            },
        },
    },
} satisfies ChartOptions<'line'>;
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

            <div v-if="data && data.performance" class="flex flex-row flex-wrap gap-3 my-4">
                <BigStatistic name="temps" :value="timeFormat.format(new Date(data.performance.duration * 100))" />
                <BigStatistic name="précision" :value="percentFormat.format(0.97)" />
                <BigStatistic name="auteur" v-if="data.performance.author" :value="data.performance.author.name" />
            </div>
            <div class="flex flex-wrap justify-center">
                <div class="relative max-w-[80vmin]">
                    <Pie :id="'misses-' + performanceId" :options="chartOptions" :data="chartData" />
                </div>
                <div class="relative min-w-[50vw] grow">
                    <Line :id="'precision-over-time-' + performanceId" :options="lineOptions" :data="lineData" />
                </div>
            </div>

            <div class='w-full flex flex-wrap justify-center'>
                <TextualButton aria-label="Réessayer" hover-color="#FAFF00">Réessayer</TextualButton>
            </div>
        </template>
    </GraphQL>
</template>
