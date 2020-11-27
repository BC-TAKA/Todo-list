var app = new Vue ({
    el: '#app',
    data: {
        todos: [],
        create: {},
        update: {},
    },
    created() {
        this.getTODOs()
    },
    methods: {
        //一覧表示機能
        async getTODOs() {
            await axios.get('http://localhost:8081/todos')
            .then((response) => {
                this.todos = response.data
            }).catch((error) => {
                console.log(error);
            })
        },
        //新規登録機能
        async doAdd() {
            await axios.post('http://localhost:8081/todos', this.create)
            .then((response) => {
                alert("登録完了しました。")
                this.getTODOs()
                this.create = {}
            }).catch((error) => {
                console.log(error);
            })
        },
        //更新機能
        async doUpdate() {
            await axios.put('http://localhost:8081/todos', this.update)
            .then((response) => {
                alert("更新が完了しました。")
                this.getTODOs()
                this.update = {}
            }).catch((error) => {
                console.log(error);
            })
        },
        //削除機能
        async doRemove(todo) {
            const id = todo.ID
            await axios.delete(`http://localhost:8081/todos?id=${id}`)
            .then((response) => {
                this.getTODOs()
            }).catch((error) => {
                console.log(error);
            })
        }
    }
})