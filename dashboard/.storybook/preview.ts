import '../styles/globals.css';

export const parameters = {
  layout: 'centered',
  backgrounds: {
    default: 'light'
  },
  actions: { argTypesRegex: '^on[A-Z].*' },
  controls: {
    matchers: {
      color: /(background|color)$/i,
      date: /Date$/
    }
  }
};
