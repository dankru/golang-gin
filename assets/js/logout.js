'use strict';

let logoutButton = document.querySelector('.logoutBtn');

logoutButton.onclick = Logout.bind(logoutButton);

async function Logout() {
  let responce = await fetch('/logout', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json;charset=utf-8',
    },
  });
}
