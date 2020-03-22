<template>
        <v-card outlined raised center-active class="home-content v-dashboard">
                <v-tabs dark active-class="active-url" class="main-tab" background-color="teal darken-3"  v-model="activeTab" centered>
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
        </v-card>
</template>

<script>
    export default {
        name: "Dashboard",
        data() {
            return {
                activeTab: `/dashboard`,
                tabs: [
                    { id: 1, name: "Profile", route: `/dashboard`, icon: `account` },
                    { id: 2, name: "Security", route: `/dashboard/security`, icon: `account-lock` },
                    { id: 3, name: "Ibans", route: `/dashboard/ibans`, icon: `cash-multiple` },
                ]
            }
        }
    }
</script>