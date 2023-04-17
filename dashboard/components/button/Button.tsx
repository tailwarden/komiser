import { ReactNode } from 'react';
import LoadingSpinner from '../icons/LoadingSpinner';

export type ButtonProps = {
  children: ReactNode;
  type?: 'button' | 'submit';
  style?:
    | 'primary'
    | 'secondary'
    | 'ghost'
    | 'bulk'
    | 'bulk-outline'
    | 'delete'
    | 'delete-ghost';
  size?: 'xs' | 'sm' | 'md' | 'lg';
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
  const xs = 'p-1';
  const sm = 'h-[2.5rem] px-3';
  const md = 'h-[3rem] px-6';
  const lg = 'h-[3.75rem] px-6';

  function handleSize() {
    let buttonSize;

    if (size === 'xs') buttonSize = xs;
    if (size === 'sm') buttonSize = sm;
    if (size === 'md') buttonSize = md;
    if (size === 'lg') buttonSize = lg;

    return buttonSize;
  }

  const base = `${handleSize()} rounded flex items-center ${
    align ? 'justify-start' : 'justify-center '
  } ${
    gap ? 'gap-3' : 'gap-2'
  }  text-sm font-medium box-border w-full sm:w-auto disabled:cursor-not-allowed ${
    transition ? 'transition-colors' : ''
  }`;

  const primary = `${base} bg-gradient-to-br from-primary bg-secondary hover:bg-primary active:from-secondary active:bg-secondary text-white disabled:from-primary disabled:bg-secondary disabled:opacity-50`;

  const secondary = `${base} bg-transparent text-primary border-[1.5px] border-primary hover:bg-komiser-130 active:bg-komiser-200 active:text-primary disabled:bg-transparent disabled:opacity-50`;

  const bulk = `${base} bg-white hover:bg-komiser-200 active:bg-komiser-300 text-secondary  disabled:bg-white disabled:opacity-50`;

  const bulkOutline = `${base} bg-transparent text-white border-[1.5px] border-white hover:bg-komiser-100/10 active:bg-transparent active:border-white/50 active:text-white disabled:bg-transparent disabled:opacity-50`;

  const ghost = `${base} bg-transparent hover:bg-black-400/10 active:bg-black-400/20 text-black-900/60  disabled:bg-transparent disabled:opacity-50`;

  const deleteStyle = `${base} border-[1.5px] border-error-600 text-error-600 hover:bg-error-100 active:bg-error-600/20 disabled:opacity-50`;

  const deleteGhostStyle = `${base} bg-error-100 text-error-600 hover:bg-error-600 hover:text-white active:bg-error-100 active:text-error-600 disabled:bg-error-600 disabled:text-white`;

  function handleStyle() {
    let buttonStyle;

    if (style === 'primary') buttonStyle = primary;
    if (style === 'secondary') buttonStyle = secondary;
    if (style === 'ghost') buttonStyle = ghost;
    if (style === 'bulk') buttonStyle = bulk;
    if (style === 'bulk-outline') buttonStyle = bulkOutline;
    if (style === 'delete-ghost') buttonStyle = deleteGhostStyle;
    if (style === 'delete') buttonStyle = deleteStyle;

    return buttonStyle;
  }

  return (
    <button
      onClick={onClick}
      type={type}
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
