<template>
    <div class="form-preview-page">
        <dataform ref="form" :schemaID="schema_id"
                  :editMode="editMode"
                  :onSuccess="onSuccess"
                  :do_render="true"
                  :permissions="{
                        c: true,
                        r: false,
                        u: false,
                        d: false,
                    }"

                  :onReady="readyForm"
                  :user_condition="null"
                  :onError="onError"></dataform>
    </div>
</template>
<script>


export default {
    props:["id"],
    name: "readOnlyView.vue",
    components: {

    },
    methods:{
        onSuccess(){

        },
        readyForm(){
            this.editMode = true;
            this.$refs.form.editModel(this.id);
        },
        onError(){

        },


    },

    data() {
        return {
            schema_id: 1050,

            editMode:false,
            formId: null,
        };
    },

};
</script>
<style lang="scss">

</style>
