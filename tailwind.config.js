module.exports = {
  content: [
    "src/**/*.{ts,html,css,scss}",
    "components/**/*.go"
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('daisyui')
  ],
};