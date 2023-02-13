let addNewsButton = document.querySelector('.addNewsButton');
let deleteNewsButton = document.querySelector('.deleteNewsButton');

let titleInput = document.querySelector('.titleInput');
let textInput = document.querySelector('.textInput');
let imageInput = document.querySelector('.imageInput');
let cards = document.querySelectorAll('.card');

const Title = /\\n +|\\n|^[\s{2,}]+|\d+-\d+-\d+|[\s{2,}]+$| {2,}|\n/g

console.log(cards)
cards.forEach(card => {
  card.isChosen = false;
  card.addEventListener('click', event => {
    if (card.isChosen == true) {
      card.style.color = 'inherit';
      card.isChosen = false;
    } else {
      card.style.color = 'red';
      card.isChosen = true;
    }
  });
});

addNewsButton.onclick = AddNews.bind(addNewsButton, titleInput, textInput, imageInput);
deleteNewsButton.addEventListener('click', deleteNews);

// Отправляем POST запрос по адресу /news с данными из полей title, text и image
// Мы отлавливаем этот сценарий в router
function AddNews(title, text, image) {
  console.log('this = ', this);

  let xhr = new XMLHttpRequest();
  xhr.open('POST', '/news');

  imagePath = image.value;
  imageName = imagePath.replace(/^.*[\\\/]/, '');

  let time = new Date();
  let data = JSON.stringify({
    Title: title.value,
    TextContent: text.value,
    Image: imageName,
  });
  xhr.send(data);
}

// Отправляем DELETE запрос по адресу /news для каждой выбранной
// карточки с новостью по-отдельности
function deleteNews() {
  console.log('Clicked on button');
  cards.forEach(card => {
    console.log(card.isChosen);
    if (card.isChosen == true) {
      text = card.textContent;
      let fixedstr = text.replace(Title, '');
      let xhr = new XMLHttpRequest();
      xhr.open("DELETE", "/news");
      let data = JSON.stringify({
        Title: fixedstr,
      });
      xhr.send(data);
    }
  });
}