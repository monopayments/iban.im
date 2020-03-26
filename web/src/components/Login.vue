<template>
    <div>
        <v-form class="form login-form" v-model="isValid">
            <v-row>
                <v-col :sm="12">
                    <h2 class="text-center">Login</h2>
                </v-col>

                <v-col :sm="12">
                    <v-text-field
                            v-model="formData.email"
                            label="Email"
                            :rules="formRules.email"
                    />
                </v-col>
                <v-col :sm="12">
                    <v-text-field
                            v-model="formData.password"
                            label="Password"
                            :rules="formRules.password"
                    />
                </v-col>

                <v-col v-if="error" :sm="12">
                    <div class="error">
                        {{error}}
                    </div>
                </v-col>
                <v-col class="df-sbc">
                    <v-btn class="ma-2" outlined to="/register">Register</v-btn>
                    <v-btn class="ma-2" :dark="isValid" :disabled="!isValid" color="primary" @click="submit">Login</v-btn>
                </v-col>
            </v-row>

        </v-form>
    </div>
</template>

<script>

    import { mapState,mapActions } from 'vuex';

    export default {
        name: "Login",
        data: () => ({
            isValid: false,
            formData: {
                email: '',
                password: '',
            },
            formRules: {
                email: [
                    v => !!v || 'Eposta zorunludur',
                    v => /.+@.+/.test(v) || 'Geçerli bir eposta adresi girin',
                ],
                password: [v => !!v || 'Şifre zorunlu alandır']
            }
        }),
        computed: {
            ...mapState(['error'])
        },
        created() {
            this.setLoaded(true);
        },
        methods: {
            ...mapActions({
                setLoaded: 'setLoaded',
            }),
            submit() {
                this.$store.dispatch('login', this.formData);
            }
        }
    }
</script>