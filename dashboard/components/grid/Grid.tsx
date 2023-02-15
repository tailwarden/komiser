import { ReactNode } from 'react';

type GridProps = {
  children: ReactNode;
  columns?: number;
};

function Grid({ children, columns = 2 }: GridProps) {
  return (
    <div className="grid grid-cols-1 gap-8 lg:grid-cols-2">{children}</div>
  );
}

export default Grid;
