var app = new Vue ({
    el: '#app',
    data: {
        todos: [],
        createTodoForm: {},
        updateTodoForm: {},
    },
    created() {
        this.getTODOs()
    },
    methods: {
        //一覧表示機能
        getTODOs() {
            axios.get('http://localhost:8081/todos')
            .then((response) => {
                this.todos = response.data
            }).catch((error) => {
                console.log(error);
            })
        },
        //新規登録機能
        doAdd() {
            axios.post('http://localhost:8081/todos', this.createTodoForm)
            .then((response) => {
                alert("登録完了しました。")
                this.getTODOs()
                this.createTodoForm = {}
            }).catch((error) => {
                console.log(error);
            })
        },
        //更新機能
        doUpdate() {
            axios.put('http://localhost:8081/todos', this.updateTodoForm)
            .then((response) => {
                alert("更新が完了しました。")
                this.getTODOs()
                this.updateTodoForm = {}
            }).catch((error) => {
                console.log(error);
            })
        },
        //削除機能
        doRemove(todo) {
            const id = todo.ID
            axios.delete(`http://localhost:8081/todos?id=${id}`)
            .then((response) => {
                this.getTODOs()
            }).catch((error) => {
                console.log(error);
            })
        }
    }
})