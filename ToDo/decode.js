//登録ボタンクリックをトリガーにnameとtodoの値を取得
function clickBtn() {
    var nameValue = document.getElementById("valueName").value;
    var todoValue = document.getElementById("valueTodo").value;
    
    //オブジェクト化
    var obj = {
        "name": nameValue,
        "TODO": todoValue
    }

    //objをjson形式に変換
    var json = JSON.stringify(obj);
    console.log(json)

        //fetchでAPIにJSONを送信する
        //API側でのエンドポイントに直す
        fetch("http://localhost:8081/regist/regist.html", {
            method:"POST",
            body:form1
        })
        .then(function(response1) {
            console.log("status=" + response1.status);
            return response1.json();
        })
        .then(function(data1) {
            console.log(JSON.stringify(data1));
        })
        .catch(function(err1) {
            console.log("err=" + err1);
        });
}