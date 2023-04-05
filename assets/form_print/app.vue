<template>
    <div class="form-preview-page">
        <div class="title" >
            <img alt="" src="/assets/lambda/images/soylYam.jpeg">
            <div class="title-text">СОЁЛ, УРЛАГИЙГ ДЭМЖИХ САНГААС САНХҮҮЖИХ ТӨСӨЛ,<br> АРГА ХЭМЖЭЭНИЙ САНАЛ ГАРГАГЧИЙН АНКЕТ</div>
            <button style="font-size: 32px; border: none; background-color: #FFFFFF; position: absolute; right: 0; margin-right: 25px" class="print-btn" @click="printForm">
                <Icon style="color: #001E35;" type="ios-print" />
            </button>
        </div>
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

            printForm(){
                var elements = document.getElementsByTagName('textarea');
                for (var i = 0; i < elements.length; i++) {
                    elements[i].style.height = 'auto';
                    elements[i].style.maxHeight =  elements[i].scrollHeight+'px';
                    elements[i].style.height =  elements[i].scrollHeight+'px';
                }

                print();
            }
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
     body, html{
        background-color: #ffffff;
     }
    .form-preview-page{
        width:297mm;
        margin: 0 auto;
        .dataform-body {
            padding: 25px 15px 15px 15px;
        }
        .dataform-header, .dataform-footer, .sub-form-add-btn, .subform-header, .clear-container, .leaflet-draw{
            display: none;
        }
        .dataform {
            padding-bottom: 60px!important;
        }
        .subform-grid{
            .action{
                display: none;
            }
        }
        .ivu-input-inner-container, .multiselect{
            pointer-events: none;
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
        .title{
            background-color: #FFFFFF;
            margin: 20px 0;
            padding: 0;
            position: relative;
            text-align: center!important;
            align-items: center !important;
            display: inline-grid;
            width: 100%;
            img{
                position: absolute;
                left: 0;
                margin-left: 15px;
                width: 10%;
            }
            .title-text{
                color: #001E35;
                font-weight: 900;
                font-size: 12px !important;
            }
        }
        textarea{
            display: block;
        }
        .subform-grid table thead th {
            font-size: 9px;
            font-weight: 500;
            text-transform: uppercase;
            padding: 4px !important;
            border: 1px solid #CCC;
            color: #565665;
        }
        .subform-grid table tbody tr td .ivu-form-item .ivu-form-item-content {
            height: auto !important;
        }
        .sub-modal-form table tbody tr td .ivu-form-item .ivu-input {
            font-size: 10px !important;
        }
    }
     table {
         page-break-inside:auto
     }

     td  {
         page-break-inside:avoid;
         page-break-after:auto
     }
     //table {page-break-inside:avoid !important; page-break-after:auto !important; display: block;}
     //th { page-break-inside:avoid !important; page-break-after:auto !important; display: block;}
     //tr  { page-break-inside:avoid !important; page-break-after:auto !important }
     //td  { page-break-inside:avoid !important; page-break-after:auto !important; display: block;}
     ////thead { display:table-header-group !important}
     ////tfoot { display:table-footer-group !important}
     //
     table{
         width: 100%!important;
         //table-layout: auto !important;
     }
    @media print {
        .form-preview-page {
            height: auto !important;
            margin-left: 0px;
            margin-right: 0px;
            width: 100% !important;
        }
        .form-preview-page .title{
            visibility: visible;
        }
        .print-btn{
            visibility: hidden;
        }
        .title{
            background-color: #FFFFFF;
            margin:  0!important;
            img{
                margin-left: 15mm!important;
            }
        }
        table{
            width: 100%!important;
            table-layout: auto !important;
        }
        .subform-grid {
            display: flex;
            flex-direction: column;
            margin-bottom: 20px !important;
            min-height: 30px;
            background: transparent;
            width: 100%!important;
        }
        @page {
            margin: 10mm 0 20mm 0;
            width:297mm;
            height: 210mm;
            size: A4 landscape !important;
            padding: 0 0 0 0;
            page-break-after: always;
        }
        body {
            margin: 0;
            color: #000;
            background-color: #fff;
            width: 100%;
            height: 100%;
            padding: 0;
        }
        .side-bar, header{
            display: none !important;
        }
        body .paper-theme main .page, .geo-result, html, body{
            height: auto !important;

        }
        .paper-theme, .page-report, .geo-result, body .paper-theme{
            margin: 0 !important;
            padding: 0 !important;
        }
        .geo-alba-logo{
            visibility: visible;
            padding: 0 10mm;
            img, h3{
                visibility: visible;
            }
        }
        .dataform, .dataform *, .top-nav {
            visibility: visible;
        }
        .dataform{
            width: 297mm !important;
            height: auto !important;
            min-height: 210mm !important;
            display: block !important;
            padding: 5mm 10mm 5mm 10mm;
            margin: 0;
        }
        .subform-grid table tbody tr td .ivu-form-item .ivu-form-item-content{
            height: 100%!important;
        }
        .subform-grid table thead th {
            padding: 4px !important;
        }
        .sub-modal-form table tbody tr td .ivu-form-item .ivu-input {
            font-size: 11px !important;
            white-space: pre !important
        }
        //table {page-break-after:always}

    }
</style>
