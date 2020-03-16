<template>
    <div>
        <v-form class="form register-form" v-model="isValid" ref="form">
            <v-row>
                <v-col :sm="12">
                    <h2 class="text-center">Register</h2>
                </v-col>
                <v-col :sm="12" :md="6">
                    <v-text-field
                            v-model="formData.first_name"
                            label="First Name"
                            :rules="formRules.first_name"
                    />
                </v-col>
                <v-col :sm="12" :md="6">
                    <v-text-field
                            v-model="formData.last_name"
                            label="Last Name"
                            :rules="formRules.last_name"
                    />
                </v-col>
                <v-col :sm="12" :md="6">
                    <v-text-field
                            v-model="formData.username"
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
                <v-col class="df-sbc">
                    <v-btn class="ma-2" to="/login" outlined>Login</v-btn>
                    <v-btn class="ma-2" color="primary" :dark="isValid" :disabled="!isValid" @click="submit">Register</v-btn>
                </v-col>
            </v-row>

        </v-form>
    </div>
</template>

<script>
    export default {
        name: "Register",
        data: () => ({
            isValid: false,
            showPassword: false,
            formData: {
                first_name: '',
                last_name: '',
                email: '',
                username: '',
                password: '',
            },
            passwordRepeat: null,
            formRules: {
                first_name: [v => !!v || 'Ad zorunlu alandır'],
                last_name: [v => !!v || 'Soyad zorunlu alandır'],
                username: [
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
            }
        },
        methods: {
            submit() {
                this.$refs.form.validate()
            }
        }
    }
</script>