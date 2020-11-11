//登録ボタンクリックをトリガーにnameとtodoの値を取得
function clickBtn() {
    var nameValue = document.getElementById("valueName").value;
    var todoValue = document.getElementById("valueTodo").value;
    
    //オブジェクト化
    var obj = {
        "name": nameValue,
        "TODO": todoValue
    }

        //fetchでJSON形式に換えてAPIに送信する
        fetch(`http://localhost:8081/todos`, {
            method: "POST",
            body: JSON.stringify(obj),
        })
        .then(function(response1) {
            console.log("status=" + response1.status);
        })
        .then(function(data1) {
            console.log(JSON.stringify(data1));
        })
        .catch(function(err1) {
            console.log("err=" + err1);
        });
}