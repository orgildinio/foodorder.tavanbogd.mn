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

        props: ["model", "rule", "label", "meta", "getSchemaByModel", "isBuilder", "editMode", "setSchemaByModel"],

        data() {
            return {
                layers: [],
                value: [],
                featureFields:[],
                layerType:'feature',
                ignoreWatching:true
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

                    }

                    this.setSchemaByModel("connection", "gis_options", this.featureFields);



                    if(!igNore){
                        this.model.form.connection = null;




                    } else {
                        Vue.set(this.$data, 'ignoreWatching', false);
                    }

                })
            },
            getLayers() {

                axios.get(ARCGISSERVER + `?f=pjson`,{transformRequest: (data, headers) => {

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
                        // this.readArcFolder(`${ARCGISSERVER}`, 0, folders);
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
          readArcFeature(url, index, index_c, mainFolders) {

            axios.get(`${url}?f=pjson`,{transformRequest: (data, headers) => {

                delete headers.common['X-CSRF-TOKEN'];
                delete headers['X-CSRF-TOKEN'];
                delete headers.common['X-Requested-With'];
              }
            }).then(response => {
              let features = [];


              if(response.data){
                  if(response.data.layers){
                      response.data.layers.map((layer, lindex) => {
                          features.push({
                              value: `${url}/${lindex}`,
                              label: layer.name,
                          });
                      });

                      mainFolders[index].children[index_c].children = features;
                      this.layers = mainFolders;
                  }

              }


            });


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
          showArcGISerror(){
            this.$Message.error('ArcGIS server тэй холбогдож чадсангүй !!!');
          },
        },
        created() {

          /*  const value = this.model.form[this.model.component];
            if (value !== null && value != "" && value != "''") {
                if(value){
                    this.value = JSON.parse(value);
                }
            }*/
        },
        mounted() {

           // this.getToken();
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
            }
        },
        watch: {
            popupFieldValue(value, oldValue){

                if(!this.ignoreWatching){
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
                if(!this.ignoreWatching) {
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
