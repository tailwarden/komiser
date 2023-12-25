/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    './app/**/*.{js,ts,jsx,tsx}',
    './pages/**/*.{js,ts,jsx,tsx}',
    './components/**/*.{js,ts,jsx,tsx}'
  ],
  theme: {
    colors: {
      inherit: 'inherit',
      transparent: 'transparent',
      current: 'currentColor',
      black: '#000000',
      white: '#FFFFFF',
      cyan: {
        50: '#ECFAFA',
        100: '#DFF5F5',
        200: '#CCF2F2',
        300: '#99E5E5',
        400: '#66D9D9',
        500: '#33CCCC',
        600: '#2AA7A7',
        700: '#208282',
        800: '#175D5D',
        900: '#0E3838',
        950: '#051313'
      },
      darkcyan: {
        50: '#EDFAFA',
        100: '#E2F6F6',
        200: '#B6EAEA',
        300: '#63CBCB',
        400: '#2EA8A8',
        500: '#008484',
        600: '#006D6E',
        700: '#065555',
        800: '#004344',
        900: '#002E2F',
        950: '#00191A'
      },
      gray: {
        50: '#F6F9F9',
        100: '#EDF1F2',
        200: '#E4E8E9',
        300: '#D0D6D6',
        400: '#B6BDBD',
        500: '#9AA3A3',
        600: '#828B8B',
        700: '#6B7272',
        800: '#535959',
        900: '#3B4040',
        950: '#0C1717'
      },
      red: {
        50: '#FFE8E8',
        100: '#F7D1D1',
        200: '#F2BABA',
        300: '#EDA3A3',
        400: '#E37575',
        500: '#DE5E5E',
        600: '#BF4F4F',
        700: '#9D4040',
        800: '#7B3131',
        900: '#592222',
        950: '#371313'
      },
      green: {
        50: '#EEFDEE',
        100: '#E1FFE3',
        200: '#CAF5CA',
        300: '#A3E7A6',
        400: '#81CF84',
        500: '#56BA5B',
        600: '#489E4E',
        700: '#3B8240',
        800: '#2E6632',
        900: '#214A24',
        950: '#142E16'
      },
      orange: {
        50: '#FFF8EB',
        100: '#FFF5DA',
        200: '#FCE0AC',
        300: '#F6C879',
        400: '#ECAD4E',
        500: '#ED8F2B',
        600: '#D5721F',
        700: '#A85924',
        800: '#844B2A',
        900: '#664029',
        950: '#422D24'
      },
      blue: {
        50: '#E8EFFD',
        100: '#D4E4FF',
        200: '#ACC7F7',
        300: '#72A1F1',
        400: '#558EEE',
        500: '#387BEB',
        600: '#2F69C6',
        700: '#2656A3',
        800: '#1D4380',
        900: '#14305D',
        950: '#0B1D3A'
      },
      purple: {
        50: '#F7F6FE',
        100: '#EDE8FC',
        200: '#E0DBFB',
        300: '#C9BDF4',
        400: '#AF99EA',
        500: '#9470E0',
        600: '#8157C5',
        700: '#714DA6',
        800: '#5F4585',
        900: '#4B3966',
        950: '#372B4A'
      },
      background: {
        DEFAULT: '#F2FFFF',
        base: '#F2F7F8',
        disabled: '#EFEDF1',
        ds: '#F5F5F5' /* ds => design-system */
      },
      komiser: {
        dark: '#009999'
      }
    },
    boxShadow: {
      right: '2px 4px 8px 0px rgba(105, 115, 114, 0.16)',
      left: '-2px 4px 8px 0px rgba(105, 115, 114, 0.16)',
      none: '0 0 #0000'
    },
    extend: {
      fontFamily: {
        sans: [
          'Noto Sans',
          'ui-sans-serif',
          'system-ui',
          '-apple-system',
          'BlinkMacSystemFont',
          'sans-serif'
        ]
      },
      transitionProperty: {
        width: 'width'
      },
      keyframes: {
        'fade-in': {
          '0%': {
            opacity: 0.5
          },
          '100%': {
            opacity: 1
          }
        },
        'fade-in-down': {
          '0%': {
            opacity: 0.5,
            transform: 'translateY(-15%)'
          },
          '100%': {
            opacity: 1,
            transform: 'translateY(0)'
          }
        },
        'fade-in-down-short': {
          '0%': {
            opacity: 0.5,
            transform: 'translateY(-5%)'
          },
          '100%': {
            opacity: 1,
            transform: 'translateY(0)'
          }
        },
        'fade-in-up': {
          '0%': {
            opacity: 0.5,
            transform: 'translateY(15%)'
          },
          '100%': {
            opacity: 1,
            transform: 'translateY(0)'
          }
        },
        'fade-in-up-short': {
          '0%': {
            opacity: 0.5,
            transform: 'translateY(5%)'
          },
          '100%': {
            opacity: 1,
            transform: 'translateY(0)'
          }
        },
        'fade-in-left': {
          '0%': {
            opacity: 0.5,
            transform: 'translateX(3%)'
          },
          '100%': {
            opacity: 1,
            transform: 'translateY(0)'
          }
        },
        'width-to-fit': {
          '0%': {
            width: '0%'
          },
          '100%': {
            width: '100%'
          }
        },
        'wide-pulse': {
          '0%': {
            transform: 'scale(1, 1)',
            opacity: 0
          },

          '50%': {
            opacity: 0.3
          },

          '100%': {
            transform: 'scale(2.25, 2.25)',
            opacity: 0
          }
        },
        scale: {
          '0%': {
            transform: 'scale(0.97)',
            opacity: 0.85
          },
          '50%': {
            transform: 'scale(1.005)'
          },
          '100%': {
            transform: 'scale(1)',
            opacity: 1
          }
        }
      },
      animation: {
        'fade-in': 'fade-in 250ms ease forwards',
        'fade-in-up': 'fade-in-up 250ms ease forwards',
        'fade-in-up-short': 'fade-in-up-short 250ms ease forwards',
        'fade-in-down': 'fade-in-down 250ms ease forwards',
        'fade-in-down-short': 'fade-in-down-short 250ms ease forwards',
        'fade-in-left': 'fade-in-left 250ms ease forwards',
        'width-to-fit': 'width-to-fit 5000ms ease-in forwards',
        'wide-pulse': 'wide-pulse 2000ms ease-in infinite',
        scale: 'scale 250ms ease forwards'
      },
      backgroundImage: {
        'dependency-graph': 'radial-gradient(#EDEBEE 2px, transparent 0)',
        'empty-cost-explorer':
          "url('/assets/img/others/empty-state-cost-explorer.png')"
      },
      width: {
        'fit-content': 'fit-content'
      }
    }
  },
  future: {
    hoverOnlyWhenSupported: true
  },
  plugins: []
};
