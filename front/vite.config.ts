import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import * as path from 'path';
import { fileURLToPath, URL } from 'node:url';

// https://vitejs.dev/config/
export default defineConfig({
    plugins: [vue()],
    server: {
        proxy: {
            '^/$': 'http://127.0.0.1:3000',
            '^/manifest.webmanifest$': 'http://127.0.0.1:3000',
            '^/sw/serviceWorker.js$': 'http://127.0.0.1:3000',
            '^/about$': 'http://127.0.0.1:3000',
            '^/checkup$': 'http://127.0.0.1:3000',
            '^/collection$': 'http://127.0.0.1:3000',
            '^/profile$': 'http://127.0.0.1:3000',
            '^/stats$': 'http://127.0.0.1:3000',
            '^/play': 'http://127.0.0.1:3000',
            '^/graphql': {
                ws: true,
                target: 'http://127.0.0.1:3000',
                changeOrigin: true,
            },
        },
    },
    build: {
        manifest: true,
        outDir: path.resolve(__dirname, '..', 'dist'),
        rollupOptions: {
            input: path.resolve(__dirname, 'main.ts'),
        },
    },
    resolve: {
        alias: {
            '@': fileURLToPath(new URL('.', import.meta.url)),
        },
    },
    test: {
        coverage: {
            provider: 'istanbul',
            reporter: ['text', 'html', 'json'],
        },
    },
});
