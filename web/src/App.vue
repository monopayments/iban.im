<template>
  <div data-app>
    <v-overlay :value="!this.$store.state.isLoaded">
      <v-progress-circular indeterminate size="64" />
    </v-overlay>
    <Nav  />
    <div class="page">
      <v-content>
        <router-view/>
      </v-content>
    </div>
  </div>
</template>

<script>

  import Nav from "./components/Nav";

export default {
  name: 'App',
  components: {
    Nav
  },

  data: () => ({
    //
  }),
};

</script>
<style lang="scss">
  @import './assets/styles/global.scss';
</style>