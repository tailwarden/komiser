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
  const colors = {
    active: {
      background: 'bg-green-100',
      text: 'text-green-400'
    },
    pending: {
      background: 'bg-orange-100',
      text: 'text-orange-400'
    },
    removed: {
      background: 'bg-rose-100',
      text: 'text-red-400'
    },
    inactive: {
      background: 'bg-zinc-100',
      text: 'text-zinc-400'
    },
    info: {
      background: 'bg-blue-100',
      text: 'text-blue-500'
    },
    new: {
      background: 'bg-sky-100',
      text: 'text-teal-600'
    },
    highlight: {
      background: 'bg-violet-100',
      text: 'text-violet-600'
    }
  };

  const handleColor = () => colors[status];

  return (
    <div
      className={`inline-flex items-start justify-start gap-2.5 rounded-3xl px-1.5 pb-1 pt-0.5 ${
        handleColor().background
      }`}
    >
      <p
        className={`${
          handleColor().text
        } font-sans text-xs font-normal ${textcase}`}
      >
        {children}
      </p>
    </div>
  );
}

export default Pill;
