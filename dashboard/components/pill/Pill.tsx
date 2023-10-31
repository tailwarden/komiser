import { ReactNode } from 'react';

export type PillProps = {
  status:
    | 'active'
    | 'pending'
    | 'removed'
    | 'inactive'
    | 'info'
    | 'new'
    | 'highlight';
  children: ReactNode; // Remove the quotes
  textcase?: 'uppercase' | 'lowercase';
};

function Pill({ status, children, textcase = 'lowercase' }: PillProps) {
  function handleBgColor() {
    let color;

    if (status === 'active') color = 'bg-green-100';
    if (status === 'pending') color = 'bg-orange-100';
    if (status === 'removed') color = 'bg-rose-100';
    if (status === 'inactive') color = 'bg-zinc-100';
    if (status === 'info') color = 'bg-blue-100';
    if (status === 'new') color = 'bg-sky-100';
    if (status === 'highlight') color = 'bg-violet-100';

    return color;
  }

  function handleFontColor() {
    let color;

    if (status === 'active') color = 'text-green-400';
    if (status === 'pending') color = 'text-orange-400';
    if (status === 'removed') color = 'text-red-400';
    if (status === 'inactive') color = 'text-zinc-400';
    if (status === 'info') color = 'text-blue-500';
    if (status === 'new') color = 'text-teal-600';
    if (status === 'highlight') color = 'text-violet-600';

    return color;
  }

  return (
    <div
      className={`inline-flex items-start justify-start gap-2.5 rounded-3xl px-1.5 pb-1 pt-0.5 ${handleBgColor()}`}
    >
      <p
        className={`${handleFontColor()} font-sans text-xs font-normal ${textcase}`}
      >
        {children}
      </p>
    </div>
  );
}

export default Pill;
