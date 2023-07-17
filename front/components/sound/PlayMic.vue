<script setup lang="ts">
import { initMicrophone, getTones } from '@/utilities/sound/microphone';
import { notes as notesRegistered, Note } from '@/utilities/sound/notes';
import { shallowRef, watch, watchEffect } from 'vue';

const props = defineProps<{
    enableCanvas: boolean;
}>();
const notes: Note[] = notesRegistered;
const stream = shallowRef<MediaStream | null>(null);
const sensitivity = shallowRef<number>(50);

watch(stream, () => {
    const setTones = () => {
        if (stream.value === null) return;
        getTones(notes, props.enableCanvas, 35);
        requestAnimationFrame(setTones);
    };
    watchEffect(() => {
        requestAnimationFrame(setTones);
    });
});
watch(sensitivity, () => null); // Required for displaying sensitivity in real time

function onClick() {
    if (stream.value === null) {
        initMicrophone(sensitivity.value).then((s: MediaStream) => (stream.value = s));
    } else {
        stream.value.getAudioTracks().forEach((track: MediaStreamTrack) => track.stop());
        stream.value = null;
    }
}
</script>

<template>
    <div class="greetings">
        <div>
            <input v-model="sensitivity" type="range" min="40" max="100"><span v-text="-sensitivity" /> Db
        </div>
        

        <button id="startBtn" type="button" @click="onClick()">{{ stream === null ? 'Start' : 'Stop' }}</button>

        <div v-if="enableCanvas">
            <span id="vOut"></span><br />
            <span id="freqs"></span><br />
            <br />
            <canvas id="visualizer"></canvas>
        </div>
    </div>
</template>
