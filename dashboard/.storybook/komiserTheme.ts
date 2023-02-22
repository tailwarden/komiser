import { create } from '@storybook/theming/create';

export default create({
  base: 'light',
  brandTitle: 'Komiser Storybook',
  brandUrl: '/',
  brandImage: './assets/img/komiser.svg',
  brandTarget: '_self',
  fontBase: '"Noto Sans", sans-serif',
  fontCode: 'monospace',
  appBg: 'white',
  appContentBg: '#F4F9F9',
  appBorderRadius: 16,
  colorPrimary: '#008484',
  colorSecondary: '#008484'
});

/** 
 *  base: 'light' | 'dark';
    colorPrimary?: string;
    colorSecondary?: string;
    appBg?: string;
    appContentBg?: string;
    appBorderColor?: string;
    appBorderRadius?: number;
    fontBase?: string;
    fontCode?: string;
    textColor?: string;
    textInverseColor?: string;
    textMutedColor?: string;
    barTextColor?: string;
    barSelectedColor?: string;
    barBg?: string;
    buttonBg?: string;
    buttonBorder?: string;
    booleanBg?: string;
    booleanSelectedBg?: string;
    inputBg?: string;
    inputBorder?: string;
    inputTextColor?: string;
    inputBorderRadius?: number;
    brandTitle?: string;
    brandUrl?: string;
    brandImage?: string;
    brandTarget?: string;
    gridCellSize?: number;
 */
