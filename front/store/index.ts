import { defineStore } from 'pinia';
import { Song } from '@/gql/types';

// Create a new store instance.
export const useSongStore = defineStore('songs', {
    state: () => ({ currentSong: {} as Song }),
    getters: {
        getCurrentSong(state): Song {
            return state.currentSong;
        },
    },
    actions: {
        setCurrentSong(song: Song) {
            this.currentSong = song;
        },
    },
});
