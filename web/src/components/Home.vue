<template>
        <v-card outlined raised class="home-content">
            <v-tabs dark active-class="active-url" class="main-tab" background-color="teal darken-3"  v-model="activeTab" centered>
                <v-tab v-for="tab of logged ? tabsLogged : tabs" :key="tab.id" :to="tab.route" exact>
                    <v-icon left>mdi-{{tab.icon}}</v-icon>
                    {{ tab.name }}
                </v-tab>

                <v-tab-item v-for="tab of tabs" :key="tab.id" :value="tab.route">
                    <div class="i-tab">
                        <router-view />
                    </div>
                </v-tab-item>
            </v-tabs>
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
                ],
                tabsLogged: [
                    { id: 1, name: "Profile", route: `/dashboard`, icon: `account` },
                    { id: 2, name: "Security", route: `/dashboard/security`, icon: `account-lock` },
                    { id: 3, name: "Ibans", route: `/dashboard/ibans`, icon: `cash-multiple` },
                    { id: 4, name: "Logout", route: `/dashboard/logout`, icon: `account-arrow-right` },
                ]
            }
        },
        computed:  {
            logged() {
                return "user" in localStorage;
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
                if(!this.logged && this.tabs.length < 4) {
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