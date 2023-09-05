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
    songId: string;
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
const stream = shallowRef<MediaStream | null>(null);
const sensitivity = ref(localStorage.getItem('mic_sensivity') ? Number(localStorage.getItem('mic_sensivity')) : 50);
const fps = 30,
    acquisitionDelay = ref(500),
    perfID = ref('');
let isRecordingNoise = true,
    isLocalStorageAccessible = false;
let startTimeStamp = 0;

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

const onClick = async () => {
    if (stream.value === null) {
        stream.value = await initMicrophone(props.songId, sensitivity.value, apolloClient, perfID);
        // if user denied access to microphone
        if (stream.value === null) return;

        if (isLocalStorageAccessible) localStorage.setItem('mic_sensivity', String(sensitivity.value));

        // start by recording noise for 5 s
        isRecordingNoise = true;
        setTimeout(() => {
            isRecordingNoise = false;
            // recording starts
            startTimeStamp = Date.now();
            console.log(notes);
            getDuration(stream, apolloClient, perfID, startTimeStamp, acquisitionDelay);
        }, 5000);
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
    isLocalStorageAccessible = true;
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
        <div>
            Réglage du délai entre chaque acquisition :
            <input v-model="acquisitionDelay" type="range" min="10" max="10000" /><span v-text="acquisitionDelay" />
            ms
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
