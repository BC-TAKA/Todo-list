fetch('http://localhost:8081/todos')
.then(response => {
    console.log(response.status); // => 200
    return response.json().then(userInfo => {
        // JSONパースされたオブジェクトが渡される
        var parent = document.getElementById("parent")
        userInfo.forEach(info => {
            var aTag = document.createElement("a")

            Object.entries(info).forEach(([key, val]) => {
                var span = document.createElement("span")
                span.innerText = `${key}: ${val}`
                aTag.appendChild(span)
            });
            parent.appendChild(aTag)
            parent.appendChild(document.createElement("br"))
        });
        // console.log(userInfo); // => {...} userInfoにデコード済みのものが入っている
        // userInfo.forEach(element => {
        //     console.log(Object.entries(element))
        //     const list = element
        //     const parent = document.getElementById("parent")
        //         const aTag = document.createElement("a")
        //         aTag.Text = list
        //         parent.appendChild(aTag)
    });
});