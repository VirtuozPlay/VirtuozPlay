<script setup lang="ts">
import { initMicrophone, getTones, getNoise } from '@/utilities/sound/microphone';
import { notes as notesRegistered, Note, notesPhaseOpposition } from '@/utilities/sound/notes';
import { shallowRef, watch, watchEffect } from 'vue';
import { getCookie, setCookie } from 'typescript-cookie';

const props = defineProps<{
    enableCanvas: boolean;
}>();
const notes: Note[] = notesRegistered;
const stream = shallowRef<MediaStream | null>(null);
const sensitivity = shallowRef<number>(getCookie('mic_sensivity') ? Number(getCookie('mic_sensivity')) : 50);
let isRecordingNoise = true;
let startTimeStamp = 0;

watch(stream, () => {
    const setTones = () => {
        if (stream.value === null) return;

        if (isRecordingNoise) {
            getNoise();
        } else {
            getTones(notes, startTimeStamp, props.enableCanvas, 35);
        }

        requestAnimationFrame(setTones);
    };
    watchEffect(() => {
        requestAnimationFrame(setTones);
    });
});
watch(sensitivity, () => null); // Required to display sensitivity in real time

function onClick() {
    if (stream.value === null) {
        // start by recording noise for 5 s
        isRecordingNoise = true;
        setTimeout(() => {
            isRecordingNoise = false;
            // recording starts
            startTimeStamp = Date.now();
            console.log(notesPhaseOpposition);
        }, 5000);
        setCookie('mic_sensivity', sensitivity.value);
        initMicrophone(sensitivity.value).then((s: MediaStream) => (stream.value = s));
    } else {
        // stop
        stream.value.getAudioTracks().forEach((track: MediaStreamTrack) => track.stop());
        stream.value = null;
    }
}
</script>

<template>
    <div class="greetings">
        <div>
            Réglage de la sensibilité : <input v-model="sensitivity" type="range" min="40" max="100" /><span
                v-text="-sensitivity"
            />
            Db
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
