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

function extractErrors(data: unknown): string[] {
    if (typeof data === 'object' && data !== null) {
        if ('errors' in data && Array.isArray(data.errors)) {
            return data.errors.map((error: unknown) => String(error));
        } else if ('errors' in data && data.errors !== null && typeof data.errors === 'object') {
            return Object.values(data.errors).map((error: unknown) => String(error));
        } else if ('error' in data) {
            return [String(data.error)];
        }
    }
    return ['Une erreur inconnue est survenue'];
}

export interface UserState {
    token: string;
    username: string;
    email: string;
}

export const useUserStore = defineStore('user', {
    state: () => ({
        user: null as UserState | null,
    }),
    actions: {
        async signUp(username: string, email: string, password: string): Promise<string[]> {
            const response = await fetch('/auth/signup', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    username,
                    email,
                    password,
                }),
            });

            if (!response.ok) {
                return extractErrors(await response.json());
            }

            this.user = (await response.json()) as UserState;

            return [];
        },

        async logIn(email: string, password: string): Promise<string[]> {
            const response = await fetch('/auth/login', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify({
                    email,
                    password,
                }),
            });

            if (!response.ok) {
                return extractErrors(await response.json());
            }

            this.user = (await response.json()) as UserState;

            return [];
        },

        async logOut(): Promise<string[]> {
            const response = await fetch('/auth/logout', {
                method: 'POST',
            });

            if (!response.ok) {
                return extractErrors(await response.json());
            }

            this.user = null;

            return [];
        },
    },
});
