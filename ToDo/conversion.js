   fetch('http://localhost:8081/todos')
    .then(response => {
        console.log(response.status); // => 200
     return response.json().then(userInfo => {
            // JSONパースされたオブジェクトが渡される
            const parent = document.getElementById("parent");

           userInfo.forEach(info => {
                const aTag = document.createElement("a");
                const btn = document.createElement("button");

                Object.entries(info).forEach(([key, val]) => {
                    const span = document.createElement("span");
                    span.innerText = `${key}: ${val}  `;
                    aTag.appendChild(span);
                    //htmlに3つの値をを渡すことで遷移先を分岐させたい
                    aTag.setAttribute('href','/detail/detail.html?' +info.ID);
                    btn.textContent = "削除";
                    btn.value = "info.ID";
                });
                parent.appendChild(aTag);
                parent.appendChild(btn);
                parent.appendChild(document.createElement("br"));
                parent.appendChild(document.createElement("br"));
                console.log(info.ID,info.Name,info.Todo);
                btn.onclick = function() {
                    //console.log(info.ID);
                    const id = info.ID;
                    console.log(id);
                    fetch("http://localhost:8081/todos?id=${id}", {
                        method: 'DELETE',
                    }).then((response) => {
                        if (response.ok) {
                            console.log("削除しました。");
                            console.log(id);
                            //表示内容を更新させる文
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
