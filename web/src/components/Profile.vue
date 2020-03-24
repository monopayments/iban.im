<template>
    <div>
        <v-form v-if="this.model" class="profile-form" ref="form">
            <v-row>

                <v-col :sm="12" :md="6">
                    <v-text-field
                            v-model="model.email"
                            label="Email"
                            disabled
                    />
                </v-col>
                <v-col :sm="12" :md="6">
                    <v-text-field
                            v-model="model.handle"
                            label="Username"
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
                <v-col :sm="12">
                    <v-textarea
                            v-model="model.bio"
                            label="Bio"
                            rows="3"
                    />
                </v-col>
                <v-col class="fr">
                    <v-btn class="ma-2" color="primary" dark @click="submit">Save</v-btn>
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
            model: null
        }),
        computed: {
            isLoaded() {
                return false;
            }
        },
        created() {
            this.model = {
                ...this.mapFields({
                    fields: ["firstName", "lastName", "handle", "bio","email"],
                    base: "profile",
                    mutation: "CHANGE_PROFILE"
                })
            };
            this.fetchProfile();
        },
        methods: {
            ...mapActions({
                fetchProfile: 'fetchProfile',
                mapFields: 'mapFields',
                changeProfile: 'changeProfile',
            }),
            submit() {
                console.log('submitted');
                this.changeProfile({
                    bio: this.bio
                });
            }
        }
    }
</script>
