<template>
    <FormItem :prop=rule :label=label>
        <div >



            <label >Газар зүйн элементийн төрөл:</label>
            <RadioGroup v-model="layerType" @on-change="saveData">
                <Radio label="Point"></Radio>
                <Radio label="Polygon"></Radio>
                <Radio label="Polyline"></Radio>
            </RadioGroup>
            <hr>
            <br>
            <div>
                <label >OBJECTID хадгаллах талбар (Маягтын)</label>
                <Select v-model="objectIdField" style="width:200px" @on-change="saveData">
                    <Option v-for="item in meta.options" :value="item" :key="item.index">{{ item }}</Option>
                </Select>
            </div>
            <div>
                <label >Газарзүйн мэдээллийн талбар (Маягтын)</label>
                <Select v-model="geoJsonField" style="width:200px" @on-change="saveData">
                    <Option v-for="item in meta.options" :value="item" :key="item.index">{{ item }}</Option>
                </Select>
            </div>
            <br>
            <hr>
        <div>
            <br>
            <table>
                <tr>
                    <td>
                        <label >ArcGIS атрибут</label>
                        <Select v-model="gisAttribute" style="width:100%">
                            <Option v-for="item in gisAttributes" :value="item.value" :key="item.value">{{ item.label }}</Option>
                        </Select>
                    </td>
                    <td>
                        <label >Маягтын талбар</label>
                        <Select v-model="formField" style="width:100%">
                            <Option v-for="item in meta.options" :value="item" :key="item.index">{{ item }}</Option>
                        </Select>
                    </td>
                    <td>
                        <br>
                        <Button type="primary" shape="circle" icon="md-add" @click="addConnection"></Button>
                    </td>
                </tr>
            </table>
            <Table border :columns="columns" :data="connections"></Table>

        </div>


        </div>
    </FormItem>
</template>

<script>

    const ARCGISSERVER = "http://103.87.255.139:6080/arcgis/rest/services/"


    export default {

        props: ["model", "rule", "label", "meta", "getSchemaByModel", "isBuilder", "editMode", "setSchemaByModel", "do_render"],

        data() {
           return {

               gisAttribute:null,
               formField:null,
               columns: [

                   {
                       title: 'ArcGIS атрибут',
                       key: 'attribute'
                   },
                   {
                       title: 'Маягтын талбар',
                       key: 'field'
                   },
                   {
                       title: 'Action',
                       key: 'action',
                       width: 150,
                       align: 'center',
                       render: (h, params) => {
                           return h('div', [

                               h('Button', {
                                   props: {
                                       type: 'error',
                                       size: 'small'
                                   },
                                   on: {
                                       click: () => {
                                           this.remove(params.index)
                                       }
                                   }
                               }, 'Устгах')
                           ]);
                       }
                   }
               ],
               connections: [],
               objectIdField:null,
               geoJsonField:null,
               layerType:null,
           }
        },
        methods: {
            remove (index) {
                this.connections.splice(index, 1);

                this.saveData();
            },
            addConnection(){
                if(this.gisAttribute && this.formField){
                    this.connections.push({
                        attribute:this.gisAttribute,
                        field:this.formField,
                    });

                    this.gisAttribute = null;
                    this.formField = null;

                    this.saveData();
                }
            },
            saveData(){
                this.model.form[this.model.component] = JSON.stringify({
                    connections: this.connections,
                    objectIdField:this.objectIdField,
                    geoJsonField:this.geoJsonField,
                    layerType:this.layerType,
                });
            }
        },
        created() {

        },
        mounted() {
            if (this.model.form[this.model.component]) {

            }
        },
        computed : {
            gisAttributes(){
                return this.meta.gis_options ? this.meta.gis_options : [];
            },
            connectionValue() {
                return this.model.form[this.model.component]
            }
        },
        watch: {
            connectionValue(value, oldValue) {

                if ((oldValue && !value) || (value && !oldValue)) {
                    if (value) {

                        let data = JSON.parse(this.model.form[this.model.component])

                        this.connections = data.connections;
                        this.objectIdField = data.objectIdField;
                        this.geoJsonField = data.geoJsonField;
                        this.layerType = data.layerType;
                    } else {

                    }
                }

            },
            // do_render(value) {
            //     if (!value) {
            //         this.connections = []
            //         this.objectIdField = null;
            //         this.geoJsonField = null;
            //         this.layerType = null;
            //         Vue.set(this.model.form, this.model.component, this.meta.default ? this.meta.default : null);
            //     }
            // },

        }

    };
</script>
