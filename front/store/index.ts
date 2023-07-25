import { defineStore } from 'pinia';
import { Song } from '@/gql/types';

// Create a new store instance.
export const useSongStore = defineStore('songs', {
    state: () => ({ currentSong: {} as Partial<Song> }),
    getters: {
        getCurrentSong(state): Partial<Song> {
            return state.currentSong;
        },
    },
    actions: {
        setCurrentSong(song: Partial<Song>) {
            this.currentSong = song;
        },
    },
});
