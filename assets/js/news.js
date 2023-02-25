let addNewsButton = document.querySelector('.addNewsBtn');
let deleteNewsButton = document.querySelector('.deleteNewsBtn');

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
async function AddNews(title, text, image) {
  console.log('this = ', this);
  
  imagePath = image.value;
  imageName = imagePath.replace(/^.*[\\\/]/, '');

  let time = new Date();
  let responce = await fetch('/news', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json;charset=utf-8',
    },
    body: JSON.stringify({
      Title: title.value,
      TextContent: text.value,
      Image: imageName,
    }),
  });
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
      let responce = fetch('/news', {
        method: 'DELETE',
        headers: {
          'Content-Type': 'application/json;charset=utf-8',
        },
        body: JSON.stringify({
          Title: fixedstr,
        }),
      });
      }
  });
}