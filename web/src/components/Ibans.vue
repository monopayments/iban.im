<template>
    <div class="iban-wrap">

        <div class="i-list">
            <v-list flat>
                <v-list-item-group v-model="selectedIndex" color="primary">
                    <v-list-item
                        v-for="(item,i) in items"
                        :key="i"
                        class="iban-item"
                    >
                        <v-list-item-icon>
                            <v-icon v-if="selectedIndex === i">mdi-minus</v-icon>
                            <v-icon v-else>mdi-plus-circle-outline</v-icon>
                            <v-icon class="q-delete" @click="remove(i)">mdi-delete</v-icon>
                        </v-list-item-icon>
                        <v-list-item-content>
                            <v-list-item-title v-text="item.handle" />
                        </v-list-item-content>
                    </v-list-item>
                </v-list-item-group>
            </v-list>
        </div>
        <div class="i-form pt-4" v-if="showForm">
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
                            label="Iban No"
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
                    <v-btn @click="cancel" class="text-none add-question" outlined>
                        <v-icon>mdi-minus</v-icon>
                        Cancel
                    </v-btn>
                    <v-btn color="primary" dark @click="save" class="text-none add-question">
                        <v-icon dark>mdi-plus</v-icon>
                        Save
                    </v-btn>
                </div>
            </v-form>
        </div>
        <v-btn class="text-none add-iban" v-if="!showForm" @click="show" rounded>
            <v-icon>mdi-plus</v-icon> Add
        </v-btn>
    </div>
</template>

<script>

    import { cloneDeep } from 'lodash';

    function reset() {
        return {
            handle: '',
            text: '',
            isPrivate: false,
            password: '',
        }
    }

    export default {
        name: "Ibans",
        data: () => ({
            items: [],
            valid: false,
            showForm: false,
            selectedIndex: undefined,
            current: reset(),
            formRules: {
                text: [
                    v => !!v || 'iban zorunlu alandır',
                ],
                handle: [
                    v => !!v || 'handle zorunlu alandır',
                    v => /^[A-Za-z0-9]*$/.test(v) || 'handle yalnızca harf ve rakam içerebilir'
                ],
            },
        }),
        computed: {
            passwordRule() {
                return () => (this.current.isPrivate && this.current.password !== '') || 'Lütfen şifre giriniz'
            },
        },
        methods: {
            remove(index) {
                this.$delete(this.items,index);
                this.selectedIndex = undefined;
                this.showForm = false;
                this.current = reset();
            },
            cancel() {
                this.showForm = false;
                this.selectedIndex = undefined;
                this.current = reset();
            },
            save() {
                if(this.selectedIndex !== undefined){
                    this.items[this.selectedIndex] = cloneDeep(this.current)
                }else{
                    this.items.push(cloneDeep(this.current))
                }
                this.current = reset();
                this.showForm = false;
                this.selectedIndex = undefined;
            },
            show() {
                this.current = reset();
                this.selectedIndex = undefined;
                const self = this;
                setTimeout(function () {
                    self.showForm = true;
                },100)
            }
        },
        watch:{
            selectedIndex (newValue,oldValue)  {
                console.log(newValue,oldValue);
                if(newValue === undefined){
                    this.current = reset();
                    this.showForm = false;
                }else{
                    this.current = cloneDeep(this.items[newValue]);
                    this.showForm = true;
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