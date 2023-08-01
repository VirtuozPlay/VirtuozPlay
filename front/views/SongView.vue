<script setup lang="ts">
import { ref } from 'vue';
import ListenPlayPause from '@/components/Playground/ListenPlayPause.vue';
import StringsFrets from '@/components/Playground/StringsFrets.vue';
import { useSongStore } from '@/store';
import { SongNote } from '@/gql/types';

const isPlaying = ref(false);
const stringsFretsRef = ref(null);

interface Position {
    string: number | undefined;
    fret: number;
    beat: number;
}

const store = useSongStore();
const title = store.currentSong.title;
const music = store.currentSong.music_path;
const audio = new Audio(music);
// const positions: Position[] = (store.currentSong.notes ?? [])
//     .filter((note?: Partial<SongNote> | null) => note != null)
//     .map((note: SongNote | null) => {
//         // null notes are filtered out, we can safely use non-null assertions
//         // const n = note as SongNote;
//         return {
//             string: note?.string,
//             fret: note?.fret,
//             beat: note?.start,
//         };
//     });

const positions: Position[] = (store.currentSong.notes ?? []).filter(
    (note?: SongNote | null) => note != null && note.beat !== undefined
) as SongNote[];
//

// const notename = (store.currentSong.notes ?? [])
//     .filter((note?: SongNote | null) => note != null)
//     .map((note: SongNote | null) => {
//         // const n = note as SongNote;
//         return {
//             note: note?.note,
//         };
//     });

const note = (store.currentSong.notes ?? []).filter((note?: SongNote | null) => note != null) as SongNote[];
const notename = note.map((note: SongNote) => ({ note: note?.note }));

const updateCurrentNoteName = () => {
    const currentNote = notename[currentIndex.value];
    if (currentNote) {
        currentNoteName.value = currentNote.note;
    } else {
        currentNoteName.value = '';
    }
};

const currentIndex = ref(0);
const currentNoteName = ref('');
const animationRunning = ref(false);
const animationPaused = ref(false);

const iconState = ref(['fas', 'headphones']);

const handleListen = () => {
    if (isPlaying.value) {
        audio.pause();
        isPlaying.value = false;
        iconState.value = ['fas', 'headphones'];
    } else {
        audio.play();
        isPlaying.value = true;
        iconState.value = ['fas', 'stop'];
    }
};

const isPosition = (string: number | undefined, fret: number): boolean => {
    const currentPosition = positions[currentIndex.value];
    return string !== undefined && string === currentPosition.string && fret === currentPosition.fret;
};

const isCurrentFret = (fret: number) => {
    return positions[currentIndex.value].fret === fret;
};

const startAnimation = () => {
    console.log('Start animation');
    if (animationRunning.value) {
        console.log('Animation already in progress -> exit');
        return;
    }
    if (animationPaused.value) {
        animationPaused.value = false;
    } else {
        currentIndex.value = 0; // restart from the beginning
    }

    const animate = () => {
        const currentPosition = positions[currentIndex.value];
        const nextPosition = positions[(currentIndex.value + 1) % positions.length];

        const start = currentPosition.beat;
        const end = nextPosition.beat;
        const duration = end - start;

        console.log(start);

        currentIndex.value = (currentIndex.value + 1) % positions.length;

        if (!animationPaused.value) {
            setTimeout(animate, duration);
        } else {
            animationRunning.value = false;
        }
        updateCurrentNoteName();
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
    currentIndex.value = 0;
};
</script>

<template>
    <main aria-label="songview section" class="mt-16 w-80vw mx-auto">
        <div>
            <h1 class="text-3xl text-center font-extrabold">{{ title }}</h1>
            <div class="flex flex-row justify-center text-3xl my-14">
                <p class="mx-5">
                    Notes : <span class="font-extrabold">{{ animationRunning ? currentNoteName : '▶' }}</span>
                </p>
                <p>⚡</p>
                <p class="mx-5">Toi : <span class="font-extrabold">B</span></p>
            </div>

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
                :on-listen="handleListen"
                :on-play="startAnimation"
                :on-pause="pauseAnimation"
                :on-stop="stopAnimation"
                :animation-running="animationRunning"
                :icon-state="iconState"
                :is-playing="isPlaying"
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
