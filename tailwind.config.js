/** @type {import('tailwindcss').Config} */
module.exports = {
  content: ["./ui/**/*.html", "./ui/**/*.templ", "./ui/**/*.go"],
  theme: {
    extend: {
      keyframes: {
        fadeInUp: {
          "0%": {
            opacity: "0",
            transform: "translateY(25px)",
          },
          "100%": {
            opacity: "1",
            transform: "translateY(0)",
          },
        },
      },
      animation: {
        fadeInUp: "fadeInUp 0.85s ease-out forwards",
      },
    },
  },
  safelist: [],
};
