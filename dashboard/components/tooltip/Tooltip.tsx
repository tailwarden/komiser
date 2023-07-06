import classNames from 'classnames';
import { ReactNode } from 'react';

type TooltipProps = {
  children: ReactNode;
  top?: 'sm' | 'md' | 'lg';
  align?: 'left' | 'center' | 'right';
  width?: 'sm' | 'md' | 'lg';
};

function Tooltip({
  children,
  top = 'md',
  align = 'left',
  width = 'md'
}: TooltipProps) {
  return (
    <div
      role="tooltip"
      className={classNames(
        'absolute left-6 top-24 z-[1000] hidden animate-fade-in-up items-center rounded-lg bg-black-900 px-4 py-2 text-xs font-medium text-black-200 opacity-0 peer-hover:flex',
        { 'top-[3rem]': top === 'sm' },
        { 'left-auto right-0': align === 'right' },
        { 'w-72': width === 'lg' }
      )}
    >
      {children}
    </div>
  );
}

export default Tooltip;
