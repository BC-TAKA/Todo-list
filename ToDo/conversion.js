var app = new Vue ({
    el: '#app',
    data: {
        list: []
    },
    created: function() {
        axios.get('http://localhost:8081/todos')
        .then(function(response) {
            this.list = response.data
        }.bind(this)).catch(function(e){
            console.error(e)
        })
    },
    methods: {
        
    }
})