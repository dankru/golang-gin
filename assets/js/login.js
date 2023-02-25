'use strict';
let loginButton = document.querySelector('.loginBtn');
let registerButton = document.querySelector('.registerBtn');
let loginInput = document.querySelector('.loginInput');
let passwordInput = document.querySelector('.passwordInput');

loginButton.onclick = Login.bind(loginButton, loginInput, passwordInput);
registerButton.onclick = Register.bind(registerButton, loginInput, passwordInput);

/**
 * @param {HTMLInputElement} login
 * @param {HTMLInputElement} password
 */
console.log('login.js');

// Отправляем POST запрос по адресу /login с данными из полей login и password
// Мы отлавливаем этот сценарий в router

async function Login(login, password) {
  let responce = await fetch('/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json;charset=utf-8',
    },
    body: JSON.stringify({
      login: login.value,
      password: password.value,
    }),
  });
  // }).then(fetch('account'));
}

async function Register(login, password) {
  let responce = await fetch('/account', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json;charset=utf-8',
    },
    body: JSON.stringify({
      login: login.value,
      password: password.value,
      Admin: false,
    }),
  });
}

// if (login.value.length < 4) {
//   login.classList.add('incorrect');
// }
// console.log('this = ', this);

// let xhr = new XMLHttpRequest();
// xhr.open('POST', '/login');

// let data = JSON.stringify({
//   login: login.value,
//   password: password.value,
// });
// xhr.send(data);
// }
// function Logout() {
//   let xhr = new XMLHttpRequest();
//   xhr.open("POST", "/logout");
//   xhr.send();
// }
