<template>
    <div v-if="profile" class="single-iban">
        <ul v-if="!current.isPrivate">
            <li><span><v-icon left>mdi-account</v-icon>Hesap Adı</span><span>{{name}}</span></li>
            <li><span><v-icon left>mdi-bank</v-icon>Handle</span><span>{{current.handle}}</span></li>
            <li><span><v-icon left>mdi-cash-multiple</v-icon>IBAN</span><span>{{current.text}}</span></li>
        </ul>
        <v-form v-else class="show-info" v-model="isValid">
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
                password: [v => !!v || 'Şifre zorunlu alandır']
            }
        }),
        computed: {
            ...mapState(['ibans','profile']),
            current() {
                return this.ibans.filter( iban => iban.handle === this.$route.params.alias)[0]
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