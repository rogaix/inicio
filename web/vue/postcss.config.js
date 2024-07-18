module.exports = {
  plugins: {
    tailwindcss: {},
    autoprefixer: {},
  },
  postcssOptions: {
    plugins: [
      require('tailwindcss'),
      require('autoprefixer'),
    ],
  },
};
