<template>
    <FormItem :label=label :prop=rule>
        <multiselect v-if="!meta.relation.multiple"
                     v-model="value"
                     :disabled="meta.disabled"
                     :options="options"
                     @search-change="searchChange"
                     track-by="value"
                     :searchable="true"
                     :allow-empty="true"
                     :placeholder="meta && meta.placeHolder !== null ? meta.placeHolder : label ? label : ''"
                     label="label">
            <template slot="singleLabel" slot-scope="{ option }">
                {{ option.label }}
            </template>

            <template slot="caret"
                      slot-scope="{ toggle }"
                      @mousedown.prevent.stop="toggle">
                <div class="caret-container">
                    <div :class="addAble ? 'multiselect__select addable-caret' : 'multiselect__select'">
                    </div>
                    <Button v-if="addAble" @click="showAddModal" type="success"
                            shape="circle" size="small"
                            icon="md-add"></Button>
                </div>
            </template>
            <template slot="clear" slot-scope="{ search }"
                      @mousedown.prevent.stop="toggle">
                <div class="clear-container">
                    <Button v-if="clearAble" @click="clearState"
                            shape="circle" size="small"
                            icon="md-close"></Button>
                </div>
            </template>
        </multiselect>

        <multiselect v-else
                     :multiple="true"
                     v-model="value"
                     :disabled="meta.disabled"
                     track-by="value"
                     :searchable="true"
                     @search-change="searchChange"
                     :placeholder="meta && meta.placeHolder !== null ? meta.placeHolder : label ? label : ''"
                     label="label"
                     :options="options">
            <template slot="caret"
                      slot-scope="{ toggle }"
                      @mousedown.prevent.stop="toggle">
                <div class="caret-container">
                    <div :class="addAble ? 'multiselect__select addable-caret' : 'multiselect__select'">
                    </div>
                    <Button v-if="addAble" @click="showAddModal" type="success"
                            shape="circle" size="small"
                            icon="md-add"></Button>
                </div>
            </template>
        </multiselect>

        <paper-modal
                :name="`add-modal-${model.component}`"
                class="add-modal"
                :min-width="200"
                :min-height="100"
                :pivot-y="0.5"
                :adaptive="true"
                :reset="true"
                :draggable="true"
                :resizable="true"
                draggable=".add-tool"
                width="600"
                height="70%"
                v-if="addAble">
            <section class="add-modal">
                <div class="add-tool">

                    <h4>{{label}}</h4>
                    <div class="add-tool-actions">


                        <a href="javascript:void(0)" @click="closeModal">
                            <i class="ti-close"></i>
                        </a>
                    </div>
                </div>

                <div class="add-body">
                    <dataform ref="form" :schemaID="meta.relation.addFrom"
                              :editMode="false"
                              :onSuccess="onSuccess"
                              :do_render="modal_show"
                              :onError="onError"></dataform>
                </div>
            </section>
        </paper-modal>
    </FormItem>

</template>

<script>

    import {isValid} from "./methods"
    import { mapGetters} from 'vuex'
    export default {
        props: ["model", "rule", "label", "meta", "disabled", "isBuilder", "relation_data", "do_render"],
        computed: {
            ...mapGetters(["attributes"]),
            selectValue() {
                return this.model.form[this.model.component];
            },
            options() {
                if(!this.isBuilder){
                    if (this.attributes.length >= 1) {

                        return this.filterOption(this.attributes)
                    }

                }
                return this.filterOption([]);

            },
        },
        data() {
            return {
                value: null,
                ignoreChange: false,
                addAble: false,
                clearAble:false,
                modal_show: false,
                searchval:null,
                //attributes:[]
            }
        },
        methods: {
            isValid: isValid,
            filterOption(options) {
                if (options) {
                    this.initialValue(options);
                    return options ? options : [];
                }
                return [];
            },
            searchChange (val) {
                this.searchval = val;
            },
            initialValue(options) {
                if (!this.ignoreChange) {
                    if (this.model.form[this.model.component]) {
                        //If multiple
                        if (this.meta.relation.multiple == true) {
                            let selectedData = this.model.form[this.model.component].toString().split(',');
                            let filtered = options.filter(item => selectedData.includes(item.value.toString()));
                            if(filtered.length >= 1)
                                this.value = filtered

                        } else {
                            let filtered = options.filter(item => item.value == this.model.form[this.model.component]);
                            this.value = filtered.length >= 1 ? filtered[0] : null;
                        }
                    } else {
                        //trigger ajillah uyd umnuh ugugdluu tseverleh
                        this.value = null;
                    }
                }
            },
            showAddModal() {
                this.modal_show = true;
                this.$modal.show(`add-modal-${this.model.component}`);
            },
            clearState()
            {
                this.value=null;
                this.clearAble=false;
            },
            closeModal() {
                this.modal_show = false;
                this.$modal.hide(`add-modal-${this.model.component}`);
            },
            //Form functions
            onSuccess(val) {

                let label = this.meta.relation.fields.map(field => val[field]);
                label = label.join(', ');

                this.relation_data.push({
                    value: val[this.meta.relation.key],
                    label: label
                });
                this.closeModal();
            },

            onError(val) {

            }
        },
        watch: {
            do_render(value) {
                if (!value) {
                    this.value = null;
                    this.clearAble=false;
                    this.ignoreChange = false;
                    Vue.set(this.model.form, this.model.component, this.meta.default ? this.meta.default : undefined);
                } else {
                    if (this.model.form[this.model.component]) {
                        let value = this.model.form[this.model.component];
                        //If multiple
                        if (this.meta.relation.multiple == true) {
                            let selectedData = value.toString().split(',');
                            this.value = this.options.filter(item => selectedData.includes(item.value.toString()));
                        } else {
                            let filtered = this.options.filter(item => item.value == value);
                            this.value = filtered.length >= 1 ? filtered[0] : null;
                        }
                        this.clearAble=true;
                    }
                }
            },
            value(val) {
                if (val) {
                    //trigger ajillah uyd ugugdluu solihgui haav
                    //this.ignoreChange = true;
                    if (this.meta.relation.multiple == true) {
                        Vue.set(this.model.form, this.model.component, val.map(vv => vv['value']).join(','));
                    } else {
                        Vue.set(this.model.form, this.model.component, val['value']);
                    }
                    this.clearAble=true;
                }
            },


        },

        created() {
            if (this.meta.relation.addAble && this.meta.relation.addFrom) {
                this.addAble = true;
            }


        }
    };
</script>

<style>
    .v-select .form-control {
        position: absolute;
        top: 0;
        width: 100% !important;
    }

    .v-select .selected-tag {
        width: 100%;
    }

</style>
