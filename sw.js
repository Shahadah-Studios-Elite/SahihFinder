const CACHE_NAME = 'sahihfinder-v1';
const ASSETS = [
  'index.html',
  'Wiki.html',
  '404.html',
  'manifest.json',
  'favicon1.png'
];

self.addEventListener('install', (event) => {
  event.waitUntil(
    caches.open(CACHE_NAME).then((cache) => {
      return cache.addAll(ASSETS);
    })
  );
});

self.addEventListener('fetch', (event) => {
  event.respondWith(
    caches.match(event.request).then((cachedResponse) => {
      return cachedResponse || fetch(event.request);
    })
  );
});
