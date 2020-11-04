
var id = document.querySelector("valueId").value;
var name = document.querySelector("#valueName").value;
var todo = document.querySelector("valueTodo").value;

/** JSON化させたいデータ */
var obj = {
    "convertId": id,
    "convertName": name,
    "convertTodo": todo
};

var json = JSON.stringify(obj);