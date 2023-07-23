<script setup lang="ts">
import { ref } from 'vue';
import ListenPlayPause from '@/components/Playground/ListenPlayPause.vue';
import StringsFrets from '@/components/Playground/StringsFrets.vue';
import { useSongStore } from '@/store';

const isPlaying = ref(false);
const stringsFretsRef = ref(null);

interface Position {
    string: number;
    fret: number;
    start: number;
}

const store = useSongStore();
console.log('getSong', store.getCurrentSong);
const title = store.getCurrentSong.title;
const positions: Position[] = store.getCurrentSong.notes.map((note: any) => ({
    string: note.string,
    fret: note.fret,
    start: note.start,
}));

console.log('positions', positions);

const currentIndex = ref(0);
const animationRunning = ref(false);
const animationPaused = ref(false);

const handleListenCanCan = () => {
    if (!isPlaying.value) {
        console.log('I listen to CanCan');
        // audio.play();
        isPlaying.value = true;
    } else {
        console.log('CanCan is paused');
        // audio.pause();
        isPlaying.value = false;
    }
};

const isPosition = (string: number, fret: number): boolean => {
    const currentPosition = positions[currentIndex.value];
    return string === currentPosition.string && fret === currentPosition.fret;
};

const startAnimation = () => {
    let currentTime = 0;
    console.log('Start animation');
    if (animationRunning.value) {
        console.log('Animation already in progress -> exit');
        return;
    }
    if (animationPaused.value) {
        animationPaused.value = false;
    } else {
        // restart the animation from the beginning
        currentIndex.value = 0;
        currentTime = 0;
    }

    const animate = () => {
        const currentPosition = positions[currentIndex.value];
        const nextPosition = positions[(currentIndex.value + 1) % positions.length];

        const start = currentPosition.start;
        const end = nextPosition.start;
        const duration = end - start;

        console.log(start);

        currentTime += duration;
        currentIndex.value = (currentIndex.value + 1) % positions.length;

        if (!animationPaused.value) {
            setTimeout(animate, duration);
        } else {
            animationRunning.value = false;
        }
    };

    animate();

    animationRunning.value = true;
};

const pauseAnimation = () => {
    console.log('Pause animation');
    animationRunning.value = false;
    animationPaused.value = true;
};

const stopAnimation = () => {
    console.log('Stop animation');
    animationRunning.value = false;
    animationPaused.value = true;
};

const isCurrentFret = (fret: number) => {
    return positions[currentIndex.value].fret === fret;
};
</script>

<template>
    <main aria-label="songview section" class="mt-16 w-80vw mx-auto">
        <div>
            <h1>{{ title }}</h1>
            <div v-for="string in 1" :key="string" class="flex flex-row flex-wrap w-full justify-around text-gray-900">
                <div v-for="fret in 14" :key="fret" :class="{ 'anim-fret': isCurrentFret(fret) }">
                    {{ fret }}
                </div>
            </div>
            <div class="relative">
                <div class="absolute top-0 left-0 w-full h-full" style="margin-top: 26px">
                    <StringsFrets ref="stringsFretsRef" :is-position="isPosition" />
                </div>
                <img alt="le manche" src="@/assets/lemanche.svg" />
            </div>

            <ListenPlayPause
                :on-listen="handleListenCanCan"
                :on-play="startAnimation"
                :on-pause="pauseAnimation"
                :on-stop="stopAnimation"
                :animation-running="animationRunning"
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
