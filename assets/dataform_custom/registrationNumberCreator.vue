<template>
    <FormItem :prop=rule :label=label>
        <Row>
            <div class="registrationNumber">
                <Col span="5">
                    <Select v-model="reg1" filterable>
                        <Option v-for="item in cityList" :value="item.value" :key="item.value">{{ item.value }}</Option>
                    </Select>
                </Col>
                <Col span="5">
                    <Select v-model="reg2" filterable>
                        <Option v-for="item in cityList" :value="item.value" :key="item.value">{{ item.value }}</Option>
                    </Select>
                </Col>
                <Col span="14">
                    <Input v-model="regNumber" placeholder="Регистерийн дугаар" style="width: 100%"/>
                </Col>
            </div>
        </Row>
    </FormItem>
</template>

<script>
export default {
    name: "registrationNumberCreator",
    props: ["model", "rule", "label", "meta", "getSchemaRelationByModel"],
    data() {
        return {
            regNumber: '',
            reg1: 'А',
            reg2: 'А',
            cityList: [
                {value: 'А'},
                {value: 'Б'},
                {value: 'В'},
                {value: 'Г'},
                {value: 'Д'},
                {value: 'Е'},
                {value: 'Ё'},
                {value: 'Ж'},
                {value: 'З'},
                {value: 'И'},
                {value: 'Й'},
                {value: 'К'},
                {value: 'Л'},
                {value: 'М'},
                {value: 'Н'},
                {value: 'О'},
                {value: 'Ө'},
                {value: 'П'},
                {value: 'Р'},
                {value: 'С'},
                {value: 'Т'},
                {value: 'У'},
                {value: 'Ү'},
                {value: 'Ф'},
                {value: 'Х'},
                {value: 'Ц'},
                {value: 'Ч'},
                {value: 'Ш'},
                {value: 'Щ'},
                {value: 'Ъ'},
                {value: 'Ы'},
                {value: 'Ь'},
                {value: 'Э'},
                {value: 'Ю'},
                {value: 'Я'},
            ],

        }
    },
    methods: {
        setValue(){
            if (this.reg1 !== null && this.reg1 != "" && this.reg2 !== null && this.reg2 != "" && this.regNumber !== null && this.regNumber >= 0 ) {
                this.model.form[this.model.component] = `${this.reg1}${this.reg2}${this.regNumber}`
            } else {
                this.model.form[this.model.component] = '';
            }
        }
    },
    mounted() {
    },
    computed: {},
    watch: {
        reg1(){
            this.setValue();
        },
        reg2(){
            this.setValue();
        },
        regNumber(){
            this.setValue();
        },
    }
}
</script>

<style scoped>


</style>
