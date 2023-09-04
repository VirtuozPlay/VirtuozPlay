<script setup lang="ts">
import { Line, Pie } from 'vue-chartjs';
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
import Shadow from '@/utilities/chart/shadow';
import { GetPostPerformanceStatsQuery } from '@/gql/queries/GetPostPerformanceStats';
import { computed, ComputedRef } from 'vue';

ChartJS.register(ArcElement, CategoryScale, LinearScale, PointElement, LineElement, Title, Tooltip, Legend, Shadow);

const props = defineProps<{
    performance: NonNullable<GetPostPerformanceStatsQuery['performance']>;
}>();

const GOOD_NOTE_THRESHOLD = 0.9;
const BAD_NOTE_THRESHOLD = 0.5;

const timeFormat = new Intl.DateTimeFormat('fr-FR', {
    minute: 'numeric',
    second: 'numeric',
});
const percentFormat = new Intl.NumberFormat('fr-FR', {
    style: 'percent',
    minimumFractionDigits: 0,
    maximumFractionDigits: 0,
});

const averagePrecision: ComputedRef<number> = computed(() => {
    return props.performance.notes.reduce((sum, note) => sum + note.precision, 0) / props.performance.notes.length;
});

const pieData: ComputedRef<ChartData<'pie'>> = computed(() => {
    const counts = props.performance.notes.reduce(
        ([good, bad, miss], note) => {
            if (note.precision >= GOOD_NOTE_THRESHOLD) {
                ++good;
            } else if (note.precision <= BAD_NOTE_THRESHOLD) {
                ++miss;
            } else {
                ++bad;
            }
            return [good, bad, miss];
        },
        [0, 0, 0]
    );
    const data = counts.map((count) => count / props.performance.notes.length);
    return {
        labels: ['Bien joué!', 'Pas terrible', 'Raté'],
        datasets: [
            {
                data,
                backgroundColor: ['#74DBA3', '#FFCC15', '#E9121C'],
            },
        ],
    };
});

const pieOptions: ChartOptions<'pie'> = {
    responsive: true,
    maintainAspectRatio: true,
    cutout: '40%',
    borderColor: '#4C324D',
    plugins: {
        legend: {
            display: false,
        },
        shadow: {
            enable: true,
            color: '#4C324D',
            offsetX: 32,
            offsetY: 32,
        },
    },
    layout: {
        padding: 16,
    },
};

const lineData: ComputedRef<ChartData<'line'>> = computed(() => {
    return {
        datasets: [
            {
                label: 'Précision',
                data: props.performance.notes.map((note) => ({
                    x: note.at,
                    y: note.precision,
                })),
                fill: false,
                borderColor: 'rgb(75, 192, 192)',
                tension: 0.1,
            },
        ],
    };
});

const lineOptions: ChartOptions<'line'> = {
    responsive: true,
    maintainAspectRatio: false,
    plugins: {
        tooltip: {
            enabled: true,
            displayColors: false,
            callbacks: {
                title: () => '',
                label: (point) => percentFormat.format(Number(point.formattedValue)),
            },
        },
    },
    scales: {
        y: {
            min: 0,
            max: 1,
            ticks: {
                stepSize: 0.2,
                callback: (value) => percentFormat.format(Number(value)),
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
                sampleSize: 10,
                callback: (value) => timeFormat.format(new Date(Number(value))),
            },
            border: {
                color: '#4C324D',
                width: 2,
            },
        },
    },
};
</script>

<template>
    <div class="flex flex-row flex-wrap gap-3 my-4">
        <BigStatistic name="temps" :value="timeFormat.format(new Date(props.performance.duration))" />
        <BigStatistic name="précision" :value="percentFormat.format(averagePrecision)" />
        <BigStatistic v-if="props.performance.author" name="auteur" :value="props.performance.author.name" />
    </div>
    <div class="flex flex-wrap justify-center my-4">
        <div class="relative max-w-[80vmin]">
            <Pie :id="'misses-' + props.performance.id" :options="pieOptions" :data="pieData" />
        </div>
        <div class="relative min-w-[50vw] grow">
            <Line :id="'precision-over-time-' + props.performance.id" :options="lineOptions" :data="lineData" />
        </div>
    </div>
</template>
