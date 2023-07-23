import { defineStore } from 'pinia';

// Create a new store instance.
export const useSongStore = defineStore('songs', {
    state: () => ({ currentSong: {} as any }),
    getters: {
        getCurrentSong: (state) => state.currentSong,
    },
    actions: {
        setCurrentSong(song: any) {
            this.currentSong = song;
        },
    },
});
