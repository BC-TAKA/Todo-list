fetch('http://localhost:8081/todos')
.then(response => {
    console.log(response.status); // => 200
    return response.json().then(userInfo => {
        // JSONパースされたオブジェクトが渡される
        const parent = document.getElementById("parent");
        userInfo.forEach(info => {
            const aTag = document.createElement("a");

            Object.entries(info).forEach(([key, val]) => {
                const span = document.createElement("span");
                span.innerText = `${key}: ${val}  `;
                aTag.appendChild(span);
                //htmlに3つの値をを渡すことで遷移先を分岐させたい
                aTag.setAttribute('href','/detail/detail.html?' +info.ID +info.Name +info.Todo);
            });
            parent.appendChild(aTag);
            parent.appendChild(document.createElement("br"));
            parent.appendChild(document.createElement("br"));
            console.log(info.ID,info.Name,info.Todo);
        });
    });
});