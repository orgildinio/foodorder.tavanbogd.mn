<template>
    <div class="cards_container" >

        <Row :gutter="20">
            <Col span="10"   v-for="projectdata in project_adv" :value="projectdata.id" :key="projectdata.id">
                <div class="cards_per_cont">
                    <div class="title_cont" :title="projectdata.garchig">
                        <p>{{subStrFunction("title", projectdata.garchig)}} &nbsp; </p>
                    </div>

                    <div class="body_cont">
                        <div class="startdate"><p><b>Эхлэх:</b> {{subStrFunction("enddate", projectdata.start_date)}}</p></div>
                        <div class="desctxt"><p>{{subStrFunction("desc", projectdata.tailbar)}} &nbsp; </p></div>
                        <div class="p_ognoo">
                            <p style="float: left;">Журамтай танилцах бол <a :href="projectdata.attach_files" target="_blank"><b> ЭНД </b></a> дарна уу</p>
                            <p style="float: right;"><b>Дуусах:  </b> {{subStrFunction("enddate", projectdata.end_date)}} &nbsp;</p>
                        </div>
                    </div>

                    <div class="footer_cont">
                        <div class="projectbtn">
<!--                            href="/admin#/p/c49b40d5-b4e2-9525-bb0a-1400edb63c43/90382fff-174e-0cfc-17b6-f80060406f3c"-->
                            <a v-if="canAdd(projectdata.id)">
                                <Button type="success"  v-if="projectdata.project_status_id == 2">Төсөл илгээсэн</Button>
                                <Button v-else :type="projectdata.editID == 0 ? 'error' : 'primary'"  @click="modalDataSet2(projectdata)" >{{ projectdata.buttonTxt }}</Button>

                            </a>
                        </div>
                        <div class="descbtn">
                            <Button @click="modalDataSet(projectdata.id)" >Дэлгэрэнгүй</Button>
<!--                            <Button @click="modalDelgerengvi(projectdata.id)" >Дэлгэрэнгүй2</Button>-->
                        </div>
                    </div>
                </div>
            </Col>

        </Row>
<!--        modal duudax - delgerengvi popup-->
        <Modal class="projectdesccont delgerengvi"
            v-model="modaldata"
               width="70%"
               title="ТӨСЛИЙН ЗАР"
               height="70%"
               footer-hide
               :styles="{top: '20px'}"
        >
            <ReadOnlyView v-if="modaldata" :id="modaldescid"></ReadOnlyView>
        </Modal>

        <!--        modal duudax 2 - nemex popup-->
        <Modal class="projectdevsccont"
               v-model="modaldata_add"
               width="1000"
               title="ТӨСӨЛ БЭЛТГЭХ"
               height="100%"
               footer-hide
               :styles="{top: '20px', 'max-height':'100vh'}"
        >
            <AddProjectModal v-if="modaldata_add"  :zar_id="currentZarID" :editID="editZarID"></AddProjectModal>
        </Modal>

        <!--        modal duudax 3 - delgerengvi 2 -->
        <Modal class="projectdevsccont"
               v-model="modaldatadelgerengvi"
               width="1000"
               title="ТӨСЛИЙН ЗАР"
               height="70%"
               footer-hide
               :styles="{top: '20px', 'max-height':'80vh'}"
        >
          <div class="zar_details">
              <h1>{{modaldesciddelgerengvi}} = {{project_adv_details.id}}</h1>
              <p></p>
              <div>
                  <table style="width:100%">
                      <tr>
                          <th>ДД</th>
                          <th>ТӨРӨЛ</th>
                      </tr>
                      <tr v-for="(zar, index) in project_adv_details.sub_project_shaardlaga" :key="index">
                          <td>{{ index+1 }}</td>
                          <td>{{ zar.project_type }}</td>
                      </tr>
                  </table>
              </div>
              <div>
                  <table style="width:100%" border>
                      <tr>
                          <th>ДД</th>
                          <th>ТӨРӨЛ</th>
                      </tr>
                      <tr v-for="(zar1, index1) in project_adv_details.sub_project_type" :key="index1">
                          <td>{{ index1+1 }}</td>
                          <td>{{ zar1.project_type_id }}</td>
                      </tr>
                  </table>
              </div>
              <div>
                  <table style="width:100%">
                      <tr>
                          <th>ДД</th>
                          <th>ТӨРӨЛ</th>
                      </tr>
                      <tr v-for="(zar2, index2) in project_adv_details.sub_project_type_other" :key="index2">
                          <td>{{ index2+1 }}</td>
                          <td>{{ zar2.project_type }}</td>
                      </tr>
                  </table>
              </div>
          </div>

