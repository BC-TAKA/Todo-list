//登録ボタンクリックをトリガーにnameとtodoの値を取得
function createTodo() {
    const nameValue = document.getElementById("valueName").value;
    const todoValue = document.getElementById("valueTodo").value;
    
    //オブジェクト化
    const obj = {
        "name": nameValue,
        "todo": todoValue
    };

    //fetchでJSON形式に換えてAPIに送信する
    fetch(`http://localhost:8081/todos`, {
        method: "POST",
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify(obj),
    }).then((response) => {
        if (response.ok) {
            alert("登録が完了しました。");
        } else {
            alert(err);
        }
    })
    .catch((err) => {
        console.error("err=" + err);
    });
}