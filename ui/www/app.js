let vm = new Vue({
    el: '#app',
    components: { VoerroTagsInput },
    data: {
        repos: [],
        langs: [],
        condition: {
            Langs: ["Config"],
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
