<script setup lang="ts">
import initMicrophone from '@/utilities/sound/microphone';
defineProps<{
    enableCanvas: boolean;
}>();
</script>

<template>
    <div class="greetings">
        <button id="startBtn" type="button" @click="initMicrophone(enableCanvas)">Start</button>

        <div v-if="enableCanvas">
            <span id="vOut"></span><br />
            <span id="freqs"></span><br />
            <br />
            <canvas id="visualizer"></canvas>
        </div>
    </div>
</template>
