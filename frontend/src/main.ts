import './style.css'

import { setupCounter } from './counter.ts'
import { getUsers } from './users.ts'
import type { User } from './types/user.ts'

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