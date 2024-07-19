module.exports = {
  transpileDependencies: true,
  configureWebpack: {
    resolve: {
      extensions: ['.ts', '.js', '.vue', '.json']
    }
  },
  css: {
    loaderOptions: {
      postcss: {
        postcssOptions: {
          plugins: [
            require('tailwindcss'),
            require('autoprefixer'),
          ],
        },
      },
    },
  },
};
