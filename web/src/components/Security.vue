<template>
    <div>
        <v-form class="profile-form" v-model="isValid" ref="form">
            <v-row>
                <v-col :md="12">
                    <h3 class="text-center">Update Password</h3>
                </v-col>
                <v-col :sm="12" :md="6">
                    <v-text-field
                            v-model="password"
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
                <v-col class="fr">
                    <v-btn class="ma-2" color="primary" :dark="isValid" :disabled="!isValid" :outlined="!isValid"  @click="submit">Save</v-btn>
                </v-col>
            </v-row>
        </v-form>
    </div>
</template>

<script>
    export default {
        name: "Security",
        data: () => ({
            isValid: false,
            showPassword: false,
            passwordRepeat: null,
            password: null,
            formRules: {
                password: [
                    v => !!v || 'Şifre zorunlu alandır',
                    v => v && v.length > 6 || 'Şifre en az 7 karakter olmalıdır'
                ]
            },
        }),
        computed: {
            passwordConfirmationRule() {
                return () => (this.password === this.passwordRepeat) || 'Şifreler eşleşmiyor'
            },
        },
        methods: {
            submit() {
                console.log('submitted');
                this.$store.dispatch('changePassword', {
                    password: this.password
                })
            }
        }
    }
</script>
