<template>
    <div class="zar-form">
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
                  :user_condition="user_condition"
                  :onError="onError"></dataform>
    </div>
</template>
<script>


export default {
    props:["zar_id", "editID"],
    name: "readOnlyView.vue",
    components: {

    },
    methods:{
        onSuccess(){
            // console.log("za ashgvi !!! ");
          window.location.reload(true);
            //window.location.href = "/admin#/p/c49b40d5-b4e2-9525-bb0a-1400edb63c43/90382fff-174e-0cfc-17b6-f80060406f3c";
        },
        readyForm(){
            this.$refs.form.setModel("form_id", this.zar_id, "Number")
            if(this.editID >= 1){

                this.editMode = true;
                this.$refs.form.editModel(this.editID);
            } else {
                this.editMode = false;

            }

        },
        onError(){

        },

    },

    data() {
        return {
            schema_id:1059,
            user_condition:[{"form_field":"user_id","user_field":"id"},{"form_field":"firstname","user_field":"id"},{"form_field":"register","user_field":"id"},{"form_field":"email","user_field":"id"},{"form_field":"phone","user_field":"id"},{"form_field":"project_owner","user_field":"id"},{"form_field":"b_lastname","user_field":"id"},{"form_field":"b_first_name","user_field":"id"},{"form_field":"lastnam","user_field":"id"},{"form_field":"user_role_id","user_field":"role"},{"form_field":"b_register","user_field":"id"},{"form_field":"b_email","user_field":"id"},{"form_field":"b_utas","user_field":"id"},{"form_field":"company_type_id","user_field":"id"},{"form_field":"pro_guitsetgegch_id","user_field":"id"}],
            editMode:false,
            formId: null,
        };
    },

};
</script>
<style lang="scss">

</style>
