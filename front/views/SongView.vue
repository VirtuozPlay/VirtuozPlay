<script setup lang="ts">
import { ref } from 'vue';
import ListenPlayPause from '@/components/Playground/ListenPlayPause.vue';
import StringsFrets from '@/components/Playground/StringsFrets.vue';
import router from '../router/index';

const isPlaying = ref(false);
const stringsFretsRef = ref(null);

interface Position {
    string: number;
    fret: number;
    time: number;
}

const data = () => {
    let title = '';
    let music = '';
    let positions: Position[] = [];

    if (router.currentRoute.value.name === 'sting') {
        title = 'sting';
        music = '/assets/music/sting/sting.mp3';
        positions = [
            { string: 4, fret: 1, time: 1000 },
            { string: 4, fret: 1, time: 1500 },
            { string: 4, fret: 2, time: 2000 },
            { string: 3, fret: 2, time: 2500 },
            { string: 3, fret: 3, time: 3000 },
            { string: 3, fret: 3, time: 4500 },
            { string: 1, fret: 4, time: 4000 },
            { string: 1, fret: 4, time: 4500 },
        ];
    } else if (router.currentRoute.value.name === 'cancan') {
        title = 'cancan';
        music = '/assets/music/cancan/cancan.mp3';
        positions = [
            { string: 1, fret: 10, time: 1000 },
            { string: 2, fret: 11, time: 200 },
            { string: 1, fret: 10, time: 2500 },
            { string: 2, fret: 11, time: 3000 },
            { string: 5, fret: 3, time: 3500 },
            { string: 6, fret: 4, time: 4500 },
            { string: 5, fret: 3, time: 4000 },
            { string: 6, fret: 4, time: 4500 },
        ];
    }

    return { title, music, positions };
};

const audio = new Audio(data().music);

let currentIndex = ref(0);
let timer: ReturnType<typeof setTimeout> | null = null;
let animationRunning = ref(false);
let animationPaused = ref(false);

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

const isPosition = (string: number, fret: number): boolean => {
    const currentPosition = data().positions[currentIndex.value];
    return string === currentPosition.string && fret === currentPosition.fret;
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
        const nextIndex = currentIndex.value + 1 >= data().positions.length ? 0 : currentIndex.value + 1;
        currentIndex.value = nextIndex;
    }, data().positions[currentIndex.value].time); // time = timeout for each execution
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

const isCurrentFret = (fret: number) => {
    return data().positions[currentIndex.value].fret === fret;
};
</script>

<template>
    <main aria-label="sting section" class="mt-16 w-80vw mx-auto">
        <div>
            <h1>{{ data().title }}</h1>
            <div v-for="string in 1" :key="string" class="flex flex-row flex-wrap w-full justify-around text-gray-900">
                <div v-for="fret in 14" :key="fret" :class="{ 'anim-fret': isCurrentFret(fret) }">
                    {{ fret }}
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
.anim-fret {
    font-weight: bold;
    font-size: 18px;
    color: #f4a11a;
}
</style>
