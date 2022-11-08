import { ReactNode } from 'react';

type TooltipProps = {
  children: ReactNode;
};

function Tooltip({ children }: TooltipProps) {
  return (
    <div className="hidden absolute opacity-0 animate-fade-in-up top-24 left-6 right-6 peer-hover:flex items-center py-2 px-4 rounded-lg bg-black-900 text-black-200 text-xs font-medium">
      {children}
    </div>
  );
}

export default Tooltip;
