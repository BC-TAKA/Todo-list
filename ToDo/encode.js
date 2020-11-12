//登録ボタンクリックをトリガーにnameとtodoの値を取得
function clickBtn() {
    const nameValue = document.getElementById("valueName").value;
    const todoValue = document.getElementById("valueTodo").value;
    
    //オブジェクト化
    var obj = {
        "name": nameValue,
        "todo": todoValue
    };

    var changeJson = JSON.stringify(obj);
    console.log(changeJson);

        //fetchでJSON形式に換えてAPIに送信する
        fetch(`http://localhost:8081/todos`, {
            method: "POST",
            headers: {
                'Content-Type': 'application/json',
            },
            body: changeJson,
        }).then(function(response1) {
            console.log("status=" + response1.status);
         })
        .then(function(data1) {
            console.log(body);
        })
        .catch(function(err1) {
            console.log("err=" + err1);
        });
}