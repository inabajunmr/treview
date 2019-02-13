let vm = new Vue({
    el: '#app',
    components: { VoerroTagsInput },
    data: {
        repos: [],
        langs: [],
        condition: {
            Langs: [],
            Span: "today",
            OnlyNew: true
        },
    },
    watch: {
        condition: {
            handler: function (val, oldVal) {
                reload(val)
            },
            deep: true
        }
    },
    computed: {
        langObjects: function () {
            var obj = {};
            for (val of this.langs) {
                obj[val] = val;
            }
            return obj
        }
    }

})
