var app = new Vue ({
    el: '#app',
    data: {
        list: []
    },
    created() {
        this.getTODOs()
    },
    methods: {
        async getTODOs() {
            try {
                let response = await axios.get('http://localhost:8081/todos')
                this.list = response.data
            } catch (e) {
                console.error(e)
            }
        }
    }
})