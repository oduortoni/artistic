import { defineConfig } from 'vite'
import { VitePWA } from 'vite-plugin-pwa'

// Function to check if a port is available
async function findBackendPort() {
  const ports = [8000, 8001, 8002, 8003, 8004, 8005];
  for (const port of ports) {
    try {
      const response = await fetch(`http://localhost:${port}/api/users`);
      if (response.status !== 404) {
        return port;
      }
    } catch (e) {
      continue;
    }
  }
  return 8000; // Default fallback
}

export default defineConfig(async () => {
  const backendPort = await findBackendPort();
  console.log(`Using backend port: ${backendPort}`);

  return {
    server: {
      host: true, // Listen on all addresses
      port: 5173,
      strictPort: true, // Don't try other ports if 5173 is taken
      proxy: {
        '/api': {
          target: `http://localhost:${backendPort}`,
          changeOrigin: true
        }
      },
      hmr: {
        clientPort: 5173,
        host: 'localhost'
      }
    },
    plugins: [
      VitePWA({
        registerType: 'autoUpdate',
        includeAssets: ['icons/*.png'],
        manifest: {
          name: 'Artistic',
          short_name: 'Artistic',
          description: 'A platform for artists to showcase their creative works',
          theme_color: '#000000',
          background_color: '#ffffff',
          display: 'standalone',
          orientation: 'any',
          start_url: '/',
          id: '/',
          scope: '/',
          icons: [
            {
              src: 'icons/icon-192x192.png',
              sizes: '192x192',
              type: 'image/png',
              purpose: 'any maskable'
            },
            {
              src: 'icons/icon-512x512.png',
              sizes: '512x512',
              type: 'image/png',
              purpose: 'any maskable'
            }
          ]
        },
        workbox: {
          globPatterns: ['**/*.{js,css,html,ico,png,svg}'],
          runtimeCaching: [
            {
              urlPattern: /^https:\/\/api\/.*/i,
              handler: 'NetworkFirst',
              options: {
                cacheName: 'api-cache',
                expiration: {
                  maxEntries: 10,
                  maxAgeSeconds: 60 * 60 * 24 // 24 hours
                },
                networkTimeoutSeconds: 10
              }
            }
          ]
        },
        devOptions: {
          enabled: true,
          type: 'module'
        }
      })
    ],
    build: {
      outDir: 'dist',
      emptyOutDir: true
    }
  }
}) 