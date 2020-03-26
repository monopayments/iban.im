<template>
    <div>

    </div>
</template>

<script>
    import {mapActions} from "vuex";

    export default {
        name: "Logout",
        created() {
            this.setLoaded(true);
            localStorage.removeItem('user');
            location.href = '/login';
        },
        methods: {
            ...mapActions({
                setLoaded: 'setLoaded',
            }),
        }
    }
</script>

<style scoped>

</style>