module.exports = {
  "devServer": {
    "port": 4881,
    // "public": "195.201.97.159:4881",
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