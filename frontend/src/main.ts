import './style.css'

import { setupCounter } from './counter.ts'
import { getUsers } from './users.ts'
import type { User } from './types/user.ts'

// Store the install prompt event
let deferredPrompt: any;

// Register service worker
if ('serviceWorker' in navigator) {
  navigator.serviceWorker.register('/sw.js')
    .then((registration) => {
      console.log('Service Worker registered with scope:', registration.scope)
    })
    .catch((error) => {
      console.error('Service Worker registration failed:', error)
    })
}

// Handle PWA installation
window.addEventListener('beforeinstallprompt', (e) => {
  // Prevent Chrome 67 and earlier from automatically showing the prompt
  e.preventDefault()
  // Store the event for later use
  deferredPrompt = e
  // Show the install button
  const installButton = document.getElementById('install-button')
  if (installButton) {
    installButton.style.display = 'block'
  }
})

document.querySelector<HTMLDivElement>('#app')!.innerHTML = `
  <div>
    <h1>Artistic</h1>
    <div class="card">
      <button id="counter" type="button"></button>
    </div>
    <div class="install-container">
      <!-- <button id="install-button" style="display: none">Install App</button> -->
      <button id="install-button">Install</button>
    </div>
    <div id="users">
      <ul><!-- users will be displayed here --></ul>
    </div>
  </div>
`

// Add install button handler
const installButton = document.getElementById('install-button')
if (installButton) {
  installButton.addEventListener('click', async () => {
    if (deferredPrompt) {
      // Show the prompt
      deferredPrompt.prompt()
      // Wait for the user to respond to the prompt
      const { outcome } = await deferredPrompt.userChoice
      console.log(`User response to the install prompt: ${outcome}`)
      // Clear the deferredPrompt variable
      deferredPrompt = null
      // Hide the install button
      installButton.style.display = 'none'
    }
  })
}

// Hide install button when PWA is installed
window.addEventListener('appinstalled', () => {
  console.log('PWA was installed')
  const installButton = document.getElementById('install-button')
  if (installButton) {
    installButton.style.display = 'none'
  }
})

setupCounter(document.querySelector<HTMLButtonElement>('#counter')!)

// Fetch and display users
getUsers().then(result => {
  const usersList = document.querySelector('#users ul')!
  usersList.innerHTML = result.users.map((user: User) => {
    return `<li class="user-item">${user.first_name} ${user.last_name}</li>`
  }).join('')
})