/** @type {import('tailwindcss').Config} */
module.exports = {
  theme: {
    extend: {
      fontFamily: {
        sans: ["Inter", "sans-serif"],
      },
    },
  },
  content: ["./views/**/*.templ"],
  plugins: [require("tailwindcss-animate"), require("@tailwindcss/forms")],
};
