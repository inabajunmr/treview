let vm = new Vue({
    el: '#app',
    components: { VoerroTagsInput },
    data: {
        repos: [],
        langs: [],
        condition: {
            Langs: [],
            CloneLangs: [],
            Span: "today",
            OnlyNew: true
        },
    },
    watch: {
        condition: {
            handler: function (val, oldVal) {
                reload(val)
                if(val.Langs.slice().sort().toString() != val.CloneLangs.slice().sort().toString()) {
                    updateConfig(val.Langs)
                    val.CloneLangs = val.Langs.slice()
                }
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
