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
  asLink?: boolean;
  href?: string;
  target?: string;
};

function Button({
  children,
  asLink = false,
  type = 'button',
  style = 'primary',
  size = 'md',
  disabled,
  loading,
  align,
  gap,
  transition = true,
  onClick,
  href,
  target
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

  const primary = `${base} font-semibold bg-gradient-to-br from-darkcyan-500 bg-darkcyan-700 hover:bg-darkcyan-500 active:from-darkcyan-700 active:bg-darkcyan-700 text-white disabled:from-darkcyan-500 disabled:bg-darkcyan-700 disabled:opacity-50`;

  const secondary = `${base} bg-transparent text-darkcyan-500 border-[1.5px] border-darkcyan-500 hover:bg-darkcyan-100 active:bg-cyan-200 active:text-darkcyan-500 disabled:bg-transparent disabled:opacity-50`;

  const ghost = `${base} bg-transparent hover:bg-gray-50 active:bg-gray-300 text-gray-950  disabled:bg-transparent disabled:opacity-50`;

  const text = `font-semibold text-sm text-darkcyan-700 hover:underline active:text-gray-950`;

  const dropdown = `text-sm font-medium flex items-center gap-2 justify-start p-2 bg-transparent text-gray-700 hover:bg-background-ds active:bg-gray-300 rounded disabled:bg-transparent disabled:opacity-50`;

  const deleteStyle = `${base} border-[1.5px] border-red-500 text-red-500 hover:bg-red-50 active:bg-red-100 disabled:opacity-50`;

  function handleStyle() {
    let buttonStyle;

    if (style === 'primary') buttonStyle = primary;
    if (style === 'secondary') buttonStyle = secondary;
    if (style === 'ghost') buttonStyle = ghost;
    if (style === 'text') buttonStyle = text;
    if (style === 'dropdown') buttonStyle = dropdown;
    if (style === 'delete') buttonStyle = deleteStyle;
    if (asLink) buttonStyle = `${buttonStyle} inline-block sm:w-fit-content`;

    return buttonStyle;
  }

  return (
    <>
      {asLink ? (
        <a
          onClick={onClick}
          className={handleStyle()}
          data-testid={style}
          href={href}
          target={target}
        >
          {loading && <LoadingSpinner />}
          {children}
        </a>
      ) : (
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
      )}
    </>
  );
}

export default Button;
