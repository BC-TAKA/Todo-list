var app = new Vue({
    el: '#app',
    data: {
        form: {},
    },
    methods: {
        async doUpdate() {
            await axios.put('http://localhost:8081/todos', this.form)
            .then((response) => {
                alert("更新完了しました。")
                this.form = {}
            }).catch((error) => {
                console.log(error);
            })
        }
    }
})
