var app = new Vue ({
    el: '#app',
    data: {
        list: [],
        form: {}
    },
    created() {
        this.getTODOs()
    },
    methods: {
        async getTODOs() {
            await axios.get('http://localhost:8081/todos')
            .then((response) => {
                this.list = response.data
            }).catch((error) => {
                console.log(error);
            })
        },
        async doAdd() {
            await axios.post('http://localhost:8081/todos', this.form)
            .then((response) => {
                alert("登録完了しました。")
                this.getTODOs()
                this.form = {}
            }).catch ((error) => {
                console.log(error);
            })
        },
        async doRemove(item) {
            const id = item.ID
            await axios.delete(`http://localhost:8081/todos?id=${id}`)
            .then((response) => {
                this.getTODOs()
            }).catch((error) => {
                console.log(error);
            })
        }
    }
})