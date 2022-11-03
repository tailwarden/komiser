/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./app/**/*.{js,ts,jsx,tsx}",
    "./pages/**/*.{js,ts,jsx,tsx}",
    "./components/**/*.{js,ts,jsx,tsx}",
  ],
  theme: {
    extend: {
      colors: {
        primary: "#009999",
        secondary: "#5000B5",
        komiser: {
          100: "#F4F9F9",
        },
        warning: {
          100: "#FFF5DA",
          600: "#EDC16B",
        },
        error: {
          100: "#FFE8E8",
          600: "#DE5E5E",
          700: "#ae4242",
          900: "#362033",
        },
        success: {
          100: "#E1FFE3",
          600: "#56BA5B",
        },
        black: {
          100: "#F6F2FB",
          150: "#E9E4EC",
          200: "#D2CADB",
          300: "#978EA1",
          400: "#635972",
          900: "#0C1717",
        },
      },
      fontFamily: {
        sans: [
          "Noto Sans",
          "ui-sans-serif",
          "system-ui",
          "-apple-system",
          "BlinkMacSystemFont",
          "sans-serif",
        ],
      },
    },
  },
  plugins: [],
};
