import { ReactNode } from 'react';

export type ButtonProps = {
  children: ReactNode;
  type?: 'button' | 'submit';
  style?:
    | 'primary'
    | 'secondary'
    | 'outline'
    | 'ghost'
    | 'delete'
    | 'delete-ghost';
  size?: 'xs' | 'sm' | 'md' | 'lg';
  disabled?: boolean;
  loading?: boolean;
  onClick?: (e?: any) => void;
};

function Button({
  children,
  type = 'button',
  style = 'primary',
  size = 'md',
  disabled,
  loading,
  onClick
}: ButtonProps) {
  const xs = 'p-3';
  const sm = 'h-[2.5rem] px-6';
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

  const base = `${handleSize()} ${
    size === 'lg' ? 'rounded' : 'rounded-lg'
  } flex items-center justify-center gap-2  text-sm font-medium box-border w-full sm:w-auto disabled:cursor-not-allowed transition-all`;

  const primary = `${base} bg-komiser-600 hover:bg-komiser-700 text-white active:bg-secondary disabled:bg-komiser-600/30`;

  const secondary = `${base} bg-black-100 hover:bg-black-200/50 active:bg-black-100 text-black-400  disabled:bg-black-100 disabled:opacity-50`;

  const outline = `${base} bg-transparent text-secondary border-2 border-secondary hover:bg-komiser-100 active:border-secondary active:text-secondary disabled:bg-transparent disabled:opacity-50`;

  const ghost = `${base} bg-transparent hover:bg-black-400/10 active:bg-black-400/20 text-black-900/60  disabled:bg-transparent disabled:opacity-50`;

  const deleteStyle = `${base} bg-error-600 text-white hover:bg-error-700 active:bg-error-600  disabled:bg-error-700 disabled:text-white/70`;

  const deleteGhostStyle = `${base} bg-transparent text-error-600 hover:bg-error-600 hover:text-white active:bg-error-100 active:text-error-600 disabled:bg-error-600 disabled:text-white`;

  function handleStyle() {
    let buttonStyle;

    if (style === 'primary') buttonStyle = primary;
    if (style === 'secondary') buttonStyle = secondary;
    if (style === 'outline') buttonStyle = outline;
    if (style === 'ghost') buttonStyle = ghost;
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
    >
      {loading && (
        <>
          <svg
            className="animate-spin h-5 w-5 text-inherit"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
          >
            <circle
              className="opacity-25"
              cx="12"
              cy="12"
              r="10"
              stroke="currentColor"
              strokeWidth="4"
            ></circle>
            <path
              className="opacity-75"
              fill="currentColor"
              d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
            ></path>
          </svg>
        </>
      )}
      {children}
    </button>
  );
}

export default Button;
