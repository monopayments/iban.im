<template>
    <div v-if="profile">
        <div class="profile-info">
            <h1>{{name}}</h1>
        </div>
        <v-list flat>
            <v-list-item-group v-model="selectedIndex" color="primary">
                <v-list-item
                        v-for="(item,i) in $store.state.ibans"
                        :key="i"
                        class="iban-item"
                >
                    <v-btn :to="`/${profile.handle}/${item.handle}`">
                        <span>{{item.handle}}</span>
                    </v-btn>
                </v-list-item>
            </v-list-item-group>
        </v-list>
    </div>
</template>

<script>
    import { mapActions } from 'vuex';
    export default {
        name: "Single",
        data: () => ({
            selectedIndex: undefined,
        }),
        computed : {
            profile() {
                return this.$store.state.profile;
            },
            ibans() {
                return this.$store.state.ibans.filter(iban => !iban.isPrivate)
            },
            name() {
                return this.profile.firstName !== "" ? this.profile.firstName + " " + this.profile.lastName : this.profile.handle;
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
            }),
        }
    }
</script>

<style>

</style>