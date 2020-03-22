<template>
    <div>
        <v-form v-if="this.$store.state.isLoaded" class="profile-form" v-model="isValid" ref="form">
            <v-row>

                <v-col :sm="12" :md="6">
                    <v-text-field
                            v-model="email"
                            label="Email"
                            disabled
                    />
                </v-col>
                <v-col :sm="12" :md="6">
                    <v-text-field
                            v-model="handle"
                            label="Username"
                            disabled
                    />
                </v-col>
                <v-col :sm="12" :md="6">
                    <v-text-field
                            v-model="firstName"
                            label="First Name"
                            :rules="formRules.firstName"
                            required
                    />
                </v-col>
                <v-col :sm="12" :md="6">
                    <v-text-field
                            v-model="lastName"
                            label="Last Name"
                            :rules="formRules.lastName"
                            required
                    />
                </v-col>
                <v-col :sm="12">
                    <v-textarea
                            v-model="bio"
                            label="Bio"
                            rows="3"
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
    import { mapActions } from 'vuex';
    import { mapFields } from "../helper"
    export default {
        name: "Profile",
        data: () => ({
            isValid: true,
            model: null,
            formRules: {
                firstName: [v => !!v || 'Ad zorunlu alandır'],
                lastName: [v => !!v || 'Soyad zorunlu alandır'],
            },
        }),
        computed: {
          ...mapFields({
              fields: ["firstName", "lastName", "handle", "bio","email"],
              base: "profile",
              mutation: "UPDATE_PROFILE"
          }),
        },
        created() {
            this.fetchProfile();
        },
        methods: {
            ...mapActions({
                fetchProfile: 'fetchProfile',
            }),
            submit() {
                console.log('submitted')
            }
        }
    }
</script>
