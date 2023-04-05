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
                schema_id:window.init.schema_id,
                id:window.init.id,
                editMode:false,
                formId: null,
            };
        },

    };
</script>
<style lang="scss">


    .form-preview-page{

        .dataform-body {
            padding: 15px
        }
        .dataform-header, .dataform-footer, .sub-form-add-btn, .subform-header, .clear-container, .leaflet-draw{
            display: none;
        }
        .subform-grid{
            .action{
                display: none;
            }
        }
        .ivu-input-inner-container, .multiselect{
            pointer-events: none;
        }
        textarea{
            max-height: 52px !important;
        }
        .ivu-form-item-label{
            background: #ccc;
            width: 100%;
            padding: 7px 15px;
            color: #333;
            text-align: left;
            font-weight: 500;
        }
        .ivu-input, .multiselect__tags{
            border-radius: 0;
        }
        .with-info-caller {
            width: calc(100% - 30px);
            display: inline-block;
        }
        .info-caller{
            width: 30px;
            display: inline-block;
            vertical-align: top;
            padding-top: 2px;
            padding-left: 2px;
            /*border-top: solid 1px #ccc;*/
            height: 20px;
            position: absolute;

        }

    }
</style>
