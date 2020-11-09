fetch('http://localhost:8081/todos')
.then(response => {
    console.log(response.status); // => 200
    return response.json().then(userInfo => {
        // JSONパースされたオブジェクトが渡される
        userInfo.forEach(element => {
            Object.entries(element).forEach(([key, val]) => {
                console.log(key, val)
            });
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