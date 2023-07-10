<script lang="ts" setup>
import PlayMic from '@/components/sound/PlayMic.vue';
</script>

<template>
    <main aria-label="checkup section" class="items-center mt-16 w-80vw flex justify-center">
        <h1>Checkup</h1>
    </main>
    <PlayMic :enable-canvas="true" />
</template>
