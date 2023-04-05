<template>
    <FormItem :prop=rule :label=label>
        <div >

            <Cascader :data="layers" v-model="value"
                      @on-change="setLayerValue"

                      :placeholder="meta && meta.placeHolder !== null ? meta.placeHolder : label"
                      :disabled="meta && meta.disabled ? meta.disabled : false"
            ></Cascader>
        </div>
    </FormItem>
</template>

<script>
    const ARCGISSERVER = "http://103.87.255.139:6080/arcgis/rest/services"


    export default {

        props: ["model", "rule", "label", "meta", "getSchemaByModel", "isBuilder", "editMode", "setSchemaByModel", "do_render"],

        data() {
            return {
                layers: [],
                value: [],
                featureFields:[],
                layerType:'feature',
                ignoreWatching:true,
                token:null,
            }
        },
        methods: {
            setLayerValue(value) {
                this.model.form[this.model.component] = JSON.stringify(value);

                this.getLayerData(value);
            },
            getLayerData(value, igNore){

                axios.get(value[value.length-1] + '?f=pjson',{transformRequest: (data, headers) => {

                        delete headers.common['X-CSRF-TOKEN'];
                        delete headers['X-CSRF-TOKEN'];
                        delete headers.common['X-Requested-With'];
                    }
                }).then(response => {



                    if(response.data.type == "Feature Layer"){

                        this.featureFields = response.data.fields.map(field=>{
                            return {
                                value:field.name,
                                label:field.name,
                            }
                        });

                        this.model.form["layer_type"] = "feature";



                        this.setSchemaByModel("popup_fields", "hidden", false);
                        this.setSchemaByModel("search_fields", "hidden", false);
                        this.setSchemaByModel("style_field", "hidden", false);
                        this.setSchemaByModel("info_template", "hidden", false);
                        this.setSchemaByModel("search_info", "hidden", false);
                        this.setSchemaByModel("show_label", "hidden", false);
                        this.setSchemaByModel("label_field", "hidden", false);
                        this.setSchemaByModel("label_color", "hidden", false);

                    } else {
                        this.featureFields = [];

                        this.model.form["layer_type"] = "mapserver";



                        this.setSchemaByModel("popup_fields", "hidden", true);
                        this.setSchemaByModel("search_fields", "hidden", true);
                        this.setSchemaByModel("style_field", "hidden", true);
                        this.setSchemaByModel("info_template", "hidden", true);
                        this.setSchemaByModel("search_info", "hidden", true);

                        this.setSchemaByModel("show_label", "hidden", true);
                        this.setSchemaByModel("label_field", "hidden", true);
                        this.setSchemaByModel("label_color", "hidden", true);
                    }


                    this.setSchemaByModel("popup_fields", "options", this.featureFields);
                    this.setSchemaByModel("search_fields", "options", this.featureFields);
                    this.setSchemaByModel("style_field", "options", this.featureFields);
                    this.setSchemaByModel("label_field", "options", this.featureFields);


                    if(!igNore){
                        this.model.form.popup_fields = null;
                        this.model.form.search_fields = null;
                        this.model.form.style_field = null;
                        this.model.form.info_template = null;
                        this.model.form.search_info = null;
                        this.model.form.show_label = null;
                        this.model.form.label_field = null;
                        this.model.form.label_color = null;




                    } else {
                        Vue.set(this.$data, 'ignoreWatching', false);
                    }

                })
            },
            showArcGISerror(){
              this.$Message.error('ArcGIS server тэй холбогдож чадсангүй !!!');
            },
            getToken(){
              axios.get("/arcgis/token").then(response => {
                  if(response.data.token){
                    this.getLayers(response.data.token)
                    this.token = response.data.token;
                  } else{
                    this.showArcGISerror();
                  }
              }).catch(()=>{
                this.showArcGISerror();
              })
            },
            getLayers() {

                axios.get(`${ARCGISSERVER}?f=json`,{transformRequest: (data, headers) => {

                        delete headers.common['X-CSRF-TOKEN'];
                        delete headers['X-CSRF-TOKEN'];
                        delete headers.common['X-Requested-With'];
                    }
                }).then(response => {


                    let folders = [];


                    if (response.data.folders.length >= 1) {
                        response.data.folders.map((folder, index) => {

                            folders.push({
                                value: `${ARCGISSERVER}${folder}`,
                                label: folder,
                                children: []
                            });
                            // let isLast = response.data.folders.length-1 == index ? true: false;
                            this.readArcFolder(`${ARCGISSERVER}/${folder}`, index, folders);

                        });
                    } else {

                        folders.push({
                            value: `${ARCGISSERVER}`,
                            label: 'root',
                            children: []
                        });
                        // let isLast = response.data.folders.length-1 == index ? true: false;
                        this.readArcFolder(`${ARCGISSERVER}`, 0, folders);
                    }





                })
            },
            readArcFolder(url, index, mainFolders) {

                axios.get(`${url}?f=pjson`,{transformRequest: (data, headers) => {
                        delete headers.common['X-CSRF-TOKEN'];
                        delete headers['X-CSRF-TOKEN'];
                        delete headers.common['X-Requested-With'];

                    }
                }).then(response => {
                    let folders = [];


                    if(response.data.services){
                      response.data.services.map((service, index_c) => {

                        let name = service.name.split('/');
                        if (service.type == 'FeatureServer') {

                          this.readArcFeature(`${url}/${name[1]}/${service.type}`, index, index_c, mainFolders);
                          folders.push({
                            value: `${url}/${name[1]}/${service.type}`,
                            label: `FeatureServer - ${name[1]}`,
                            children: []
                          });
                        } else {
                          folders.push({
                            value: `${url}/${name[1]}/${service.type}`,
                            label: `${service.type} - ${name[1]}`,
                          });
                        }
                      });

                      mainFolders[index].children = folders;

                      this.layers = mainFolders;
                    }


                });

            },
            readArcFeature(url, index, index_c, mainFolders, token) {

                axios.get(`${url}?f=pjson`,{transformRequest: (data, headers) => {

                        delete headers.common['X-CSRF-TOKEN'];
                        delete headers['X-CSRF-TOKEN'];
                        delete headers.common['X-Requested-With'];
                    }
                }).then(response => {
                    let features = [];

                    response.data.layers.map((layer, lindex) => {
                        features.push({
                            value: `${url}/${lindex}`,
                            label: layer.name,
                        });
                    });

                    mainFolders[index].children[index_c].children = features;
                    this.layers = mainFolders;

                });


            }
        },
        created() {

            const value = this.model.form[this.model.component];
            if (value !== null && value != "" && value != "''") {
                if(value){
                    this.value = JSON.parse(value);


                }
            }
        },
        mounted() {
            // console.log(this.isBuilder)
            // if(!this.isBuilder){
            //    this.getToken();
            // }
            this.getLayers()
        },
        computed : {
            popupFieldValue(){

                return this.model.form.popup_fields;
            },
            searchFieldValue(){
                return this.model.form.search_fields;
            },
            layerValue() {
                return this.model.form[this.model.component];
            },

        },
        watch: {
            do_render(value, oldValue){
                if(oldValue == true && value === false){
                    Vue.set(this.$data, 'ignoreWatching', true);
                }
            },
            popupFieldValue(value, oldValue){
                if(!this.ignoreWatching && this.do_render === true){
                    let fields = [];
                    if (value !== null && value != "" && value != "''") {
                        if(value){

                            fields = value.split(',');
                        }
                    }
                    let infoValue = "";
                    fields.forEach(field => {

                        infoValue = infoValue + "" + field + " " + "<b>{" + field + "}</b> <br/>";

                    });


                    this.model.form.info_template = infoValue;
                }

            },
            searchFieldValue(value, oldValue){
                if(!this.ignoreWatching  && this.do_render === true) {
                    let fields = [];
                    if (value !== null && value != "" && value != "''") {
                        if (value)
                            fields = value.split(',');
                    }
                    let searchValue = "";
                    fields.forEach((field, index) => {
                        if (index >= fields.length - 1) {
                            searchValue = searchValue + "" + field + " " + "{" + field + "}";
                        } else
                            searchValue = searchValue + "" + field + " " + "{" + field + "}, ";

                    });

                    this.model.form.search_info = searchValue;
                }
            },
            layerValue(value, oldValue) {

                if ((oldValue && !value) || (value && !oldValue)) {
                    if(value){

                        this.value = JSON.parse(value);
                        this.ignoreWatching = true;
                        this.getLayerData(this.value, true);
                    }

                }
            }
        }

    };
</script>