<!--            <AddProjectModal v-if="modaldata_add"  :zar_id="projectaddid"></AddProjectModal>-->
        </Modal>

    </div>
</template>
<script>
import {GET_PROJECT_ZAR, GET_PROJECT_USER_CHECK, GET_PROJECT_ZAR_BY_ID, GET_ZAR_BY_ID} from "../graphql/queries";
import ReadOnlyView from "./readOnlyView";
import AddProjectModal from "./addProjectModal";
export default {
    name: "prjctadv",
    data() {
        return {
            user_id: window.init.user.id,
            project_adv: [],
            project_adv_details: {},
            registered_projects: [],
            modaldata: false,
            modaldata_add: false,
            modaldatadelgerengvi: false,
            modaldescid: 0,
            projectaddid: 0,
            modaldesciddelgerengvi: 0,
            editZarID: 0,
            currentZarID: 0,
        };
    },
    components:{
        ReadOnlyView:ReadOnlyView,
        AddProjectModal: AddProjectModal,
    },
    mounted() {
        this.getProjectUserRegCeck();

    },
    methods: {
        changeShowResult(projectdata_id){
            if(this.project_adv.length >= 1){
                let index = this.project_adv.findIndex(result=>result.id == projectdata_id);
                if(index >= 0){
                    return true;
                } else {
                    return  false;
                }
            } else {
                return false;
            }
        },
        canAdd(zariD){
            let foundIndex = this.project_adv.findIndex(userProject=>userProject.form_id == zariD);
            if(foundIndex >= 0){ return false; }
            else { return true; }
        },

        //this.registered_projects - d tusul bvrtgvvlex table ees login xiisen user id gaar oroltsox xvselt ilgeesen tusliin id nuudiig array bolgoj butsaana
        async getProjectUserRegCeck(){
            let registered_projects = [];
            await this.$apollo.query({
                query: GET_PROJECT_USER_CHECK,
                variables: {user_id: window.init.user.id.toString()}
            }).then(({data}) => {



                this.getProjectAdv(data.z_project_user_check);
            });
        },

        // delgerengvi buttonii popup
        modalDataSet(descid){
            this.modaldata = true;
            this.modaldescid = descid;
        },
        modalDelgerengvi(descid){
            this.modaldatadelgerengvi = true;
            this.modaldesciddelgerengvi = descid;
            this.getZarDetails(descid);
        },

        // nemex button ii poup
        modalDataSet2(projectZar){

            this.currentZarID = projectZar.id;
            this.editZarID = projectZar.editID;
            this.modaldata_add = true;


        },

        // card deerxi text iin urt taaruulax xeseg
        subStrFunction(type, strings){
            let zalgaas = "";
            if(type === "title"){
                let strlenght = strings.length;
                if(strlenght > 75){ zalgaas = " ..."; }
                return strings.substr(0, 75) + zalgaas;
            }
            if (type === "desc"){
                let strlenght = strings.length;
                if(strlenght > 210){ zalgaas = " ..."; }
                return strings.substr(0, 210) + zalgaas;
            }
            if (type === "enddate" && strings !== null ){ return strings.substr(0, 10); }
        },

        // tusliin zar table oroltsox xvselt ilgeesen tusliin id nuudaas bwal xasaj push xiine - fill data
        async  getProjectAdv(userProjects){
             await  this.$apollo.query({
                query: GET_PROJECT_ZAR,
                 // variables: {user_id: this.user_id.toString()}
            }).then(({data}) => {


                 this.project_adv = data.tbl_project_zar.map(zar=>{
                     let buttonTxt = "Нэмэх";
                     let editID = 0;
                     let project_status_id = 0;

                     let userProjectIndex = userProjects.findIndex(userProject=>userProject.form_id == zar.id);
                     if(userProjectIndex >= 0){
                         buttonTxt = "Засах" ;
                         editID = userProjects[userProjectIndex].id
                         project_status_id = userProjects[userProjectIndex].project_status_id
                     }

                     return {...zar, buttonTxt:buttonTxt, editID:editID, project_status_id:project_status_id}
                 });
            });
        },


        async  getZarDetails(project_ids){
            console.log( "project_ids >> ");
            console.log( project_ids);
            await  this.$apollo.query({
                query: GET_PROJECT_ZAR_BY_ID,
                variables: {z_id: project_ids.toString()}
            }).then(({data}) => { this.project_adv_details  = data.tbl_project_zar[0]; console.log( this.project_adv_details ) });
        },

    },
}
</script>
<style scoped>

</style>
