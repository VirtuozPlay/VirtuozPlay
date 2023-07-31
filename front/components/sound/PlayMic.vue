<script setup lang="ts">
import { initMicrophone, getTones, getNoise, getDuration } from '@/utilities/sound/microphone';
import { notes } from '@/utilities/sound/notes';
import { ref, shallowRef, watch, watchEffect } from 'vue';
import { GraphQLWsLink } from '@apollo/client/link/subscriptions';
import { createClient } from 'graphql-ws';
import { ApolloClient } from '@apollo/client/core/ApolloClient';
import { InMemoryCache } from '@apollo/client/cache/inmemory/inMemoryCache';

const props = defineProps<{
    enableCanvas: boolean;
}>();
// Websocket client
const apolloClient = new ApolloClient({
    cache: new InMemoryCache(),
    link: new GraphQLWsLink(
        createClient({
            url: 'ws://127.0.0.1:5173/graphql',
        })
    ),
});
const perfID = ref('');
const stream = shallowRef<MediaStream | null>(null);
const sensitivity = shallowRef<number>(
    localStorage.getItem('mic_sensivity') ? Number(localStorage.getItem('mic_sensivity')) : 50
);
let isRecordingNoise = true,
    isLocalStorageAcessible = false;
let startTimeStamp = 0;
const fps = 30;

watch(stream, () => {
    const setTones = () => {
        if (stream.value === null) return;

        if (isRecordingNoise) {
            getNoise();
        } else {
            getTones(startTimeStamp, props.enableCanvas, 35);
        }

        /* This program requires too much ressources.
         * By default, requestAnimationFrame is based on monitor
         * refresh rate. Get sound 60/120 times per second is overkill.
         */
        setTimeout(() => {
            requestAnimationFrame(setTones);
        }, 1000 / fps);
    };
    watchEffect(() => {
        requestAnimationFrame(setTones);
    });
});
watch(sensitivity, () => null); // Required to display sensitivity in real time

const onClick = () => {
    if (stream.value === null) {
        // start by recording noise for 5 s
        isRecordingNoise = true;
        setTimeout(() => {
            isRecordingNoise = false;
            // recording starts
            startTimeStamp = Date.now();
            console.log(notes);
            getDuration(stream, apolloClient, perfID, startTimeStamp);
        }, 5000);

        if (isLocalStorageAcessible) localStorage.setItem('mic_sensivity', String(sensitivity.value));

        initMicrophone(sensitivity.value, apolloClient, perfID).then((s: MediaStream) => (stream.value = s));
    } else {
        // stop
        stream.value.getAudioTracks().forEach((track: MediaStreamTrack) => track.stop());
        stream.value = null;
    }
};

// check if Local Storage is available
try {
    localStorage.setItem('testing local storage', 'test');
    localStorage.removeItem('testing local storage');
    isLocalStorageAcessible = true;
} catch (e) {
    console.log('Local storage is not available');
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
