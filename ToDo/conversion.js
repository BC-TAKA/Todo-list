fetch('http://localhost:8081/todos')
.then(response => {
    console.log(response.status); // => 200
    return response.json().then(userInfo => {
        // JSONパースされたオブジェクトが渡される
        const parent = document.getElementById("parent");

        //一覧をpタグに格納し、削除ボタンとともに表示
        userInfo.forEach(info => {
            const pTag = document.createElement("p");

            Object.entries(info).forEach(([key, val]) => {
                const span = document.createElement("span");
                span.innerText = `${key}: ${val}  `;
                pTag.appendChild(span);
                btn.textContent = "削除";
                btn.value = "info.ID";
            });
            parent.appendChild(pTag);
            parent.appendChild(btn);
            parent.appendChild(document.createElement("br"));
            parent.appendChild(document.createElement("br"));
            console.log(info.ID,info.Name,info.Todo);
            btn.onclick = function() {
                const id = info.ID;
                console.log(id);
                fetch(`http://localhost:8081/todos?id=${id}`, {
                    method: 'DELETE',
                }).then((response) => {
                    if (response.ok) {
                        console.log("削除しました。");
                        console.log(id);
                    } else {
                        console.log("エラーです。");
                    }
                }).catch((err) => {
                    console.log(err);
                });
            }
        });
    });
});
