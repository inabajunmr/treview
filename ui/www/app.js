let vm = new Vue({
    el: '#app',
    components: { VoerroTagsInput },
    data: {
        repos: [],
        langs: [],
        condition: {
            Lang: "Config",
            Span: "today",
            OnlyNew: true
        },
        selectedTags: [
            'tags',
            'selected',
            'by',
            'default',
        ],
        // ALTERNATIVELY
        selectedTags: 'tags,selected,by,default',    
        watch: {
            condition: {
                handler: function (val, oldVal) {
                    reload(val)
                },
                deep: true
            }
        }        
    }
})
