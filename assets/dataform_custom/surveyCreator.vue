<template>
    <FormItem :prop=rule :label=label>
        <div id="surveyCreatorContainer"></div>

    </FormItem>
</template>

<script>
// import * as SurveyCreator from "survey-creator";
// import "survey-creator/survey-creator.css";

// import * as SurveyKo from "survey-knockout";
// import * as widgets from "surveyjs-widgets";
// import "bootstrap/dist/css/bootstrap.css";

//
// widgets.select2(SurveyKo);
// widgets.inputmask(SurveyKo);
// widgets.jquerybarrating(SurveyKo);
// widgets.jqueryuidatepicker(SurveyKo);
// widgets.nouislider(SurveyKo);
// widgets.select2tagbox(SurveyKo);
// widgets.sortablejs(SurveyKo);
// widgets.autocomplete(SurveyKo);
// widgets.bootstrapslider(SurveyKo);
var mainColor = "#1152BB";
var mainHoverColor = "#1152BB";
var textColor = "#4a4a4a";
var headerColor = "#1152BB";
var headerBackgroundColor = "#4a4a4a";
var bodyContainerBackgroundColor = "#f8f8f8";




// SurveyKo.Serializer.addProperty("question", "tag:number");

export default {

    props: ["model", "rule", "label", "meta", "do_render"],
    methods: {
        setData(data){
            this.model.form[this.model.component] = data;
        },

        init(){

            if(this.surveyCreator == null){

                let svOptions = {
                    showLogicTab: true,
                    showJSONEditorTab:false,
                    questionTypes: [
                        'text',
                        'checkbox',
                        'comment',
                        // 'imagepicker',
                        'ranking',
                        // 'image',
                        'dropdown',
                        // 'file',
                        // 'html',
                        'matrix',
                        'matrixdropdown',
                        'matrixdynamic',
                        'multipletext',
                        'panel',
                        // 'paneldynamic',
                        'radiogroup',
                        'rating',
                        'boolean',
                        // 'expression',
                        'signaturepad',
                        'flowpanel'
                    ],

                };

                SurveyCreator.localization.locales["mn"] = window.init.translation_mn;


                SurveyCreator.localization.currentLocale = "mn";

                Survey.surveyLocalization.locales["mn"] = window.init.translation_mn_survey;
                Survey.surveyLocalization.defaultLocale = "mn";

                // Survey
                //     .StylesManager
                //     .applyTheme("modern");

//remove a property to the page object. You can't set it in JSON as well
                Survey
                    .Serializer
                    .removeProperty("panelbase", "visibleIf");
//remove a property from the base question class and as result from all questions
                Survey
                    .Serializer
                    .removeProperty("question", "visibleIf");

               // Survey.StylesManager.applyTheme("modern");
                Survey
                    .StylesManager
                    .applyTheme("modern");

                Survey
                    .StylesManager
                    .applyTheme("modern");


             /*   var defaultThemeColorsSurvey = Survey
                    .StylesManager
                    .ThemeColors["default"];
                defaultThemeColorsSurvey["$main-color"] = mainColor;
                defaultThemeColorsSurvey["$main-hover-color"] = mainHoverColor;
                defaultThemeColorsSurvey["$text-color"] = textColor;
                defaultThemeColorsSurvey["$header-color"] = headerColor;
                defaultThemeColorsSurvey["$header-background-color"] = headerBackgroundColor;
                defaultThemeColorsSurvey["$body-container-background-color"] = bodyContainerBackgroundColor;

                var defaultThemeColorsEditor = SurveyCreator
                    .StylesManager
                    .ThemeColors["modern"];
                defaultThemeColorsEditor["$primary-color"] = mainColor;
                defaultThemeColorsEditor["$secondary-color"] = mainColor;
                defaultThemeColorsEditor["$primary-hover-color"] = mainHoverColor;
                defaultThemeColorsEditor["$primary-text-color"] = textColor;
                defaultThemeColorsEditor["$selection-border-color"] = mainColor;

                Survey
                    .StylesManager
                    .applyTheme();
                SurveyCreator
                    .StylesManager
                    .applyTheme();*/

                this.surveyCreator = new SurveyCreator.SurveyCreator(

                    svOptions
                );
                Survey.Serializer.findProperty("survey", "locale").visible = false;
                Survey.Serializer.findProperty("survey", "locale").visible = false;
                // this.surveyCreator.showToolbox = "right";
                //  this.surveyCreator.showPropertyGrid = "right";
                //  this.surveyCreator.rightContainerActiveItem("toolbox");
                this.surveyCreator.saveSurveyFunc = ()=> {
                    this.setData(JSON.stringify(this.surveyCreator.text));
                };

                this.surveyCreator.onModified.add(() => {
                    this.setData(JSON.stringify(this.surveyCreator.text));

                });

                this.surveyCreator
                    .onShowingProperty
                    .add((sender, options)=> {

                        if (options.obj.getType() == "survey") {

                            // options.canShow = options.property.name == "title";
                            if(
                                options.property.name == "title" ||
                                options.property.name == "description" ||
                                options.property.name == "questionsOrder" ||
                                options.property.name == "questionErrorLocation" ||
                                options.property.name == "questionStartIndex" ||
                                options.property.name == "questionTitlePattern" ||
                                options.property.name == "showQuestionNumbers" ||
                                options.property.name == "questionDescriptionLocation" ||
                                options.property.name == "questionTitleLocation" ||
                                options.property.name == "questionsOnPageMode" ||
                                options.property.name == "pages"

                            ){
                                options.canShow = true;
                            } else {
                                options.canShow = false;
                            }
                        } else if (options.obj.getType() == "page") {
                            //console.log(options.property.name);
                            // options.canShow = options.property.name == "title";
                            if(
                                options.property.name == "enableIf" ||
                                options.property.name == "requiredIf"
                            ){
                                options.canShow = false;
                            } else {
                                options.canShow = true;
                            }
                        } else {
                            if(options.property.name == "textUpdateMode"){
                                options.canShow = false;
                            }
                            console.log(options.property.name);


                        }
                    });
                this.surveyCreator.render("surveyCreatorContainer");
                this.surveyCreator.JSON = {};
            }

        }
    },
    data(){
        return {
            surveyCreator:null
        }
    },

    mounted() {
        let root = document.documentElement;
        root.style.setProperty('(--primary', mainColor);
        root.style.setProperty('(--primary', mainColor);
        this.init();

    },
    computed: {
        myValue() {
            return this.model.form[this.model.component];
        },
        surveyTitle() {
            return this.model.form["title"];
        },
        surveyDescription() {
            return this.model.form["description"];
        },
    },
    watch:{
        myValue(value, oldValue) {
            if(oldValue && !value || value && !oldValue){
                this.surveyCreator.text = JSON.parse(value);
            }
        },
        do_render(value){
            if(!value){
                this.surveyCreator.text = undefined;

            }
        },
        surveyTitle(value){
            if(this.surveyCreator && value){
                this.surveyCreator.JSON = {
                    ...this.surveyCreator.JSON, title:value
                }
            }
        },
        surveyDescription(value){
            this.surveyCreator.JSON = {
                ...this.surveyCreator.JSON, description:value
            }
        },
    }

};
</script>
<style>
.svc-creator__banner, .svc-logo-image{
    display: none;
}
#surveyCreatorContainer *{
    font-family: Roboto, sans-serif;
}
.sv-popup {
    z-index: 10000000 !important;

}
.svd-designer-container--left-side {
    /*width: 30%;*/
    /*max-width: 30%;*/
    /*flex-basis: 30%;*/
}

</style>
