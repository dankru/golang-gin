
let loginButton = document.querySelector(".loginBtn");
let logoutButton = document.querySelector(".logoutBtn");
let loginInput = document.querySelector(".loginInput");
let loginPass = document.querySelector(".loginPass");


loginButton.onclick = Login.bind(loginButton, loginInput, loginPass);
logoutButton.onclick = Logout.bind(logoutButton);

/**
 * @param {HTMLInputElement} login
 * @param {HTMLInputElement} password
 */
console.log("login.js");

// Отправляем POST запрос по адресу /login с данными из полей login и password 
// Мы отлавливаем этот сценарий в router  
function Login (login, password) {
  if (login.value.length < 4) {
    login.classList.add("incorrect");
  }
  console.log("this = ", this);

  let xhr = new XMLHttpRequest();
  xhr.open("POST", "/login");

  let data = JSON.stringify({
    login: login.value,
    password: password.value
  })
  xhr.send(data);
}

function Logout() {
  let xhr = new XMLHttpRequest();
  xhr.open("POST", "/logout");
  xhr.send();  
}