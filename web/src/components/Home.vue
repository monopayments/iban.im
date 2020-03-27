<template>
        <v-card outlined class="home-content">
            <v-tabs dark active-class="active-url" class="main-tab" background-color="teal darken-3 mono-bg"  v-model="activeTab" centered>
                <v-tab v-for="tab of tabs" :key="tab.id" :to="tab.route" exact>
                    <v-icon left>mdi-{{tab.icon}}</v-icon>
                    {{ tab.name }}
                </v-tab>

                <v-tab-item v-for="tab of tabs" :key="tab.id" :value="tab.route">
                    <div class="i-tab">
                        <router-view />
                    </div>
                </v-tab-item>
            </v-tabs>
            <div style="display: none">

                Shorten, create and share memorable links for IBANS.

                <b>IBAN.im/nickname/bank</b>

            </div>
        </v-card>
</template>

<script>
    export default {
        name: "Home",
        data() {
            return {
                activeTab: `/`,
                tabs: [
                    { id: 1, name: "Home", route: `/`, icon: `home` },
                    { id: 2, name: "Login", route: `/login`, icon: `account` },
                    { id: 3, name: "Register", route: `/register`, icon: `account-plus` },
                ]
            }
        },
        created() {
            if("username" in this.$route.params){
                console.log(this.$route.params);
            }else{
                this.$store.state.isLoaded = true
            }
        },
        watch : {
            '$store.state.profile'(profile) {
                if(this.tabs.length < 4) {
                    const route = `/${profile.handle}`;
                    this.tabs.push({
                        id: 4,
                        name: profile.handle,
                        icon: `account`,
                        route
                    });
                    this.activeTab = route;
                }

            },
        }
    }
</script>