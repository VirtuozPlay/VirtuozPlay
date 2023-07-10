<script setup lang="ts">
import { ref } from 'vue';
import ListenPlayPause from '@/components/Playground/ListenPlayPause.vue';

const audio = new Audio('assets/music/cancan/cancan.mp3');
const isPlaying = ref(false);

const handleListenCanCan = () => {
    if (!isPlaying.value) {
        console.log("J'écoute la musique CanCan");
        audio.play();
        isPlaying.value = true;
    } else {
        console.log('La musique CanCan est en pause');
        audio.pause();
        isPlaying.value = false;
    }
};
const handlePlayCanCan = () => {
    console.log('Je lance la musique CanCan');
};

const handlePauseCanCan = () => {
    console.log('Pause CanCan !');
};
</script>

<template>
    <main aria-label="checkup section" class="mt-16 w-80vw mx-auto">
        <div>
            <h1>Can Can</h1>
            <img alt="le manche" src="@/assets/lemanche.svg" />
            <ListenPlayPause :onListen="handleListenCanCan" :onPlay="handlePlayCanCan" :onPause="handlePauseCanCan" />
        </div>
    </main>
</template>
