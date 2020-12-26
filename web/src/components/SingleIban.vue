<template>
    <div v-if="profile" class="single-iban">
        <div class="single-iban__button">
            <v-tooltip right>
                <template v-slot:activator="{ on, attrs }">
                    <v-btn 
                        icon
                        color="pink"
                        v-bind="attrs"
                        v-on="on"
                        :to="{name: 'home.single', params: {username: $route.params.username}}">
                        <v-icon>mdi-keyboard-return</v-icon>
                    </v-btn>
                
                </template>
                <span>Profile</span>
            </v-tooltip>
        </div>
        <ul v-if="current && !current.isPrivate">
            <li><span><v-icon left>mdi-account</v-icon>IBAN name</span><span>{{name}}</span></li>
            <li><span><v-icon left>mdi-bank</v-icon>Handle</span><span>{{current.handle}}</span></li>
            <li>
                <span>
                    <v-icon left>
                        mdi-cash-multiple
                    </v-icon>
                    IBAN
                </span>
                <span>
                    {{current.text}}
                </span>
                <span>
                    <v-icon 
                        left 
                        style="cursor: pointer;"
                        v-clipboard:copy="current.text"
                        v-clipboard:success="onCopy"
                        v-clipboard:error="onError">
                        mdi-content-copy
                    </v-icon>
                </span>
            </li>
            <li>
                <span>
                    {{current.description}}
                </span>
            </li>
        </ul>
        <v-form v-else-if="current && current.isPrivate" class="show-info" v-model="isValid">
            <v-row>
                <v-col :md="6" :sm="12">
                    <v-text-field
                            v-model="formData.password"
                            label="Password"
                            :rules="formRules.password"
                    />
                </v-col>
                <v-col :md="6" :sm="12">
                    <div class="show-submit">
                        <v-btn class="ma-2" :dark="isValid" :disabled="!isValid" color="primary" @click="showInfo">Show</v-btn>
                    </div>
                </v-col>
            </v-row>
        </v-form>
        <div v-else>
            <b>An account named <i>{{$route.params.alias}}</i> was not found</b>
        </div>
    </div>
</template>

<script>
    import {mapActions, mapState} from 'vuex';
    export default {
        name: "SingleIban",
        data: () => ({
            canShow: false,
            isValid: false,
            formData: {
                password: '',
            },
            formRules: {
                password: [v => !!v || 'Password is required']
            }
        }),
        computed: {
            ...mapState(['ibans','profile']),
            current() {
                return this.ibans.filter( iban => iban.handle.toLowerCase() === this.$route.params.alias.toLowerCase())[0]
            },
            name() {
                return `${this.profile.firstName} ${this.profile.lastName}`
            }
        },
        created() {
            this.fetchSingleProfile({
                username : this.$route.params.username
            });
        },
        methods: {
            ...mapActions({
                fetchSingleProfile: 'fetchSingleProfile',
                checkShowPassword: 'checkShowPassword',
            }),
            showInfo() {
                this.checkShowPassword({
                    id       : this.current.id,
                    password : this.formData.password
                });
            },
            onCopy() {
                // TODO: can be added alert library or something
                alert('Iban was copied to clipboard!')
            },
            onError() {
                alert('Something went wrong!')
            }
        },
        watch : {
            '$store.state.canShow'(ok) {
                if(this.current.isPrivate && ok) {
                    this.current.isPrivate = false;
                }
            },
        }
    }
</script>

<style>

</style>