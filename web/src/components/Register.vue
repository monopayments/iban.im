<template>
    <div>
        <v-form class="form register-form" v-model="isValid" ref="form">
            <v-row>
                <v-col :sm="12">
                    <h2 class="text-center">Register</h2>
                </v-col>
                <v-col :sm="12" :md="6">
                    <v-text-field
                            v-model="formData.firstName"
                            label="First Name"
                            :rules="formRules.firstName"
                    />
                </v-col>
                <v-col :sm="12" :md="6">
                    <v-text-field
                            v-model="formData.lastName"
                            label="Last Name"
                            :rules="formRules.lastName"
                    />
                </v-col>
                <v-col :sm="12" :md="6">
                    <v-text-field
                            v-model="formData.handle"
                            label="UserName"
                            :rules="formRules.username"
                    />
                </v-col>
                <v-col :sm="12" :md="6">
                    <v-text-field
                            v-model="formData.email"
                            label="Email"
                            type="email"
                            :rules="formRules.email"
                    />
                </v-col>
                <v-col :sm="12" :md="6">
                    <v-text-field
                            v-model="formData.password"
                            label="Password"
                            :rules="formRules.password"
                            :type="showPassword ? 'text' : 'password'"
                            :append-icon="showPassword ? 'visibility' : 'visibility_off'"
                            @click:append="showPassword = !showPassword"
                    />
                </v-col>
                <v-col :sm="12" :md="6">
                    <v-text-field
                            v-model="passwordRepeat"
                            label="Password Again"
                            :rules="[passwordConfirmationRule]"
                            :type="showPassword ? 'text' : 'password'"
                            :append-icon="showPassword ? 'visibility' : 'visibility_off'"
                            @click:append="showPassword = !showPassword"
                    />
                </v-col>
                <v-col :sm="12">
                    <v-checkbox
                            v-model="formData.visible"
                            label=" Email Adresimi Göster"
                    />
                </v-col>
                <v-col v-if="error" :sm="12">
                    <div class="error">
                        {{error}}
                    </div>
                </v-col>
                <v-col class="df-sbc">
                    <v-btn class="ma-2" to="/login" outlined>Login</v-btn>
                    <v-btn class="ma-2" color="primary" :dark="isValid" :disabled="!isValid" @click="submit">Register</v-btn>
                </v-col>
            </v-row>

        </v-form>
    </div>
</template>

<script>
    import { mapState } from 'vuex';

    export default {
        name: "Register",
        data: () => ({
            isValid: false,
            showPassword: false,
            formData: {
                firstName: '',
                lastName: '',
                email: '',
                handle: '',
                password: '',
                visible: false,
            },
            passwordRepeat: null,
            formRules: {
                firstName: [v => !!v || 'Ad zorunlu alandır'],
                lastName: [v => !!v || 'Soyad zorunlu alandır'],
                handle: [
                    v => !!v || 'Kullanıcı adı zorunlu alandır',
                    v => /^[A-Za-z0-9]*$/.test(v) || 'kullanıcı adı yalnızca harf ve rakam içerebilir'
                ],
                email: [
                    v => !!v || 'Eposta zorunludur',
                    v => /.+@.+/.test(v) || 'Geçerli bir eposta adresi girin',
                ],
                password: [
                    v => !!v || 'Şifre zorunlu alandır',
                    v => v.length > 6 || 'Şifre en az 7 karakter olmalıdır'
                ]
            },
        }),
        computed: {
            passwordConfirmationRule() {
                return () => (this.formData.password === this.passwordRepeat) || 'Şifreler eşleşmiyor'
            },
            ...mapState(['error'])
        },
        methods: {
            submit() {

                this.$store.dispatch('register', this.formData)
            }
        }
    }
</script>