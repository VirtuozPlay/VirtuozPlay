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
        },
    },
    build: {
        manifest: true,
        outDir: path.resolve(__dirname, 'dist'),
        rollupOptions: {
            input: path.resolve(__dirname, 'assets/main.ts'),
        },
    },
    resolve: {
        alias: {
            '@': fileURLToPath(new URL('./assets', import.meta.url)),
        },
    },
    test: {
        coverage: {
            provider: 'istanbul',
            reporter: ['text', 'html', 'json'],
        },
    },
});
