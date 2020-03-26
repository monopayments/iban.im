<template>
    <div>
        To use the app, you'll need to
        <br>
        <router-link to="/login">
            Login
        </router-link>
        or
        <router-link to="/register" >
            Register
        </router-link>
    </div>
</template>

<script>
    export default {
        name: "Main",
        data: () => ({

        }),
        created() {
            this.$store.state.isLoaded = true
        }
    }
</script>