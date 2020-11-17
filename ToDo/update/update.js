function clickbtn() {
    //全てstringとして取り出している
    const idValue = document.getElementById("idValue").value;
    const nameValue = document.getElementById("nameValue").value;
    const todoValue = document.getElementById("todoValue").value;

    //入力されたIDが0より大きい整数なら入力内容をオブジェクト化
    if (isFinite(idValue) && 0 < idValue) {
        var obj = {
            "id": idValue,
            "name": nameValue,
            "todo": todoValue
        };
    } else {
        console.log("IDが正しくありません");
        return
    }

    fetch('http://localhost:8081/todos', {
        method: "PUT",
        body: JSON.stringify(obj),
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