<template>
    <v-card outlined class="home-content">
        <v-tabs show-arrows dark active-class="active-url" class="main-tab" background-color="teal darken-3 mono-bg"  v-model="activeTab" centered>
            <v-tab v-for="tab of tabs" :key="tab.id" :to="tab.route" exact>
                <v-icon left>mdi-{{tab.icon}}</v-icon>
                {{ tab.name }}
            </v-tab>

            <v-tab-item v-for="tab of tabs" :key="tab.id" :value="tab.route">
                <router-view class="i-tab" />
            </v-tab-item>
        </v-tabs>
    </v-card>
</template>

<script>

    const publicRoutes = [
        { id: 1, name: "Home", route: `/`, icon: `home` },
        { id: 2, name: "Login", route: `/login`, icon: `account` },
        { id: 3, name: "Register", route: `/register`, icon: `account-plus` },
    ];

    const privateRoutes = [
        { id: 1, name: "Profile", route: `/dashboard`, icon: `account` },
        { id: 2, name: "Security", route: `/dashboard/security`, icon: `account-lock` },
        { id: 3, name: "IBANs", route: `/dashboard/ibans`, icon: `cash-multiple` },
        { id: 4, name: "Logout", route: `/dashboard/logout`, icon: `account-arrow-right` },
    ];

    export default {
        name: "Home",
        data() {
            return {
                activeTab: null,
                tabs: publicRoutes
            }
        },
        created() {
            console.log('created');
            console.log(this.$store.state.logged);
            if("username" in this.$route.params){
                console.log(this.$route.params);
            }else{
                this.$store.state.isLoaded = true
            }
            if(this.$store.state.logged) {
                this.tabs = privateRoutes;
            }
            // if(this.$store.state.logged && this.activeTab === '') {
            //     this.activeTab = '/dashboard'
            // }
        },
        watch : {
            '$store.state.logged'(value) {
                if(value) {
                    this.tabs = privateRoutes;
                }else{
                    this.tabs = publicRoutes;
                }
            },
        }
    }
</script>