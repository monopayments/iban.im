<template>
    <div class="iban-wrap">
        <h3 class="text-center">IBANs</h3>
        <div class="i-list">
            <v-list flat>
                <v-list-item-group v-model="selectedIndex" color="primary">
                    <v-list-item
                        v-for="(item,i) in ibans"
                        :key="i"
                        class="iban-item"
                    >
                        <v-list-item-icon>
                            <v-icon v-if="selectedIndex === i">mdi-minus</v-icon>
                            <v-icon v-else>mdi-plus-circle-outline</v-icon>
                        </v-list-item-icon>
                        <v-list-item-content>
                            <v-list-item-title v-text="item.handle" />
                        </v-list-item-content>
                    </v-list-item>
                </v-list-item-group>
            </v-list>
        </div>
        <v-dialog v-model="dialog" max-width="600px">
            <template v-slot:activator="{ on }">
                <v-btn class="text-none add-iban" v-on="on" @click="show" rounded>
                    <v-icon>mdi-plus</v-icon> Add
                </v-btn>
            </template>
            <v-card>
                <v-card-text>
                    <v-container>
                        <div class="i-form pt-4">
                            <v-form ref="form" v-model="valid">
                                <div class="form-item">
                                    <v-text-field
                                            v-model="current.handle"
                                            label="Handle"
                                            :rules="formRules.handle" />
                                </div>
                                <div class="form-item">
                                    <v-text-field
                                            v-model="current.text"
                                            label="IBAN No"
                                            :rules="formRules.text" />
                                </div>
                                <div class="form-item">
                                    <v-checkbox
                                            v-model="current.isPrivate"
                                            label=" Private"
                                    />
                                </div>
                                <div class="form-item" v-if="current.isPrivate">
                                    <v-text-field
                                            v-model="current.password"
                                            label="Password"
                                            :rules="[passwordRule]" />
                                </div>
                                <div class="form-item form-buttons">
                                    <v-btn v-if="current.id === ''" @click="cancel" class="text-none add-question" outlined>
                                        <v-icon>mdi-minus</v-icon>
                                        Cancel
                                    </v-btn>
                                    <v-btn v-else @click="remove" class="text-none add-question" outlined>
                                        <v-icon>mdi-minus</v-icon>
                                        Delete
                                    </v-btn>
                                    <v-btn color="primary" dark @click="save" class="text-none add-question">
                                        <v-icon dark>mdi-plus</v-icon>
                                        Save
                                    </v-btn>
                                </div>
                            </v-form>
                        </div>
                    </v-container>
                </v-card-text>
            </v-card>
        </v-dialog>

    </div>
</template>

<script>
    import { mapActions,mapState } from 'vuex';
    import { cloneDeep } from 'lodash';

    function reset() {
        return {
            id: "",
            handle: '',
            text: '',
            isPrivate: false,
            password: '',
        }
    }

    export default {
        name: "Ibans",
        data: () => ({
            dialog: false,
            valid: false,
            showForm: false,
            selectedIndex: undefined,
            current: reset(),
            formRules: {
                text: [
                    v => !!v || 'IBAN is required',
                ],
                handle: [
                    v => !!v || 'IBAN handle is required',
                    v => /^[A-Za-z0-9]*$/.test(v) || 'please only use a-z,A-Z or 0-9'
                ],
            },
        }),
        computed: {
            ...mapState(['ibans']),
            passwordRule() {
                return () => (this.current.isPrivate && this.current.password !== '') || 'Please provide password'
            },
        },
        created() {
            this.fetchProfile();
            this.fetchIbans();
        },
        methods: {
            ...mapActions({
                fetchProfile: 'fetchProfile',
                fetchIbans: 'fetchIbans',
                ibanUpdate: 'ibanUpdate',
                ibanDelete: 'ibanDelete',
            }),
            remove() {
                console.log(this.selectedIndex);
                const r = confirm('Are you sure?');
                if (r !== true) {
                    return;
                }

                this.ibanDelete(this.ibans[this.selectedIndex].id).then((data) => {
                    if(data.errors){
                        alert(data.errors[0].message);
                        return;
                    }
                    if(data.data.ibanDelete.ok){
                        this.$delete(this.ibans,this.selectedIndex);
                        this.selectedIndex = undefined;
                        this.dialog = false;
                        this.current = reset();
                    }else{
                        alert(data.data.ibanDelete.msg);
                    }
                })
            },
            cancel() {
                this.dialog = false;
                this.selectedIndex = undefined;
                this.current = reset();
            },
            save() {
                // yoksa null hatası veriyor
                if(!this.current.isPrivate){
                    this.current.password = '';
                }
                const process = this.current.id === "" ? "ibanNew" : "ibanUpdate";
                this.ibanUpdate(this.current).then((data) => {
                    if(data.errors){
                        alert(data.errors[0].message);
                        return;
                    }
                    if(!data.data[process].ok){
                        alert(data.data[process].msg);
                    }else{
                        this.current.id = data.data[process].iban.id;
                        if(this.selectedIndex !== undefined) {
                            this.ibans[this.selectedIndex] = cloneDeep(this.current)
                        }else{
                            this.ibans.push(cloneDeep(this.current));
                        }
                        this.current = reset();
                        this.dialog = false;
                        this.selectedIndex = undefined;
                    }
                });

            },
            show() {
                this.current = reset();
                this.selectedIndex = undefined;
                const self = this;
                setTimeout(function () {
                    self.dialog = true;
                },100)
            }
        },
        watch:{
            selectedIndex (newValue,oldValue)  {
                console.log(newValue,oldValue);
                if(newValue === undefined){
                    this.current = reset();
                    this.dialog = false;
                }else{
                    this.current = cloneDeep(this.ibans[newValue]);
                    this.dialog = true;
                }
            }
        }
    }
</script>

<style scoped>

    .form-buttons {
        display: flex;
        justify-content: space-between;
    }

    .iban-item{
        position: relative;
    }
    .mdi-delete{
        position: absolute !important;
        right: 0;
    }

</style>