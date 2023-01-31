import { ReactNode } from 'react';

type TooltipProps = {
  children: ReactNode;
};

function Tooltip({ children }: TooltipProps) {
  return (
    <div className="absolute top-24 left-6 z-20 hidden max-w-[calc(100%-48px)] animate-fade-in-up items-center rounded-lg bg-black-900 py-2 px-4 text-xs font-medium text-black-200 opacity-0 peer-hover:flex">
      {children}
    </div>
  );
}

export default Tooltip;
