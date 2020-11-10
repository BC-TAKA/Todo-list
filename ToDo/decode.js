function clickBtn() {
    const form1 = new FormData(document.getElementById("form1"));
        console.log(form1)

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