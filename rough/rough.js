let xhr = new XMLHttpRequest();
xhr.open("GET", "http://127.0.0.1:8000/api/solution")
console.log(xhr.send())