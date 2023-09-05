<script setup lang="ts">
import { ref } from 'vue';
import ListenPlayPause from '@/components/Playground/ListenPlayPause.vue';
import StringsFrets from '@/components/Playground/StringsFrets.vue';
import { useSongStore } from '@/store';
import { SongNote } from '@/gql/types';

const isPlaying = ref(false);
const stringsFretsRef = ref(null);

const store = useSongStore();
const title = store.currentSong.title;
const music = store.currentSong.music_path;
const audio = new Audio(music);

// filter -> exclude empty elements from the list
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

console.log('chordData', chordDatas);

type Positions = {
    time: Time[];
};

type Time = {
    [stringNumber: number]: Fret[];
};

type Fret = number;

const positions: Positions = {
    time: [],
};

for (let i = 0; i < chordDatas.length; i++) {
    positions.time.push({
        1: [],
        2: [],
        3: [],
        4: [],
        5: [],
        6: [],
    });
}

chordDatas.forEach((chordData, index) => {
    const { chords } = chordData;
    chords.forEach((chord) => {
        const { string, fret } = chord;
        positions.time[index][string].push(fret);
    });
});

console.log('positions', positions);

// isPosition checks if the fret is currently present on the string at currentIndex.value
// It passes the array of true false to the StringsFrets component which will display the green circles on the guitar strings
const isPosition = (string: number, fret: number): boolean => {
    const currentPosition = positions.time[currentIndex.value];
    if (currentPosition && string in currentPosition) {
        const stings = currentPosition[string];
        const isFret = stings.includes(fret);
        if (isFret) {
            console.log('isFret', isFret, stings, fret);
        }
        return isFret;
    }
    return false;
};

// Function to display the strip above the guitar neck
// If the freight is currently displayed it uses the CSS class anim-fret in the template v-for loop
const isCurrentFret = (fret: number): boolean => {
    const currentPosition = positions.time[currentIndex.value];
    if (currentPosition) {
        // Check for each string if fret is present
        for (let string = 1; string <= 6; string++) {
            if (currentPosition[string] && currentPosition[string].includes(fret)) {
                return true;
            }
        }
    }

    return false;
};

// Function to manage animation
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
