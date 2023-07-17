const sw = self as unknown as ServiceWorkerGlobalScope;

console.log('Service worker loaded!');

// Create/install cache
sw.addEventListener('install', (event: ExtendableEvent) => {
    // The promise that skipWaiting() returns can be safely ignored.
    void sw.skipWaiting();
    event.waitUntil(
        caches
            .open('VirtuozPlay')
            .then((cache) => cache.addAll(['/', '/manifest.webmanifest', '/serviceWorker.js']))
            .catch((err) => console.error(err))
    );
});

// Claim control instantly.
sw.addEventListener('activate', (_evt: ExtendableEvent) => sw.clients.claim());

// Load from cache first, then fallback to network if not found.
sw.addEventListener('fetch', (event: FetchEvent) => {
    caches.match(event.request).then((res) => res || fetch(event.request));
});
