import './style.css'

import { setupCounter } from './counter.ts'
import { getUsers } from './users.ts'
import type { User } from './types/user.ts'

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

// Log when PWA is installable
window.addEventListener('beforeinstallprompt', (e) => {
  // Prevent Chrome 67 and earlier from automatically showing the prompt
  e.preventDefault()
  console.log('PWA is installable!')
})

document.querySelector<HTMLDivElement>('#app')!.innerHTML = `
  <div>
    <h1>Artistic</h1>
    <div class="card">
      <button id="counter" type="button"></button>
    </div>
    <div id="users">
      <ul><!-- users will be displayed here --></ul>
    </div>
  </div>
`

setupCounter(document.querySelector<HTMLButtonElement>('#counter')!)

// Fetch and display users
getUsers().then(result => {
  const usersList = document.querySelector('#users ul')!
  usersList.innerHTML = result.users.map((user: User) => {
    return `<li class="user-item">${user.first_name} ${user.last_name}</li>`
  }).join('')
})