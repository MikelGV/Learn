/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx}",
    "./components/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      fontFamily: {
        bebasneue: "'Bebas neue', 'serif'",
        poppins: "'Poppins', 'serif'"
      }
    },
  },
  plugins: [],
}
