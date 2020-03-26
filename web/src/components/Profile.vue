<template>
    <div>
        <v-form v-if="this.model" v-model="isValid" class="profile-form" ref="form">
            <v-row>
                <v-col :md="12">
                    <h3 class="text-center">Profile</h3>
                </v-col>
                <v-col :sm="12" :md="12">
                    <v-text-field
                            v-model="model.email"
                            label="Email"
                            disabled
                    />
                </v-col>

                <v-col :sm="12" :md="6">
                    <v-text-field
                            v-model="model.firstName"
                            label="First Name"
                            disabled
                    />
                </v-col>
                <v-col :sm="12" :md="6">
                    <v-text-field
                            v-model="model.lastName"
                            label="Last Name"
                            disabled
                    />
                </v-col>
                <v-col :sm="12" :md="6">
                    <v-text-field
                            v-model="model.handle"
                            label="Username"
                            :rules="formRules.handle"
                    />
                </v-col>
                <v-col :sm="12">
                    <v-textarea
                            v-model="model.bio"
                            label="Bio"
                            rows="3"
                    />
                </v-col>
                <v-col v-if="error" :sm="12">
                    <div class="error">
                        {{error}}
                    </div>
                </v-col>
                <v-col class="fr">
                    <v-btn class="ma-2" color="primary" :disabled="!isValid" :dark="isValid" @click="submit">Save</v-btn>
                </v-col>
            </v-row>
        </v-form>
    </div>
</template>

<script>
    import { mapActions } from 'vuex';
    export default {
        name: "Profile",
        data: () => ({
            isValid: false,
            error: null,
            model: {
                email: '',
                firstName: '',
                lastName: '',
                handle: '',
                bio: '',
            },
            formRules: {
                handle: [
                    v => !!v || 'Kullanıcı adı zorunlu alandır',
                    v => /^[A-Za-z0-9]*$/.test(v) || 'kullanıcı adı yalnızca harf ve rakam içerebilir'
                ],
            },
        }),
        created() {
            this.fetchProfile();
        },
        methods: {
            ...mapActions({
                fetchProfile: 'fetchProfile',
                changeProfile: 'changeProfile',
            }),
            submit() {
                console.log('submitted');
                this.error = null;
                this.changeProfile({
                    bio: this.model.bio,
                    handle: this.model.handle
                });
            }
        },
        watch: {
            '$store.state.profile'(value) {
                this.model = value;
            },
            '$store.state.error'(value) {
                this.error = value;
            }
        }
    }
</script>
