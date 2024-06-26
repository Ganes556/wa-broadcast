/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ['../view/**/*.html', '../view/**/*.templ', '../view/**/*.go'],
  theme: {
    extend: {},
  },
  daisyui: {
    themes: ['light', 'night'],
  },
  plugins: [require('daisyui')],
};
