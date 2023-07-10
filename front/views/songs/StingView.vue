<script setup lang="ts">
import ListenPlayPause from '@/components/Playground/ListenPlayPause.vue';

const handleListenSting = () => {
    console.log("J'écoute la musique Sting");
};

const handlePlaySting = () => {
    console.log('Je lance la musique Sting');
};

const handlePauseSting = () => {
    console.log('Pause Sting !');
};
</script>

<template>
    <main aria-label="checkup section" class="items-center mt-16 w-80vw flex justify-center">
        <h1>Hello</h1>
        <img alt="VirtuozPlay Logo" src="@/assets/lemanche.svg" width="128" height="128" />
        <ListenPlayPause :onListen="handleListenSting" :onPlay="handlePlaySting" :onPause="handlePauseSting" />
    </main>
</template>
