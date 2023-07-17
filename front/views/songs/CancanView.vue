<script setup lang="ts">
import { ref } from 'vue';
import ListenPlayPause from '@/components/Playground/ListenPlayPause.vue';
import StringsFrets from '@/components/Playground/StringsFrets.vue';

const audio = new Audio('../../../assets/music/cancan/cancan.mp3');
const isPlaying = ref(false);
const stringsFretsRef = ref(null);

interface Position {
    row: number;
    column: number;
    time: number;
}

let currentIndex = ref(0);
let timer: ReturnType<typeof setTimeout> | null = null;
let animationRunning = ref(false);
let animationPaused = ref(false);

const positions: Position[] = [
    { row: 1, column: 10, time: 1000 },
    { row: 2, column: 11, time: 1500 },
    { row: 1, column: 10, time: 2000 },
    { row: 2, column: 11, time: 2500 },
    { row: 5, column: 3, time: 3000 },
    { row: 6, column: 4, time: 4500 },
    { row: 5, column: 3, time: 4000 },
    { row: 6, column: 4, time: 4500 },
];

const handleListenCanCan = () => {
    if (!isPlaying.value) {
        console.log('I listen to CanCan');
        audio.play();
        isPlaying.value = true;
    } else {
        console.log('CanCan is paused');
        audio.pause();
        isPlaying.value = false;
    }
};

const isPosition = (row: number, column: number): boolean => {
    const currentPosition = positions[currentIndex.value];
    return row === currentPosition.row && column === currentPosition.column;
};
const startAnimation = () => {
    if (animationRunning.value) {
        console.log('Animation already in progress -> exit');
        return;
    }
    if (animationPaused.value) {
        animationPaused.value = false;
    } else {
        // restart the animation from the beginning
        currentIndex.value = 0;
    }

    animationRunning.value = true;
    console.log('Start animation');
    timer = setInterval(() => {
        // at the end of the song -> loop the music
        const nextIndex = currentIndex.value + 1 >= positions.length ? 0 : currentIndex.value + 1;
        currentIndex.value = nextIndex;
    }, positions[currentIndex.value].time); // time = timeout for each execution
};

const pauseAnimation = () => {
    console.log('Pause animation');
    animationRunning.value = false;
    clearInterval(timer!);
    timer = null;
    animationPaused.value = true;
};

const stopAnimation = () => {
    console.log('Stop animation');
    animationRunning.value = false;
    clearInterval(timer!);
    currentIndex.value = 0;
};

const isCurrentColumn = (column: number) => {
    return positions[currentIndex.value].column === column;
};
</script>

<template>
    <main aria-label="checkup section" class="mt-16 w-80vw mx-auto">
        <div>
            <h1>Can Can</h1>
            <div v-for="row in 1" :key="row" class="flex-container w-full test">
                <div v-for="column in 14" :key="column" :class="{ 'bold-column': isCurrentColumn(column) }">
                    {{ column }}
                </div>
            </div>
            <div class="relative">
                <div class="absolute top-0 left-0 w-full h-full" style="margin-top: 26px">
                    <StringsFrets :isPosition="isPosition" ref="stringsFretsRef" />
                </div>
                <img alt="le manche" src="@/assets/lemanche.svg" />
            </div>
            <ListenPlayPause
                :onListen="handleListenCanCan"
                :onPlay="startAnimation"
                :onPause="pauseAnimation"
                :onStop="stopAnimation"
                :animationRunning="animationRunning"
            />
        </div>
    </main>
</template>

<style>
.flex-container {
    display: flex;
    flex-flow: row wrap;
    width: 100%;
    justify-content: space-around;
    padding: 0;
    margin: 0;
}

.test {
    color: black;
}

.bold-column {
    font-weight: bold;
    font-size: 18px;
    color: blueviolet;
}
</style>
