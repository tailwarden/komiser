import { ReactNode } from 'react';

type GridProps = {
  children: ReactNode;
  columns?: number;
  gap?: 'sm' | 'md' | 'lg';
};

function Grid({ children, columns = 2, gap = 'lg' }: GridProps) {
  return (
    <div
      className={`grid grid-cols-1 ${gap === 'sm' ? 'gap-4' : ''} ${
        gap === 'md' ? 'gap-6' : ''
      } ${gap === 'lg' ? 'gap-8' : ''} lg:grid-cols-2`}
    >
      {children}
    </div>
  );
}

export default Grid;
