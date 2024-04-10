/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './**/*.{html,js,ts,jsx,tsx}',
    './**/*.go',
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/typography'),
    require("daisyui")
  ],
}

