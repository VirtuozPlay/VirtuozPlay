<script lang="ts" setup>
import PlayMic from '@/components/sound/PlayMic.vue';
import { ref } from 'vue';

const selectedSong = ref('QyPqpmFqWC4uGInmkodgP');
</script>

<template>
    <main aria-label="checkup section" class="items-center mt-16 w-80vw flex justify-center">
        <h1>Checkup</h1>
    </main>
    <div>
        <label>Song ID : </label>
        <input v-model="selectedSong" type="text" class="w-1/4 bg-blue-200 rounded-lg" />
    </div>
    <PlayMic :song-id="selectedSong" :enable-canvas="true" />
</template>
