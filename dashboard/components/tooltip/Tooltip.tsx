import classNames from 'classnames';
import { ReactNode } from 'react';

type TooltipProps = {
  children: ReactNode;
  top?: 'xs' | 'sm' | 'md' | 'lg';
  bottom?: 'xs' | 'sm' | 'md' | 'lg';
  align?: 'left' | 'center' | 'right';
  width?: 'sm' | 'md' | 'lg' | 'xl';
};

function Tooltip({
  children,
  top = 'md',
  align = 'left',
  width = 'md',
  bottom
}: TooltipProps) {
  return (
    <div
      role="tooltip"
      className={classNames(
        'absolute z-[1000] hidden animate-fade-in-up items-center rounded-lg bg-gray-950 px-4 py-2 text-xs font-medium text-gray-300 opacity-0 peer-hover:flex',
        { 'top-0': top === 'xs' && !bottom },
        { 'top-[3rem]': top === 'sm' && !bottom },
        { 'top-24': top === 'md' && !bottom },
        { 'top-36': top === 'lg' && !bottom },
        { 'bottom-0': bottom === 'xs' },
        { 'bottom-[3rem]': bottom === 'sm' },
        { 'bottom-24': bottom === 'md' },
        { 'bottom-36': bottom === 'lg' },
        { 'left-6': align === 'left' },
        { 'right-6': align === 'right' },
        { 'w-72': width === 'lg' },
        { 'w-[30.5rem]': width === 'xl' }
      )}
    >
      {children}
    </div>
  );
}

export default Tooltip;
