module.exports = {
  content: [
    "src/**/*.{ts,html,css,scss}",
    "components/**/*.go",
    "dist/**/*.html",
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('daisyui'),
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
    require('@tailwindcss/aspect-ratio'),
    require('@tailwindcss/container-queries'),
  ],
};