import { ReactNode } from 'react';
import LoadingSpinner from '../icons/LoadingSpinner';

export type ButtonProps = {
  children: ReactNode;
  type?: 'button' | 'submit';
  style?: 'primary' | 'secondary' | 'ghost' | 'text' | 'dropdown' | 'delete';
  size?: 'xxs' | 'xs' | 'sm' | 'md' | 'lg';
  disabled?: boolean;
  loading?: boolean;
  align?: 'left';
  gap?: 'md';
  transition?: boolean;
  onClick?: (e?: any) => void;
};

function Button({
  children,
  type = 'button',
  style = 'primary',
  size = 'md',
  disabled,
  loading,
  align,
  gap,
  transition = true,
  onClick
}: ButtonProps) {
  const xxs = 'p-1';
  const xs = 'py-2 px-4';
  const sm = 'py-2.5 px-6';
  const md = 'py-3.5 px-6';
  const lg = 'py-4 px-6';

  function handleSize() {
    let buttonSize;

    if (size === 'xxs') buttonSize = xxs;
    if (size === 'xs') buttonSize = xs;
    if (size === 'sm') buttonSize = sm;
    if (size === 'md') buttonSize = md;
    if (size === 'lg') buttonSize = lg;

    return buttonSize;
  }

  const base = `${handleSize()} rounded flex items-center ${
    align ? 'justify-start' : 'justify-center '
  } ${
    gap ? 'gap-3' : 'gap-1'
  }  text-sm font-medium box-border w-full sm:w-auto disabled:cursor-not-allowed ${
    transition ? 'transition-colors' : ''
  }`;

  const primary = `${base} font-semibold bg-gradient-to-br from-primary bg-secondary hover:bg-primary active:from-secondary active:bg-secondary text-white disabled:from-primary disabled:bg-secondary disabled:opacity-50`;

  const secondary = `${base} bg-transparent text-primary border-[1.5px] border-primary hover:bg-komiser-130 active:bg-komiser-200 active:text-primary disabled:bg-transparent disabled:opacity-50`;

  const ghost = `${base} bg-transparent hover:bg-black-100 active:bg-black-400/20 text-black-800  disabled:bg-transparent disabled:opacity-50`;

  const text = `font-semibold text-sm text-komiser-700 hover:underline active:text-black-800`;

  const dropdown = `text-sm font-medium flex items-center gap-2 justify-start p-2 bg-transparent text-black-400 hover:bg-black-150 active:bg-black-200 rounded disabled:bg-transparent disabled:opacity-50`;

  const deleteStyle = `${base} border-[1.5px] border-error-600 text-error-600 hover:bg-error-100 active:bg-error-600/20 disabled:opacity-50`;

  function handleStyle() {
    let buttonStyle;

    if (style === 'primary') buttonStyle = primary;
    if (style === 'secondary') buttonStyle = secondary;
    if (style === 'ghost') buttonStyle = ghost;
    if (style === 'text') buttonStyle = text;
    if (style === 'dropdown') buttonStyle = dropdown;
    if (style === 'delete') buttonStyle = deleteStyle;

    return buttonStyle;
  }

  return (
    <button
      type={type}
      onClick={onClick}
      className={handleStyle()}
      disabled={disabled || loading}
      data-testid={style}
    >
      {loading && <LoadingSpinner />}
      {children}
    </button>
  );
}

export default Button;
