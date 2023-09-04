<script setup lang="ts">
import { ref } from 'vue';
import ListenPlayPause from '@/components/Playground/ListenPlayPause.vue';
import StringsFrets from '@/components/Playground/StringsFrets.vue';
import { useSongStore } from '@/store';
import { SongNote } from '@/gql/types';

const isPlaying = ref(false);
const stringsFretsRef = ref(false);

const store = useSongStore();
const title = store.currentSong.title;
const music = store.currentSong.music_path;
const audio = new Audio(music);

// const positions: Position[] = (store.currentSong.notes ?? []).filter(
//     (note?: SongNote | null) => note != null && note.beat !== undefined
// ) as SongNote[];

// const isPosition = (string: number | undefined, fret: number): boolean => {
//     const currentPosition = positions[currentIndex.value];
//     return string !== undefined && string === currentPosition.string && fret === currentPosition.fret;
// };
// const isCurrentFret = (fret: number) => {
//     return positions[currentIndex.value].fret === fret;
// };

// filter = exclut les éléments vides de la liste
const notes = (store.currentSong.notes ?? []).filter((notes?: SongNote | null) => notes != null) as SongNote[];
console.log('notes', notes);

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

interface Chord {
    note: string | undefined;
    fret: number | undefined;
    string: number | undefined;
    // alter: number | undefined;
    // octave: number | undefined;
    abscissa: number | undefined;
}

interface MergeData {
    measure: number;
    beat: number;
    duration: number;
    chords: Chord[];
}

interface ChordTable {
    measure: number;
    beat: number;
    duration: number;
    start: number;
    end: number;
    chords: Chord[];
}

interface Position {
    frets: (number | undefined)[];
    strings: (number | undefined)[];
}

const currentIndex = ref(0);
const animationRunning = ref(false);
const animationPaused = ref(false);
const currentNoteName = ref([]);

function transformAndMergeData(): MergeData[] {
    const transformedData = notes.map((item: SongNote) => ({
        measure: item.measure,
        beat: item.beat,
        duration: item.duration,
        chords: [
            {
                note: item.note,
                fret: item.fret,
                string: item.string,
                abscissa: item.abscissa,
            },
        ],
    }));

    const mergedData = transformedData.reduce<MergeData[]>((acc, item, index) => {
        if (index > 0 && item.chords[0].abscissa === acc[acc.length - 1].chords[0].abscissa) {
            acc[acc.length - 1].chords.push(item.chords[0]);
            acc[acc.length - 1].duration += item.duration;
        } else {
            acc.push({
                measure: item.measure,
                beat: item.beat,
                duration: item.duration,
                chords: [item.chords[0]],
            });
        }
        return acc;
    }, []);

    return mergedData;
}
const mergedData = transformAndMergeData();

function calculateStartEnd(mergedData: ChordTable[]) {
    let currentTime = 0;

    mergedData.forEach((item) => {
        item.start = currentTime;
        item.end = currentTime + item.duration * 1000;

        currentTime = item.end;
    });
    return mergedData;
}
const chordDatas = calculateStartEnd(mergedData as ChordTable[]);
console.log('chordDatas', chordDatas);

const notenames = chordDatas.map((item) => item.chords.map((chord) => chord.note));
console.log('notenames', notenames);

const updateCurrentNoteName = () => {
    const currentNote = notenames[currentIndex.value];
    console.log('currentNote', currentNote);
    if (currentNote) {
        currentNoteName.value = currentNote;
    } else {
        currentNoteName.value = [];
    }
};

const positions: Position[] = chordDatas.map((item: ChordTable) => {
    const myfrets = item.chords.map((chord) => chord.fret);
    const mystrings = item.chords.map((chord) => chord.string);
    console.log('myfrets', myfrets);
    console.log('mystrings', mystrings);
    return {
        frets: myfrets,
        strings: mystrings,
    };
});

console.log('positions', positions);

const isPosition = (strings: number[], frets: number[]): boolean => {
    const currentPosition = positions[currentIndex.value];
    console.log('strings in current position', strings, currentPosition.strings.includes(strings));
    console.log('fret in current position', frets, currentPosition.frets.includes(frets));
    return strings !== undefined && currentPosition.strings.includes(strings) && currentPosition.frets.includes(frets);
};

const isCurrentFret = (frets: number): boolean => {
    const currentPosition = positions[currentIndex.value];
    return currentPosition && currentPosition.frets.includes(frets);
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
        currentIndex.value = (currentIndex.value + 1) % chordDatas.length;
        console.log('chordDatas.length', chordDatas.length);
        const currentData = chordDatas[currentIndex.value];
        console.log('currentData', currentData);
        const nextData = chordDatas[(currentIndex.value + 1) % chordDatas.length];

        const start = currentData.start;
        console.log('start', start);
        const end = nextData.start;
        console.log('end', end);
        const duration = end - start;

        if (!animationPaused.value) {
            console.log('duration', duration);
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
                    Notes :
                    <span class="font-extrabold">{{ animationRunning ? currentNoteName.join(' / ') : '▶' }}</span>
                </p>
                <p>⚡</p>
                <p class="mx-5">Toi : <span class="font-extrabold">B</span></p>
            </div>

            <div
                v-for="strings in 1"
                :key="strings"
                class="flex flex-row flex-wrap w-full justify-around text-gray-900"
            >
                <div v-for="frets in 14" :key="frets" :class="{ 'anim-fret': isCurrentFret(frets) }">
                    {{ frets }}
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
