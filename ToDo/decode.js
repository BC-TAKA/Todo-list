//登録ボタンのクリックをトリガーに、nameとtodoを取得
document.getElementById("todoRegist").onclick = function() {
    var registName = document.getElementsByName(registName).value;
    var registTodo = document.getElementsByName(registTodo).value;

    //JSON化してAPIに渡したいデータ
    var obj = {
        "name": registName,
        "todo": registTodo
    };

    //データをJSON文字列に変換
    var json = JSON.stringify(obj);
    CheckJson(json)
}