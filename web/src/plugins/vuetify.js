import Vue from "vue";
import "../../node_modules/material-design-icons-iconfont/dist/material-design-icons.css"; // Ensure you are using css-loader
import Vuetify from "vuetify/lib"
import colors from "vuetify/lib/util/colors";

Vue.use(Vuetify);

export default new Vuetify({
  icons: {
    iconfont: "md", // default - only for display purposes
  },
  theme: {
    themes: {
      dark: {
        primary: colors.red.darken1, // #E53935
        secondary: colors.red.lighten4, // #FFCDD2
        accent: colors.indigo.base, // #3F51B5
      },
    },
  },
});
