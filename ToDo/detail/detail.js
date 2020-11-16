//選択されたaタグのURLを取得
var url = document.location.href;

//URLと末尾のIDを分割して配列に格納・・・IDはresult[1]に入る
var result = url.split('?');

//変数idにIDを格納
id = result[1];
console.log(id);

fetch(`http://localhost:8081/search?id=${id}`, {
    method: 'GET',
}).then((response) => {
    if (response.ok) {
        console.log(id);
    } else {
        console.log("エラーです。");
    }
}).catch((err) => {
    console.log(err);
});