let vm = new Vue({
    el: '#app',
    data: {
        repos: [],
        langs: [],
        condition: {
            Lang: "Config",
            Span: "today",
            OnlyNew: true
        }
    },
    watch: {
        condition: {
            handler: function (val, oldVal) {
                reload(val)
            },
            deep: true
        }
    }
})
