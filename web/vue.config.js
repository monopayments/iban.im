module.exports = {
  "devServer": {
    "port": 4881,
    proxy: {
      '^/graph': {
        target: 'http://localhost:4880',
        changeOrigin: true
      },
    }
  },
  "transpileDependencies": [
    "vuetify"
  ]
};