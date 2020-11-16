function clickbtn() {
    const idValue = document.getElementById("idValue").value;
    const nameValue = document.getElementById("nameValue").value;
    const todoValue = document.getElementById("todoValue").value;

    //入力されたIDが整数なら入力内容をオブジェクト化
    if (isFinite(idValue)) {
        var obj = {
            "id": idValue,
            "name": nameValue,
            "todo": todoValue
        };
    } else {
        console.log("IDが正しくありません");
    }
    // console.log(obj);
    var changeJson = JSON.stringify(obj);
    console.log(changeJson);

    fetch('http://localhost:8081/todos', {
        method: "PUT",
        body: changeJson,
    }).then(function(response1) {
        console.log(response1.status);
    })
    .then(function(data1) {
        console.log(body);
    })
    .catch(function(err1) {
        console.log("err =" + err1);
    });
}



// //選択されたaタグのURLを取得
// var url = document.location.href;

// //URLと末尾のIDを分割して配列に格納・・・IDはresult[1]に入る
// var result = url.split('?');

// //変数idにIDを格納
// id = result[1];
// console.log(id);

// fetch(`http://localhost:8081/search?id=${id}`, {
//     method: 'GET',
// }).then((response) => {
//     if (response.ok) {
//         console.log(id);
//     } else {
//         console.log("エラーです。");
//     }
// }).catch((err) => {
//     console.log(err);
// });